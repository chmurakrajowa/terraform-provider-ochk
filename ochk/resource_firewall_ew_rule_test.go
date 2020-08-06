package ochk

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk"
	"testing"
	"time"
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
		CheckDestroy: testAccFirewallEWRuleResourceDoesntExist(displayName),
	})
}

func TestAccFirewallEWRuleResource_withPositions(t *testing.T) {
	randDisplayName := generateRandName()
	displayNameMiddle := randDisplayName + "-middle"
	displayNameBefore := randDisplayName + "-before"
	displayNameAfter := randDisplayName + "-after"

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFirewallEWRuleResourceConfigWithOrder(displayNameBefore, displayNameMiddle, displayNameAfter),
				Check: resource.ComposeTestCheckFunc(
					checkFuncWithRetry(30, testAccFirewallEWRuleCheckRulesOrder("data.ochk_security_policy.default", displayNameBefore, displayNameMiddle, displayNameAfter)),
				),
			},
		},
		CheckDestroy: resource.ComposeTestCheckFunc(
			testAccFirewallEWRuleResourceDoesntExist(displayNameBefore),
			testAccFirewallEWRuleResourceDoesntExist(displayNameMiddle),
			testAccFirewallEWRuleResourceDoesntExist(displayNameAfter)),
	})
}

func testAccFirewallEWRuleCheckRulesOrder(securityPolicyResourceName string, displayNames ...string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		ctx := context.Background()

		client := testAccProvider.Meta().(*sdk.Client)

		securityPolicyResource, ok := state.RootModule().Resources[securityPolicyResourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", securityPolicyResourceName)
		}

		securityPolicies, err := client.FirewallEWRules.List(ctx, securityPolicyResource.Primary.ID)
		if err != nil {
			return err
		}

		for i := 0; i < len(securityPolicies); {
			if securityPolicies[i].DisplayName == displayNames[0] {
				for j := 1; j < len(displayNames) && i+j < len(securityPolicies); {
					if securityPolicies[i+j].DisplayName != displayNames[j] {
						return fmt.Errorf("invalid order of security policies, expected %s but is %s", displayNames[j], securityPolicies[i+j].DisplayName)
					}
					i++
					j++
				}

				return nil
			}

			i++
		}

		return fmt.Errorf("invalid order of security policies, policy %s not found", displayNames[0])
	}
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

func testAccFirewallEWRuleResourceConfigWithOrder(displayNameBefore string, displayNameMiddle string, displayNameAfter string) string {
	source := generateRandName()
	destination := generateRandName()

	return fmt.Sprintf(`
data "ochk_security_policy" "default" {
  display_name = "devel"
}

data "ochk_service" "http" {
  display_name = "http"
}

resource "ochk_security_group" "source-before" {
  display_name = "%[1]s-before"

  members {
    id = "e1e2f617-014c-4119-bac8-49fa4a93db47"
    type = "VIRTUAL_MACHINE"
  }
}

resource "ochk_security_group" "source-middle" {
  display_name = "%[1]s-middle"

  members {
    id = "e1e2f617-014c-4119-bac8-49fa4a93db47"
    type = "VIRTUAL_MACHINE"
  }
}
resource "ochk_security_group" "source-after" {
  display_name = "%[1]s-after"

  members {
    id = "e1e2f617-014c-4119-bac8-49fa4a93db47"
    type = "VIRTUAL_MACHINE"
  }
}

resource "ochk_security_group" "destination-before" {
  display_name = "%[2]s-before"
  
  members {
    id = "e1e2f617-014c-4119-bac8-49fa4a93db47"
    type = "VIRTUAL_MACHINE"
  }
}

resource "ochk_security_group" "destination-middle" {
  display_name = "%[2]s-middle"
  
  members {
    id = "e1e2f617-014c-4119-bac8-49fa4a93db47"
    type = "VIRTUAL_MACHINE"
  }
}

resource "ochk_security_group" "destination-after" {
  display_name = "%[2]s-after"
  
  members {
    id = "e1e2f617-014c-4119-bac8-49fa4a93db47"
    type = "VIRTUAL_MACHINE"
  }
}

resource "ochk_firewall_ew_rule" "middle" {
  display_name = %[3]q
  security_policy_id = data.ochk_security_policy.default.id

  services = [data.ochk_service.http.id]
  source = [ochk_security_group.source-middle.id]
  destination = [ochk_security_group.destination-middle.id]
}

resource "ochk_firewall_ew_rule" "before" {
  display_name = %[4]q
  security_policy_id = data.ochk_security_policy.default.id

  services = [data.ochk_service.http.id]
  source = [ochk_security_group.source-before.id]
  destination = [ochk_security_group.destination-before.id]
  position {
 	rule_id = ochk_firewall_ew_rule.middle.id
	revise_operation = "BEFORE"
  }
}

resource "ochk_firewall_ew_rule" "after" {
  display_name = %[5]q
  security_policy_id = data.ochk_security_policy.default.id

  services = [data.ochk_service.http.id]
  source = [ochk_security_group.source-after.id]
  destination = [ochk_security_group.destination-after.id]
  position {
 	rule_id = ochk_firewall_ew_rule.middle.id
	revise_operation = "AFTER"
  }
}
`, source, destination, displayNameMiddle, displayNameBefore, displayNameAfter)
}

func checkFuncWithRetry(checkCount int, checkFunc resource.TestCheckFunc) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for i := 0; i < checkCount; i++ {
			err := checkFunc(s)
			if err != nil {
				if i == checkCount-1 {
					return err
				}

				time.Sleep(time.Second)
				continue
			}

			break
		}

		return nil
	}
}

func testAccFirewallEWRuleResourceDoesntExist(displayName string) resource.TestCheckFunc {
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

		firewallRule, err := client.FirewallEWRules.ListByDisplayName(ctx, securityPolicies[0].SecurityPolicyID, displayName)
		if err != nil {
			return err
		}

		if len(firewallRule) > 0 {
			return fmt.Errorf("firewall EW rule %s still exists", firewallRule[0].RuleID)
		}

		return nil
	}
}
