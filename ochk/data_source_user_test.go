package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccUserDataSource_read(t *testing.T) {
	resourceName := "data.ochk_user.default"
	userName := "devel-admin"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUserDataSourceConfig(userName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "devel-admin"),
					resource.TestCheckResourceAttr(resourceName, "email_address", "info@ochk.pl"),
					resource.TestCheckResourceAttr(resourceName, "description", "Devel Admin"),
					resource.TestCheckResourceAttr(resourceName, "first_name", "Devel"),
					resource.TestCheckResourceAttr(resourceName, "last_name", "Admin"),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "locked", "false"),
					resource.TestCheckResourceAttr(resourceName, "user_principal_name", "devel-admin@vsphere.local"),
					resource.TestCheckResourceAttr(resourceName, "principal_id", "252"),
					resource.TestCheckResourceAttr(resourceName, "principal_name", "devel-admin"),
					resource.TestCheckResourceAttr(resourceName, "principal_domain", "vsphere.local"),
					resource.TestCheckResourceAttr(resourceName, "principal_type", "USER"),
				),
			},
		},
	})
}

func testAccUserDataSourceConfig(displayName string) string {
	return fmt.Sprintf(`
data "ochk_user" "default" {
  name = %[1]q
}
`, displayName)
}
