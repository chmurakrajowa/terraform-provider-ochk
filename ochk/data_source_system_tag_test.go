package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccSystemTagDataSource_read(t *testing.T) {
	resourceName := "data.ochk_system_tag.res-st1"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemTagDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", testData.SystemTagName),
				),
			},
		},
	})
}

func testAccSystemTagDataSourceConfig() string {
	return fmt.Sprintf(`
data "ochk_system_tag" "res-st1" {
	display_name = %[1]q
}
`, testData.SystemTagName)
}
