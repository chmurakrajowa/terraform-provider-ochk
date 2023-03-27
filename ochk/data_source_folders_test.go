package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccFoldersDataSource_read(t *testing.T) {
	resourceName := "data.ochk_folders.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFoldersDataSourceConfig(testData.Project1Name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "folders.0.folder_id"),
					resource.TestCheckResourceAttrSet(resourceName, "folders.0.folder_name"),
					resource.TestCheckResourceAttrSet(resourceName, "folders.0.folder_path"),
				),
			},
		},
	})
}

func testAccFoldersDataSourceConfig(projectName string) string {
	return fmt.Sprintf(`
data "ochk_project" "project-folder-1" {
	 display_name = %[1]q
}
data "ochk_folders" "default" {
	project_id = data.ochk_project.project-folder-1.id
}

`, projectName)
}
