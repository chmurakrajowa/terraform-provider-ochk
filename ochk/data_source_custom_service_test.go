package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccCustomServiceDataSource_read(t *testing.T) {
	resourceName := "data.ochk_custom_service.http"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomServiceDataSourceConfig(testData.CustomService1DisplayName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", testData.CustomService1DisplayName),
					resource.TestCheckResourceAttr(resourceName, "ports.0.protocol", "TCP"),
					testStringInSet(resourceName, "ports.0.source", "443"),
					testStringInSet(resourceName, "ports.0.destination", "1-65535"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_by"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_at"),
				),
			},
		},
	})
}

func testAccCustomServiceDataSourceConfig(displayName string) string {
	return fmt.Sprintf(`
data "ochk_custom_service" "http" {
  display_name = %[1]q
}
`, displayName)
}
