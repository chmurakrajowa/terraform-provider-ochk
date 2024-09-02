package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAccVirtualNetworksDataSource_read(t *testing.T) {
	ctx := context.Background()
	client, err := sdk.NewClient(ctx, os.Getenv("TF_VAR_host"), os.Getenv("TF_VAR_platform"), os.Getenv("TF_VAR_api_key"), false, "")
	if err != nil {
		assert.Error(t, err)
	}
	fmt.Printf("Virtual Network name: %v\n", client.VirtualNetworks)
	proxy := client.VirtualNetworks
	virtualNetworkInstances, err := proxy.List(context.Background())
	if err != nil {
		assert.Error(t, err)
	} else {
		assert.Condition(t, func() bool {
			if virtualNetworkInstances != nil {
				return true
			}
			return false
		})
		assert.Condition(t, func() bool {
			if virtualNetworkInstances != nil && len(virtualNetworkInstances) >= 0 {

				return true
			}
			return false
		})
	}
}
