package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

type BillingTagTestData struct {
	ResourceName string
	DisplayName  string
}

func (c *BillingTagTestData) ToString() string {
	return executeTemplateToString(`
resource "ochk_billing_tag" "{{ .ResourceName}}" {
  display_name = "{{ .DisplayName}}"
}
`, c)
}

func (c *BillingTagTestData) FullResourceName() string {
	return "ochk_billing_tag." + c.ResourceName
}

func TestAccBillingTagResource_create_update(t *testing.T) {
	billingTag := BillingTagTestData{
		ResourceName: "default",
		DisplayName:  generateShortRandName(devTestDataPrefix),
	}

	billingTagUpdated := billingTag
	billingTagUpdated.DisplayName += "-upd"

	BillingTagResourceName := billingTag.FullResourceName()
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: billingTag.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(BillingTagResourceName, "display_name", billingTag.DisplayName),
				),
			},
			{
				ResourceName:      BillingTagResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: billingTagUpdated.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(BillingTagResourceName, "display_name", billingTagUpdated.DisplayName),
				),
			},
		},
		CheckDestroy: testAccBillingTagResourceNotExists(billingTagUpdated.DisplayName),
	})
}

func testAccBillingTagResourceNotExists(displayName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		proxy := testAccProvider.Meta().(*sdk.Client).BillingTags

		billingTags, err := proxy.ListBillingTagsByTagName(context.Background(), displayName)
		if err != nil {
			return err
		}

		if len(billingTags) > 0 {
			return fmt.Errorf("billing_tag %d still exists", billingTags[0].BillingTagID)
		}

		return nil
	}
}
