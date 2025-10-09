package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/openapi/v3"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenProjects(t *testing.T) {
	cases := []struct {
		expanded  []openapi.ProjectInstance
		flattened []map[strfmt.UUID]interface{}
	}{
		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []openapi.ProjectInstance{
				{
					ProjectId: NewNullableString("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					Name:      NewNullableString("test1"),
				},
				{
					ProjectId: NewNullableString("791bf702-22fb-4c76-bebb-1fee7ee75607"),
					Name:      NewNullableString("test2"),
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
