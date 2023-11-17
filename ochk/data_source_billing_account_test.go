package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSource_read(t *testing.T) {
	resourceName := "data.ochk_account.rest-act1"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAcctDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", testData.AccountName),
					resource.TestCheckResourceAttrPair(resourceName, "projects.0.id", "data.ochk_project.project-acct-1", "id"),
					resource.TestCheckResourceAttrSet(resourceName, "projects.0.display_name"),
				),
			},
		},
	})
}

func testAcctDataSourceConfig() string {
	return fmt.Sprintf(`
data "ochk_project" "project-acct-1" {
  display_name = "`+testData.Project1Name+`"
}

resource "ochk_billing_account" "rest-act1" {
  display_name = %[1]q

  projects {
    project_id = data.ochk_project.project-acct-1.id
  }
}
`, testData.AccountName)
}
