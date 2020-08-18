package ochk

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenExpandSecurityGroupMembers(t *testing.T) {
	cases := []struct {
		expanded  []*models.SecurityGroupMember
		flattened []map[string]interface{}
	}{

		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},

		// include display_name
		{
			expanded: []*models.SecurityGroupMember{
				{
					ID:          "afdb07d8-d0d2-11ea-87d0-0242ac130003",
					MemberType:  "LOGICAL_PORT",
					DisplayName: "MyLogicaPort",
				},
			},
			flattened: []map[string]interface{}{
				{
					"id":           "afdb07d8-d0d2-11ea-87d0-0242ac130003",
					"type":         "LOGICAL_PORT",
					"display_name": "MyLogicaPort",
				},
			},
		},

		// empty display name
		{
			expanded: []*models.SecurityGroupMember{
				{
					ID:         "afdb07d8-d0d2-11ea-87d0-0242ac130003",
					MemberType: "LOGICAL_PORT",
				},
			},
			flattened: []map[string]interface{}{
				{
					"id":   "afdb07d8-d0d2-11ea-87d0-0242ac130003",
					"type": "LOGICAL_PORT",
				},
			},
		},

		// multiple members
		{
			expanded: []*models.SecurityGroupMember{
				{
					ID:         "afdb07d8-d0d2-11ea-87d0-0242ac130003",
					MemberType: "LOGICAL_PORT",
				},
				{
					ID:         "f437d690-d0d2-11ea-87d0-0242ac130003",
					MemberType: "VIRTUAL_MACHINE",
				},
			},
			flattened: []map[string]interface{}{
				{
					"id":   "afdb07d8-d0d2-11ea-87d0-0242ac130003",
					"type": "LOGICAL_PORT",
				},
				{
					"id":   "f437d690-d0d2-11ea-87d0-0242ac130003",
					"type": "VIRTUAL_MACHINE",
				},
			},
		},
	}

	for _, c := range cases {
		flattenedSetType := schema.NewSet(securityGroupMembersHash, mapSliceToInterfaceSlice(c.flattened)).List()
		outFlattened := flattenSecurityGroupMembers(c.expanded).List()
		assert.EqualValues(t, flattenedSetType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, c.flattened)

		flattenedInterfaceSlice := mapSliceToInterfaceSlice(c.flattened)
		outExpanded := expandSecurityGroupMembers(flattenedInterfaceSlice)
		assert.EqualValues(t, c.expanded, outExpanded, "Error matching output and expanded: %#v vs %#v", outExpanded, c.expanded)
	}
}
