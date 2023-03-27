package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccIPCollectionsDataSource_read(t *testing.T) {
	resourceName := "data.ochk_ip_collections.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIPCollectionsSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "ip_collections.0.ip_collection_id"),
					resource.TestCheckResourceAttrSet(resourceName, "ip_collections.0.display_name"),
					resource.TestCheckResourceAttrSet(resourceName, "ip_collections.0.ip_addresses.0"),
					resource.TestCheckResourceAttrSet(resourceName, "ip_collections.0.project_id"),
				),
			},
		},
	})
}

func testAccIPCollectionsSourceConfig() string {
	return fmt.Sprintf(`
data "ochk_ip_collections" "default" {
}
`)
}
