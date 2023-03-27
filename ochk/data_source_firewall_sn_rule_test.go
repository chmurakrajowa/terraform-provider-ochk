package ochk

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

type FirewallRuleSNDataSourceTestData struct {
	ResourceName string
	DisplayName  string
	VPCName      string
	VRFName      string
}

func (c *FirewallRuleSNDataSourceTestData) ToString() string {
	return executeTemplateToString(`

data "ochk_vrf" "vrf" {
	display_name = "{{ .VRFName}}"
}	

data "ochk_project" "project-1" {
  display_name = "`+testData.Project1Name+`"
}

data "ochk_vpc" "vpc" {
  display_name = "{{ .VPCName}}"
  vrf_id = data.ochk_vrf.vrf.id
}

data "ochk_firewall_sn_rule" "{{ .ResourceName}}" {
  display_name = "{{ .DisplayName}}"
  vpc_id = data.ochk_vpc.vpc.id
}
`, c)
}

func (c *FirewallRuleSNDataSourceTestData) FullResourceName() string {
	return "data.ochk_firewall_sn_rule." + c.ResourceName
}

func TestAccFirewallRuleSNDataSource_read(t *testing.T) {
	dataSourceProject := "data.ochk_project.project-1"
	dataSourceVpc := "data.ochk_vpc.vpc"

	FirewallRule := &FirewallRuleSNDataSourceTestData{
		ResourceName: "default",
		DisplayName:  testData.FirewallSNRuleName,
		VPCName:      testData.VPC,
		VRFName:      testData.VRF,
	}

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: FirewallRule.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(FirewallRule.FullResourceName(), "display_name", FirewallRule.DisplayName),
					resource.TestCheckResourceAttrPair(FirewallRule.FullResourceName(), "vpc_id", dataSourceVpc, "id"),
					resource.TestCheckResourceAttrPair(FirewallRule.FullResourceName(), "project_id", dataSourceProject, "id"),
					resource.TestCheckResourceAttrSet(FirewallRule.FullResourceName(), "action"),
					resource.TestCheckResourceAttrSet(FirewallRule.FullResourceName(), "direction"),
					resource.TestCheckResourceAttrSet(FirewallRule.FullResourceName(), "disabled"),
					resource.TestCheckResourceAttrSet(FirewallRule.FullResourceName(), "ip_protocol"),
					resource.TestCheckResourceAttrSet(FirewallRule.FullResourceName(), "priority"),
					resource.TestCheckResourceAttrSet(FirewallRule.FullResourceName(), "created_by"),
					resource.TestCheckResourceAttrSet(FirewallRule.FullResourceName(), "created_at"),
					resource.TestCheckResourceAttrSet(FirewallRule.FullResourceName(), "modified_by"),
					resource.TestCheckResourceAttrSet(FirewallRule.FullResourceName(), "modified_at"),
				),
			},
		},
	})
}
