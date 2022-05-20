package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccBillingTagDataSource_read(t *testing.T) {
	resourceName := "data.ochk_billing_tag.res-bt1"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccBillingTagDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", testData.BillingTagName),
				),
			},
		},
	})
}

func testAccBillingTagDataSourceConfig() string {
	return fmt.Sprintf(`
data "ochk_billing_tag" "res-bt1" {
	display_name = %[1]q
}
`, testData.BillingTagName)
}
