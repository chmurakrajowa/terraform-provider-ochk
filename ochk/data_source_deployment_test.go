package ochk

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

type DeploymentDataSourceTestData struct {
	ResourceName string
	DisplayName  string
}

func (c *DeploymentDataSourceTestData) ToString() string {
	return executeTemplateToString(`
data "ochk_deployment" "{{ .ResourceName}}" {
  display_name = "{{ .DisplayName}}"
}
`, c)
}

func (c *DeploymentDataSourceTestData) FullResourceName() string {
	return "data.ochk_deployment." + c.ResourceName
}

func TestAccDeploymentDataSource_read(t *testing.T) {
	resourceName := "data.ochk_deployment.example"

	deployment := DeploymentDataSourceTestData{
		ResourceName: "example",
		DisplayName:  testData.Deployment1DisplayName,
	}

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: deployment.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", testData.Deployment1DisplayName),
				),
			},
		},
	})
}
