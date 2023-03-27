package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccSecurityGroupsDataSource_read(t *testing.T) {
	resourceName := "data.ochk_security_groups.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityGroupsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "security_groups.0.security_group_id"),
					resource.TestCheckResourceAttrSet(resourceName, "security_groups.0.display_name"),
				),
			},
		},
	})
}

func testAccSecurityGroupsDataSourceConfig() string {
	return fmt.Sprintf(`
data "ochk_security_groups" "default" {
}
`)
}
