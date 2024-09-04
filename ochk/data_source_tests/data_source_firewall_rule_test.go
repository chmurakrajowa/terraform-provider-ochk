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

func TestAccFirewallRuleDataSource_read(t *testing.T) {
	ctx := context.Background()
	client, err := sdk.NewClient(ctx, os.Getenv("TF_VAR_host"), os.Getenv("TF_VAR_platform"), os.Getenv("TF_VAR_api_key"), false, "")
	if err != nil {
		assert.Error(t, err)
	}
	fmt.Printf("FirewallRule name: %v\n", client.FirewallRules)
	proxy := client.FirewallRules
	projectId := strfmt.UUID("3cda830c-b37f-46dc-be54-f649d31bec66")
	securityGroupId := strfmt.UUID("ff765340-2c9e-4dc8-b565-bd1d84793f5d")
	ruleId := strfmt.UUID("06f8f3bc-93d7-4d84-98bd-270e2b9b847f")

	rulesInstance, err := proxy.Read(context.Background(), projectId, securityGroupId, ruleId)
	if err != nil {
		assert.Error(t, err)
	} else {
		assert.Condition(t, func() bool {
			if rulesInstance != nil {
				return true
			}
			return false
		})
	}
}
