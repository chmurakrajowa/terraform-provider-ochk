package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccSecurityGroupDataSource_read(t *testing.T) {
	resourceName := "data.ochk_security_group.one_member"
	displayName := generateRandName()

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityGroupDataSourceConfig(displayName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", displayName),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_by"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_at"),
					resource.TestCheckResourceAttrPair(resourceName, "members.0.id", "data.ochk_virtual_machine.default", "id"),
					resource.TestCheckResourceAttr(resourceName, "members.0.type", "VIRTUAL_MACHINE"),
					resource.TestCheckResourceAttrSet(resourceName, "members.0.display_name"),
				),
			},
		},
	})
}

func testAccSecurityGroupDataSourceConfig(displayName string) string {
	return fmt.Sprintf(`
data "ochk_virtual_machine" "default" {
	display_name = %[2]q
}

resource "ochk_security_group" "one_member" {
  display_name = %[1]q

  members {
    id = data.ochk_virtual_machine.default.id
    type = "VIRTUAL_MACHINE"
  }
}

data "ochk_security_group" "one_member" {
  display_name = ochk_security_group.one_member.display_name
  depends_on = [ochk_security_group.one_member]
}
`, displayName, testData.LegacyVirtualMachineDisplayName)
}
