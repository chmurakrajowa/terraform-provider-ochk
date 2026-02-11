package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenTags(t *testing.T) {
	TagIdValue := int32(10)
	TagIdValue2 := int32(20)
	var virtualMachines []*string

	cases := []struct {
		expanded  []openapi.Tag
		flattened []map[strfmt.UUID]interface{}
	}{
		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []openapi.Tag{
				{
					TagId:                  &TagIdValue,
					TagValue:               NewNullableString("Tag1"),
					ProjectId:              NewNullableString("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					RelatedVirtualMachines: virtualMachines,
				},
				{
					TagId:                  &TagIdValue2,
					TagValue:               NewNullableString("Tag2"),
					ProjectId:              NewNullableString("e1675817-f1a1-45c1-988b-ec2f142867aa"),
					RelatedVirtualMachines: virtualMachines,
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
