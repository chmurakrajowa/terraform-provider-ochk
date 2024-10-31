package ochk

import (
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

import (
	"context"
	"testing"
)

func testAccFirewallRuleCreateResourceId(firewallRuleResourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		firewallRuleResource := s.RootModule().Resources[firewallRuleResourceName]
		resourceID := firewallRuleResource.Primary.Attributes["project_id"] + "/" + firewallRuleResource.Primary.ID
		return resourceID, nil
	}
}

func TestAccFirewallRuleResource_create_update(t *testing.T) {
	resourceName := "ochk_firewall_rule.default"
	displayName := generateRandName(devTestDataPrefix)
	//displayNameUpdated := displayName + "-upd"

	dataSourceSecurityGroup := "data.ochk_security_group.sg-1"
	dataSourceProject := "data.ochk_project.project-1"

	//projectName := testData.Project1Name
	//securityGroupName := testData.SecurityGroupName

	description := "TEST"
	ether_type := "IPv4"
	direction := "EGRESS"
	protocol := "TCP"
	port_range_min := 20008
	port_range_max := 20020
	remote_ip_prefix := "200.200.200.1/32"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFirewallRuleResourceConfig(displayName, description, ether_type, direction, protocol, port_range_min, port_range_max, remote_ip_prefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(resourceName, "project_id", dataSourceProject, "id"),
					resource.TestCheckResourceAttrPair(resourceName, "security_group_id", dataSourceSecurityGroup, "id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", displayName),
					resource.TestCheckResourceAttr(resourceName, "description", "TEST"),
					resource.TestCheckResourceAttr(resourceName, "direction", "EGRESS"),
					resource.TestCheckResourceAttr(resourceName, "ether_type", "IPv4"),
					resource.TestCheckResourceAttr(resourceName, "protocol", "TCP"),
					resource.TestCheckResourceAttr(resourceName, "port_range_min", "20008"),
					resource.TestCheckResourceAttr(resourceName, "port_range_max", "20020"),
					resource.TestCheckResourceAttr(resourceName, "remote_ip_prefix", "200.200.200.1/32"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_by"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_at"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccFirewallRuleCreateResourceId(resourceName),
			},
		},
		CheckDestroy: testAccFirewallRuleResourceNotExists(displayName),
	})
}

func testAccFirewallRuleResourceConfig(displayName string, description string, ether_type string, direction string, protocol string, port_range_min int, port_range_max int, remote_ip_prefix string) string {
	return fmt.Sprintf(`

data "ochk_project" "project-1" {
  display_name = "`+testData.Project1Name+`"
}

data "ochk_security_group" "sg-1" {
  display_name = "`+testData.SecurityGroupName+`"
  project_id = data.ochk_project.project-1.id
}

resource "ochk_firewall_rule" "default" {
  display_name = %[1]q
  security_group_id = data.ochk_security_group.sg-1.id
  project_id = data.ochk_project.project-1.id
  description = %[2]q
  ether_type = %[3]q
  direction = %[4]q
  protocol = %[5]q
  port_range_min = %[6]d
  port_range_max = %[6]d
  remote_ip_prefix = %[8]q
}
`, displayName, description, ether_type, direction, protocol, port_range_min, port_range_max, remote_ip_prefix)
}

func testAccFirewallRuleResourceNotExists(displayName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ctx := context.Background()

		client := testAccProvider.Meta().(*sdk.Client)

		projects, err := client.Projects.ListByName(ctx, testData.Project1Name)
		if err != nil {
			return err
		}

		if len(projects) != 1 {
			return fmt.Errorf("wrong number of projects")
		}

		security_groups, err := client.SecurityGroups.ListByDisplayName(ctx, testData.SecurityGroupName)
		if err != nil {
			return err
		}

		if len(security_groups) != 1 {
			return fmt.Errorf("wrong number of security groups")
		}

		firewallRule, err := client.FirewallRules.ListByName(ctx, projects[0].ProjectID, security_groups[0].ID, displayName)
		if err != nil {
			return err
		}

		if len(firewallRule) > 0 {
			return fmt.Errorf("firewall rule %s still exists", firewallRule[0].RuleID)
		}

		return nil
	}
}
