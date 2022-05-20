package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccBackupListDataSource_read(t *testing.T) {
	resourceName := "data.ochk_backup_list.example"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccBackupListDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", testData.BackupListName),
				),
			},
		},
	})
}

func testAccBackupListDataSourceConfig() string {
	return fmt.Sprintf(`
data "ochk_backup_plan" "default-backup-plan" {
	display_name = %[1]q
}
data "ochk_backup_list" "example" {
  backup_plan_id = data.ochk_backup_plan.default-backup-plan.id
  display_name = %[2]q
}
`, testData.BackupPlanName, testData.BackupListName)
}
