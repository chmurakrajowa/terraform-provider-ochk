package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk"
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
}

func (c *VirtualNetworkTestData) ToString() string {
	return ochk.CastTemplateToString(`
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

	project := ProjectDataSourceTestData{
		ResourceName: ochk.GenerateRandName(ochk.DevTestDataOpenstackPrefix),
		DisplayName:  ochk.TestOpenstackData.Project1Name,
	}

	virtualNetwork := VirtualNetworkTestData{
		ResourceName:      "default",
		DisplayName:       ochk.GenerateRandName(ochk.DevTestDataOpenstackPrefix),
		ProjectID:         strfmt.UUID(ochk.CastDataResourceID(&project)),
		SubnetNetworkCidr: "10.18.1.0/24",
		IpamEnabled:       true,
	}

	configInitial := project.ToString() + virtualNetwork.ToString()

	resourceName := virtualNetwork.FullResourceName()
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: ochk.TestAccProviderFactories,

		Steps: []resource.TestStep{
			{
				Config: configInitial,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", virtualNetwork.DisplayName),
					resource.TestCheckResourceAttrPair(resourceName, "project_id", project.FullResourceName(), "id"),
					resource.TestCheckResourceAttr(resourceName, "vpc_id", virtualNetwork.RouterRefID.String()),
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
		ResourceName: ochk.GenerateRandName(ochk.DevTestDataOpenstackPrefix),
		DisplayName:  ochk.TestOpenstackData.Project1Name,
	}

	virtualNetwork := VirtualNetworkTestData{
		ResourceName:      "default",
		DisplayName:       ochk.GenerateRandName(ochk.DevTestDataOpenstackPrefix),
		IpamEnabled:       true,
		SubnetNetworkCidr: "10.16.1.0/24",
		ProjectID:         strfmt.UUID(ochk.CastDataResourceID(&project)),
	}

	configInitial := project.ToString() + virtualNetwork.ToString()

	resourceName := virtualNetwork.FullResourceName()
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: ochk.TestAccProviderFactories,
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
		ResourceName: ochk.GenerateRandName(ochk.DevTestDataOpenstackPrefix),
		DisplayName:  ochk.TestOpenstackData.Project1Name,
	}

	router1 := RouterTestData{
		ResourceName: "router1",
		DisplayName:  ochk.GenerateRandName(ochk.DevTestDataOpenstackPrefix),
		ParentRouter: ochk.TestOpenstackData.VRF,
		ProjectName:  ochk.TestOpenstackData.Project1Name,
	}

	router2 := RouterTestData{
		ResourceName: "router2",
		DisplayName:  ochk.GenerateRandName(ochk.DevTestDataOpenstackPrefix),
		ParentRouter: ochk.TestOpenstackData.VRF,
		ProjectName:  ochk.TestOpenstackData.Project1Name,
	}

	virtualNetwork := VirtualNetworkTestData{
		ResourceName:      "default",
		DisplayName:       ochk.GenerateRandName(ochk.DevTestDataOpenstackPrefix),
		IpamEnabled:       true,
		SubnetNetworkCidr: "10.16.1.0/24",
		RouterRefID:       strfmt.UUID(ochk.CastDataResourceID(&router1)),
		ProjectID:         strfmt.UUID(ochk.CastDataResourceID(&project)),
	}

	configInitial := project.ToString() + router1.ToString("n-test1") + virtualNetwork.ToString()

	virtualNetworkUpdated := virtualNetwork
	virtualNetworkUpdated.RouterRefID = strfmt.UUID(ochk.CastDataResourceID(&router2))

	configUpdated := project.ToString() + router1.ToString("n-test1") + router2.ToString("n-test2") + virtualNetworkUpdated.ToString()

	resourceName := virtualNetwork.FullResourceName()

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: ochk.TestAccProviderFactories,
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
		proxy := ochk.TestAccProviderDefinition.Meta().(*sdk.Client).VirtualNetworks

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
