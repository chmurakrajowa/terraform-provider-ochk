package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenTags(t *testing.T) {
	cases := []struct {
		expanded  []*models.Tag
		flattened []map[strfmt.UUID]interface{}
	}{
		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []*models.Tag{
				{
					TagID:     10,
					TagValue:  "Tag1",
					ProjectID: strfmt.UUID("e1675817-f1a1-45c1-988b-ec2f142867e0"),
				},
				{
					TagID:     20,
					TagValue:  "Tag2",
					ProjectID: strfmt.UUID("e1675817-f1a1-45c1-988b-ec2f142867aa"),
				},
			},
			flattened: []map[strfmt.UUID]interface{}{
				{
					"tag_id":       "10",
					"display_name": "Tag1",
					"project_id":   strfmt.UUID("e1675817-f1a1-45c1-988b-ec2f142867e0"),
				},
				{
					"tag_id":       "20",
					"display_name": "Tag2",
					"project_id":   strfmt.UUID("e1675817-f1a1-45c1-988b-ec2f142867aa"),
				},
			},
		},
	}
	for _, c := range cases {
		flattenedType := mapSliceToInterfaceSlice(c.flattened)
		outFlattened := mapSliceToInterfaceSlice(flattenTags(c.expanded))
		assert.EqualValues(t, flattenedType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, flattenedType)
	}
}
