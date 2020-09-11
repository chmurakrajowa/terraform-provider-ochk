package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strconv"
	"testing"
)

type VirtualNetworkTestData struct {
	ResourceName string
	DisplayName  string
	IpamEnabled  bool
}

func (c *VirtualNetworkTestData) ToString() string {
	return executeTemplateToString(`
resource "ochk_virtual_network" "{{ .ResourceName}}" {
	display_name = "{{.DisplayName}}"
	ipam_enabled = {{.IpamEnabled}}
}
`, c)
}

func (c *VirtualNetworkTestData) FullResourceName() string {
	return "ochk_subtenant." + c.ResourceName
}

func TestAccVirtualNetworkResource_create(t *testing.T) {
	virtualNetwork := VirtualNetworkTestData{
		ResourceName: "default",
		DisplayName:  "my-display-name",
	}

	configInitial := virtualNetwork.ToString()

	resourceName := virtualNetwork.FullResourceName()
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: configInitial,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", virtualNetwork.DisplayName),
					resource.TestCheckResourceAttr(resourceName, "ipam_enabled", strconv.FormatBool(virtualNetwork.IpamEnabled)),
				),
			},
		},
		CheckDestroy: testAccVirtualNetworkResourceNotExists(virtualNetwork.DisplayName),
	})
}

func testAccVirtualNetworkResourceNotExists(displayName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		proxy := testAccProvider.Meta().(*sdk.Client).VirtualNetworks

		virtualNetworks, err := proxy.ListByDisplayName(context.Background(), displayName)
		if err != nil {
			return err
		}

		if len(virtualNetworks) > 0 {
			return fmt.Errorf("virtual network %s still exists", virtualNetworks[0].VirtualNetworkID)
		}

		return nil
	}
}
