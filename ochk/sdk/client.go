package sdk

import (
	"context"
	"crypto/tls"
	"fmt"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client/vidm_controller"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/models"
	"net/http"
)

type Client struct {
	logger         *FileLogger
	SecurityGroups SecurityGroupsProxy
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

	params := vidm_controller.GetTokenUsingPOSTParams{
		VidmTokenRequest: &models.VIDMTokenRequest{
			Tenant:   tenant,
			Password: password,
			Username: username,
		},
		Context:    ctx,
		HTTPClient: httpClient,
	}

	authResponse, err := ochkClient.VidmController.GetTokenUsingPOST(&params)
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
	}, nil
}

func mapToSchemes(insecure bool) []string {
	if insecure {
		return []string{"http"}
	}

	return []string{"https"}
}
