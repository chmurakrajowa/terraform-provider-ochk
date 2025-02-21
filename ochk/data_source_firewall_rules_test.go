package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestFirewallRulesDataSource_read(t *testing.T) {
	resourceName := "data.ochk_firewall_rules.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testFirewallRulesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "firewall_rules.0.rule_id"),
					resource.TestCheckResourceAttrSet(resourceName, "firewall_rules.0.name"),
				),
			},
		},
	})
}

func testFirewallRulesDataSourceConfig() string {
	return fmt.Sprintf(`
data "ochk_project" "default" {
	display_name = "` + testData.Project1Name + `"
}

data "ochk_security_group" "default" {
  display_name = "` + testData.SecurityGroupName + `"
  project_id = data.ochk_project.default.id
}

data "ochk_firewall_rules" "default" {
  project_id = data.ochk_project.default.id
  security_group_id = data.ochk_security_group.default.id
}
`)
}
