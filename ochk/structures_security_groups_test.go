package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenSecurityGroups(t *testing.T) {
	cases := []struct {
		expanded  []openapi.SecurityGroup
		flattened []map[strfmt.UUID]interface{}
	}{
		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []openapi.SecurityGroup{
				{
					Id:          NewNullableString("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					ProjectId:   NewNullableString("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					DisplayName: NewNullableString("test1"),
				},
				{
					Id:          NewNullableString("791bf702-22fb-4c76-bebb-1fee7ee75607"),
					ProjectId:   NewNullableString("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					DisplayName: NewNullableString("test2"),
				},
			},
			flattened: []map[strfmt.UUID]interface{}{
				{
					"security_group_id": strfmt.UUID("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					"project_id":        strfmt.UUID("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					"display_name":      "test1",
				},
				{
					"security_group_id": strfmt.UUID("791bf702-22fb-4c76-bebb-1fee7ee75607"),
					"project_id":        strfmt.UUID("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					"display_name":      "test2",
				},
			},
		},
	}
	for _, c := range cases {
		flattenedType := mapSliceToInterfaceSlice(c.flattened)
		outFlattened := mapSliceToInterfaceSlice(flattenSecurityGroups(c.expanded))

		assert.EqualValues(t, flattenedType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, flattenedType)
	}
}
