package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccFloatingIPVmsDataSource_read(t *testing.T) {
	resourceName := "data.ochk_floating_ip_vms.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFloatingIPVmsSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "floating_ip_vms.0.virtual_machine_id"),
					resource.TestCheckResourceAttrSet(resourceName, "floating_ip_vms.0.virtual_machine_name"),
					resource.TestCheckResourceAttrSet(resourceName, "floating_ip_vms.0.osc_port_id"),
					resource.TestCheckResourceAttrSet(resourceName, "floating_ip_vms.0.mac_address"),
					resource.TestCheckResourceAttrSet(resourceName, "floating_ip_vms.0.network_name"),
					resource.TestCheckResourceAttrSet(resourceName, "floating_ip_vms.0.ip_address"),
				),
			},
		},
	})
}

func testAccFloatingIPVmsSourceConfig() string {
	return fmt.Sprintf(`

data "ochk_floating_ip_vms" "default" {
}
`)
}
