package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

type AccountTestData struct {
	ResourceName string
	DisplayName  string
	Projects     []AccountProjectsTestData
}

type AccountProjectsTestData struct {
	ID string
}

func (c *AccountTestData) ToString() string {
	return executeTemplateToString(`


data "ochk_project" "proj-1" {
  display_name = "`+testData.Project1Name+`"
}

resource "ochk_billing_account" "{{ .ResourceName}}" {
  display_name = "{{ .DisplayName}}"
  account_description = "{{ .DisplayName}}"
  projects {
	 	 project_id = data.ochk_project.proj-1.id
  }
}
`, c)
}

func (c *AccountTestData) FullResourceName() string {
	fmt.Printf("FullResourceName >>>>  %v\n", c.ResourceName)
	return "ochk_billing_account." + c.ResourceName
}

func TestAccountResource_create(t *testing.T) {

	account := AccountTestData{
		ResourceName: "account_one_project",
		DisplayName:  generateRandName(devTestDataPrefix),
		Projects:     nil,
	}

	/* Account with one project and with updated display_name */
	accountUpdated := account
	accountUpdated.DisplayName += "-upd"
	fmt.Printf("Account full name: %v\n", account.DisplayName)
	accountResourceName := account.FullResourceName()
	fmt.Printf("Account accountResourceName  full name: %v\n", accountResourceName)

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: account.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(accountResourceName, "display_name", account.DisplayName),
					resource.TestCheckResourceAttrSet(accountResourceName, "account_description"),
					resource.TestCheckResourceAttrSet(accountResourceName, "discount"),
					resource.TestCheckResourceAttrSet(accountResourceName, "alarms"),
					resource.TestCheckResourceAttrSet(accountResourceName, "cost"),
					resource.TestCheckResourceAttr(accountResourceName, "projects.0.project_id", "3cda830c-b37f-46dc-be54-f649d31bec66"),
				),
			},
			{
				ResourceName:      accountResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: accountUpdated.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(accountResourceName, "display_name", accountUpdated.DisplayName),
					resource.TestCheckResourceAttrPair(accountResourceName, "projects.0.project_id", "data.ochk_project.proj-1", "id"),
				),
			},
		},
		CheckDestroy: testAccountResourceNotExists(accountUpdated.DisplayName),
	})
}

func testAccountResourceNotExists(displayName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		proxy := testAccProvider.Meta().(*sdk.Client).Accounts

		accounts, err := proxy.ListAccountByName(context.Background(), displayName)
		if err != nil {
			return err
		}

		if len(accounts) > 0 {
			return fmt.Errorf("account %s still exists", accounts[0].AccountID)
		}

		return nil
	}
}
