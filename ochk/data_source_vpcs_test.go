package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccVPCsDataSource_read(t *testing.T) {
	resourceName := "data.ochk_vpcs.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVPCsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "vpcs.0.vpc_id"),
					resource.TestCheckResourceAttrSet(resourceName, "vpcs.0.vrf_id"),
					resource.TestCheckResourceAttrSet(resourceName, "vpcs.0.display_name"),
					resource.TestCheckResourceAttrSet(resourceName, "vpcs.0.project_id"),
					resource.TestCheckResourceAttrSet(resourceName, "vpcs.0.folder_path"),
				),
			},
		},
	})
}

func testAccVPCsDataSourceConfig() string {
	return fmt.Sprintf(`
data "ochk_vpcs" "default" {
}
`)
}
