package ochk

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

type BackupPlanDataSourceTestData struct {
	ResourceName string
	DisplayName  string
}

func (c *BackupPlanDataSourceTestData) ToString() string {
	return executeTemplateToString(`
data "ochk_backup_plan" "{{ .ResourceName}}" {
  display_name = "{{ .DisplayName}}"
}
`, c)
}

func (c *BackupPlanDataSourceTestData) FullResourceName() string {
	return "data.ochk_backup_plan." + c.ResourceName
}

func TestBackupPlanDataSource_read(t *testing.T) {
	backupPlan := BackupPlanDataSourceTestData{
		ResourceName: "default",
		DisplayName:  testData.BackupPlanName,
	}

	resourceName := backupPlan.FullResourceName()

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: backupPlan.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", backupPlan.DisplayName),
				),
			},
		},
	})
}
