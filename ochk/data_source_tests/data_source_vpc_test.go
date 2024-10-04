package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAccRouterDataSource_read(t *testing.T) {

	ctx := context.Background()
	client, err := sdk.NewClient(ctx, os.Getenv("TF_VAR_host"), os.Getenv("TF_VAR_platform"), os.Getenv("TF_VAR_api_key"), false, "", os.Getenv("TF_VAR_platform_type"))
	if err != nil {
		assert.Error(t, err)
	}
	fmt.Printf("VPC name: %v\n", client.Routers)
	proxy := client.Routers
	vpcInstance, err := proxy.Read(context.Background(), "7de4d3e7-978e-4b15-9cd3-970012c969ac")
	if err != nil {
		assert.Error(t, err)
	} else {
		assert.Condition(t, func() bool {
			if vpcInstance.DisplayName == "nieUsuwac222" {
				return true
			}
			return false
		})
	}

}
