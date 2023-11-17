package sdk

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/w_s_o2_logout_token"
	wso2controller "github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/w_s_o2_token"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
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
	key                string
	apiClientTransport httptransport.Runtime
}

var clientMutex sync.Mutex

func Logout() error {

	clientMutex.Lock()
	defer clientMutex.Unlock()

	for key, caschedClient := range clientCache {

		ochkClient := client.New(&caschedClient.c.apiClientTransport, strfmt.Default)

		httpClient := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}

		params := w_s_o2_logout_token.LogoutTokenUsingPOSTParams{
			Context:    *caschedClient.ctx,
			HTTPClient: httpClient,
		}

		authResponse, err := ochkClient.WsO2LogoutToken.LogoutTokenUsingPOST(&params)
		if err != nil {
			return fmt.Errorf("error while logout: %+v", err)
		}

		if !authResponse.Payload.Success {
			return fmt.Errorf("logout token failed: %s", authResponse.Payload.Messages)
		}
		delete(clientCache, key)

	}

	return nil
}

func NewClient(ctx context.Context, host string, platform string, username string, password string, insecure bool, debugLogFile string) (*Client, error) {

	clientMutex.Lock()
	defer clientMutex.Unlock()

	if c := getClientFromCache(host, platform, username, insecure, debugLogFile); c != nil {
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

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	apiClientTransport := httptransport.New(host, client.DefaultBasePath, mapToSchemes(insecure))
	apiClientTransport.SetDebug(true)
	apiClientTransport.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	if defaultLogger != nil {
		apiClientTransport.SetLogger(defaultLogger)
	}

	ochkClient := client.New(apiClientTransport, strfmt.Default)

	params := wso2controller.GetTokenUsingPOSTParams{
		WsoTokenRequest: &models.WSOTokenRequest{
			Platform: platform,
			Password: password,
			Username: username,
		},
		Context:    ctx,
		HTTPClient: httpClient,
	}

	authResponse, err := ochkClient.WsO2Token.GetTokenUsingPOST(&params)
	if err != nil {
		return nil, fmt.Errorf("error while retrieving auth token: %+v", err)
	}

	if !authResponse.Payload.Success {
		return nil, fmt.Errorf("retrieving auth token failed: %s", authResponse.Payload.Messages)
	}

	apiClientAuthTransport := httptransport.New(host, client.DefaultBasePath, mapToSchemes(insecure))
	apiClientAuthTransport.SetDebug(true)
	if defaultLogger != nil {
		apiClientAuthTransport.SetLogger(defaultLogger)
	}
	apiClientAuthTransport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", authResponse.Payload.Token)

	authClient := client.New(apiClientAuthTransport, strfmt.Default)

	c := &Client{
		SecurityGroups: SecurityGroupsProxy{
			httpClient: httpClient,
			service:    authClient.SecurityGroups,
		},
		FirewallEWRules: FirewallEWRulesProxy{
			httpClient: httpClient,
			service:    authClient.FirewallRulesew,
		},
		FirewallSNRules: FirewallSNRulesProxy{
			httpClient: httpClient,
			service:    authClient.FirewallRulessn,
		},
		Services: ServicesProxy{
			httpClient: httpClient,
			service:    authClient.DefaultServices,
		},
		Routers: RoutersProxy{
			httpClient: httpClient,
			service:    authClient.Routers,
		},
		VirtualMachines: VirtualMachinesProxy{
			httpClient: httpClient,
			service:    authClient.VirtualMachines,
		},
		Projects: ProjectsProxy{
			httpClient: httpClient,
			service:    authClient.Projects,
		},
		VirtualNetworks: VirtualNetworksProxy{
			httpClient: httpClient,
			service:    authClient.VirtualNetworks,
		},
		Requests: RequestsProxy{
			httpClient: httpClient,
			service:    authClient.Requests,
		},
		IPCollections: IPCollectionsProxy{
			httpClient: httpClient,
			service:    authClient.IPCollections,
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
			service:    authClient.KmsKeyManagement,
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
			service:    authClient.NatRules,
		},
		Folders: FoldersProxy{
			httpClient: httpClient,
			service:    authClient.Folder,
		},
		PublicIPAddresses: PublicIPAddressProxy{
			httpClient: httpClient,
			service:    authClient.IPamPublicIPAllocations,
		},
		Snapshots: SnapshotsProxy{
			httpClient: httpClient,
			service:    authClient.VirtualMachines,
		},
		Accounts: AccountsProxy{
			httpClient: httpClient,
			service:    authClient.BillingAccounts,
		},
	}

	c.apiClientTransport = *apiClientAuthTransport
	c.key = cacheClient(c, host, platform, username, insecure, debugLogFile, &ctx)
	return c, nil
}

type cachedClient struct {
	c         *Client
	cacheTime time.Time
	ctx       *context.Context
}

var clientCacheLifetime = time.Minute * 5
var clientCache = map[string]cachedClient{}

func clientCacheKey(host string, platform string, username string, insecure bool, file string) string {
	return fmt.Sprintf("%s_%s_%s_%t_%s", host, platform, username, insecure, file)
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

func getClientFromCache(host string, platform string, username string, insecure bool, file string) *Client {
	key := clientCacheKey(host, platform, username, insecure, file)
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
