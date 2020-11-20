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
	Subtenants               []string
	DNSSearchSuffix          string
	DNSSuffix                string
	GatewayAddress           string
	PrimaryDNSAddress        string
	PrimaryWinsAddress       string
	RouterRefID              string
	SecondaryDNSAddress      string
	SecondaryWinsAddress     string
	SubnetMask               string
	SubnetGatewayAddressCidr string
	SubnetNetworkCidr        string
}

func (c *VirtualNetworkTestData) ToString() string {
	return executeTemplateToString(`
resource "ochk_virtual_network" "{{ .ResourceName}}" {
	display_name = "{{.DisplayName}}"
	ipam_enabled = "{{.IpamEnabled}}"
	subtenants = {{ StringsToTFList .Subtenants}}
	dns_search_suffix = "{{ .DNSSearchSuffix }}"
	dns_suffix = "{{ .DNSSuffix }}"
	primary_dns_address = "{{ .PrimaryDNSAddress }}"
	primary_wins_address = "{{ .PrimaryWinsAddress }}"
	router = {{ StringTFValue .RouterRefID }}
	secondary_dns_address = "{{ .SecondaryDNSAddress }}"
	secondary_wins_address = "{{ .SecondaryWinsAddress }}"
	subnet_network_cidr = "{{ .SubnetNetworkCidr }}"
}
`, c)
}

func (c *VirtualNetworkTestData) FullResourceName() string {
	return "ochk_virtual_network." + c.ResourceName
}

func TestAccVirtualNetworkResource_create_minimal(t *testing.T) {
	subtenant1 := SubtenantDataSourceTestData{
		ResourceName: "subtenant1",
		Name:         testData.Subtenant1Name,
	}

	virtualNetwork := VirtualNetworkTestData{
		ResourceName: "default",
		DisplayName:  generateRandName(),
		Subtenants:   []string{testDataResourceID(&subtenant1)},
	}

	configInitial := subtenant1.ToString() + virtualNetwork.ToString()

	resourceName := virtualNetwork.FullResourceName()
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,

		Steps: []resource.TestStep{
			{
				Config: configInitial,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", virtualNetwork.DisplayName),
					resource.TestCheckResourceAttr(resourceName, "ipam_enabled", strconv.FormatBool(virtualNetwork.IpamEnabled)),
					resource.TestCheckResourceAttrPair(resourceName, "subtenants.0", subtenant1.FullResourceName(), "id"),
				),
			},
		},
		CheckDestroy: testAccVirtualNetworkResourceNotExists(virtualNetwork.DisplayName),
	})
}

func TestAccVirtualNetworkResource_createWithIpamAndSubnet(t *testing.T) {
	subtenant1 := SubtenantDataSourceTestData{
		ResourceName: "subtenant1",
		Name:         testData.Subtenant2Name,
	}

	virtualNetwork := VirtualNetworkTestData{
		ResourceName:         "default",
		DisplayName:          generateRandName(),
		Subtenants:           []string{testDataResourceID(&subtenant1)},
		IpamEnabled:          true,
		PrimaryDNSAddress:    "192.168.1.6",
		SecondaryDNSAddress:  "192.168.1.2",
		DNSSuffix:            "test.lcl",
		DNSSearchSuffix:      "test.lcl,prod.lcl",
		PrimaryWinsAddress:   "192.168.1.3",
		SecondaryWinsAddress: "192.168.1.3",
		SubnetNetworkCidr:    "10.16.1.0/24",
	}

	configInitial := subtenant1.ToString() + virtualNetwork.ToString()

	resourceName := virtualNetwork.FullResourceName()
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: configInitial,
				Check: resource.ComposeTestCheckFunc(
					//resource.TestCheckResourceAttr(resourceName, "display_name", virtualNetwork.DisplayName),
					resource.TestCheckResourceAttr(resourceName, "primary_dns_address", virtualNetwork.PrimaryDNSAddress),
					resource.TestCheckResourceAttr(resourceName, "secondary_dns_address", virtualNetwork.SecondaryDNSAddress),
					resource.TestCheckResourceAttr(resourceName, "dns_suffix", virtualNetwork.DNSSuffix),
					resource.TestCheckResourceAttr(resourceName, "dns_search_suffix", virtualNetwork.DNSSearchSuffix),
					resource.TestCheckResourceAttr(resourceName, "primary_wins_address", virtualNetwork.PrimaryWinsAddress),
					resource.TestCheckResourceAttr(resourceName, "secondary_wins_address", virtualNetwork.SecondaryWinsAddress),
					resource.TestCheckResourceAttrPair(resourceName, "subtenants.0", subtenant1.FullResourceName(), "id"),
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
	subtenant1 := SubtenantDataSourceTestData{
		ResourceName: "subtenant1",
		Name:         testData.Subtenant3Name,
	}

	subtenant2 := SubtenantDataSourceTestData{
		ResourceName: "subtenant2",
		Name:         testData.Subtenant4Name,
	}

	router1 := RouterTestData{
		ResourceName: "router1",
		DisplayName:  generateRandName(),
	}

	router2 := RouterTestData{
		ResourceName: "router2",
		DisplayName:  generateRandName(),
	}

	virtualNetwork := VirtualNetworkTestData{
		ResourceName:         "default",
		DisplayName:          generateRandName(),
		Subtenants:           []string{testDataResourceID(&subtenant1)},
		IpamEnabled:          true,
		PrimaryDNSAddress:    "192.168.1.6",
		SecondaryDNSAddress:  "192.168.1.2",
		DNSSuffix:            "test.lcl",
		DNSSearchSuffix:      "test.lcl,prod.lcl",
		PrimaryWinsAddress:   "192.168.1.3",
		SecondaryWinsAddress: "192.168.1.3",
		SubnetNetworkCidr:    "10.16.1.0/24",
		RouterRefID:          testDataResourceID(&router1),
	}

	configInitial := subtenant1.ToString() + subtenant2.ToString() + router1.ToString() + virtualNetwork.ToString()

	virtualNetworkUpdated := virtualNetwork
	virtualNetworkUpdated.Subtenants = []string{testDataResourceID(&subtenant2)}
	virtualNetworkUpdated.RouterRefID = testDataResourceID(&router2)

	configUpdated := subtenant1.ToString() + subtenant2.ToString() + router1.ToString() + router2.ToString() + virtualNetworkUpdated.ToString()

	resourceName := virtualNetwork.FullResourceName()

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: configInitial,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", virtualNetwork.DisplayName),
					resource.TestCheckResourceAttr(resourceName, "primary_dns_address", virtualNetwork.PrimaryDNSAddress),
					resource.TestCheckResourceAttr(resourceName, "secondary_dns_address", virtualNetwork.SecondaryDNSAddress),
					resource.TestCheckResourceAttr(resourceName, "dns_suffix", virtualNetwork.DNSSuffix),
					resource.TestCheckResourceAttr(resourceName, "dns_search_suffix", virtualNetwork.DNSSearchSuffix),
					resource.TestCheckResourceAttr(resourceName, "primary_wins_address", virtualNetwork.PrimaryWinsAddress),
					resource.TestCheckResourceAttr(resourceName, "secondary_wins_address", virtualNetwork.SecondaryWinsAddress),
					resource.TestCheckResourceAttrPair(resourceName, "subtenants.0", subtenant1.FullResourceName(), "id"),
					resource.TestCheckResourceAttr(resourceName, "subnet_network_cidr", virtualNetwork.SubnetNetworkCidr),
					resource.TestCheckResourceAttrSet(resourceName, "gateway_address"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_mask"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_gateway_address_cidr"),
					resource.TestCheckResourceAttrPair(resourceName, "router", router1.FullResourceName(), "id"),
					resource.TestCheckResourceAttrPair(resourceName, "subtenants.0", subtenant1.FullResourceName(), "id"),
				),
			},
			{
				Config: configUpdated,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", virtualNetworkUpdated.DisplayName),
					resource.TestCheckResourceAttr(resourceName, "primary_dns_address", virtualNetworkUpdated.PrimaryDNSAddress),
					resource.TestCheckResourceAttr(resourceName, "secondary_dns_address", virtualNetworkUpdated.SecondaryDNSAddress),
					resource.TestCheckResourceAttr(resourceName, "dns_suffix", virtualNetworkUpdated.DNSSuffix),
					resource.TestCheckResourceAttr(resourceName, "dns_search_suffix", virtualNetworkUpdated.DNSSearchSuffix),
					resource.TestCheckResourceAttr(resourceName, "primary_wins_address", virtualNetworkUpdated.PrimaryWinsAddress),
					resource.TestCheckResourceAttr(resourceName, "secondary_wins_address", virtualNetworkUpdated.SecondaryWinsAddress),
					resource.TestCheckResourceAttr(resourceName, "subnet_network_cidr", virtualNetwork.SubnetNetworkCidr),
					resource.TestCheckResourceAttr(resourceName, "subnet_network_cidr", virtualNetwork.SubnetNetworkCidr),
					resource.TestCheckResourceAttrSet(resourceName, "gateway_address"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_mask"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_gateway_address_cidr"),
					resource.TestCheckResourceAttrPair(resourceName, "router", router2.FullResourceName(), "id"),
					resource.TestCheckResourceAttrPair(resourceName, "subtenants.0", subtenant2.FullResourceName(), "id"),
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
