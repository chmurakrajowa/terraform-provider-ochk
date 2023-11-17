package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

type SnapshotTestData struct {
	ResourceName     string
	Name             string
	VirtualMachineID string
	PowerState       string
}

func (c *SnapshotTestData) ToString() string {
	return executeTemplateToString(`

data "ochk_virtual_machine" "vm-snaps-1" {
  display_name = "`+testData.VirtualMachineDisplayName+`"
}

resource "ochk_snapshot" "{{ .ResourceName}}" {
  display_name = "{{ .Name}}"
  virtual_machine_id = data.ochk_virtual_machine.vm-snaps-1.id
}
`, c)
}

func (c *SnapshotTestData) FullResourceName() string {
	return "ochk_snapshot." + c.ResourceName
}

func TestAccSnapshotResource_create(t *testing.T) {
	virtualMachine := VirtualMachineDataSourceTestData{
		ResourceName: "default",
		DisplayName:  testData.VirtualMachineDisplayName,
	}

	snapshot := SnapshotTestData{
		ResourceName:     "default",
		Name:             generateShortRandName(devTestDataPrefix),
		VirtualMachineID: testDataResourceID(&virtualMachine),
	}

	SnapshotResourceName := snapshot.FullResourceName()
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: snapshot.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(SnapshotResourceName, "display_name", snapshot.Name),
					resource.TestCheckResourceAttrPair(SnapshotResourceName, "virtual_machine_id", "data.ochk_virtual_machine.vm-snaps-1.id", "id"),
				),
			},
			{
				ResourceName:      SnapshotResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
		CheckDestroy: testAccSnapshotResourceNotExists(snapshot.Name, snapshot.VirtualMachineID),
	})
}

func testAccSnapshotResourceNotExists(Name string, VirtualMachineID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		proxy := testAccProvider.Meta().(*sdk.Client).Snapshots

		snapshots, err := proxy.ListSnapshotsByName(context.Background(), VirtualMachineID, Name)
		if err != nil {
			return err
		}

		if len(snapshots) > 0 {
			return fmt.Errorf("snapshot %s still exists", snapshots[0].SnapshotID)
		}

		return nil
	}
}
