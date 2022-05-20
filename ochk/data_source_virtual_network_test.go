package ochk

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

type VirtualNetworkDataSourceTestData struct {
	ResourceName string
	DisplayName  string
}

func (c *VirtualNetworkDataSourceTestData) ToString() string {
	return executeTemplateToString(`
data "ochk_virtual_network" "{{ .ResourceName}}" {
  display_name = "{{ .DisplayName}}"
}
`, c)
}

func (c *VirtualNetworkDataSourceTestData) FullResourceName() string {
	return "data.ochk_virtual_network." + c.ResourceName
}

func TestAccVirtualNetworkDatasource(t *testing.T) {
	virtualNetwork := VirtualNetworkDataSourceTestData{
		ResourceName: "default",
		DisplayName:  testData.VirtualNetwork1DisplayName,
	}

	resourceName := virtualNetwork.FullResourceName()
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,

		Steps: []resource.TestStep{
			{
				Config: virtualNetwork.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", virtualNetwork.DisplayName),
				),
			},
		},
	})
}
