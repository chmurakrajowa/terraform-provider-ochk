package ochk

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk"
	"log"
	"testing"
)

func TestAccSecurityGroupResource_create(t *testing.T) {
	name := fmt.Sprintf("tf-test-%s", acctest.RandStringFromCharSet(8, acctest.CharSetAlphaNum))

	//TODO zbyt wiele razy jest wołany POST /vidm/token HTTP/1.1, coś jest nie tak
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccSecurityGroupResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityGroupResourceConfig(name),
				Check:  testAccSecurityGroupResourceExists("ochk_security_group.test"),
			},
		},
	})
}

func testAccSecurityGroupResourceExists(resourceID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceID]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceID)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("resource %s has no ID set", resourceID)
		}

		proxy := testAccProvider.Meta().(*sdk.Client).SecurityGroups

		securityGroup, err := proxy.Read(context.Background(), rs.Primary.ID)
		if err != nil {
			return err
		}

		displayName := rs.Primary.Attributes["display_name"]

		if securityGroup.DisplayName != displayName {
			return fmt.Errorf("security group display name does not match: actual %s, expected %s", securityGroup.DisplayName, displayName)
		}

		//TODO jak porównywać members?

		return nil
	}
}

func testAccSecurityGroupResourceConfig(name string) string {
	// TODO config should refer to vm from datasource
	return fmt.Sprintf(`
resource "ochk_security_group" "test" {
  display_name = %[1]q

  members {
    id = "e1e2f617-014c-4119-bac8-49fa4a93db47"
    type = "VIRTUAL_MACHINE"
  }
}
`, name)
}

func testAccSecurityGroupResourceDestroy() resource.TestCheckFunc {
	return func(state *terraform.State) error {
		for _, rs := range state.RootModule().Resources {
			if rs.Type != "ochk_security_group" {
				continue
			}

			log.Printf("checking security group %s exists", rs.Primary.ID)
			proxy := testAccProvider.Meta().(*sdk.Client).SecurityGroups

			exists, err := proxy.Exists(context.Background(), rs.Primary.ID)
			if err != nil {
				return err
			}

			if exists {
				return fmt.Errorf("security group %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}
}
