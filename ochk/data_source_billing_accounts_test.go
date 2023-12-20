package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAcctsDataSource_read(t *testing.T) {
	resourceName := "data.ochk_accounts.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "accounts.0.account_id"),
					resource.TestCheckResourceAttrSet(resourceName, "accounts.0.display_name"),
				),
			},
		},
	})
}

func testAccsDataSourceConfig() string {
	return fmt.Sprintf(`
data "ochk_billing_accounts" "default" {
}
`)
}
