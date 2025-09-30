package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/openapi/v3"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenExpandSecurityGroupMembers(t *testing.T) {

	cases := []struct {
		expanded  []openapi.SecurityGroup
		flattened []map[strfmt.UUID]interface{}
	}{

		{
			expanded:  nil,
			flattened: nil,
		},

		{
			expanded: []openapi.SecurityGroup{
				{
					Id:          strfmt.UUID("dcbf922a-a2fc-401f-ac1f-5159c15d4b8b"),
					ProjectId:   strfmt.UUID("dcbf922a-a2fc-401f-ac1f-5159c15d4b8b"),
					DisplayName: "MySG1",
				},
			},
			flattened: []map[strfmt.UUID]interface{}{
				{
					"security_group_id": strfmt.UUID("dcbf922a-a2fc-401f-ac1f-5159c15d4b8b"),
					"project_id":        strfmt.UUID("dcbf922a-a2fc-401f-ac1f-5159c15d4b8b"),
					"display_name":      "MySG1",
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
