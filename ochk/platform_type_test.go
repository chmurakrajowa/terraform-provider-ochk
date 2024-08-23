package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAccPlatformType_read(t *testing.T) {

	ctx := context.Background()
	client, err := sdk.NewClient(ctx, os.Getenv("TF_VAR_host"), os.Getenv("TF_VAR_platform"), os.Getenv("TF_VAR_api_key"), false, "")
	if err != nil {
		assert.Error(t, err)
	}
	proxy := client.PlatformType
	platformType, err := proxy.Read(context.Background())
	if err != nil {
		assert.Error(t, err)
	} else {
		assert.Condition(t, func() bool {
			if platformType == "VMWARE" {
				return true
			}
			if platformType == "OPENSTACK" {
				return true
			}
			return false
		})
	}
}
