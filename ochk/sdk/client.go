package sdk

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/runtime/logger"
	"github.com/go-openapi/strfmt"
	"net/http"
	"sync"
	"time"
)

type Client struct {
	FirewallEWRules    FirewallEWRulesProxy
	FirewallSNRules    FirewallSNRulesProxy
	Requests           RequestsProxy
	Routers            RoutersProxy
	SecurityGroups     SecurityGroupsProxy
	Services           ServicesProxy
	Projects           ProjectsProxy
	VirtualMachines    VirtualMachinesProxy
	VirtualNetworks    VirtualNetworksProxy
	IPCollections      IPCollectionsProxy
	Deployments        DeploymentsProxy
	CustomServices     CustomServicesProxy
	KMSKeys            KMSKeysProxy
	BackupPlans        BackupPlansProxy
	BackupLists        BackupListsProxy
	Tags               TagsProxy
	Nats               NatProxy
	Folders            FoldersProxy
	PublicIPAddresses  PublicIPAddressProxy
	Snapshots          SnapshotsProxy
	Accounts           AccountsProxy
	PlatformType       PlatformTypeProxy
	key                string
	PType              models.PlatformType
	apiClientTransport httptransport.Runtime
}

var clientMutex sync.Mutex

type myTransport struct {
}

var PLATFORM = ""
var API_KEY = ""

func assign(platform string, api_key string) {
	PLATFORM = platform
	API_KEY = api_key
}

func (t *myTransport) RoundTrip(req *http.Request) (*http.Response, error) {

	req.Header.Add("platform", PLATFORM)
	req.Header.Add("x-api-key", API_KEY)
	return http.DefaultTransport.RoundTrip(req)
}

func NewClient(ctx context.Context, host string, platform string, api_key string, insecure bool, debugLogFile string) (*Client, error) {

	clientMutex.Lock()
	defer clientMutex.Unlock()
	assign(platform, api_key)

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

	apiClientTransport := httptransport.New(host, client.DefaultBasePath, mapToSchemes(insecure))
	apiClientTransport.SetDebug(true)
	apiClientTransport.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	if defaultLogger != nil {
		apiClientTransport.SetLogger(defaultLogger)
	}

	apiClientAuthTransport := httptransport.New(host, client.DefaultBasePath, mapToSchemes(insecure))
	apiClientAuthTransport.SetDebug(true)
	if defaultLogger != nil {
		apiClientAuthTransport.SetLogger(defaultLogger)
	}

	authClient := client.New(apiClientAuthTransport, strfmt.Default)

	c := &Client{
		SecurityGroups: SecurityGroupsProxy{
			httpClient: httpClient,
			service:    authClient.SecurityGroup,
		},
		FirewallEWRules: FirewallEWRulesProxy{
			httpClient: httpClient,
			service:    authClient.DfwRule,
		},
		FirewallSNRules: FirewallSNRulesProxy{
			httpClient: httpClient,
			service:    authClient.GfwRule,
		},
		Services: ServicesProxy{
			httpClient: httpClient,
			service:    authClient.DefaultServices,
		},
		Routers: RoutersProxy{
			httpClient: httpClient,
			service:    authClient.Router,
		},
		VirtualMachines: VirtualMachinesProxy{
			httpClient: httpClient,
			service:    authClient.VirtualMachine,
		},
		Projects: ProjectsProxy{
			httpClient: httpClient,
			service:    authClient.Projects,
		},
		VirtualNetworks: VirtualNetworksProxy{
			httpClient: httpClient,
			service:    authClient.VirtualNetwork,
		},
		Requests: RequestsProxy{
			httpClient: httpClient,
			service:    authClient.Requests,
		},
		IPCollections: IPCollectionsProxy{
			httpClient: httpClient,
			service:    authClient.IPCollection,
		},
		Deployments: DeploymentsProxy{
			httpClient: httpClient,
			service:    authClient.Deployments,
		},
		CustomServices: CustomServicesProxy{
			httpClient: httpClient,
			service:    authClient.CustomServices,
		},
		KMSKeys: KMSKeysProxy{
			httpClient: httpClient,
			service:    authClient.Key,
		},
		BackupPlans: BackupPlansProxy{
			httpClient: httpClient,
			service:    authClient.Backups,
		},
		BackupLists: BackupListsProxy{
			httpClient: httpClient,
			service:    authClient.Backups,
		},
		Tags: TagsProxy{
			httpClient: httpClient,
			service:    authClient.Tags,
		},
		Nats: NatProxy{
			httpClient: httpClient,
			service:    authClient.NatRule,
		},
		Folders: FoldersProxy{
			httpClient: httpClient,
			service:    authClient.Folder,
		},
		PublicIPAddresses: PublicIPAddressProxy{
			httpClient: httpClient,
			service:    authClient.PublicIP,
		},
		Snapshots: SnapshotsProxy{
			httpClient: httpClient,
			service:    authClient.VirtualMachineSnapshot,
		},
		Accounts: AccountsProxy{
			httpClient: httpClient,
			service:    authClient.Accounts,
		},
		PlatformType: PlatformTypeProxy{
			httpClient: httpClient,
			service:    authClient.Identification,
		},
	}

	c.apiClientTransport = *apiClientAuthTransport

	platformType, err := checkPlatformType(ctx, c)

	if err != nil {
		return nil, err
	}

	c.PType = platformType
	c.key = cacheClient(c, host, platform, api_key, insecure, debugLogFile, &ctx)
	return c, nil
}

type cachedClient struct {
	c         *Client
	cacheTime time.Time
	ctx       *context.Context
}

func checkPlatformType(ctx context.Context, c *Client) (models.PlatformType, error) {
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
