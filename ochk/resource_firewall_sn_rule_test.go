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
	dataSourceRouter := "data.ochk_router.default"
	dataSourceService := "data.ochk_service.http"
	resourceRuleSource := "ochk_security_group.source"
	resourceRuleDestination := "ochk_security_group.destination"

	sourceDisplayName := generateRandName()
	destinationDisplayName := generateRandName()
	ruleDisplayName := generateRandName()

	routerName := "T1"
	routerNameUpdated := "T0"

	ruleDisplayNameUpdated := ruleDisplayName + "-updated"
	actionUpdated := "DROP"
	directionUpdated := "OUT"
	ipProtocolUpdated := "IPV4"

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFirewallSNRuleResourceConfig(routerName, sourceDisplayName, destinationDisplayName) +
					testAccFirewallSNRuleResourceConfigNoPosition(ruleDisplayName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", ruleDisplayName),
					resource.TestCheckResourceAttr(resourceName, "action", "ALLOW"),
					resource.TestCheckResourceAttr(resourceName, "direction", "IN_OUT"),
					resource.TestCheckResourceAttr(resourceName, "ip_protocol", "IPV4_IPV6"),
					resource.TestCheckResourceAttrPair(resourceName, "source.0", resourceRuleSource, "id"),
					resource.TestCheckResourceAttrPair(resourceName, "destination.0", resourceRuleDestination, "id"),
					resource.TestCheckResourceAttrPair(resourceName, "services.0", dataSourceService, "id"),
					resource.TestCheckResourceAttrPair(resourceName, "scope.0", dataSourceRouter, "id"),
					resource.TestCheckNoResourceAttr(resourceName, "position"),
				),
			},
			{
				Config: testAccFirewallSNRuleResourceConfig(routerNameUpdated, sourceDisplayName, destinationDisplayName) +
					testAccFirewallSNRuleResourceConfigNoPosition(ruleDisplayName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(resourceName, "scope.0", dataSourceRouter, "id"),
				),
			},
			{
				Config: testAccFirewallSNRuleResourceConfig(routerNameUpdated, sourceDisplayName, destinationDisplayName) +
					testAccFirewallSNRuleResourceConfigNoPositionUpdate(ruleDisplayNameUpdated, actionUpdated, ipProtocolUpdated, directionUpdated),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", ruleDisplayNameUpdated),
					resource.TestCheckResourceAttr(resourceName, "action", actionUpdated),
					resource.TestCheckResourceAttr(resourceName, "ip_protocol", ipProtocolUpdated),
					resource.TestCheckResourceAttr(resourceName, "direction", directionUpdated),
				),
			},
		},
		CheckDestroy: testAccFirewallSNRuleResourceNotExists(ruleDisplayName),
	})
}

func testAccFirewallSNRuleResourceConfig(router string, source string, destination string) string {
	return fmt.Sprintf(`
locals {
	routerDisplayName = %[1]q
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
  display_name = %[2]q

  members {
    id = "e1e2f617-014c-4119-bac8-49fa4a93db47"
    type = "VIRTUAL_MACHINE"
  }
}

resource "ochk_security_group" "destination" {
  display_name = %[3]q
  
  members {
    id = "e1e2f617-014c-4119-bac8-49fa4a93db47"
    type = "VIRTUAL_MACHINE"
  }
}
`, router, source, destination)
}

func testAccFirewallSNRuleResourceConfigNoPosition(displayName string) string {
	return fmt.Sprintf(`
resource "ochk_firewall_sn_rule" "no_position" {
  display_name = %[1]q
  gateway_policy_id = data.ochk_gateway_policy.default.id

  services = [data.ochk_service.http.id]
  source = [ochk_security_group.source.id]
  destination = [ochk_security_group.destination.id]
  scope = [data.ochk_router.default.id]
}
`, displayName)
}

func testAccFirewallSNRuleResourceConfigNoPositionUpdate(displayName string, action string, ipProtocol string, direction string) string {
	return fmt.Sprintf(`
resource "ochk_firewall_sn_rule" "no_position" {
  display_name = %[1]q
  action = %[2]q 
  ip_protocol = %[3]q
  direction = %[4]q
  gateway_policy_id = data.ochk_gateway_policy.default.id

  services = [data.ochk_service.http.id]
  source = [ochk_security_group.source.id]
  destination = [ochk_security_group.destination.id]
  scope = [data.ochk_router.default.id]
}
`, displayName, action, ipProtocol, direction)
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
