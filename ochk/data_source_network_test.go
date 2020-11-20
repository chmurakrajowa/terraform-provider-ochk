package ochk

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

type NetworkDataSourceTestData struct {
	ResourceName string
	Name         string
}

func (c *NetworkDataSourceTestData) ToString() string {
	return executeTemplateToString(`
data "ochk_network" "{{ .ResourceName}}" {
  name = "{{ .Name}}"
}
`, c)
}

func (c *NetworkDataSourceTestData) FullResourceName() string {
	return "data.ochk_network." + c.ResourceName
}

func TestAccNetworkDataSourceTestDataDataSource_read(t *testing.T) {
	network := &NetworkDataSourceTestData{
		ResourceName: "default",
		Name:         testData.Network1Name,
	}

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: network.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(network.FullResourceName(), "name", network.Name),
					resource.TestCheckResourceAttr(network.FullResourceName(), "network_type", "OpaqueNetwork"),
					resource.TestCheckResourceAttrSet(network.FullResourceName(), "network"),
				),
			},
		},
	})
}
