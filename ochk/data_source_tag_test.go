package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccTagDataSource_read(t *testing.T) {
	resourceName := "data.ochk_tag.res-bt1"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTagDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", testData.TagName),
					resource.TestCheckResourceAttrPair(resourceName, "project_id", "data.ochk_project.project-tag-1", "id"),
				),
			},
		},
	})
}

func testAccTagDataSourceConfig() string {
	return fmt.Sprintf(`
data "ochk_tag" "res-bt1" {
	display_name = %[1]q
}

data "ochk_project" "project-tag-1" {
	 display_name = %[2]q
}
`, testData.TagName, testData.Project1Name)
}
