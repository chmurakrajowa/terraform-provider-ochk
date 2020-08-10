package ochk

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk"
	"testing"
)

func TestAccFirewallSNRuleResource_noPosition(t *testing.T) {
	resourceName := "ochk_firewall_sn_rule.no_position"
	displayName := generateRandName()

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFirewallSNRuleResourceConfig(displayName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", displayName),
					resource.TestCheckResourceAttr(resourceName, "action", "ALLOW"),
					resource.TestCheckResourceAttr(resourceName, "direction", "IN_OUT"),
					resource.TestCheckResourceAttr(resourceName, "ip_protocol", "IPV4_IPV6"),
					resource.TestCheckNoResourceAttr(resourceName, "position"),
				),
			},
		},
		CheckDestroy: testAccFirewallSNRuleResourceNotExists(displayName),
	})
}

func testAccFirewallSNRuleResourceConfig(displayName string) string {
	source := generateRandName()
	destination := generateRandName()

	return fmt.Sprintf(`
locals {
	routerDisplayName = "T1"
}

data "ochk_gateway_policy" "default" {
  display_name = local.routerDisplayName
}

data "ochk_router" "default" {
  display_name = local.routerDisplayName
}

data "ochk_service" "http" {
  display_name = "http"
}

resource "ochk_security_group" "source" {
  display_name = %[1]q

  members {
    id = "e1e2f617-014c-4119-bac8-49fa4a93db47"
    type = "VIRTUAL_MACHINE"
  }
}

resource "ochk_security_group" "destination" {
  display_name = %[2]q
  
  members {
    id = "e1e2f617-014c-4119-bac8-49fa4a93db47"
    type = "VIRTUAL_MACHINE"
  }
}

resource "ochk_firewall_sn_rule" "no_position" {
  display_name = %[3]q
  gateway_policy_id = data.ochk_gateway_policy.default.id

  services = [data.ochk_service.http.id]
  source = [ochk_security_group.source.id]
  destination = [ochk_security_group.destination.id]
  scope = [data.ochk_router.default.id]
}
`, source, destination, displayName)
}

func testAccFirewallSNRuleResourceNotExists(displayName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ctx := context.Background()

		client := testAccProvider.Meta().(*sdk.Client)

		gatewayPolicies, err := client.GatewayPolicy.ListByDisplayName(ctx, "T1")
		if err != nil {
			return err
		}

		if len(gatewayPolicies) != 1 {
			return fmt.Errorf("wrong number of gateway policies")
		}

		firewallRule, err := client.FirewallSNRules.ListByDisplayName(ctx, gatewayPolicies[0].GatewayPolicyID, displayName)
		if err != nil {
			return err
		}

		if len(firewallRule) > 0 {
			return fmt.Errorf("firewall SN rule %s still exists", firewallRule[0].RuleID)
		}

		return nil
	}
}
