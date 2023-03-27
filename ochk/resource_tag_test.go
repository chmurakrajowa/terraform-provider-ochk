package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

type TagTestData struct {
	ResourceName string
	DisplayName  string
}

func (c *TagTestData) ToString() string {
	return executeTemplateToString(`

data "ochk_project" "project-tag-1" {
  display_name = "`+testData.Project1Name+`"
}

resource "ochk_tag" "{{ .ResourceName}}" {
  display_name = "{{ .DisplayName}}"
  project_id = data.ochk_project.project-tag-1.id
}
`, c)
}

func (c *TagTestData) FullResourceName() string {
	return "ochk_tag." + c.ResourceName
}

func TestAccTagResource_create_update(t *testing.T) {
	tag := TagTestData{
		ResourceName: "default",
		DisplayName:  generateShortRandName(devTestDataPrefix),
	}

	tagUpdated := tag
	tagUpdated.DisplayName += "-upd"

	TagResourceName := tag.FullResourceName()
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: tag.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(TagResourceName, "display_name", tag.DisplayName),
					resource.TestCheckResourceAttrPair(TagResourceName, "project_id", "data.ochk_project.project-tag-1", "id"),
				),
			},
			{
				ResourceName:      TagResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: tagUpdated.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(TagResourceName, "display_name", tagUpdated.DisplayName),
				),
			},
		},
		CheckDestroy: testAccTagResourceNotExists(tagUpdated.DisplayName),
	})
}

func testAccTagResourceNotExists(displayName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		proxy := testAccProvider.Meta().(*sdk.Client).Tags

		tags, err := proxy.ListTagsByTagName(context.Background(), displayName)
		if err != nil {
			return err
		}

		if len(tags) > 0 {
			return fmt.Errorf("tag %d still exists", tags[0].TagID)
		}

		return nil
	}
}
