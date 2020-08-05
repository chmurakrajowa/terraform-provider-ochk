package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccSecurityGroupDataSource_read(t *testing.T) {
	resourceName := "data.ochk_security_group.one_member"
	displayName := fmt.Sprintf("tf-acc_test-%s", acctest.RandStringFromCharSet(8, acctest.CharSetAlphaNum))

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityGroupDataSourceConfig(displayName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", displayName),
					resource.TestCheckResourceAttr(resourceName, "members.0.id", "e1e2f617-014c-4119-bac8-49fa4a93db47"),
					resource.TestCheckResourceAttr(resourceName, "members.0.type", "VIRTUAL_MACHINE"),
					resource.TestCheckResourceAttrSet(resourceName, "members.0.display_name"),
				),
			},
		},
		CheckDestroy: testAccSecurityGroupResourceExists(displayName),
	})
}

func testAccSecurityGroupDataSourceConfig(displayName string) string {
	return fmt.Sprintf(`
resource "ochk_security_group" "one_member" {
  display_name = %[1]q

  members {
    id = "e1e2f617-014c-4119-bac8-49fa4a93db47"
    type = "VIRTUAL_MACHINE"
  }
}

data "ochk_security_group" "one_member" {
  display_name = ochk_security_group.one_member.display_name
}
`, displayName)
}
