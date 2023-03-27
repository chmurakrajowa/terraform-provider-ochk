package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenBackupPlans(t *testing.T) {
	cases := []struct {
		expanded  []*models.BackupPlan
		flattened []map[string]interface{}
	}{
		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []*models.BackupPlan{
				{
					BackupPlanID:   "e1675817-f1a1-45c1-988b-ec2f142867e0",
					BackupPlanName: "Platinium 1",
				},
				{
					BackupPlanID:   "791bf702-22fb-4c76-bebb-1fee7ee75607",
					BackupPlanName: "Platinium 2",
				},
			},
			flattened: []map[string]interface{}{
				{
					"backup_plan_id": "e1675817-f1a1-45c1-988b-ec2f142867e0",
					"display_name":   "Platinium 1",
				},
				{
					"backup_plan_id": "791bf702-22fb-4c76-bebb-1fee7ee75607",
					"display_name":   "Platinium 2",
				},
			},
		},
	}
	for _, c := range cases {
		flattenedType := mapSliceToInterfaceSlice(c.flattened)
		outFlattened := mapSliceToInterfaceSlice(flattenBackupPlans(c.expanded))
		assert.EqualValues(t, flattenedType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, flattenedType)
	}
}
