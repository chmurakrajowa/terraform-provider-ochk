package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSource_read(t *testing.T) {
	resourceName := "data.ochk_billing_account.rest-act1"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAcctDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", testData.AccountName),
					resource.TestCheckResourceAttrSet(resourceName, "account_description"),
				),
			},
		},
	})
}

func testAcctDataSourceConfig() string {
	return fmt.Sprintf(`

data "ochk_billing_account" "rest-act1" {
  display_name = %[1]q
}
`, testData.AccountName)
}
