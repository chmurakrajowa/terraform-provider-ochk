package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

type FloatingIpTestData struct {
	ResourceName string
	Description  string
	DisplayName  string
	VmPortId     string
	VmFixedIp    string
	VmName       string
}

//func (c *FloatingIpTestData) ToString() string {
//
//	return executeTemplateToString(`resource "ochk_floating_ip_address" "floating_ip_test2"  {
//		display_name = "{{.DisplayName}}"
//		description = "{{.Description}}"
//		vm_port_id = "{{.VmPortId}}"
//	}`, c)
//}

func (c *FloatingIpTestData) FullResourceName() string {
	return "ochk_floating_ip_address." + c.ResourceName
}

func TestUnitFloating_Ip_Resource_create(t *testing.T) {

	floating_ip := FloatingIpTestData{
		ResourceName: "floating_ip_test2",
		Description:  "tf-test-description",
		DisplayName:  "FLOATING_IP_TEST",
		VmPortId:     "d3c6bda6-172d-408b-a3ce-4f025f467e31",
		VmFixedIp:    "192.168.2.51",
		VmName:       "Redhat8",
	}

	//configInitial := floating_ip.ToString()

	fmt.Printf("Floating Ip address full name: %v\n", floating_ip.DisplayName)
	fmt.Printf("Floating Ip virtual machine port: %v\n", floating_ip.VmPortId)

	//fmt.Printf("ConfigInitial: %v\n", configInitial)

	floating_ip_updated := floating_ip

	floating_ip_updated.Description += " - updated"
	floating_ip_updated.DisplayName += "-updated"
	//floating_ip_updated.VmName = "Redhat8"
	//floating_ip_updated.VmFixedIp = "192.168.2.51"
	//configUpdated := floating_ip_updated.ToString()
	//fmt.Printf("####### configUpdated >>>>> %s\n", configUpdated)

	projectResourceName := floating_ip.FullResourceName()

	//fmt.Printf("####### configInitial >>>>> %s\n", configInitial)

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFloatingIpAddressResourceConfig(floating_ip.DisplayName, floating_ip.Description, floating_ip.VmPortId, floating_ip.VmFixedIp, floating_ip.VmName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(projectResourceName, "display_name", floating_ip.DisplayName),
					resource.TestCheckResourceAttr(projectResourceName, "description", floating_ip.Description),
					resource.TestCheckResourceAttr(projectResourceName, "vm_fixed_ip", floating_ip.VmFixedIp),
				),
			},
			{
				ResourceName: projectResourceName,
				ImportState:  true,
				//ImportStateVerify: true,  // to run after create endpoint to get vm_fixed_ip
			},
			{
				Config: testAccFloatingIpAddressResourceConfig(floating_ip_updated.DisplayName, floating_ip_updated.Description, floating_ip_updated.VmPortId, floating_ip.VmFixedIp, floating_ip.VmName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(projectResourceName, "display_name", floating_ip_updated.DisplayName),
					resource.TestCheckResourceAttr(projectResourceName, "description", floating_ip_updated.Description),
					resource.TestCheckResourceAttr(projectResourceName, "vm_fixed_ip", floating_ip_updated.VmFixedIp),
				),
			},
		},
		CheckDestroy: testAccFloatingIpAddressResourceNotExists(floating_ip.DisplayName),
	})
}

func testAccFloatingIpAddressResourceNotExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		proxy := testAccProvider.Meta().(*sdk.Client).FloatingIPAddresses
		floatingIp, err := proxy.ListByName(context.Background(), name)
		if err != nil {
			return err
		}

		if len(floatingIp) > 0 {
			return fmt.Errorf("project %s still exists", floatingIp[0].FloatingIPID)
		}
		return nil
	}
}

func testAccFloatingIpAddressResourceConfig(displayName string, description string, vm_port_id string, vm_fixed_ip string, vm_name string) string {
	return fmt.Sprintf(`
resource "ochk_floating_ip_address" "floating_ip_test2" {
  display_name = %[1]q
  description = %[2]q
  vm_port_id = %[3]q
  vm_fixed_ip = %[4]q
  vm_name = %[5]q
}
`, displayName, description, vm_port_id, vm_fixed_ip, vm_name)
}
