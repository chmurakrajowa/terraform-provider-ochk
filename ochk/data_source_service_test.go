package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccServiceDataSource_read(t *testing.T) {
	resourceName := "data.ochk_service.http"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccServiceDataSourceConfig("http"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", "http"),
				),
			},
		},
	})
}

func testAccServiceDataSourceConfig(displayName string) string {
	return fmt.Sprintf(`
data "ochk_service" "http" {
  display_name = %[1]q
}
`, displayName)
}
