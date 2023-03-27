package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccRouterDataSource_read(t *testing.T) {
	resourceName := "data.ochk_vpc.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRouterDataSourceConfig(testData.VPC),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", testData.VPC),
					resource.TestCheckResourceAttrPair(resourceName, "project_id", "data.ochk_project.project-vpc-1", "id"),
					resource.TestCheckResourceAttr(resourceName, "folder_path", "/"),
					resource.TestCheckResourceAttrPair(resourceName, "vrf_id", "data.ochk_vrf.vrf", "id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_by"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_at"),
				),
			},
		},
	})
}

func testAccRouterDataSourceConfig(displayName string) string {
	return fmt.Sprintf(`

data "ochk_project" "project-vpc-1" {
	 display_name = %[3]q
}

data "ochk_vrf" "vrf" {
	display_name = %[2]q
}

data "ochk_vpc" "default" {
  display_name = %[1]q
  vrf_id = data.ochk_vrf.vrf.id
}
`, displayName, testData.VRF, testData.Project1Name)
}
