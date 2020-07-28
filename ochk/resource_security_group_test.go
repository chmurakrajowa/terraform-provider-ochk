package ochk

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client/security_groups"
	"log"
	"testing"
)

func TestAccSecurityGroupResource_create(t *testing.T) {
	name := fmt.Sprintf("tf-test-%s", acctest.RandStringFromCharSet(8, acctest.CharSetAlphaNum))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccSecurityGroupResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityGroupResourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccSecurityGroupResourceExists("ochk_security_group.test"),
					//todo check display_name
				),
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

		client := testAccProvider.Meta().(*sdk.Client)
		exists, err := checkSecurityGroupExists(client, rs.Primary.ID)
		if err != nil {
			return err
		}

		if !exists {
			return fmt.Errorf("security group %s does not exist", rs.Primary.ID)
		}

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
			client := testAccProvider.Meta().(*sdk.Client)
			exists, err := checkSecurityGroupExists(client, rs.Primary.ID)
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

func checkSecurityGroupExists(client *sdk.Client, id string) (bool, error) {
	_, err := client.GetOchk().SecurityGroups.SecurityGroupGetUsingGET(&security_groups.SecurityGroupGetUsingGETParams{
		GroupID:    id,
		Context:    context.Background(),
		HTTPClient: client.GetHTTPClient(),
	})

	if err != nil {
		var notFoundErr *security_groups.SecurityGroupGetUsingGETNotFound
		if ok := errors.As(err, &notFoundErr); ok {
			return false, nil
		}

		return false, fmt.Errorf("error getting security group: %w", err)
	}

	return true, nil
}
