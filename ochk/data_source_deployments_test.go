package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDeploymentsDataSource_read(t *testing.T) {
	resourceName := "data.ochk_deployments.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeploymentsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "deployments.0.deployment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "deployments.0.display_name"),
					resource.TestCheckResourceAttrSet(resourceName, "deployments.0.deployment_type"),
				),
			},
		},
	})
}

func testAccDeploymentsDataSourceConfig() string {
	return fmt.Sprintf(`
data "ochk_deployments" "default" {
}
`)
}
