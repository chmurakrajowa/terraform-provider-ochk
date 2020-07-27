package sdk

import (
	"context"
	"crypto/tls"
	"fmt"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client/vidm_controller"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/models"
	"net/http"
	"os"
)

type Client struct {
	client     *client.Ochk
	httpClient *http.Client
	logger     *FileLogger
}

func getEnvOrDefault(env string, defaultValue string) string {
	if envValue := os.Getenv(env); envValue != "" {
		return envValue
	}

	return defaultValue
}

func NewClient(host string, tenant string, username string, password string, insecure bool, debugLogFile string) (*Client, error) {
	host = getEnvOrDefault("OCHK_HOST", host)
	tenant = getEnvOrDefault("OCHK_TENANT", tenant)
	username = getEnvOrDefault("OCHK_USERNAME", username)
	password = getEnvOrDefault("OCHK_PASSWORD", password)

	ctx := context.Background()

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
	apiClientTransport.Debug = true

	ochkClient := client.New(apiClientTransport, nil)

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
	apiClientAuthTransport.Debug = true
	if logger != nil {
		apiClientAuthTransport.SetLogger(logger)
	}
	apiClientAuthTransport.DefaultAuthentication = httptransport.APIKeyAuth("token", "header", authResponse.Payload.Token)

	authClient := client.New(apiClientAuthTransport, nil)

	return &Client{
		client: authClient,
		httpClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
		logger: logger,
	}, nil
}

func (c *Client) GetOchk() *client.Ochk {
	return c.client
}

func (c *Client) GetHTTPClient() *http.Client {
	return c.httpClient
}

func mapToSchemes(insecure bool) []string {
	if insecure {
		return []string{"http"}
	}

	return []string{"https"}
}
