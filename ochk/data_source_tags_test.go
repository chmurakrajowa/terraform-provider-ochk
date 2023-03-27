package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccTagsDataSource_read(t *testing.T) {
	resourceName := "data.ochk_tags.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTagsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "tags.0.display_name"),
					resource.TestCheckResourceAttrSet(resourceName, "tags.0.tag_id"),
					resource.TestCheckResourceAttrSet(resourceName, "tags.0.project_id"),
				),
			},
		},
	})
}

func testAccTagsDataSourceConfig() string {
	return fmt.Sprintf(`
data "ochk_tags" "default" {
}
`)
}
