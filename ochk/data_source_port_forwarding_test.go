package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccPortForwardingDataSource_read(t *testing.T) {
	resourceName := "data.ochk_port_forwarding.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPortForwardingDataSourceConfig(testData.PortForwardingName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", testData.PortForwardingName),
					resource.TestCheckResourceAttrSet(resourceName, "floating_ip_id"),
					resource.TestCheckResourceAttrSet(resourceName, "port_forwarding_id"),
					resource.TestCheckResourceAttrSet(resourceName, "description"),
					resource.TestCheckResourceAttrSet(resourceName, "public_address"),
					resource.TestCheckResourceAttrSet(resourceName, "external_port"),
					resource.TestCheckResourceAttrSet(resourceName, "external_port_range"),
					resource.TestCheckResourceAttrSet(resourceName, "internal_ip_address"),
					resource.TestCheckResourceAttrSet(resourceName, "internal_port"),
					resource.TestCheckResourceAttrSet(resourceName, "internal_port_range"),
					resource.TestCheckResourceAttrSet(resourceName, "protocol"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_by"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_at"),
				),
			},
		},
	})
}

func testAccPortForwardingDataSourceConfig(PortFwdDisplayName string) string {
	return fmt.Sprintf(`
data "ochk_floating_ip_address" "floating_ip" {
     display_name = "`+testData.FloatingIpAddressName+`"
}

data "ochk_port_forwarding" "default" {
  display_name = %[1]q
  floating_ip_id = data.ochk_floating_ip_address.floating_ip.id
}
`, PortFwdDisplayName)
}
