package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccCustomServicesDataSource_read(t *testing.T) {
	resourceName := "data.ochk_custom_services.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomServicesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "custom_services.0.custom_service_id"),
					resource.TestCheckResourceAttrSet(resourceName, "custom_services.0.display_name"),
					resource.TestCheckResourceAttrSet(resourceName, "custom_services.0.project_id"),
				),
			},
		},
	})
}

func testAccCustomServicesDataSourceConfig() string {
	return fmt.Sprintf(`
data "ochk_custom_services" "default" {
}
`)
}
