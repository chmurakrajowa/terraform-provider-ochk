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
					TagID:                  10,
					TagValue:               "Tag1",
					ProjectID:              strfmt.UUID("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					RelatedVirtualMachines: []strfmt.UUID(nil),
				},
				{
					TagID:                  20,
					TagValue:               "Tag2",
					ProjectID:              strfmt.UUID("e1675817-f1a1-45c1-988b-ec2f142867aa"),
					RelatedVirtualMachines: []strfmt.UUID(nil),
				},
			},
			flattened: []map[strfmt.UUID]interface{}{
				{
					"tag_id":                   "10",
					"display_name":             "Tag1",
					"project_id":               strfmt.UUID("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					"related_virtual_machines": []strfmt.UUID(nil),
				},
				{
					"tag_id":                   "20",
					"display_name":             "Tag2",
					"project_id":               strfmt.UUID("e1675817-f1a1-45c1-988b-ec2f142867aa"),
					"related_virtual_machines": []strfmt.UUID(nil),
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
