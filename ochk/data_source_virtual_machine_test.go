package ochk

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

type VirtualMachineDataSourceTestData struct {
	ResourceName string
	DisplayName  string
}

func (c *VirtualMachineDataSourceTestData) ToString() string {
	return executeTemplateToString(`
data "ochk_virtual_machine" "{{ .ResourceName}}" {
  display_name = "{{ .DisplayName}}"
}
`, c)
}

func (c *VirtualMachineDataSourceTestData) FullResourceName() string {
	return "data.ochk_virtual_machine." + c.ResourceName
}

func TestAccVirtualMachineDataSource_read(t *testing.T) {
	virtualMachine := &VirtualMachineDataSourceTestData{
		ResourceName: "default",
		DisplayName:  testData.LegacyVirtualMachineDisplayName,
	}

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: virtualMachine.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(virtualMachine.FullResourceName(), "display_name", virtualMachine.DisplayName),
					resource.TestCheckResourceAttrSet(virtualMachine.FullResourceName(), "host_id"),
				),
			},
		},
	})
}
