package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccVirtualMachineDataSource_read(t *testing.T) {
	resourceName := "data.ochk_virtual_machine.default"
	displayName := "devel0000001157"

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVirtualMachineDataSourceConfig(displayName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", displayName),
				),
			},
		},
	})
}

func testAccVirtualMachineDataSourceConfig(displayName string) string {
	return fmt.Sprintf(`
data "ochk_virtual_machine" "default" {
  display_name = %[1]q
}
`, displayName)
}
