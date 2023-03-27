package ochk

import "github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"

func flattenBackupPlans(in []*models.BackupPlan) []map[string]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[string]interface{}

	for _, v := range in {
		m := make(map[string]interface{})
		m["backup_plan_id"] = v.BackupPlanID
		m["display_name"] = v.BackupPlanName
		out = append(out, m)
	}
	return out
}
