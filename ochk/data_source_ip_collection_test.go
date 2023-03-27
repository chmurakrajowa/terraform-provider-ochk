package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccIPCollectionDataSource_read(t *testing.T) {
	resourceName := "data.ochk_ip_collection.example"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIPCollectionDataSourceConfig(testData.IPCollection1DisplayName, testData.Project1Name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", testData.IPCollection1DisplayName),
					resource.TestCheckResourceAttrPair(resourceName, "project_id", "data.ochk_project.project-ipc-1", "id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_by"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_at"),
					resource.TestCheckResourceAttr(resourceName, "ip_addresses.0", "1.0.0.1"),
				),
			},
		},
	})
}

func testAccIPCollectionDataSourceConfig(displayName string, projectName string) string {
	return fmt.Sprintf(`
data "ochk_ip_collection" "example" {
  display_name = %[1]q
}

data "ochk_project" "project-ipc-1" {
	 display_name = %[2]q
}
`, displayName, projectName)
}
