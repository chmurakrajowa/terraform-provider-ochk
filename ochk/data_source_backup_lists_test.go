package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccBackupListsDataSource_read(t *testing.T) {
	resourceName := "data.ochk_backup_lists.default"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccBackupListsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "backup_lists.0.backup_list_id"),
					resource.TestCheckResourceAttrSet(resourceName, "backup_lists.0.display_name"),
				),
			},
		},
	})
}

func testAccBackupListsDataSourceConfig() string {
	return fmt.Sprintf(`
data "ochk_backup_plan" "backup_plan" {
  display_name = "` + testData.BackupPlanName + `"
}

data "ochk_backup_lists" "default" {
  backup_plan_id = data.ochk_backup_plan.backup_plan.id
}
`)
}
