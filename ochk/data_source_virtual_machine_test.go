package ochk

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

type VirtualMachineDataSourceTestData struct {
	ResourceName string
	DisplayName  string
}

func (c *VirtualMachineDataSourceTestData) ToString(projectName string) string {
	return executeTemplateToString(`
data "ochk_virtual_machine" "{{ .ResourceName}}" {
  display_name = "{{ .DisplayName}}"
}

data "ochk_project" "project`+projectName+`-1" {
	 display_name = "`+testData.Project1Name+`"
}
`, c)
}

func (c *VirtualMachineDataSourceTestData) FullResourceName() string {
	return "data.ochk_virtual_machine." + c.ResourceName
}

func TestAccVirtualMachineDataSource_read(t *testing.T) {
	virtualMachine := &VirtualMachineDataSourceTestData{
		ResourceName: "default",
		DisplayName:  testData.VirtualMachineDisplayName,
	}

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: virtualMachine.ToString("-vm"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(virtualMachine.FullResourceName(), "display_name", virtualMachine.DisplayName),
					resource.TestCheckResourceAttr(virtualMachine.FullResourceName(), "folder_path", "/"),
					resource.TestCheckResourceAttrPair(virtualMachine.FullResourceName(), "project_id", "data.ochk_project.project-vm-1", "id"),
				),
			},
		},
	})
}
