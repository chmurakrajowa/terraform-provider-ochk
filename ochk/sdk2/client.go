package sdk

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/openapi/v3"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/runtime/logger"
	"net/http"
	"sync"
	"time"
)

type Client struct {
	FloatingIPAddresses FloatingIPAddressProxy
	FloatingIPVms       FloatingIPVmsProxy
	FirewallRules       FirewallRulesProxy
	FirewallEWRules     FirewallEWRulesProxy
	FirewallSNRules     FirewallSNRulesProxy
	Requests            RequestsProxy
	Routers             RoutersProxy
	SecurityGroups      SecurityGroupsProxy
	Services            ServicesProxy
	Projects            ProjectsProxy
	VirtualMachines     VirtualMachinesProxy
	VirtualNetworks     VirtualNetworksProxy
	IPCollections       IPCollectionsProxy
	Deployments         DeploymentsProxy
	CustomServices      CustomServicesProxy
	KMSKeys             KMSKeysProxy
	BackupPlans         BackupPlansProxy
	BackupLists         BackupListsProxy
	Tags                TagsProxy
	Nats                NatProxy
	PortForwarding      PortsForwardingProxy
	Folders             FoldersProxy
	PublicIPAddresses   PublicIPAddressProxy
	Snapshots           SnapshotsProxy
	Accounts            AccountsProxy
	PlatformType        PlatformTypeProxy
	key                 string
	PType               openapi.PlatformType
	apiClientTransport  httptransport.Runtime
}

var clientMutex sync.Mutex

type myTransport struct {
}

var PLATFORM = ""
var API_KEY = ""
var PLATFORM_TYPE = ""

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

var E1000 = "ERROR{1000}: Check input variables. Selected platform: \"%s\" is not from indicated virtualization platform: \"%s\"."

func assign(platform_type string, platform string, api_key string) {
	PLATFORM = platform
	API_KEY = api_key
	PLATFORM_TYPE = platform_type
}

func (t *myTransport) RoundTrip(req *http.Request) (*http.Response, error) {

	req.Header.Add("platform", PLATFORM)
	req.Header.Add("x-api-key", API_KEY)
	return http.DefaultTransport.RoundTrip(req)
}

func NewClient(ctx context.Context, host string, platform string, api_key string, insecure bool, debugLogFile string, platformType string) (*Client, error) {

	clientMutex.Lock()
	defer clientMutex.Unlock()
	assign(platformType, platform, api_key)

	if c := getClientFromCache(host, platform, api_key, insecure, debugLogFile); c != nil {
		return c, nil
	}

	var defaultLogger logger.Logger = NewStdErrLogger()
	if debugLogFile != "" {
		fileLogger := NewFileLogger(debugLogFile)
		if err := fileLogger.Init(); err != nil {
			return nil, fmt.Errorf("error initializing file logger: %v", err)
		}
		defaultLogger = fileLogger
	}

	//httpClient := &http.Client{
	//	Transport: &http.Transport{
	//		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//	},
	//}

	httpClient := &http.Client{
		Transport: &myTransport{},
	}

	apiClientTransport := httptransport.New(host, DefaultBasePath, mapToSchemes(insecure))
	apiClientTransport.SetDebug(true)
	apiClientTransport.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	if defaultLogger != nil {
		apiClientTransport.SetLogger(defaultLogger)
	}

	apiClientAuthTransport := httptransport.New(host, DefaultBasePath, mapToSchemes(insecure))
	apiClientAuthTransport.SetDebug(true)
	if defaultLogger != nil {
		apiClientAuthTransport.SetLogger(defaultLogger)
	}

	configuration := openapi.NewConfiguration()
	apiClient := openapi.NewAPIClient(configuration)

	c := &Client{
		SecurityGroups: SecurityGroupsProxy{
			httpClient: httpClient,
			service:    apiClient.SecurityGroupAPI,
		},
		FirewallEWRules: FirewallEWRulesProxy{
			httpClient: httpClient,
			service:    apiClient.DfwRuleAPI,
		},
		FirewallSNRules: FirewallSNRulesProxy{
			httpClient: httpClient,
			service:    apiClient.GfwRuleAPI,
		},
		FirewallRules: FirewallRulesProxy{
			httpClient: httpClient,
			service:    apiClient.FirewallRuleAPI,
		},
		Services: ServicesProxy{
			httpClient: httpClient,
			service:    apiClient.DefaultServicesAPI,
		},
		Routers: RoutersProxy{
			httpClient: httpClient,
			service:    apiClient.RouterAPI,
		},
		VirtualMachines: VirtualMachinesProxy{
			httpClient: httpClient,
			service:    apiClient.VirtualMachineAPI,
		},
		Projects: ProjectsProxy{
			httpClient: httpClient,
			service:    apiClient.ProjectsAPI,
		},
		VirtualNetworks: VirtualNetworksProxy{
			httpClient: httpClient,
			service:    apiClient.VirtualNetworkAPI,
		},
		Requests: RequestsProxy{
			httpClient: httpClient,
			service:    apiClient.RequestsAPI,
		},
		IPCollections: IPCollectionsProxy{
			httpClient: httpClient,
			service:    apiClient.IpCollectionAPI,
		},
		Deployments: DeploymentsProxy{
			httpClient: httpClient,
			service:    apiClient.DeploymentsAPI,
		},
		CustomServices: CustomServicesProxy{
			httpClient: httpClient,
			service:    apiClient.CustomServicesAPI,
		},
		KMSKeys: KMSKeysProxy{
			httpClient: httpClient,
			service:    apiClient.KeyAPI,
		},
		BackupPlans: BackupPlansProxy{
			httpClient: httpClient,
			service:    apiClient.BackupsAPI,
		},
		BackupLists: BackupListsProxy{
			httpClient: httpClient,
			service:    apiClient.BackupsAPI,
		},
		Tags: TagsProxy{
			httpClient: httpClient,
			service:    apiClient.TagsAPI,
		},
		Nats: NatProxy{
			httpClient: httpClient,
			service:    apiClient.NatRuleAPI,
		},
		PortForwarding: PortsForwardingProxy{
			httpClient: httpClient,
			service:    apiClient.PortForwardingAPI,
		},
		Folders: FoldersProxy{
			httpClient: httpClient,
			service:    apiClient.FolderAPI,
		},
		PublicIPAddresses: PublicIPAddressProxy{
			httpClient: httpClient,
			service:    apiClient.PublicIpAPI,
		},
		FloatingIPAddresses: FloatingIPAddressProxy{
			httpClient: httpClient,
			service:    apiClient.FloatingIpAPI,
		},
		FloatingIPVms: FloatingIPVmsProxy{
			httpClient: httpClient,
			service:    apiClient.FloatingIpVmsAPI,
		},
		Snapshots: SnapshotsProxy{
			httpClient: httpClient,
			service:    apiClient.VirtualMachineSnapshotAPI,
		},
		Accounts: AccountsProxy{
			httpClient: httpClient,
			service:    apiClient.AccountsAPI,
		},
		PlatformType: PlatformTypeProxy{
			httpClient: httpClient,
			service:    apiClient.IdentificationAPI,
		},
	}

	c.apiClientTransport = *apiClientAuthTransport

	platformTypeAPI, err := checkPlatformType(ctx, c)

	if err != nil {
		return nil, err
	}

	if string(platformTypeAPI) != PLATFORM_TYPE {
		return nil, fmt.Errorf(E1000, PLATFORM, PLATFORM_TYPE)
	}
	c.PType = platformTypeAPI
	c.key = cacheClient(c, host, platform, api_key, insecure, debugLogFile, &ctx)
	return c, nil
}

type cachedClient struct {
	c         *Client
	cacheTime time.Time
	ctx       *context.Context
}

func checkPlatformType(ctx context.Context, c *Client) (openapi.PlatformType, error) {
	proxy := c.PlatformType
	platformType, err := proxy.Read(ctx)
	if err != nil {
		return "UNKNOWN", fmt.Errorf("error checking platform type. : %v", err)
	}
	return platformType, nil
}

var clientCacheLifetime = time.Minute * 5
var clientCache = map[string]cachedClient{}

func clientCacheKey(host string, platform string, api_key string, insecure bool, file string) string {
	return fmt.Sprintf("%s_%s_%s_%t_%s", host, platform, api_key, insecure, file)
}

func getClientFromCacheByKey(key string) *Client {
	if clientFromCache, ok := clientCache[key]; ok {
		if time.Since(clientFromCache.cacheTime) > clientCacheLifetime {
			//log.Printf("Evicting expired client from cache by key: %s", key)
			return nil
		}

		//log.Printf("Returning client from cache by key: %s", key)
		return clientFromCache.c
	}

	return nil
}

func getClientFromCache(host string, platform string, api_key string, insecure bool, file string) *Client {
	key := clientCacheKey(host, platform, api_key, insecure, file)
	return getClientFromCacheByKey(key)

}

func cacheClient(c *Client, host string, platform string, username string, insecure bool, debugLogFile string, context *context.Context) string {
	key := clientCacheKey(host, platform, username, insecure, debugLogFile)
	//log.Printf("Putting client into cache by key: %s", key)
	clientCache[key] = cachedClient{
		c:         c,
		cacheTime: time.Now(),
		ctx:       context,
	}
	return key
}

func mapToSchemes(insecure bool) []string {
	if insecure {
		return []string{"http"}
	}

	return []string{"https"}
}
