package ochk

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"testing"
)

type SecurityGroupTestData struct {
	ResourceName string
	DisplayName  string
	Members      []SecurityGroupMemberTestData
}

type SecurityGroupMemberTestData struct {
	ID   string
	Type string
}

func (c *SecurityGroupTestData) ToString() string {
	return executeTemplateToString(`
resource "ochk_security_group" "{{ .ResourceName}}" {
  display_name = "{{ .DisplayName}}"

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
		DisplayName:  "devel0000000344",
	}

	securityGroup := SecurityGroupTestData{
		ResourceName: "one_member",
		DisplayName:  generateRandName(),
		Members: []SecurityGroupMemberTestData{
			{
				ID:   virtualMachine.FullResourceName() + ".id",
				Type: "VIRTUAL_MACHINE",
			},
		},
	}
	configOneMember := virtualMachine.ToString() + securityGroup.ToString()

	/* Security group with one member with updated display_name */
	securityGroupUpdated := securityGroup
	securityGroupUpdated.DisplayName += "-updated"
	configOneMemberUpdated := virtualMachine.ToString() + securityGroupUpdated.ToString()

	/* Security group with two members and updated display_name */
	securityGroupTwoMembers := securityGroupUpdated
	virtualMachine2 := VirtualMachineDataSourceTestData{
		ResourceName: "default2",
		DisplayName:  testDataVirtualMachine1DisplayName,
	}
	securityGroupTwoMembers.Members = append(securityGroupTwoMembers.Members, SecurityGroupMemberTestData{
		ID:   virtualMachine2.FullResourceName() + ".id",
		Type: "VIRTUAL_MACHINE",
	})
	configTwoMembers := securityGroupTwoMembers.ToString() + virtualMachine.ToString() + virtualMachine2.ToString()

	securityGroupResourceName := securityGroup.FullResourceName()
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: configOneMember,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(securityGroupResourceName, "display_name", securityGroup.DisplayName),
					resource.TestCheckResourceAttrPair(securityGroupResourceName, "members.0.id", virtualMachine.FullResourceName(), "id"),
					resource.TestCheckResourceAttr(securityGroupResourceName, "members.0.type", securityGroup.Members[0].Type),
					resource.TestCheckResourceAttrSet(securityGroupResourceName, "members.0.display_name"),
				),
			},
			{
				Config: configOneMemberUpdated,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(securityGroupResourceName, "display_name", securityGroupUpdated.DisplayName),
				),
			},
			{
				Config: configTwoMembers,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(securityGroupResourceName, "display_name", securityGroupTwoMembers.DisplayName),
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
