package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccAutoNatDataSource_read(t *testing.T) {
	resourceName := "data.ochk_auto_nat.auto-nat"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAutoNatDataSourceConfig(testData.AutoNatName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", testData.AutoNatName),
					resource.TestCheckResourceAttrSet(resourceName, "description"),
					resource.TestCheckResourceAttrSet(resourceName, "vrf_id"),
					resource.TestCheckResourceAttr(resourceName, "nat_type", "AUTO"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_by"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_at"),
					resource.TestCheckResourceAttr(resourceName, "action", "SNAT"),
					resource.TestCheckResourceAttrSet(resourceName, "enabled"),
					resource.TestCheckResourceAttrSet(resourceName, "priority"),
					resource.TestCheckResourceAttrSet(resourceName, "source_network"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_network_id"),
				),
			},
		},
	})
}

func TestAccDnatDataSource_read(t *testing.T) {
	resourceName := "data.ochk_manual_nat.dnat"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccManualNatDataSourceConfig(testData.DnatName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", testData.DnatName),
					resource.TestCheckResourceAttrSet(resourceName, "description"),
					resource.TestCheckResourceAttrSet(resourceName, "vrf_id"),
					resource.TestCheckResourceAttr(resourceName, "nat_type", "MANUAL"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_by"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_at"),
					resource.TestCheckResourceAttr(resourceName, "action", "DNAT"),
					resource.TestCheckResourceAttrSet(resourceName, "enabled"),
					resource.TestCheckResourceAttrSet(resourceName, "priority"),
					resource.TestCheckResourceAttrSet(resourceName, "service_id"),
					resource.TestCheckResourceAttrSet(resourceName, "translated_ports"),
					resource.TestCheckResourceAttrSet(resourceName, "source_network"),
					resource.TestCheckResourceAttrSet(resourceName, "destination_network"),
					resource.TestCheckResourceAttrSet(resourceName, "translated_network"),
				),
			},
		},
	})
}

func testAccAutoNatDataSourceConfig(autoNatDisplayName string) string {
	return fmt.Sprintf(`
data "ochk_auto_nat" "auto-nat" {
  display_name = %[1]q
}
`, autoNatDisplayName)
}

func testAccManualNatDataSourceConfig(natDisplayName string) string {
	return fmt.Sprintf(`
data "ochk_manual_nat" "dnat" {
  display_name = %[1]q
}
`, natDisplayName)
}
