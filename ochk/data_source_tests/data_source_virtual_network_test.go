package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAccVirtualNetworkDatasource(t *testing.T) {
	checkVariables := CheckInputVariables()
	if checkVariables == "" {
		ctx := context.Background()
		client, err := sdk.NewClient(ctx, os.Getenv("TF_VAR_host"), os.Getenv("TF_VAR_platform"), os.Getenv("TF_VAR_api_key"), false, "", os.Getenv("TF_VAR_platform_type"))
		if err != nil {
			assert.Error(t, err)
		}
		fmt.Printf("Virtual Network name: %v\n", client.VirtualNetworks)
		proxy := client.VirtualNetworks
		virtualNetworkInstance, err := proxy.Read(context.Background(), "e58f31c1-7402-4a67-9a3a-138523cf071a")
		if err != nil {
			assert.Error(t, err)
		} else {
			assert.Condition(t, func() bool {
				if virtualNetworkInstance.DisplayName == "admin-priv-vrf" {
					return true
				}
				return false
			})
		}
	} else {
		fmt.Printf("ERROR:: %s", checkVariables)
	}

}
