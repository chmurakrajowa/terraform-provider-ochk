package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccBackupPlansDataSource_read(t *testing.T) {
	resourceName := "data.ochk_backup_plans.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccBackupPlansDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "backup_plans.0.backup_plan_id"),
					resource.TestCheckResourceAttrSet(resourceName, "backup_plans.0.display_name"),
				),
			},
		},
	})
}

func testAccBackupPlansDataSourceConfig() string {
	return fmt.Sprintf(`
data "ochk_backup_plans" "default" {
}
`)
}
