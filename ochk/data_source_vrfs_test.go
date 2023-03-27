package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccVRFsDataSource_read(t *testing.T) {
	resourceName := "data.ochk_vrfs.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVRFsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "vrfs.0.vrf_id"),
					resource.TestCheckResourceAttrSet(resourceName, "vrfs.0.display_name"),
				),
			},
		},
	})
}

func testAccVRFsDataSourceConfig() string {
	return fmt.Sprintf(`
data "ochk_vrfs" "default" {
}
`)
}
