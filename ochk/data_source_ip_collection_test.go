package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccIPCollectionDataSource_read(t *testing.T) {
	resourceName := "data.ochk_ip_collection.example"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIPCollectionDataSourceConfig(testDataIPCollection1DisplayName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", testDataIPCollection1DisplayName),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_by"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_at"),
					resource.TestCheckResourceAttr(resourceName, "addresses.0.address", "100.100.100.100"),
				),
			},
		},
	})
}

func testAccIPCollectionDataSourceConfig(displayName string) string {
	return fmt.Sprintf(`
data "ochk_ip_collection" "example" {
  display_name = %[1]q
}
`, displayName)
}
