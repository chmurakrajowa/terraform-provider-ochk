package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccSnapshotsDataSource_read(t *testing.T) {
	resourceName := "data.ochk_snapshots.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSnapshotsDataSourceConfig(testData.VirtualMachine1DisplayName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "snapshots.0.snapshot_id"),
					resource.TestCheckResourceAttrSet(resourceName, "snapshots.0.display_name"),
				),
			},
		},
	})
}

func testAccSnapshotsDataSourceConfig(virtualMachineName string) string {
	return fmt.Sprintf(`
data "ochk_virtual_machine" "virtual-machine-1" {
	 display_name = %[1]q
}
data "ochk_snapshots" "default" {
	virtual_machine_id = data.ochk_virtual_machine.virtual-machine-1.id
}

`, virtualMachineName)
}
