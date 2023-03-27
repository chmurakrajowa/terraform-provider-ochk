package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDeploymentTypesDataSource_read(t *testing.T) {
	resourceName := "data.ochk_deployment_types.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeploymentTypesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "deployment_types.0.deployment_type"),
				),
			},
		},
	})
}

func testAccDeploymentTypesDataSourceConfig() string {
	return fmt.Sprintf(`
data "ochk_deployment_types" "default" {
}
`)
}
