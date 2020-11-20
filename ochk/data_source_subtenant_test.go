package ochk

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

type SubtenantDataSourceTestData struct {
	ResourceName string
	Name         string
}

func (c *SubtenantDataSourceTestData) ToString() string {
	return executeTemplateToString(`
data "ochk_subtenant" "{{ .ResourceName}}" {
  name = "{{ .Name}}"
}
`, c)
}

func (c *SubtenantDataSourceTestData) FullResourceName() string {
	return "data.ochk_subtenant." + c.ResourceName
}

func TestAccSubtenantDataSource_read(t *testing.T) {
	subtenant := &SubtenantDataSourceTestData{
		ResourceName: "default",
		Name:         testData.Subtenant1Name,
	}

	user := UserDataSourceTestData{
		ResourceName: "default",
		Name:         testData.User1Name,
	}

	subtenantNetwork := NetworkDataSourceTestData{
		ResourceName: "default",
		Name:         testData.SubtenantNetworkName,
	}

	config := user.ToString() + subtenantNetwork.ToString() + subtenant.ToString()

	resourceName := subtenant.FullResourceName()
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", subtenant.Name),
					//FIXME waiting for backend fix
					//resource.TestCheckResourceAttr(resourceName, "description", subtenant.Description),
					resource.TestCheckResourceAttr(resourceName, "email", "email1@example.com"),
					resource.TestCheckResourceAttr(resourceName, "memory_reserved_size_mb", "30000"),
					resource.TestCheckResourceAttr(resourceName, "storage_reserved_size_gb", "400"),
					resource.TestCheckResourceAttrPair(resourceName, "users.0", user.FullResourceName(), "id"),
					resource.TestCheckResourceAttrPair(resourceName, "networks.0", subtenantNetwork.FullResourceName(), "id"),
				),
			},
		},
	})
}
