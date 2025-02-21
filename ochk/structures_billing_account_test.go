package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenExpandAccountProjects(t *testing.T) {
	cases := []struct {
		expanded  []*models.AccountProjectInstance
		flattened []map[strfmt.UUID]interface{}
	}{
		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []*models.AccountProjectInstance{
				{
					ProjectID: "e1675817-f1a1-45c1-988b-ec2f142867aa",
				},
			},
			flattened: []map[strfmt.UUID]interface{}{
				{
					"project_id": strfmt.UUID("e1675817-f1a1-45c1-988b-ec2f142867aa"),
				},
			},
		},
	}
	for _, c := range cases {
		flattenedInterfaceSlice := mapSliceToInterfaceSlice(c.flattened)
		outExpanded := expandAcctProjects(flattenedInterfaceSlice)
		assert.EqualValues(t, c.expanded, outExpanded, "Error matching output and expanded: %#v vs %#v", outExpanded, c.expanded)
	}
}
