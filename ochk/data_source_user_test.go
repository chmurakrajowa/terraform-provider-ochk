package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccUserDataSource_read(t *testing.T) {
	resourceName := "data.ochk_user.default"
	userName := "devel-jpuser"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUserDataSourceConfig(userName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", userName),
					resource.TestCheckResourceAttr(resourceName, "email_address", "test@ochk.pl"),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
					resource.TestCheckResourceAttr(resourceName, "first_name", "Test"),
					resource.TestCheckResourceAttr(resourceName, "last_name", "Test"),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "locked", "false"),
					resource.TestCheckResourceAttr(resourceName, "user_principal_name", "devel-jpuser@under.test"),
					resource.TestCheckResourceAttr(resourceName, "principal_id", "487"),
					resource.TestCheckResourceAttr(resourceName, "principal_name", "devel-jpuser"),
					resource.TestCheckResourceAttr(resourceName, "principal_domain", "under.test"),
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
