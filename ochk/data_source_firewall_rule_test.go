package ochk

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

type FirewallRuleDataSourceTestData struct {
	ResourceName      string
	DisplayName       string
	ProjectName       string
	SecurityGroupName string
}

func (c *FirewallRuleDataSourceTestData) ToString() string {
	return executeTemplateToString(`

data "ochk_project" "project-1" {
  display_name = "`+testData.Project1Name+`"
}

data "ochk_security_group" "sg-1" {
  display_name = "{{ .SecurityGroupName}}"
  project_id = data.ochk_project.project-1.id
}

data "ochk_firewall_rule" "{{ .ResourceName}}" {
  display_name = "{{ .DisplayName}}"
  project_id = data.ochk_project.project-1.id
  security_group_id = data.ochk_security_group.sg-1.id
}
`, c)
}

func (c *FirewallRuleDataSourceTestData) FullResourceName() string {
	return "data.ochk_firewall_rule." + c.ResourceName
}

func TestAccFirewallRuleDataSource_read(t *testing.T) {
	dataSourceProject := "data.ochk_project.project-1"
	dataSourceSecurityGroup := "data.ochk_security_group.sg1"

	FirewallRule := &FirewallRuleDataSourceTestData{
		ResourceName:      "default",
		DisplayName:       testData.FirewallRuleName,
		ProjectName:       testData.Project1Name,
		SecurityGroupName: testData.SecurityGroupName,
	}

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: FirewallRule.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(FirewallRule.FullResourceName(), "display_name", FirewallRule.DisplayName),
					resource.TestCheckResourceAttrPair(FirewallRule.FullResourceName(), "project_id", dataSourceProject, "id"),
					resource.TestCheckResourceAttrPair(FirewallRule.FullResourceName(), "security_group_id", dataSourceSecurityGroup, "id"),
					resource.TestCheckResourceAttrSet(FirewallRule.FullResourceName(), "ether_type"),
					resource.TestCheckResourceAttrSet(FirewallRule.FullResourceName(), "direction"),
					resource.TestCheckResourceAttrSet(FirewallRule.FullResourceName(), "protocol"),
					resource.TestCheckResourceAttrSet(FirewallRule.FullResourceName(), "port_range_min"),
					resource.TestCheckResourceAttrSet(FirewallRule.FullResourceName(), "port_range_max"),
					resource.TestCheckResourceAttrSet(FirewallRule.FullResourceName(), "remote_ip_prefix"),
					resource.TestCheckResourceAttrSet(FirewallRule.FullResourceName(), "created_by"),
					resource.TestCheckResourceAttrSet(FirewallRule.FullResourceName(), "created_at"),
					resource.TestCheckResourceAttrSet(FirewallRule.FullResourceName(), "modified_by"),
					resource.TestCheckResourceAttrSet(FirewallRule.FullResourceName(), "modified_at"),
				),
			},
		},
	})
}
