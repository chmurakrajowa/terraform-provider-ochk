package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccFloatingIPsDataSource_read(t *testing.T) {
	resourceName := "data.ochk_floating_ip_addresses.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFloatingIPsSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "floating_ip_addresses.0.floating_ip_id"),
					resource.TestCheckResourceAttrSet(resourceName, "floating_ip_addresses.0.public_address"),
					resource.TestCheckResourceAttrSet(resourceName, "floating_ip_addresses.0.display_name"),
				),
			},
		},
	})
}

func testAccFloatingIPsSourceConfig() string {
	return fmt.Sprintf(`

data "ochk_floating_ip_addresses" "default" {
}
`)
}
