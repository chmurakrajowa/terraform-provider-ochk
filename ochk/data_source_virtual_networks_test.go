package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccVirtualNetworksDataSource_read(t *testing.T) {
	resourceName := "data.ochk_virtual_networks.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVirtualNetworksDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "virtual_networks.0.virtual_network_id"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_networks.0.display_name"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_networks.0.folder_path"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_networks.0.ipam_enabled"),
					//Fixme: first element sometimes is L2 segment without vpc_id.
					//resource.TestCheckResourceAttrSet(resourceName, "virtual_networks.0.vpc_id"),
				),
			},
		},
	})
}

func testAccVirtualNetworksDataSourceConfig() string {
	return fmt.Sprintf(`
data "ochk_virtual_networks" "default" {
}
`)
}
