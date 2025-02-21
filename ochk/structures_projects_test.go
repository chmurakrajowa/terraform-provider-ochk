package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenProjects(t *testing.T) {
	cases := []struct {
		expanded  []*models.ProjectInstance
		flattened []map[strfmt.UUID]interface{}
	}{
		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []*models.ProjectInstance{
				{
					ProjectID: "e1675817-f1a1-45c1-988b-ec2f142867e0",
					Name:      "test1",
				},
				{
					ProjectID: "791bf702-22fb-4c76-bebb-1fee7ee75607",
					Name:      "test2",
				},
			},
			flattened: []map[strfmt.UUID]interface{}{
				{
					"project_id":   strfmt.UUID("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					"display_name": "test1",
				},
				{
					"project_id":   strfmt.UUID("791bf702-22fb-4c76-bebb-1fee7ee75607"),
					"display_name": "test2",
				},
			},
		},
	}
	for _, c := range cases {
		flattenedType := mapSliceToInterfaceSlice(c.flattened)
		outFlattened := mapSliceToInterfaceSlice(flattenProjects(c.expanded))
		assert.EqualValues(t, flattenedType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, flattenedType)
	}
}
