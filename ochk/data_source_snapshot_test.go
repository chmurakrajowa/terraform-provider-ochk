package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccSnapshotDataSource_read(t *testing.T) {
	resourceName := "data.ochk_snapshot.snap-def"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSnapshotDataSourceConfig(testData.VirtualMachine1DisplayName, testData.SnapshotName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", testData.SnapshotName),
					resource.TestCheckResourceAttrPair(resourceName, "virtual_machine_id", "data.ochk_virtual_machine.virtual-machine-1", "id"),
					resource.TestCheckResourceAttrSet(resourceName, "snapshot_description"),
					resource.TestCheckResourceAttrSet(resourceName, "power_state"),
					resource.TestCheckResourceAttrSet(resourceName, "parent_id"),
				),
			},
		},
	})
}

func testAccSnapshotDataSourceConfig(virtualMachineName string, snapshotName string) string {
	return fmt.Sprintf(`
data "ochk_virtual_machine" "virtual-machine-1" {
	 display_name = %[1]q
}
data "ochk_snapshot" "snap-def" {
	virtual_machine_id = data.ochk_virtual_machine.virtual-machine-1.id
	display_name = %[2]q
}
`, virtualMachineName, snapshotName)
}
