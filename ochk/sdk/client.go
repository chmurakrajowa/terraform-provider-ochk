package sdk

import (
	"context"
	"crypto/tls"
	"fmt"
	openapi "github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/runtime/logger"
	"log"
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
	AvailablePublicIp   AvailablePublicIpProxy
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
var HOST = ""

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	//DefaultHost string = "localhost" // set if provider is run locally build from code
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

var E1000 = "ERROR{1000}: Check input variables. Selected platform: \"%s\" is not from indicated virtualization platform: \"%s\"."

func assign(platform_type string, platform string, api_key string, host string) {
	PLATFORM = platform
	API_KEY = api_key
	PLATFORM_TYPE = platform_type
	HOST = host
}

func (t *myTransport) RoundTrip(req *http.Request) (*http.Response, error) {

	req.Header.Add("platform", PLATFORM)
	req.Header.Add("x-api-key", API_KEY)
	return http.DefaultTransport.RoundTrip(req)
}

func NewClient(ctx context.Context, host string, platform string, api_key string, insecure bool, debugLogFile string, platformType string) (*Client, error) {

	clientMutex.Lock()
	defer clientMutex.Unlock()
	assign(platformType, platform, api_key, host)

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
	configuration.HTTPClient = httpClient
	if insecure {
		configuration.Servers = openapi.ServerConfigurations{
			{
				URL: "http://" + HOST,
			},
		}
		log.Printf("Base URL: %+v", configuration.Servers)

		configuration.Scheme = "http"
	} else {
		configuration.Servers = openapi.ServerConfigurations{
			{
				URL: HOST,
			},
		}
		configuration.Scheme = "https"

	}
	configuration.Debug = true

	authClient := openapi.NewAPIClient(configuration)

	c := &Client{
		SecurityGroups: SecurityGroupsProxy{
			httpClient: httpClient,
			service:    authClient.SecurityGroupAPI,
		},
		FirewallEWRules: FirewallEWRulesProxy{
			httpClient: httpClient,
			service:    authClient.DfwRuleAPI,
		},
		FirewallSNRules: FirewallSNRulesProxy{
			httpClient: httpClient,
			service:    authClient.GfwRuleAPI,
		},
		FirewallRules: FirewallRulesProxy{
			httpClient: httpClient,
			service:    authClient.FirewallRuleAPI,
		},
		Services: ServicesProxy{
			httpClient: httpClient,
			service:    authClient.DefaultServicesAPI,
		},
		Routers: RoutersProxy{
			httpClient: httpClient,
			service:    authClient.RouterAPI,
		},
		VirtualMachines: VirtualMachinesProxy{
			httpClient: httpClient,
			service:    authClient.VirtualMachineAPI,
		},
		Projects: ProjectsProxy{
			httpClient: httpClient,
			service:    authClient.ProjectsAPI,
		},
		VirtualNetworks: VirtualNetworksProxy{
			httpClient: httpClient,
			service:    authClient.VirtualNetworkAPI,
		},
		Requests: RequestsProxy{
			httpClient: httpClient,
			service:    authClient.RequestsAPI,
		},
		IPCollections: IPCollectionsProxy{
			httpClient: httpClient,
			service:    authClient.IpCollectionAPI,
		},
		Deployments: DeploymentsProxy{
			httpClient: httpClient,
			service:    authClient.DeploymentsAPI,
		},
		CustomServices: CustomServicesProxy{
			httpClient: httpClient,
			service:    authClient.CustomServicesAPI,
		},
		KMSKeys: KMSKeysProxy{
			httpClient: httpClient,
			service:    authClient.KeyAPI,
		},
		BackupPlans: BackupPlansProxy{
			httpClient: httpClient,
			service:    authClient.BackupsAPI,
		},
		BackupLists: BackupListsProxy{
			httpClient: httpClient,
			service:    authClient.BackupsAPI,
		},
		Tags: TagsProxy{
			httpClient: httpClient,
			service:    authClient.TagsAPI,
		},
		Nats: NatProxy{
			httpClient: httpClient,
			service:    authClient.NatRuleAPI,
		},
		PortForwarding: PortsForwardingProxy{
			httpClient: httpClient,
			service:    authClient.PortForwardingAPI,
		},
		Folders: FoldersProxy{
			httpClient: httpClient,
			service:    authClient.FolderAPI,
		},
		PublicIPAddresses: PublicIPAddressProxy{
			httpClient: httpClient,
			service:    authClient.PublicIpAPI,
		},
		AvailablePublicIp: AvailablePublicIpProxy{
			httpClient: httpClient,
			service:    authClient.AvailablePublicIpAPI,
		},
		FloatingIPAddresses: FloatingIPAddressProxy{
			httpClient: httpClient,
			service:    authClient.FloatingIpAPI,
		},
		FloatingIPVms: FloatingIPVmsProxy{
			httpClient: httpClient,
			service:    authClient.FloatingIpVmsAPI,
		},
		Snapshots: SnapshotsProxy{
			httpClient: httpClient,
			service:    authClient.VirtualMachineSnapshotAPI,
		},
		Accounts: AccountsProxy{
			httpClient: httpClient,
			service:    authClient.AccountsAPI,
		},
		PlatformType: PlatformTypeProxy{
			httpClient: httpClient,
			service:    authClient.IdentificationAPI,
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
