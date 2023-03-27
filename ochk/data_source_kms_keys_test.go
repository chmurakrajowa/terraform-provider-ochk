package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccKeysDataSource_read(t *testing.T) {
	resourceName := "data.ochk_kms_keys.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccKeysSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "kms_keys.0.kms_key_id"),
					resource.TestCheckResourceAttrSet(resourceName, "kms_keys.0.display_name"),
					resource.TestCheckResourceAttrSet(resourceName, "kms_keys.0.key_usage.0"),
					resource.TestCheckResourceAttrSet(resourceName, "kms_keys.0.state"),
				),
			},
		},
	})
}

func testAccKeysSourceConfig() string {
	return fmt.Sprintf(`
data "ochk_kms_keys" "default" {
}
`)
}
