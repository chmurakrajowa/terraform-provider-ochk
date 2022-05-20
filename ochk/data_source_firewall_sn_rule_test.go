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

data "ochk_router" "vrf" {
	display_name = "{{ .VRFName}}"
}	

data "ochk_router" "vpc" {
  display_name = "{{ .VPCName}}"
  parent_router_id = data.ochk_router.vrf.id
}

data "ochk_firewall_sn_rule" "{{ .ResourceName}}" {
  display_name = "{{ .DisplayName}}"
  router_id = data.ochk_router.vpc.id
}
`, c)
}

func (c *FirewallRuleSNDataSourceTestData) FullResourceName() string {
	return "data.ochk_firewall_sn_rule." + c.ResourceName
}

func TestAccFirewallRuleSNDataSource_read(t *testing.T) {
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
					//resource.TestCheckResourceAttrSet(FirewallRule.FullResourceName(), "router_id"),
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
