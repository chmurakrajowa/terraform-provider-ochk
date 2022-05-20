package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

type SystemTagTestData struct {
	ResourceName string
	DisplayName  string
}

func (c *SystemTagTestData) ToString() string {
	return executeTemplateToString(`
resource "ochk_system_tag" "{{ .ResourceName}}" {
  display_name = "{{ .DisplayName}}"
}
`, c)
}

func (c *SystemTagTestData) FullResourceName() string {
	return "ochk_system_tag." + c.ResourceName
}

func TestAccSystemTagResource_create_update(t *testing.T) {
	SystemTag := SystemTagTestData{
		ResourceName: "default",
		DisplayName:  generateShortRandName(devTestDataPrefix),
	}

	SystemTagUpdated := SystemTag
	SystemTagUpdated.DisplayName += "-upd"

	SystemTagResourceName := SystemTag.FullResourceName()
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: SystemTag.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(SystemTagResourceName, "display_name", SystemTag.DisplayName),
				),
			},
			{
				ResourceName:      SystemTagResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: SystemTagUpdated.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(SystemTagResourceName, "display_name", SystemTagUpdated.DisplayName),
				),
			},
		},
		CheckDestroy: testAccSystemTagResourceNotExists(SystemTagUpdated.DisplayName),
	})
}

func testAccSystemTagResourceNotExists(displayName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		proxy := testAccProvider.Meta().(*sdk.Client).SystemTags

		SystemTags, err := proxy.ListSystemTagsByTagName(context.Background(), displayName)
		if err != nil {
			return err
		}

		if len(SystemTags) > 0 {
			return fmt.Errorf("system_tag %d still exists", SystemTags[0].SystemTagID)
		}

		return nil
	}
}
