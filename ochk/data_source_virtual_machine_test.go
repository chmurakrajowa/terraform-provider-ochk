package ochk

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
	"text/template"
)

var virtualMachineDataSourceConfigTemplate *template.Template

type VirtualMachineDataSourceTestData struct {
	ResourceName string
	DisplayName  string
}

func (c *VirtualMachineDataSourceTestData) ToString() string {
	return executeTemplateToString(virtualMachineDataSourceConfigTemplate, c)
}

func (c *VirtualMachineDataSourceTestData) FullResourceName() string {
	return "data.ochk_virtual_machine." + c.ResourceName
}

func init() {
	virtualMachineDataSourceConfigTemplate = createNewTemplate("VirtualMachineDataSourceTemplate", `
data "ochk_virtual_machine" "{{ .ResourceName}}" {
  display_name = "{{ .DisplayName}}"
}
`)
}

func TestAccVirtualMachineDataSource_read(t *testing.T) {
	virtualMachine := VirtualMachineDataSourceTestData{
		ResourceName: "default",
		DisplayName:  testDataVirtualMachine1DisplayName,
	}

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: virtualMachine.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(virtualMachine.FullResourceName(), "display_name", virtualMachine.DisplayName),
				),
			},
		},
	})
}
