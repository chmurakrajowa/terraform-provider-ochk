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
	ResourceName             string
	DisplayName              string
	IpamEnabled              bool
	GatewayAddress           string
	RouterRefID              string
	SubnetMask               string
	SubnetGatewayAddressCidr string
	SubnetNetworkCidr        string
	ProjectID                string
}

func (c *VirtualNetworkTestData) ToString() string {
	return executeTemplateToString(`
resource "ochk_virtual_network" "{{.ResourceName}}" {
	display_name = "{{.DisplayName}}"
	ipam_enabled = "{{.IpamEnabled}}"
	vpc_id = {{ StringTFValue .RouterRefID }}
    project_id = {{ StringTFValue .ProjectID }}
	subnet_network_cidr = "{{ .SubnetNetworkCidr }}"
}
`, c)
}

func (c *VirtualNetworkTestData) FullResourceName() string {
	return "ochk_virtual_network." + c.ResourceName
}

func TestAccVirtualNetworkResource_create_minimal(t *testing.T) {

	project := ProjectDataSourceTestData{
		ResourceName: generateRandName(devTestDataPrefix),
		DisplayName:  testData.Project1Name,
	}

	virtualNetwork := VirtualNetworkTestData{
		ResourceName: "default",
		DisplayName:  generateRandName(devTestDataPrefix),
		ProjectID:    testDataResourceID(&project),
	}

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
					resource.TestCheckResourceAttr(resourceName, "vpc_id", virtualNetwork.RouterRefID),
					resource.TestCheckResourceAttrSet(resourceName, "folder_path"),
					resource.TestCheckResourceAttr(resourceName, "ipam_enabled", strconv.FormatBool(virtualNetwork.IpamEnabled)),
				),
			},
		},
		CheckDestroy: testAccVirtualNetworkResourceNotExists(virtualNetwork.DisplayName),
	})
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
		ProjectID:         testDataResourceID(&project),
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
		RouterRefID:       testDataResourceID(&router1),
		ProjectID:         testDataResourceID(&project),
	}

	configInitial := project.ToString() + router1.ToString("n-test1") + virtualNetwork.ToString()

	virtualNetworkUpdated := virtualNetwork
	virtualNetworkUpdated.RouterRefID = testDataResourceID(&router2)

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
