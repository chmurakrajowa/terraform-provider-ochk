package sdk

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client"
	vidmcontroller "github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/v_id_m"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/runtime/logger"
	"github.com/go-openapi/strfmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Client struct {
	FirewallEWRules FirewallEWRulesProxy
	FirewallSNRules FirewallSNRulesProxy
	GatewayPolicy   GatewayPolicyProxy
	LogicalPorts    LogicalPortsProxy
	Networks        NetworksProxy
	Requests        RequestsProxy
	Routers         RoutersProxy
	SecurityGroups  SecurityGroupsProxy
	SecurityPolicy  SecurityPolicyProxy
	Services        ServicesProxy
	Subtenants      SubtenantsProxy
	Users           UsersProxy
	VirtualMachines VirtualMachinesProxy
	VirtualNetworks VirtualNetworksProxy
	IPCollections   IPCollectionsProxy
	Deployments     DeploymentsProxy
	CustomServices  CustomServicesProxy
	KMSKeys         KMSKeysProxy
}

var clientMutex sync.Mutex

func NewClient(ctx context.Context, host string, tenant string, username string, password string, insecure bool, debugLogFile string) (*Client, error) {
	clientMutex.Lock()
	defer clientMutex.Unlock()

	if c := getClientFromCache(host, tenant, username, insecure, debugLogFile); c != nil {
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

	params := vidmcontroller.GetTokenUsingPOSTParams{
		VidmTokenRequest: &models.VIDMTokenRequest{
			Tenant:   tenant,
			Password: password,
			Username: username,
		},
		Context:    ctx,
		HTTPClient: httpClient,
	}

	authResponse, err := ochkClient.VIDm.GetTokenUsingPOST(&params)
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
	apiClientAuthTransport.DefaultAuthentication = httptransport.APIKeyAuth("token", "header", authResponse.Payload.Token)

	authClient := client.New(apiClientAuthTransport, strfmt.Default)

	c := &Client{
		SecurityGroups: SecurityGroupsProxy{
			httpClient: httpClient,
			service:    authClient.SecurityGroups,
		},
		SecurityPolicy: SecurityPolicyProxy{
			httpClient: httpClient,
			service:    authClient.SecurityPolicies,
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
		GatewayPolicy: GatewayPolicyProxy{
			httpClient: httpClient,
			service:    authClient.GatewayPolicies,
		},
		Routers: RoutersProxy{
			httpClient: httpClient,
			service:    authClient.Routers,
		},
		VirtualMachines: VirtualMachinesProxy{
			httpClient:    httpClient,
			service:       authClient.VirtualMachines,
			legacyService: authClient.VirtualMachinesnsx,
		},
		LogicalPorts: LogicalPortsProxy{
			httpClient: httpClient,
			service:    authClient.LogicalPorts,
		},
		Users: UsersProxy{
			httpClient: httpClient,
			service:    authClient.Users,
		},
		Subtenants: SubtenantsProxy{
			httpClient: httpClient,
			service:    authClient.Subtenants,
		},
		Networks: NetworksProxy{
			httpClient: httpClient,
			service:    authClient.Networks,
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
	}

	cacheClient(c, host, tenant, username, insecure, debugLogFile)

	return c, nil
}

type cachedClient struct {
	c         *Client
	cacheTime time.Time
}

var clientCacheLifetime = time.Minute * 5
var clientCache = map[string]cachedClient{}

func clientCacheKey(host string, tenant string, username string, insecure bool, file string) string {
	return fmt.Sprintf("%s_%s_%s_%t_%s", host, tenant, username, insecure, file)
}

func getClientFromCache(host string, tenant string, username string, insecure bool, file string) *Client {
	key := clientCacheKey(host, tenant, username, insecure, file)
	if clientFromCache, ok := clientCache[key]; ok {
		if time.Since(clientFromCache.cacheTime) > clientCacheLifetime {
			log.Printf("Evicting expired client from cache by key: %s", key)
			return nil
		}

		log.Printf("Returning client from cache by key: %s", key)
		return clientFromCache.c
	}

	return nil
}

func cacheClient(c *Client, host string, tenant string, username string, insecure bool, debugLogFile string) {
	key := clientCacheKey(host, tenant, username, insecure, debugLogFile)
	log.Printf("Putting client into cache by key: %s", key)
	clientCache[key] = cachedClient{
		c:         c,
		cacheTime: time.Now(),
	}
}

func mapToSchemes(insecure bool) []string {
	if insecure {
		return []string{"http"}
	}

	return []string{"https"}
}
