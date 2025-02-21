package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

func testAccPortForwardingResourceConfig(display_name string, protocol string, internal_port_id string, internal_ip_address string, internal_port int, public_address string, external_port int) string {
	return fmt.Sprintf(`

data "ochk_floating_ip_address" "floating_ip" {
     display_name = "`+testData.FloatingIpAddressName+`"
}

resource "ochk_port_forwarding" "default" {
  display_name = %[1]q
  floating_ip_id = data.ochk_floating_ip_address.floating_ip.id
  protocol = %[2]q
  internal_port_id = %[3]q
  internal_ip_address = %[4]q
  internal_port = %[5]d
  public_address = %[6]q
  external_port = %[7]d
}
`, display_name, protocol, internal_port_id, internal_ip_address, internal_port, public_address, external_port)
}

func TestAccPortForwardingResource_create_update(t *testing.T) {
	resourceName := "ochk_port_forwarding.default"
	display_name := generateRandName(devTestDataPrefix)
	display_name_updated := display_name + "-upd"

	dataSourceFloatingIp := "ochk_floating_ip_address.floating_ip"

	protocol := "tcp"
	internal_port_id := "d2b24e3e-b6c1-44ef-a227-479b8f4bbc21"
	internal_ip_address := "192.168.2.209"
	internal_port := 1219
	internal_port_upd := 1216
	public_address := "203.0.113.29"
	external_port := 1221
	external_port_upd := 1227

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPortForwardingResourceConfig(display_name, protocol, internal_port_id, internal_ip_address, internal_port, public_address, external_port),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(resourceName, "floating_ip_id", dataSourceFloatingIp, "id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", display_name),
					resource.TestCheckResourceAttr(resourceName, "protocol", "tcp"),
					resource.TestCheckResourceAttr(resourceName, "internal_port_id", "11752b30-9749-4534-8ec1-2b0e9670b92b"),
					resource.TestCheckResourceAttr(resourceName, "internal_ip_address", "10.10.123.185"),
					resource.TestCheckResourceAttr(resourceName, "internal_port", "1219"),
					resource.TestCheckResourceAttr(resourceName, "public_address", "203.0.113.29"),
					resource.TestCheckResourceAttr(resourceName, "external_port", "1221"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_by"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_at"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccPortForwardingResourceConfig(display_name_updated, protocol, internal_port_id, internal_ip_address, internal_port_upd, public_address, external_port_upd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", display_name_updated),
					resource.TestCheckResourceAttr(resourceName, "internal_port", "1216"),
					resource.TestCheckResourceAttr(resourceName, "external_port", "1227"),
				),
			},
		},
		CheckDestroy: resource.ComposeTestCheckFunc(
			testAccPortForwardingResourceNotExists(display_name, strfmt.UUID(dataSourceFloatingIp)),
			testAccPortForwardingResourceNotExists(display_name_updated, strfmt.UUID(dataSourceFloatingIp)),
		),
	})
}

func testAccPortForwardingResourceNotExists(displayName string, floatingIpId strfmt.UUID) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		proxy := testAccProvider.Meta().(*sdk.Client).PortForwarding

		PortFwd, err := proxy.ListByName(context.Background(), floatingIpId, displayName)
		if err != nil {
			return err
		}

		if len(PortFwd) > 0 {
			return fmt.Errorf("port_forwarding %s still exists", PortFwd[0].PortForwardingID)
		}
		return nil
	}
}
