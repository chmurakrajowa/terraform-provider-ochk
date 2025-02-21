package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
)

func flattenBackupPlans(in []*models.BackupPlan) []map[strfmt.UUID]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[strfmt.UUID]interface{}

	for _, v := range in {
		m := make(map[strfmt.UUID]interface{})
		m["backup_plan_id"] = v.BackupPlanID
		m["display_name"] = v.BackupPlanName
		out = append(out, m)
	}
	return out
}
