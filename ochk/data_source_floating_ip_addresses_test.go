package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAccFloatingIPAddressesDataSource_read(t *testing.T) {
	ctx := context.Background()
	client, err := sdk.NewClient(ctx, os.Getenv("TF_VAR_host"), os.Getenv("TF_VAR_platform"), os.Getenv("TF_VAR_api_key"), false, "", os.Getenv("TF_VAR_platform_type"))
	if err != nil {
		assert.Error(t, err)
	}

	fmt.Printf("Public ip client %v\n", client.PublicIPAddresses)
	fmt.Printf("Floating ip client: %v\n", client.FloatingIPAddresses)
	proxy := client.FloatingIPAddresses
	floatingIPInstances, err := proxy.List(context.Background())
	if err != nil {
		assert.Error(t, err)
	} else {
		assert.Condition(t, func() bool {
			if floatingIPInstances != nil {
				return true
			}
			return false
		})
		assert.Condition(t, func() bool {
			if floatingIPInstances != nil && len(floatingIPInstances) >= 0 {

				return true
			}
			return false
		})
	}
}
