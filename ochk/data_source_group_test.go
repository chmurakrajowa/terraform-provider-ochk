package ochk

/*

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccGroupDataSource_read(t *testing.T) {
	resourceName := "data.ochk_group.admin-InfraAdm1"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep {
			{
				Config: testAccGroupDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", testData.InfraAdminGroup),
					resource.TestCheckResourceAttrSet(resourceName, "description"),
					resource.TestCheckResourceAttrSet(resourceName, "group_type"),
					resource.TestCheckResourceAttrSet(resourceName, "domain"),
					resource.TestCheckResourceAttrSet(resourceName, "users"),
					resource.TestCheckResourceAttrSet(resourceName, "child_groups"),
				),
			},
		},
	})
}

func testAccGroupDataSourceConfig() string {
	return fmt.Sprintf(`
data "ochk_group" "admin-InfraAdm" {
	name = %[1]q
}
`, testData.InfraAdminGroup)
}
*/
