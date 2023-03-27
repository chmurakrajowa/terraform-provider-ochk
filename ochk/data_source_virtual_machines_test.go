package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccVirtualMachinesDataSource_read(t *testing.T) {
	resourceName := "data.ochk_virtual_machines.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVirtualMachinesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "virtual_machines.0.virtual_machine_id"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_machines.0.display_name"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_machines.0.folder_path"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_machines.0.project_id"),
				),
			},
		},
	})
}

func testAccVirtualMachinesDataSourceConfig() string {
	return fmt.Sprintf(`
data "ochk_virtual_machines" "default" {
}
`)
}
