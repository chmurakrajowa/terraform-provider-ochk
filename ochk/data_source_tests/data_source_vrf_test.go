package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAccVRFDataSource_read(t *testing.T) {
	checkVariables := CheckInputVariables()
	if checkVariables == "" {
		ctx := context.Background()
		client, err := sdk.NewClient(ctx, os.Getenv("TF_VAR_host"), os.Getenv("TF_VAR_platform"), os.Getenv("TF_VAR_api_key"), false, "", os.Getenv("TF_VAR_platform_type"))
		if err != nil {
			assert.Error(t, err)
		}
		fmt.Printf("VRF name: %v\n", client.Routers)
		proxy := client.Routers
		vrfInstance, err := proxy.Read(context.Background(), "5df9b4f3-8f50-4a98-95d5-8ef2304f4362")
		if err != nil {
			assert.Error(t, err)
		} else {
			assert.Condition(t, func() bool {
				if vrfInstance.DisplayName == "VRF-default" {
					return true
				}
				return false
			})
		}
	} else {
		fmt.Printf("ERROR:: %s", checkVariables)
	}
}
