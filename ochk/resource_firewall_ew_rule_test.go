package ochk

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk"
	"testing"
)

func TestAccFirewallEWRuleResource_noPosition(t *testing.T) {
	resourceName := "ochk_firewall_ew_rule.no_position"
	displayName := generateRandName()

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFirewallEWRuleResourceConfig(displayName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", displayName),
					resource.TestCheckResourceAttr(resourceName, "action", "ALLOW"),
					resource.TestCheckResourceAttr(resourceName, "direction", "IN_OUT"),
					resource.TestCheckResourceAttr(resourceName, "ip_protocol", "IPV4_IPV6"),
					resource.TestCheckNoResourceAttr(resourceName, "position"),
				),
			},
		},
		CheckDestroy: testAccFirewallEWRuleResourceNotExists(displayName),
	})
}

func testAccFirewallEWRuleResourceConfig(displayName string) string {
	source := generateRandName()
	destination := generateRandName()

	return fmt.Sprintf(`
data "ochk_security_policy" "default" {
  display_name = "devel"
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

resource "ochk_firewall_ew_rule" "no_position" {
  display_name = %[3]q
  security_policy_id = data.ochk_security_policy.default.id

  services = [data.ochk_service.http.id]
  source = [ochk_security_group.source.id]
  destination = [ochk_security_group.destination.id]
}
`, source, destination, displayName)
}

func testAccFirewallEWRuleResourceNotExists(displayName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ctx := context.Background()

		client := testAccProvider.Meta().(*sdk.Client)

		securityPolicies, err := client.SecurityPolicy.ListByDisplayName(ctx, "devel")
		if err != nil {
			return err
		}

		if len(securityPolicies) != 1 {
			return fmt.Errorf("wrong number of security policies")
		}

		firewallRule, err := client.FirewallEWFRules.ListByDisplayName(ctx, securityPolicies[0].SecurityPolicyID, displayName)
		if err != nil {
			return err
		}

		if len(firewallRule) > 0 {
			return fmt.Errorf("firewall EW rule %s still exists", firewallRule[0].RuleID)
		}

		return nil
	}
}
