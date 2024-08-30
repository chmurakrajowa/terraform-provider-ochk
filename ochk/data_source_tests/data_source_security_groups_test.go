package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAccSecurityGroupsDataSource_read(t *testing.T) {
	ctx := context.Background()
	client, err := sdk.NewClient(ctx, os.Getenv("TF_VAR_host"), os.Getenv("TF_VAR_platform"), os.Getenv("TF_VAR_api_key"), false, "")
	if err != nil {
		assert.Error(t, err)
	}
	fmt.Printf("Project name: %v\n", client.Projects)
	proxy := client.Projects
	securityGroupInstances, err := proxy.List(context.Background())
	if err != nil {
		assert.Error(t, err)
	} else {
		assert.Condition(t, func() bool {
			if securityGroupInstances != nil {
				return true
			}
			return false
		})
		assert.Condition(t, func() bool {
			if securityGroupInstances != nil && len(securityGroupInstances) >= 0 {

				return true
			}
			return false
		})
	}

}
