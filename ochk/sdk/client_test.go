package sdk

import (
	"context"
	"crypto/tls"
	"fmt"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client/security_groups"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client/vidm_controller"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client/virtual_machines"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"net/http"
	"testing"
)

func TestHttpClient(t *testing.T) {
	t.SkipNow()

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

	fmt.Printf("%+v\n", getTokenUsingPOSTOKResponse.Payload)

	apiClientAuthTransport := httptransport.New(client.DefaultHost, client.DefaultBasePath, []string{"https"})
	apiClientAuthTransport.Debug = true
	apiClientAuthTransport.DefaultAuthentication = httptransport.APIKeyAuth("token", "header", getTokenUsingPOSTOKResponse.Payload.Token)
	iasAPIAuthClient := client.New(apiClientAuthTransport, nil)

	securityGroupListUsingGETResponse, err := iasAPIAuthClient.SecurityGroups.SecurityGroupListUsingGET(&security_groups.SecurityGroupListUsingGETParams{
		Context:    ctx,
		HTTPClient: &iasHttpClient,
	})

	if !assert.NoError(t, err) {
		return
	}

	fmt.Printf("%+v\n", securityGroupListUsingGETResponse.Payload)

	virtualMachineListUsingGETResponse, err := iasAPIAuthClient.VirtualMachines.VirtualMachineListUsingGET(&virtual_machines.VirtualMachineListUsingGETParams{
		Context:    ctx,
		HTTPClient: &iasHttpClient,
	})
	if !assert.NoError(t, err) {
		return
	}
	fmt.Printf("%+v\n", virtualMachineListUsingGETResponse.Payload)

	vmList := virtualMachineListUsingGETResponse.Payload.VirtualMachineCollection
	if len(vmList) > 0 {
		virtualMachineUsingGETResponse, err := iasAPIAuthClient.VirtualMachines.VirtualMachineGetUsingGET(&virtual_machines.VirtualMachineGetUsingGETParams{
			VirtualMachineID: vmList[0].VirtualMachineID,
			Context:          ctx,
			HTTPClient:       &iasHttpClient,
		})
		if !assert.NoError(t, err) {
			return
		}
		fmt.Printf("%+v\n", virtualMachineUsingGETResponse.Payload)

		securityGroupCreateUsingPUTOKResponse, securityGroupCreateUsingPUTCreatedResponse, err := iasAPIAuthClient.SecurityGroups.SecurityGroupCreateUsingPUT(&security_groups.SecurityGroupCreateUsingPUTParams{
			SecurityGroup: &models.SecurityGroup{

				DisplayName: fmt.Sprintf("sg-tf-test-%d", rand.Intn(100000)),
				Members: []*models.SecurityGroupMember{
					{
						ID:         vmList[0].VirtualMachineID,
						MemberType: "VIRTUAL_MACHINE",
					},
				},
			},
			Context:    ctx,
			HTTPClient: &iasHttpClient,
		})

		if !assert.NoError(t, err) {
			return
		}

		if securityGroupCreateUsingPUTOKResponse != nil {
			fmt.Printf("securityGroupCreateUsingPUTOKResponse: %+v", securityGroupCreateUsingPUTOKResponse)
		}
		if securityGroupCreateUsingPUTCreatedResponse != nil {
			fmt.Printf("securityGroupCreateUsingPUTCreatedResponse: %+v", securityGroupCreateUsingPUTCreatedResponse)

			securityGroupID := securityGroupCreateUsingPUTCreatedResponse.Payload.SecurityGroup.ID

			securityGroupListUsingGETResponse, err := iasAPIAuthClient.SecurityGroups.SecurityGroupGetUsingGET(&security_groups.SecurityGroupGetUsingGETParams{
				GroupID:    securityGroupID,
				Context:    ctx,
				HTTPClient: &iasHttpClient,
			})

			if !assert.NoError(t, err) {
				return
			}

			fmt.Printf("securityGroup %s: %+v\n", securityGroupID, securityGroupListUsingGETResponse.Payload)
		}
	}
}
