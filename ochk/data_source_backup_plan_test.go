package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccBackupPlanDataSource_read(t *testing.T) {
	resourceName := "data.ochk_backup_plan.example"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccBackupPlanDataSourceConfig(testData.BackupPlanName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", testData.BackupPlanName),
				),
			},
		},
	})
}

func testAccBackupPlanDataSourceConfig(displayName string) string {
	return fmt.Sprintf(`
data "ochk_backup_plan" "example" {
  display_name = %[1]q
}
`, displayName)
}
