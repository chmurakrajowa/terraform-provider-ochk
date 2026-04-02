package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

type PublicIpTestData struct {
	ResourceName      string
	Description       string
	DisplayName       string
	PublicAddressIp   string
	PublicAddressIpId string
}

func (c *PublicIpTestData) FullResourceName() string {
	return "ochk_public_ip_address." + c.ResourceName
}

func TestUnitPublic_Ip_Resource_create(t *testing.T) {

	public_ip := PublicIpTestData{
		ResourceName:      "public_ip_test",
		Description:       "tf-test-description",
		DisplayName:       "PUBLIC_IP_TEST",
		PublicAddressIp:   "203.0.113.93",
		PublicAddressIpId: "f09ba8de-35a6-46e5-8173-18740589a5c3",
	}

	fmt.Printf("Public Ip address full name: %v\n", public_ip.DisplayName)

	//fmt.Printf("ConfigInitial: %v\n", configInitial)

	public_ip_updated := public_ip

	public_ip_updated.Description += " - updated"
	public_ip_updated.DisplayName += "-updated"
	//fmt.Printf("####### configUpdated >>>>> %s\n", configUpdated)

	ResourceName := public_ip.ResourceName

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPublicIpAddressResourceConfig(public_ip.DisplayName, public_ip.Description, public_ip.PublicAddressIp, public_ip.PublicAddressIpId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(ResourceName, "display_name", public_ip.DisplayName),
					resource.TestCheckResourceAttr(ResourceName, "description", public_ip.Description),
					resource.TestCheckResourceAttr(ResourceName, "public_address_ip", public_ip.PublicAddressIp),
					resource.TestCheckResourceAttr(ResourceName, "public_address_ip_id", public_ip.PublicAddressIpId),
				),
			},
			{
				ResourceName: ResourceName,
				ImportState:  true,
				//ImportStateVerify: true,  // to run after create endpoint to get vm_fixed_ip
			},
			{
				Config: testAccPublicIpAddressResourceConfig(public_ip_updated.DisplayName, public_ip_updated.Description, public_ip_updated.PublicAddressIp, public_ip.PublicAddressIpId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(ResourceName, "display_name", public_ip_updated.DisplayName),
					resource.TestCheckResourceAttr(ResourceName, "description", public_ip_updated.Description),
					resource.TestCheckResourceAttr(ResourceName, "public_address_ip", public_ip.PublicAddressIp),
					resource.TestCheckResourceAttr(ResourceName, "public_address_ip_id", public_ip.PublicAddressIpId),
				),
			},
		},
		CheckDestroy: testAccPublicIpAddressResourceNotExists(public_ip.DisplayName),
	})
}

func testAccPublicIpAddressResourceNotExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		proxy := testAccProvider.Meta().(*sdk.Client).PublicIPAddresses
		publicIp, err := proxy.ListByName(context.Background(), name)
		if err != nil {
			return err
		}

		if len(publicIp) > 0 {
			return fmt.Errorf("public ip address %s still exists", publicIp[0].PublicIpAddress.IpAddressId)
		}
		return nil
	}
}

func testAccPublicIpAddressResourceConfig(displayName string, description string, publicAddressIp string, publicAddressIpId string) string {
	return fmt.Sprintf(`
resource "ochk_public_ip_address" "public_ip_test" {
  display_name = %[1]q
  description = %[2]q
  public_address_ip = %[3]q
  public_address_ip_id = %[4]q
}
`, displayName, description, publicAddressIp, publicAddressIpId)
}
