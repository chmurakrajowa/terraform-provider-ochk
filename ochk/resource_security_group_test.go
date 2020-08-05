package ochk

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk"
	"testing"
)

func TestAccSecurityGroupResource_create(t *testing.T) {
	resourceName := "ochk_security_group.one_member"
	displayName := generateRandName()

	//TODO zbyt wiele razy jest wołany POST /vidm/token HTTP/1.1, coś jest nie tak
	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityGroupResourceConfig(displayName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", displayName),
					resource.TestCheckResourceAttr(resourceName, "members.0.id", "e1e2f617-014c-4119-bac8-49fa4a93db47"),
					resource.TestCheckResourceAttr(resourceName, "members.0.type", "VIRTUAL_MACHINE"),
					resource.TestCheckResourceAttrSet(resourceName, "members.0.display_name"),
				),
			},
		},
		CheckDestroy: testAccSecurityGroupResourceNotExists(displayName),
	})
}

func testAccSecurityGroupResourceNotExists(displayName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		proxy := testAccProvider.Meta().(*sdk.Client).SecurityGroups

		securityGroups, err := proxy.ListByDisplayName(context.Background(), displayName)
		if err != nil {
			return err
		}

		if len(securityGroups) > 0 {
			return fmt.Errorf("security group %s still exists", securityGroups[0].ID)
		}

		return nil
	}
}

func testAccSecurityGroupResourceConfig(name string) string {
	// TODO config should refer to vm from datasource
	return fmt.Sprintf(`
resource "ochk_security_group" "one_member" {
  display_name = %[1]q

  members {
    id = "e1e2f617-014c-4119-bac8-49fa4a93db47"
    type = "VIRTUAL_MACHINE"
  }
}
`, name)
}
