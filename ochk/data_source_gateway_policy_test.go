package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccGatewayPolicyDataSource_read(t *testing.T) {
	resourceName := "data.ochk_gateway_policy.default"
	displayName := "T1"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGatewayPolicyDataSourceConfig(displayName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", displayName),
				),
			},
		},
	})
}

func testAccGatewayPolicyDataSourceConfig(displayName string) string {
	return fmt.Sprintf(`
data "ochk_gateway_policy" "default" {
  display_name = %[1]q
}
`, displayName)
}
