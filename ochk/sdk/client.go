package sdk

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client"
	vidmcontroller "github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/v_id_m"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"net/http"
)

type Client struct {
	logger          *FileLogger
	SecurityGroups  SecurityGroupsProxy
	SecurityPolicy  SecurityPolicyProxy
	GatewayPolicy   GatewayPolicyProxy
	FirewallEWRules FirewallEWRulesProxy
	FirewallSNRules FirewallSNRulesProxy
	Services        ServicesProxy
	Routers         RoutersProxy
	VirtualMachines VirtualMachinesProxy
	IPSets          IPSetsProxy
}

func NewClient(ctx context.Context, host string, tenant string, username string, password string, insecure bool, debugLogFile string) (*Client, error) {
	var logger *FileLogger
	if debugLogFile != "" {
		logger = NewFileLogger(debugLogFile)
		if err := logger.Init(); err != nil {
			return nil, fmt.Errorf("error initializing file logger: %v", err)
		}
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
	if logger != nil {
		apiClientAuthTransport.SetLogger(logger)
	}
	apiClientAuthTransport.DefaultAuthentication = httptransport.APIKeyAuth("token", "header", authResponse.Payload.Token)

	authClient := client.New(apiClientAuthTransport, strfmt.Default)

	return &Client{
		logger: logger,
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
			httpClient: httpClient,
			service:    authClient.VirtualMachines,
		},
		IPSets: IPSetsProxy{
			httpClient: httpClient,
			service:    authClient.IPSets,
		},
	}, nil
}

func mapToSchemes(insecure bool) []string {
	if insecure {
		return []string{"http"}
	}

	return []string{"https"}
}
