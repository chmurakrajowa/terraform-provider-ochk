package ochk

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

type FloatingIpDataSourceTestData struct {
	ResourceName      string
	DisplayName       string
	FloatingIpAddress string
}

func (c *FloatingIpDataSourceTestData) ToString() string {
	return executeTemplateToString(`
data "ochk_floating_ip_address" "{{ .ResourceName}}" {
  display_name = "{{ .DisplayName}}"
}
`, c)
}

func (c *FloatingIpDataSourceTestData) FullResourceName() string {
	return "data.ochk_floating_ip_address." + c.ResourceName
}

func TestAccFloatingIpDataSource_read(t *testing.T) {
	resourceName := "data.ochk_floating_ip_address.example"

	floatingIpAddress := FloatingIpDataSourceTestData{
		ResourceName:      "example",
		DisplayName:       testData.FloatingIpAddressName,
		FloatingIpAddress: testData.FloatingIpAddress,
	}

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: floatingIpAddress.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", testData.FloatingIpAddressName),
					resource.TestCheckResourceAttr(resourceName, "public_address", testData.FloatingIpAddress),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				),
			},
		},
	})
}
