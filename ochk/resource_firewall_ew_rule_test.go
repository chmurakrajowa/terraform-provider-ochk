package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"testing"
)

func testAccirewallEWRuleCreateResourceId(firewallRuleResourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		firewallRuleResource := s.RootModule().Resources[firewallRuleResourceName]
		resourceID := firewallRuleResource.Primary.Attributes["vpc_id"] + "/" + firewallRuleResource.Primary.ID
		return resourceID, nil
	}
}

func TestAccFirewallEWRuleResource_create_update(t *testing.T) {
	resourceName := "ochk_firewall_ew_rule.default"
	displayName := generateRandName(devTestDataPrefix)
	displayNameUpdated := displayName + "-upd"
	action := "ALLOW"
	actionUpdated := "DROP"
	direction := "IN_OUT"
	directionUpdated := "OUT"
	ipProtocol := "IPV4_IPV6"
	ipProtocolUpdated := "IPV4"
	dataSourceRouter := "ochk_vpc.default"

	source := generateRandName(devTestDataPrefix)
	destination := generateRandName(devTestDataPrefix)

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFirewallEWRuleResourceConfig(displayName, source, destination, action, ipProtocol, direction, 1000, "data.ochk_custom_service.custom_service1.id"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", displayName),
					resource.TestCheckResourceAttrPair(resourceName, "project_id", "data.ochk_project.project-1", "id"),
					resource.TestCheckResourceAttr(resourceName, "action", action),
					resource.TestCheckResourceAttr(resourceName, "direction", direction),
					resource.TestCheckResourceAttr(resourceName, "ip_protocol", ipProtocol),
					resource.TestCheckResourceAttr(resourceName, "priority", "1000"),
					resource.TestCheckResourceAttrPair(resourceName, "custom_services.0", "data.ochk_custom_service.custom_service1", "id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_by"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_at"),
					resource.TestCheckResourceAttrPair(resourceName, "vpc_id", dataSourceRouter, "id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccirewallEWRuleCreateResourceId(resourceName),
			},
			{
				Config: testAccFirewallEWRuleResourceConfig(displayNameUpdated, source, destination, actionUpdated, ipProtocolUpdated, directionUpdated, 2000, "data.ochk_custom_service.custom_service2.id"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", displayNameUpdated),
					resource.TestCheckResourceAttrPair(resourceName, "project_id", "data.ochk_project.project-1", "id"),
					resource.TestCheckResourceAttr(resourceName, "action", actionUpdated),
					resource.TestCheckResourceAttr(resourceName, "direction", directionUpdated),
					resource.TestCheckResourceAttr(resourceName, "ip_protocol", ipProtocolUpdated),
					resource.TestCheckResourceAttr(resourceName, "priority", "2000"),
					resource.TestCheckResourceAttrPair(resourceName, "custom_services.0", "data.ochk_custom_service.custom_service2", "id"),
				),
			},
		},
		CheckDestroy: testAccFirewallEWRuleResourceDoesntExist(displayName),
	})
}

func TestAccFirewallEWRuleResource_withPriority(t *testing.T) {
	randDisplayName := generateRandName(devTestDataPrefix)
	displayNameMiddle := randDisplayName + "mid"
	displayNameBefore := randDisplayName + "bef"
	displayNameAfter := randDisplayName + "aft"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFirewallEWRuleResourceConfigWithOrder(displayNameBefore, displayNameMiddle, displayNameAfter),
				Check: resource.ComposeTestCheckFunc(
					testAccFirewallEWRuleCheckRulesOrder("ochk_vpc.default", displayNameBefore, displayNameMiddle, displayNameAfter),
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

		if !checkOrderOfSecurityPolicies(securityPolicies, displayNames) {
			securityPoliciesDisplayNames := transformToStringSlice(len(securityPolicies), func(idx int) string {
				return securityPolicies[idx].DisplayName
			})
			return fmt.Errorf("security policies not found in expected order: %+v, security polices: %+v", displayNames, securityPoliciesDisplayNames)
		}

		return nil
	}
}

func TestCheckOrderOfSecurityPolicies(t *testing.T) {
	assert.NotPanics(t, func() {
		expectedOrder := []string{"a", "b", "c"}

		matchingCases := [][]string{
			{"a", "b", "c"},
			{"0", "a", "b", "c"},
			{"a", "b", "c", "d"},
			{"1", "2", "3", "4", "a", "b", "c"},
			{"1", "2", "3", "4", "a", "b", "c", "5", "6", "7"},
		}

		notMatchingCases := [][]string{
			{"a", "b"},
			{"0", "a", "b"},
			{"b", "c", "d"},
			{"1", "2", "3", "4"},
			{},
		}

		stringArrayToDFWRule := func(strArr []string) []*models.DFWRule {
			var result []*models.DFWRule
			for i := 0; i < len(strArr); i++ {
				result = append(result, &models.DFWRule{DisplayName: strArr[i]})
			}
			return result
		}

		for i := 0; i < len(matchingCases); i++ {
			assert.True(t, checkOrderOfSecurityPolicies(stringArrayToDFWRule(matchingCases[i]), expectedOrder))
		}

		for i := 0; i < len(notMatchingCases); i++ {
			assert.False(t, checkOrderOfSecurityPolicies(stringArrayToDFWRule(notMatchingCases[i]), expectedOrder))
		}
	})
}

func checkOrderOfSecurityPolicies(securityPolicies []*models.DFWRule, displayNames []string) bool {
	indexOfFirstElement := findIndexOfFirstMatch(len(securityPolicies), func(idx int) bool {
		return securityPolicies[idx].DisplayName == displayNames[0]
	})

	if indexOfFirstElement == -1 {
		return false
	}

	if indexOfFirstElement+len(displayNames) > len(securityPolicies) {
		return false
	}

	leftSlice := securityPolicies[indexOfFirstElement : indexOfFirstElement+len(displayNames)]
	return slicesEqual(len(leftSlice), len(displayNames), func(lhsIdx int, rhsIdx int) bool {
		return leftSlice[lhsIdx].DisplayName == displayNames[rhsIdx]
	})
}

func testAccFirewallEWRuleResourceConfig(displayName string, source string, destination string, action string, ipProtocol string, direction string, priority int64, customServiceResourceID string) string {
	return fmt.Sprintf(`

locals {
	routerDisplayName = %[1]q
}

data "ochk_vrf" "vrf-default" {
	display_name = %[13]q
}

data "ochk_project" "project-1" {
  display_name = "`+testData.Project1Name+`"
}

resource "ochk_vpc" "default" {
  display_name = local.routerDisplayName
  vrf_id = data.ochk_vrf.vrf-default.id
  project_id = data.ochk_project.project-1.id
}

data "ochk_service" "http" {
  display_name = "http"
}

data "ochk_virtual_machine" "default" {
	display_name = %[7]q
}

data "ochk_custom_service" "custom_service1" {
	display_name = %[9]q
}

data "ochk_custom_service" "custom_service2" {
	display_name = %[10]q
}

resource "ochk_security_group" "source" {
  display_name = %[1]q
  project_id = data.ochk_project.project-1.id

  members {
    id = data.ochk_virtual_machine.default.id
    type = "VIRTUAL_MACHINE"
  }
}

resource "ochk_security_group" "destination" {
  display_name = %[2]q
  project_id = data.ochk_project.project-1.id
  
  members {
    id = data.ochk_virtual_machine.default.id
    type = "VIRTUAL_MACHINE"
  }
}

resource "ochk_firewall_ew_rule" "default" {
  display_name = %[3]q
  vpc_id = ochk_vpc.default.id
  project_id = data.ochk_project.project-1.id

  services = [data.ochk_service.http.id]
  custom_services = [%[11]s]
  source = [ochk_security_group.source.id]
  destination = [ochk_security_group.destination.id]

  action = %[4]q 
  ip_protocol = %[5]q
  direction = %[6]q

  priority = %[8]d
}
`, source, destination, displayName, action, ipProtocol, direction, testData.VirtualMachineDisplayName, priority, testData.CustomService1DisplayName, testData.CustomService2DisplayName, customServiceResourceID, testData.VPC, testData.VRF)
}

func testAccFirewallEWRuleResourceConfigWithOrder(displayNameBefore string, displayNameMiddle string, displayNameAfter string) string {
	source := generateRandName(devTestDataPrefix)
	destination := generateRandName(devTestDataPrefix)

	return fmt.Sprintf(`

locals {
	routerDisplayName = %[1]q
}

data "ochk_vrf" "vrf" {
	display_name = %[7]q
}

data "ochk_project" "project-1" {
  display_name = "`+testData.Project1Name+`"
}

resource "ochk_vpc" "default" {
  display_name = local.routerDisplayName
  vrf_id = data.ochk_vrf.vrf.id
  project_id = data.ochk_project.project-1.id
}

data "ochk_service" "http" {
  display_name = "http"
}

data "ochk_virtual_machine" "default" {
	display_name = %[6]q
}

resource "ochk_security_group" "source-before" {
  display_name = "%[1]sbef"
  project_id = data.ochk_project.project-1.id

  members {
    id = data.ochk_virtual_machine.default.id
    type = "VIRTUAL_MACHINE"
  }
}

resource "ochk_security_group" "source-middle" {
  display_name = "%[1]smid"
  project_id = data.ochk_project.project-1.id

  members {
    id = data.ochk_virtual_machine.default.id
    type = "VIRTUAL_MACHINE"
  }
}

resource "ochk_security_group" "source-after" {
  display_name = "%[1]saft"
  project_id = data.ochk_project.project-1.id

  members {
    id = data.ochk_virtual_machine.default.id
    type = "VIRTUAL_MACHINE"
  }
}

resource "ochk_security_group" "destination-before" {
  display_name = "%[2]sbef"
  project_id = data.ochk_project.project-1.id
  
  members {
    id = data.ochk_virtual_machine.default.id
    type = "VIRTUAL_MACHINE"
  }
}

resource "ochk_security_group" "destination-middle" {
  display_name = "%[2]smid"
  project_id = data.ochk_project.project-1.id
  
  members {
    id = data.ochk_virtual_machine.default.id
    type = "VIRTUAL_MACHINE"
  }
}

resource "ochk_security_group" "destination-after" {
  display_name = "%[2]saft"
  project_id = data.ochk_project.project-1.id
  
  members {
    id = data.ochk_virtual_machine.default.id
    type = "VIRTUAL_MACHINE"
  }
}

resource "ochk_firewall_ew_rule" "middle" {
  display_name = %[3]q
  project_id = data.ochk_project.project-1.id
  vpc_id = ochk_vpc.default.id

  services = [data.ochk_service.http.id]
  source = [ochk_security_group.source-middle.id]
  destination = [ochk_security_group.destination-middle.id]
  priority = 20001
}

resource "ochk_firewall_ew_rule" "before" {
  display_name = %[4]q
  project_id = data.ochk_project.project-1.id
  vpc_id = ochk_vpc.default.id

  services = [data.ochk_service.http.id]
  source = [ochk_security_group.source-before.id]
  destination = [ochk_security_group.destination-before.id]
  priority = 20000
}

resource "ochk_firewall_ew_rule" "after" {
  display_name = %[5]q
  project_id = data.ochk_project.project-1.id
  vpc_id = ochk_vpc.default.id

  services = [data.ochk_service.http.id]
  source = [ochk_security_group.source-after.id]
  destination = [ochk_security_group.destination-after.id]
  priority = 20002
}
`, source, destination, displayNameMiddle, displayNameBefore, displayNameAfter, testData.VirtualMachineDisplayName, testData.VRF)
}

func testAccFirewallEWRuleResourceDoesntExist(displayName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ctx := context.Background()

		client := testAccProvider.Meta().(*sdk.Client)

		routers, err := client.Routers.ListByDisplayName(ctx, testData.VPC)

		if err != nil {
			return err
		}

		if len(routers) != 1 {
			return fmt.Errorf("wrong number of routers")
		}

		firewallRule, err := client.FirewallEWRules.ListByDisplayName(ctx, routers[0].RouterID, displayName)
		if err != nil {
			return err
		}

		if len(firewallRule) > 0 {
			return fmt.Errorf("firewall EW rule %s still exists", firewallRule[0].RuleID)
		}

		return nil
	}
}
