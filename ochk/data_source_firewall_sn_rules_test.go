package ochk

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccFirewallSNRulesDataSource_read(t *testing.T) {
	resourceName := "data.ochk_firewall_sn_rules.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFirewallSNRulesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "firewall_sn_rules.0.firewall_sn_rule_id"),
					resource.TestCheckResourceAttrSet(resourceName, "firewall_sn_rules.0.project_id"),
					resource.TestCheckResourceAttrSet(resourceName, "firewall_sn_rules.0.display_name"),
				),
			},
		},
	})
}

func testAccFirewallSNRulesDataSourceConfig() string {
	return fmt.Sprintf(`
data "ochk_vrf" "default" {
	display_name = %q
}

data "ochk_vpc" "default" {
  display_name = %q
  vrf_id = data.ochk_vrf.default.id
}

data "ochk_firewall_sn_rules" "default" {
  vpc_id = data.ochk_vpc.default.id
}
`, testData.VRF, testData.VPC)
}
