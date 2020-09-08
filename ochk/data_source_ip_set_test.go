package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccIPSetDataSource_read(t *testing.T) {
	resourceName := "data.ochk_ip_set.example"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIPSetDataSourceConfig(testDataIPSet1DisplayName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", testDataIPSet1DisplayName),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_by"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_at"),
					resource.TestCheckResourceAttr(resourceName, "addresses.0.address", "10.10.10.1"),
				),
			},
		},
	})
}

func testAccIPSetDataSourceConfig(displayName string) string {
	return fmt.Sprintf(`
data "ochk_ip_set" "example" {
  display_name = %[1]q
}
`, displayName)
}
