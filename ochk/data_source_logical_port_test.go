package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccLogicalPortDataSource_read(t *testing.T) {
	resourceName := "data.ochk_logical_port.example"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLogicalPortDataSourceConfig(testDataLogicalPort1DisplayName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", testDataLogicalPort1DisplayName),
				),
			},
		},
	})
}

func testAccLogicalPortDataSourceConfig(displayName string) string {
	return fmt.Sprintf(`
data "ochk_logical_port" "example" {
  display_name = %[1]q
}
`, displayName)
}
