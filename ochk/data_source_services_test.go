package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccServicesDataSource_read(t *testing.T) {
	resourceName := "data.ochk_services.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccServicesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "services.0.service_id"),
					resource.TestCheckResourceAttrSet(resourceName, "services.0.display_name"),
				),
			},
		},
	})
}

func testAccServicesDataSourceConfig() string {
	return fmt.Sprintf(`
data "ochk_services" "default" {
}
`)
}
