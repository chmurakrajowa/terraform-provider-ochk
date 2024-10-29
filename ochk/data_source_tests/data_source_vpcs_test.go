package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAccVPCsDataSource_read(t *testing.T) {
	checkVariables := CheckInputVariables()
	if checkVariables == "" {
		ctx := context.Background()
		client, err := sdk.NewClient(ctx, os.Getenv("TF_VAR_host"), os.Getenv("TF_VAR_platform"), os.Getenv("TF_VAR_api_key"), false, "", os.Getenv("TF_VAR_platform_type"))
		if err != nil {
			assert.Error(t, err)
		}
		fmt.Printf("VPC name: %v\n", client.Routers)
		proxy := client.Routers
		vpcInstances, err := proxy.List(context.Background())
		if err != nil {
			assert.Error(t, err)
		} else {
			assert.Condition(t, func() bool {
				if vpcInstances != nil {
					return true
				}
				return false
			})
			assert.Condition(t, func() bool {
				if vpcInstances != nil && len(vpcInstances) >= 0 {

					return true
				}
				return false
			})
		}
	} else {
		fmt.Printf("ERROR:: %s", checkVariables)
	}
}
