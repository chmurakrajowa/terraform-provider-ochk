package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccSecurityPolicyDataSource_read(t *testing.T) {
	resourceName := "data.ochk_security_policy.default"

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityPolicyDataSourceConfig("devel"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", "devel"),
				),
			},
		},
	})
}

func testAccSecurityPolicyDataSourceConfig(displayName string) string {
	return fmt.Sprintf(`
data "ochk_security_policy" "default" {
  display_name = %[1]q
}
`, displayName)
}
