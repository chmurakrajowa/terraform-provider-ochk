package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strconv"
	"testing"
)

type VirtualNetworkTestData struct {
	ResourceName             string
	DisplayName              string
	IpamEnabled              bool
	GatewayAddress           string
	RouterRefID              strfmt.UUID
	SubnetMask               string
	SubnetGatewayAddressCidr string
	SubnetNetworkCidr        string
	ProjectID                strfmt.UUID
	DNSServer                []DNSServerTestData
}

type DNSServerTestData struct {
	ID      strfmt.UUID
	Address string
}

func (c *VirtualNetworkTestData) ToString() string {
	return executeTemplateToString(`
resource "ochk_virtual_network" "{{.ResourceName}}" {
	display_name = "{{.DisplayName}}"
	ipam_enabled = "{{.IpamEnabled}}"
	vpc_id = {{ UuidTFValue .RouterRefID }}
    project_id = {{ UuidTFValue .ProjectID }}
	subnet_network_cidr = "{{ .SubnetNetworkCidr }}"
   {{range $dns_server := .DNSServer}}
   dns_servers {
     address = "{{ $dns_server.Address }}"
   }
   {{end}}
}
`, c)
}

func (c *VirtualNetworkTestData) ToStringVMware() string {
	return executeTemplateToString(`
resource "ochk_virtual_network" "{{.ResourceName}}" {
	display_name = "{{.DisplayName}}"
	ipam_enabled = "{{.IpamEnabled}}"
	vpc_id = {{ UuidTFValue .RouterRefID }}
    project_id = {{ UuidTFValue .ProjectID }}
	subnet_network_cidr = "{{ .SubnetNetworkCidr }}"
}
`, c)
}

func (c *VirtualNetworkTestData) FullResourceName() string {
	return "ochk_virtual_network." + c.ResourceName
}

func TestAccVirtualNetworkResource_create_minimal(t *testing.T) {
	platformType := checkPlatformType()

	project := ProjectDataSourceTestData{
		ResourceName: generateRandName(devTestDataPrefix),
		DisplayName:  testData.Project1Name,
	}

	virtualNetwork := VirtualNetworkTestData{
		ResourceName: "default",
		DisplayName:  generateRandName(devTestDataPrefix),
		ProjectID:    strfmt.UUID(testDataResourceID(&project)),
		DNSServer: []DNSServerTestData{
			{
				Address: "8.8.8.8",
			},
		},
	}

	if platformType == "OPENSTACK" {
		configInitial := project.ToString() + virtualNetwork.ToString()

		resourceName := virtualNetwork.FullResourceName()
		resource.ParallelTest(t, resource.TestCase{
			ProviderFactories: testAccProviderFactories,

			Steps: []resource.TestStep{
				{
					Config: configInitial,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(resourceName, "display_name", virtualNetwork.DisplayName),
						resource.TestCheckResourceAttrPair(resourceName, "project_id", project.FullResourceName(), "id"),
						resource.TestCheckResourceAttr(resourceName, "vpc_id", virtualNetwork.RouterRefID.String()),
						resource.TestCheckResourceAttrSet(resourceName, "folder_path"),
						resource.TestCheckResourceAttr(resourceName, "ipam_enabled", strconv.FormatBool(virtualNetwork.IpamEnabled)),
						resource.TestCheckResourceAttr(resourceName, "members.0.type", virtualNetwork.DNSServer[0].Address),
					),
				},
			},
			CheckDestroy: testAccVirtualNetworkResourceNotExists(virtualNetwork.DisplayName),
		})
	} else {
		configInitial := project.ToString() + virtualNetwork.ToStringVMware()

		resourceName := virtualNetwork.FullResourceName()
		resource.ParallelTest(t, resource.TestCase{
			ProviderFactories: testAccProviderFactories,

			Steps: []resource.TestStep{
				{
					Config: configInitial,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(resourceName, "display_name", virtualNetwork.DisplayName),
						resource.TestCheckResourceAttrPair(resourceName, "project_id", project.FullResourceName(), "id"),
						resource.TestCheckResourceAttr(resourceName, "vpc_id", virtualNetwork.RouterRefID.String()),
						resource.TestCheckResourceAttrSet(resourceName, "folder_path"),
						resource.TestCheckResourceAttr(resourceName, "ipam_enabled", strconv.FormatBool(virtualNetwork.IpamEnabled)),
						resource.TestCheckResourceAttr(resourceName, "members.0.type", virtualNetwork.DNSServer[0].Address),
					),
				},
			},
			CheckDestroy: testAccVirtualNetworkResourceNotExists(virtualNetwork.DisplayName),
		})
	}

}

func TestAccVirtualNetworkResource_createWithIpamAndSubnet(t *testing.T) {

	project := ProjectDataSourceTestData{
		ResourceName: generateRandName(devTestDataPrefix),
		DisplayName:  testData.Project1Name,
	}

	virtualNetwork := VirtualNetworkTestData{
		ResourceName:      "default",
		DisplayName:       generateRandName(devTestDataPrefix),
		IpamEnabled:       true,
		SubnetNetworkCidr: "10.16.1.0/24",
		ProjectID:         strfmt.UUID(testDataResourceID(&project)),
	}

	configInitial := project.ToString() + virtualNetwork.ToString()

	resourceName := virtualNetwork.FullResourceName()
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: configInitial,
				Check: resource.ComposeTestCheckFunc(
					//resource.TestCheckResourceAttr(resourceName, "display_name", virtualNetwork.DisplayName),
					resource.TestCheckResourceAttr(resourceName, "subnet_network_cidr", virtualNetwork.SubnetNetworkCidr),
					resource.TestCheckResourceAttrSet(resourceName, "gateway_address"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_mask"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_gateway_address_cidr"),
				),
			},
		},
		CheckDestroy: testAccVirtualNetworkResourceNotExists(virtualNetwork.DisplayName),
	})
}

func TestAccVirtualNetworkResource_createAndUpdateWithIpamSubnetRouter(t *testing.T) {

	project := ProjectDataSourceTestData{
		ResourceName: generateRandName(devTestDataPrefix),
		DisplayName:  testData.Project1Name,
	}

	router1 := RouterTestData{
		ResourceName: "router1",
		DisplayName:  generateRandName(devTestDataPrefix),
		ParentRouter: testData.VRF,
		ProjectName:  testData.Project1Name,
	}

	router2 := RouterTestData{
		ResourceName: "router2",
		DisplayName:  generateRandName(devTestDataPrefix),
		ParentRouter: testData.VRF,
		ProjectName:  testData.Project1Name,
	}

	virtualNetwork := VirtualNetworkTestData{
		ResourceName:      "default",
		DisplayName:       generateRandName(devTestDataPrefix),
		IpamEnabled:       true,
		SubnetNetworkCidr: "10.16.1.0/24",
		RouterRefID:       strfmt.UUID(testDataResourceID(&router1)),
		ProjectID:         strfmt.UUID(testDataResourceID(&project)),
	}

	configInitial := project.ToString() + router1.ToString("n-test1") + virtualNetwork.ToString()

	virtualNetworkUpdated := virtualNetwork
	virtualNetworkUpdated.RouterRefID = strfmt.UUID(testDataResourceID(&router2))

	configUpdated := project.ToString() + router1.ToString("n-test1") + router2.ToString("n-test2") + virtualNetworkUpdated.ToString()

	resourceName := virtualNetwork.FullResourceName()

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: configInitial,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", virtualNetwork.DisplayName),
					resource.TestCheckResourceAttrPair(resourceName, "project_id", "data.ochk_project.project-n-test1", "id"),
					resource.TestCheckResourceAttr(resourceName, "folder_path", "/"),
					resource.TestCheckResourceAttrSet(resourceName, "gateway_address"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_mask"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_gateway_address_cidr"),
					resource.TestCheckResourceAttrPair(resourceName, "vpc_id", router1.FullResourceName(), "id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: configUpdated,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", virtualNetworkUpdated.DisplayName),
					resource.TestCheckResourceAttrPair(resourceName, "project_id", "data.ochk_project.project-n-test1", "id"),
					resource.TestCheckResourceAttr(resourceName, "subnet_network_cidr", virtualNetwork.SubnetNetworkCidr),
					resource.TestCheckResourceAttr(resourceName, "subnet_network_cidr", virtualNetwork.SubnetNetworkCidr),
					resource.TestCheckResourceAttrSet(resourceName, "gateway_address"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_mask"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_gateway_address_cidr"),
					resource.TestCheckResourceAttrPair(resourceName, "vpc_id", router2.FullResourceName(), "id"),
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
