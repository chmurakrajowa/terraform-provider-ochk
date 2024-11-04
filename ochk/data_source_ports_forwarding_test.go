package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccPortsForwardingDataSource_read(t *testing.T) {
	resourceName := "data.ochk_ports_forwarding.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPortsForwardingSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "ports_forwarding.0.port_forwarding_id"),
					resource.TestCheckResourceAttrSet(resourceName, "ports_forwarding.0.display_name"),
					resource.TestCheckResourceAttrSet(resourceName, "ports_forwarding.0.floating_ip_id"),
					resource.TestCheckResourceAttrSet(resourceName, "ports_forwarding.0.internal_port_id"),
				),
			},
		},
	})
}

func testAccPortsForwardingSourceConfig() string {
	return fmt.Sprintf(`

data "ochk_floating_ip_address" "floating_ip" {
     display_name = "` + testData.FloatingIpAddressName + `"
}

data "ochk_ports_forwarding" "default" {
	 floating_ip_id = data.ochk_floating_ip_address.floating_ip.id
}
`)
}
