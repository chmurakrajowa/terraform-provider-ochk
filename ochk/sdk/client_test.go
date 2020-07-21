package sdk

import (
	"context"
	"crypto/tls"
	"fmt"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client/vidm_controller"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHttpClient(t *testing.T) {

	ctx := context.Background()
	iasHttpClient := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	/*iasHttpClient := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	*/

	params := vidm_controller.GetTokenUsingPOSTParams{
		VidmTokenRequest: &models.VIDMTokenRequest{
			Password: "zaq1@WSX",
			Tenant:   "devel",
			Username: "devel-admin",
		},
		Context:    ctx,
		HTTPClient: &iasHttpClient,
	}

	apiClientTransport := httptransport.New(client.DefaultHost, client.DefaultBasePath, []string{"https"})
	apiClientTransport.Debug = true
	//apiClientTransportConfig.DefaultAuthentication = httptransport.APIKeyAuth("token", )

	iasAPIClient := client.New(apiClientTransport, nil)
	getTokenUsingPOSTOKResponse, err := iasAPIClient.VidmController.GetTokenUsingPOST(&params)
	if !assert.NoError(t, err) {
		return
	}

	fmt.Printf("%+v", getTokenUsingPOSTOKResponse.Payload)
}
