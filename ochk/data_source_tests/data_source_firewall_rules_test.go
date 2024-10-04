package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAccFirewallRulesDataSource_read(t *testing.T) {
	ctx := context.Background()
	client, err := sdk.NewClient(ctx, os.Getenv("TF_VAR_host"), os.Getenv("TF_VAR_platform"), os.Getenv("TF_VAR_api_key"), false, "", os.Getenv("TF_VAR_platform_type"))
	if err != nil {
		assert.Error(t, err)
	}
	fmt.Printf("FirewallRules name: %v\n", client.FirewallRules)
	proxy := client.FirewallRules
	projectId := strfmt.UUID("3cda830c-b37f-46dc-be54-f649d31bec66")
	securityGroupId := strfmt.UUID("ff765340-2c9e-4dc8-b565-bd1d84793f5d")
	rulesInstances, err := proxy.List(context.Background(), projectId, securityGroupId)
	if err != nil {
		assert.Error(t, err)
	} else {
		assert.Condition(t, func() bool {
			if rulesInstances != nil {
				return true
			}
			return false
		})
		assert.Condition(t, func() bool {
			if rulesInstances != nil && len(rulesInstances) >= 0 {

				return true
			}
			return false
		})
	}
}
