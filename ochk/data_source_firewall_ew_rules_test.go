package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccFirewallEWRulesDataSource_read(t *testing.T) {
	resourceName := "data.ochk_firewall_ew_rules.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFirewallEWRulesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "firewall_ew_rules.0.firewall_ew_rule_id"),
					resource.TestCheckResourceAttrSet(resourceName, "firewall_ew_rules.0.project_id"),
					resource.TestCheckResourceAttrSet(resourceName, "firewall_ew_rules.0.display_name"),
				),
			},
		},
	})
}

func testAccFirewallEWRulesDataSourceConfig() string {
	return fmt.Sprintf(`
data "ochk_vrf" "default" {
	display_name = "` + testData.VRF + `"
}

data "ochk_vpc" "default" {
  display_name = "` + testData.VPC + `"
  vrf_id = data.ochk_vrf.default.id
}

data "ochk_firewall_ew_rules" "default" {
  vpc_id = data.ochk_vpc.default.id
}
`)
}
