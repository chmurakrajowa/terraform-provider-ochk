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
