package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccRouterDataSource_read(t *testing.T) {
	resourceName := "data.ochk_router.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRouterDataSourceConfig("T1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", "T1"),
					resource.TestCheckResourceAttrSet(resourceName, "parent_t0_id"),
					resource.TestCheckResourceAttrSet(resourceName, "router_type"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_by"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_at"),
				),
			},
		},
	})
}

func testAccRouterDataSourceConfig(displayName string) string {
	return fmt.Sprintf(`
data "ochk_router" "default" {
  display_name = %[1]q
}
`, displayName)
}
