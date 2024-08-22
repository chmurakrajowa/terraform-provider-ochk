package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

type SecurityGroupTestData struct {
	ResourceName string
	DisplayName  string
	Members      []SecurityGroupMemberTestData
}

type SecurityGroupMemberTestData struct {
	ID   strfmt.UUID
	Type string
}

func (c *SecurityGroupTestData) ToString(projNameSuffix string) string {
	return executeTemplateToString(`

data "ochk_project" "proj-1`+projNameSuffix+`" {
  display_name = "`+testData.Project1Name+`"
}

resource "ochk_security_group" "{{ .ResourceName}}" {
  display_name = "{{ .DisplayName}}"
  project_id = data.ochk_project.proj-1`+projNameSuffix+`.id

  {{range $member := .Members}}
  members {
    id   = {{ $member.ID }}
    type = "{{ $member.Type }}"
  }
  {{end}}
}
`, c)
}

func (c *SecurityGroupTestData) FullResourceName() string {
	return "ochk_security_group." + c.ResourceName
}

func TestAccSecurityGroupResource_create(t *testing.T) {
	/* Security group with one member */
	virtualMachine := VirtualMachineDataSourceTestData{
		ResourceName: "default",
		DisplayName:  testData.VirtualMachineDisplayName,
	}

	securityGroup := SecurityGroupTestData{
		ResourceName: "one_member",
		DisplayName:  generateRandName(devTestDataPrefix),
		Members: []SecurityGroupMemberTestData{
			{
				ID:   strfmt.UUID(testDataResourceID(&virtualMachine)),
				Type: "VIRTUAL_MACHINE",
			},
		},
	}
	configOneMember := virtualMachine.ToString("-sc1") + securityGroup.ToString("-one-mbmr")

	/* Security group with one member with updated display_name */
	securityGroupUpdated := securityGroup
	securityGroupUpdated.DisplayName += "-upd"
	configOneMemberUpdated := virtualMachine.ToString("-sc1-upd") + securityGroupUpdated.ToString("-one-mbmr-upd")

	/* Security group with two members and updated display_name */
	securityGroupTwoMembers := securityGroupUpdated
	virtualMachine2 := VirtualMachineDataSourceTestData{
		ResourceName: "default2",
		DisplayName:  testData.VirtualMachine2DisplayName,
	}
	securityGroupTwoMembers.Members = append(securityGroupTwoMembers.Members, SecurityGroupMemberTestData{
		ID:   strfmt.UUID(testDataResourceID(&virtualMachine2)),
		Type: "VIRTUAL_MACHINE",
	})
	configTwoMembers := securityGroupTwoMembers.ToString("-two-mbmrs") + virtualMachine.ToString("-sc21") + virtualMachine2.ToString("-sg22")

	securityGroupResourceName := securityGroup.FullResourceName()
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: configOneMember,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(securityGroupResourceName, "display_name", securityGroup.DisplayName),
					resource.TestCheckResourceAttrPair(securityGroupResourceName, "project_id", "data.ochk_project.proj-1-one-mbmr", "id"),
					resource.TestCheckResourceAttrPair(securityGroupResourceName, "members.0.id", virtualMachine.FullResourceName(), "id"),
					resource.TestCheckResourceAttr(securityGroupResourceName, "members.0.type", securityGroup.Members[0].Type),
					resource.TestCheckResourceAttrSet(securityGroupResourceName, "members.0.display_name"),
					resource.TestCheckResourceAttrSet(securityGroupResourceName, "created_by"),
					resource.TestCheckResourceAttrSet(securityGroupResourceName, "created_at"),
					resource.TestCheckResourceAttrSet(securityGroupResourceName, "modified_by"),
					resource.TestCheckResourceAttrSet(securityGroupResourceName, "modified_at"),
				),
			},
			{
				ResourceName:      securityGroupResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: configOneMemberUpdated,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(securityGroupResourceName, "display_name", securityGroupUpdated.DisplayName),
					resource.TestCheckResourceAttrPair(securityGroupResourceName, "project_id", "data.ochk_project.proj-1-one-mbmr-upd", "id"),
				),
			},
			{
				Config: configTwoMembers,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(securityGroupResourceName, "display_name", securityGroupTwoMembers.DisplayName),
					resource.TestCheckResourceAttrPair(securityGroupResourceName, "project_id", "data.ochk_project.proj-1-two-mbmrs", "id"),
					//TODO this relies on certain ordering in backend, should be fixed to any ordering
					resource.TestCheckResourceAttrPair(securityGroupResourceName, "members.0.id", virtualMachine.FullResourceName(), "id"),
					resource.TestCheckResourceAttr(securityGroupResourceName, "members.0.type", securityGroupTwoMembers.Members[0].Type),
					resource.TestCheckResourceAttrSet(securityGroupResourceName, "members.0.display_name"),
					resource.TestCheckResourceAttrPair(securityGroupResourceName, "members.1.id", virtualMachine2.FullResourceName(), "id"),
					resource.TestCheckResourceAttr(securityGroupResourceName, "members.1.type", securityGroupTwoMembers.Members[1].Type),
					resource.TestCheckResourceAttrSet(securityGroupResourceName, "members.1.display_name"),
				),
			},
		},
		CheckDestroy: testAccSecurityGroupResourceNotExists(securityGroupTwoMembers.DisplayName),
	})
}

func testAccSecurityGroupResourceNotExists(displayName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		proxy := testAccProvider.Meta().(*sdk.Client).SecurityGroups

		securityGroups, err := proxy.ListByDisplayName(context.Background(), displayName)
		if err != nil {
			return err
		}

		if len(securityGroups) > 0 {
			return fmt.Errorf("security group %s still exists", securityGroups[0].ID)
		}

		return nil
	}
}
