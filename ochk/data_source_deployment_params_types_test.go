package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDeploymentParamsTypesDataSource_read(t *testing.T) {
	resourceName := "data.ochk_deployment_params_types.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeploymentParamsTypesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "deployment_params_types.0.deployment_param_type"),
				),
			},
		},
	})
}

func testAccDeploymentParamsTypesDataSourceConfig() string {
	return fmt.Sprintf(`
data "ochk_deployment_params_types" "default" {
}
`)
}
