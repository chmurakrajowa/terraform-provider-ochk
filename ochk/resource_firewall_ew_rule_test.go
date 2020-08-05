package ochk

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk"
	"log"
	"testing"
)

func TestAccFirewallEWRuleResource_create(t *testing.T) {
	name := fmt.Sprintf("tf-test-%s", acctest.RandStringFromCharSet(8, acctest.CharSetAlphaNum))

	//TODO zbyt wiele razy jest wołany POST /vidm/token HTTP/1.1, coś jest nie tak
	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccFirewallEWRuleResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccFirewallEWRuleResourceConfig(name),
				//TODO [M.K.] IMO nie ma sensu sprawdzać czy dany zasób istnieje, on jest pobierany na końcu create
				//Trzeba po prostu sprawdzić czy state się jakoś zgadza
				Check:  testAccFirewallEWRuleResourceExists("ochk_firewall_ew_rule.test"),
			},
		},
	})
}

func testAccFirewallEWRuleResourceExists(resourceID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceID]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceID)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("resource %s has no ID set", resourceID)
		}

		securityPolicyID, ok := rs.Primary.Attributes["security_policy_id"]
		if !ok {
			return fmt.Errorf("resource %s has no security_policy_id set", securityPolicyID)
		}

		proxy := testAccProvider.Meta().(*sdk.Client).FirewallEWFRules

		_, err := proxy.Read(context.Background(), securityPolicyID, rs.Primary.ID)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccFirewallEWRuleResourceConfig(name string) string {
	return fmt.Sprintf(`
resource "ochk_firewall_ew_rule" "test" {
  display_name = %[1]q

  members {
    id = "e1e2f617-014c-4119-bac8-49fa4a93db47"
    type = "VIRTUAL_MACHINE"
  }
}
`, name)
}

func testAccFirewallEWRuleResourceDestroy() resource.TestCheckFunc {
	return func(state *terraform.State) error {
		for _, rs := range state.RootModule().Resources {
			if rs.Type != "ochk_firewall_ew_rule" {
				continue
			}

			log.Printf("checking security group %s exists", rs.Primary.ID)
			proxy := testAccProvider.Meta().(*sdk.Client).SecurityGroups

			exists, err := proxy.Exists(context.Background(), rs.Primary.ID)
			if err != nil {
				return err
			}

			if exists {
				return fmt.Errorf("security group %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}
}
