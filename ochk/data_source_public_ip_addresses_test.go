package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccPublicIPAddressesDataSource_read(t *testing.T) {
	resourceName := "data.ochk_public_ip_addresses.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPublicIPAddressesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "public_ip_addresses.0.public_ip_address_id"),
					resource.TestCheckResourceAttrSet(resourceName, "public_ip_addresses.0.ip_address"),
					resource.TestCheckResourceAttrSet(resourceName, "public_ip_addresses.0.display_name"),
				),
			},
		},
	})
}

func testAccPublicIPAddressesDataSourceConfig() string {
	return fmt.Sprintf(`
data "ochk_public_ip_addresses" "default" {
}
`)
}
