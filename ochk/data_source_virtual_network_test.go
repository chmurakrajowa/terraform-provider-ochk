package ochk

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

type VirtualNetworkDataSourceTestData struct {
	ResourceName string
	DisplayName  string
	FolderPath   string
}

func (c *VirtualNetworkDataSourceTestData) ToString(projectName string) string {
	return executeTemplateToString(`
data "ochk_virtual_network" "{{ .ResourceName}}" {
  display_name = "{{ .DisplayName}}"
}

data "ochk_project" "project`+projectName+`" {
	 display_name = "`+testData.Project1Name+`"
}
`, c)
}

func (c *VirtualNetworkDataSourceTestData) FullResourceName() string {
	return "data.ochk_virtual_network." + c.ResourceName
}

func TestAccVirtualNetworkDatasource(t *testing.T) {
	platformType := checkPlatformType()

	virtualNetwork := VirtualNetworkDataSourceTestData{
		ResourceName: "default",
		DisplayName:  testData.VirtualNetwork1DisplayName,
		FolderPath:   "/",
	}

	resourceName := virtualNetwork.FullResourceName()
	if platformType == "VMWARE" {
		resource.ParallelTest(t, resource.TestCase{
			ProviderFactories: testAccProviderFactories,
			Steps: []resource.TestStep{
				{
					Config: virtualNetwork.ToString("-vnet"),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(resourceName, "display_name", virtualNetwork.DisplayName),
						resource.TestCheckResourceAttr(resourceName, "folder_path", virtualNetwork.FolderPath),
						resource.TestCheckResourceAttrPair(resourceName, "project_id", "data.ochk_project.project-vnet", "id"),
						resource.TestCheckResourceAttrSet(resourceName, "ipam_enabled"),
						resource.TestCheckResourceAttrSet(resourceName, "gateway_address"),
						resource.TestCheckResourceAttrSet(resourceName, "subnet_mask"),
						resource.TestCheckResourceAttrSet(resourceName, "subnet_gateway_address_cidr"),
						resource.TestCheckResourceAttrSet(resourceName, "subnet_network_cidr"),
						resource.TestCheckResourceAttrSet(resourceName, "vpc_id"),
					),
				},
			},
		})
	} else {
		resource.ParallelTest(t, resource.TestCase{
			ProviderFactories: testAccProviderFactories,
			Steps: []resource.TestStep{
				{
					Config: virtualNetwork.ToString("-vnet"),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(resourceName, "display_name", virtualNetwork.DisplayName),
						resource.TestCheckResourceAttr(resourceName, "folder_path", virtualNetwork.FolderPath),
						resource.TestCheckResourceAttrPair(resourceName, "project_id", "data.ochk_project.project-vnet", "id"),
						resource.TestCheckResourceAttrSet(resourceName, "ipam_enabled"),
						resource.TestCheckResourceAttrSet(resourceName, "gateway_address"),
						resource.TestCheckResourceAttrSet(resourceName, "subnet_mask"),
						resource.TestCheckResourceAttrSet(resourceName, "subnet_gateway_address_cidr"),
						resource.TestCheckResourceAttrSet(resourceName, "subnet_network_cidr"),
						resource.TestCheckResourceAttrSet(resourceName, "vpc_id"),
						resource.TestCheckResourceAttrSet(resourceName, "dns_servers.0.id"),
						resource.TestCheckResourceAttrSet(resourceName, "dns_servers.0.address"),
					),
				},
			},
		})
	}

}
