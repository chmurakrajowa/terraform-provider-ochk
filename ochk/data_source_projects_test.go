package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccProjectsDataSource_read(t *testing.T) {
	resourceName := "data.ochk_projects.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "projects.0.project_id"),
					resource.TestCheckResourceAttrSet(resourceName, "projects.0.display_name"),
				),
			},
		},
	})
}

func testAccProjectsDataSourceConfig() string {
	return fmt.Sprintf(`
data "ochk_projects" "default" {
}
`)
}
