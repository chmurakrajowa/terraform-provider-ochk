package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccManualNatsDataSource_read(t *testing.T) {
	resourceName := "data.ochk_manual_nats.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccManualNatsSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "manual_nats.0.manual_nat_id"),
					resource.TestCheckResourceAttrSet(resourceName, "manual_nats.0.display_name"),
					resource.TestCheckResourceAttrSet(resourceName, "manual_nats.0.action"),
					resource.TestCheckResourceAttrSet(resourceName, "manual_nats.0.enabled"),
					resource.TestCheckResourceAttrSet(resourceName, "manual_nats.0.vrf_id"),
					resource.TestCheckResourceAttrSet(resourceName, "manual_nats.0.source_network"),
					//resource.TestCheckResourceAttrSet(resourceName, "manual_nats.0.destination_network"),
				),
			},
		},
	})
}

func testAccManualNatsSourceConfig() string {
	return fmt.Sprintf(`
data "ochk_manual_nats" "default" {
}
`)
}

func TestAccAutoNatsDataSource_read(t *testing.T) {
	resourceName := "data.ochk_auto_nats.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAutoNatsSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "auto_nats.0.auto_nat_id"),
					resource.TestCheckResourceAttrSet(resourceName, "auto_nats.0.display_name"),
					resource.TestCheckResourceAttrSet(resourceName, "auto_nats.0.virtual_network_id"),
					resource.TestCheckResourceAttrSet(resourceName, "auto_nats.0.enabled"),
					resource.TestCheckResourceAttrSet(resourceName, "auto_nats.0.vrf_id"),
				),
			},
		},
	})
}

func testAccAutoNatsSourceConfig() string {
	return fmt.Sprintf(`
data "ochk_auto_nats" "default" {
}
`)
}
