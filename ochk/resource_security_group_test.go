package ochk

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk"
	"testing"
	"text/template"
)

var securityGroupConfigTemplate *template.Template

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
	return executeTemplateToString(securityGroupConfigTemplate, c)
}

func (c *SecurityGroupTestData) FullResourceName() string {
	return "ochk_security_group." + c.ResourceName
}

func init() {
	securityGroupConfigTemplate = createNewTemplate("SecurityGroupConfigTemplate", `
resource "ochk_security_group" "{{ .ResourceName}}" {
  display_name = "{{ .DisplayName}}"

  {{range $member := .Members}}
  members {
    id   = {{ $member.ID }}
    type = "{{ $member.Type }}"
  }
  {{end}}
}
`)
}

func TestAccSecurityGroupResource_create(t *testing.T) {
	virtualMachine := VirtualMachineDataSourceTestData{
		ResourceName: "default",
		DisplayName:  testDataVirtualMachine1DisplayName,
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

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: securityGroup.ToString() + virtualMachine.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(securityGroup.FullResourceName(), "display_name", securityGroup.DisplayName),
					resource.TestCheckResourceAttrPair(securityGroup.FullResourceName(), "members.0.id", virtualMachine.FullResourceName(), "id"),
					resource.TestCheckResourceAttr(securityGroup.FullResourceName(), "members.0.type",  securityGroup.Members[0].Type),
					resource.TestCheckResourceAttrSet(securityGroup.FullResourceName(), "members.0.display_name"),
				),
			},
			{
				PreConfig: func() {
					securityGroup.DisplayName += "updated"
				},
				Config: securityGroup.ToString() + virtualMachine.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(securityGroup.FullResourceName(), "display_name", securityGroup.DisplayName),
				),
			},
		},
		CheckDestroy: testAccSecurityGroupResourceNotExists(securityGroup.DisplayName),
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
