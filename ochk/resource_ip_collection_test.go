package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

type IPCollectionTestData struct {
	ResourceName string
	DisplayName  string
	IPAddresses  []string
}

func (c *IPCollectionTestData) ToString() string {
	return executeTemplateToString(`
resource "ochk_ip_collection" "{{ .ResourceName}}" {
  display_name = "{{ .DisplayName}}"
  ip_addresses  = {{ StringsToTFList .IPAddresses}}
}
`, c)
}

func (c *IPCollectionTestData) FullResourceName() string {
	return "ochk_ip_collection." + c.ResourceName
}

func TestAccIPCollectionResource_create_update(t *testing.T) {
	ipCollection := IPCollectionTestData{
		ResourceName: "default",
		DisplayName:  generateRandName(),
		IPAddresses:  []string{"10.10.10.10"},
	}

	ipCollectionUpdated := ipCollection
	ipCollectionUpdated.DisplayName += "-updated"
	ipCollectionUpdated.IPAddresses = []string{"192.168.1.10"}

	IPCollectionResourceName := ipCollection.FullResourceName()
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ipCollection.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(IPCollectionResourceName, "display_name", ipCollection.DisplayName),
					resource.TestCheckResourceAttr(IPCollectionResourceName, "ip_addresses.0", ipCollection.IPAddresses[0]),
					resource.TestCheckResourceAttrSet(IPCollectionResourceName, "created_by"),
					resource.TestCheckResourceAttrSet(IPCollectionResourceName, "created_at"),
					resource.TestCheckResourceAttrSet(IPCollectionResourceName, "modified_by"),
					resource.TestCheckResourceAttrSet(IPCollectionResourceName, "modified_at"),
				),
			},
			{
				Config: ipCollectionUpdated.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(IPCollectionResourceName, "display_name", ipCollectionUpdated.DisplayName),
					resource.TestCheckResourceAttr(IPCollectionResourceName, "ip_addresses.0", ipCollectionUpdated.IPAddresses[0]),
					resource.TestCheckResourceAttrSet(IPCollectionResourceName, "created_by"),
					resource.TestCheckResourceAttrSet(IPCollectionResourceName, "created_at"),
					resource.TestCheckResourceAttrSet(IPCollectionResourceName, "modified_by"),
					resource.TestCheckResourceAttrSet(IPCollectionResourceName, "modified_at"),
				),
			},
		},
		CheckDestroy: testAccIPCollectionResourceNotExists(ipCollectionUpdated.DisplayName),
	})
}

func testAccIPCollectionResourceNotExists(displayName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		proxy := testAccProvider.Meta().(*sdk.Client).IPCollections

		ipCollections, err := proxy.ListByDisplayName(context.Background(), displayName)
		if err != nil {
			return err
		}

		if len(ipCollections) > 0 {
			return fmt.Errorf("router %s still exists", ipCollections[0].ID)
		}

		return nil
	}
}
