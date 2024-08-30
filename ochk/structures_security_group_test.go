package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenExpandSecurityGroupMembers(t *testing.T) {

	cases := []struct {
		expanded  []*models.SecurityGroupMember
		flattened []map[strfmt.UUID]interface{}
	}{

		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},

		// include display_name

		//{
		//	expanded: []*models.SecurityGroupMember{
		//		{
		//			ID:          strfmt.UUID("dcbf922a-a2fc-401f-ac1f-5159c15d4b8b"),
		//			MemberType:  "VIRTUAL_MACHINE",
		//			DisplayName: "MyVirtualMachine",
		//		},
		//	},
		//	flattened: []map[strfmt.UUID]interface{}{
		//		{
		//			"id":           strfmt.UUID("dcbf922a-a2fc-401f-ac1f-5159c15d4b8b"),
		//			"type":         models.SecurityGroupMemberType("VIRTUAL_MACHINE"),
		//			"display_name": "MyVirtualMachine",
		//		},
		//	},
		//},

		// empty display name

		//{
		//	expanded: []*models.SecurityGroupMember{
		//		{
		//			ID:         "afdb07d8-d0d2-11ea-87d0-0242ac130003",
		//			MemberType: models.SecurityGroupMemberType("LOGICAL_PORT"),
		//		},
		//	},
		//	flattened: []map[strfmt.UUID]interface{}{
		//		{
		//			"id":   strfmt.UUID("afdb07d8-d0d2-11ea-87d0-0242ac130003"),
		//			"type": models.SecurityGroupMemberType("LOGICAL_PORT"),
		//		},
		//	},
		//},

		// multiple members

		//{
		//	expanded: []*models.SecurityGroupMember{
		//		{
		//			ID:         "afdb07d8-d0d2-11ea-87d0-0242ac130003",
		//			MemberType: models.SecurityGroupMemberType("LOGICAL_PORT"),
		//		},
		//		{
		//			ID:         "f437d690-d0d2-11ea-87d0-0242ac130003",
		//			MemberType: models.SecurityGroupMemberType("VIRTUAL_MACHINE"),
		//		},
		//	},
		//	flattened: []map[strfmt.UUID]interface{}{
		//		{
		//			"id":   strfmt.UUID("afdb07d8-d0d2-11ea-87d0-0242ac130003"),
		//			"type": models.SecurityGroupMemberType("LOGICAL_PORT"),
		//		},
		//		{
		//			"id":   strfmt.UUID("f437d690-d0d2-11ea-87d0-0242ac130003"),
		//			"type": models.SecurityGroupMemberType("VIRTUAL_MACHINE"),
		//		},
		//	},
		//},
	}

	for _, c := range cases {
		//flattenedSetType := schema.NewSet(securityGroupMembersHash, mapSliceToInterfaceSlice(c.flattened)).List()
		//outFlattened := flattenSecurityGroupMembers(c.expanded).List()
		//assert.EqualValues(t, flattenedSetType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, c.flattened)

		flattenedInterfaceSlice := mapSliceToInterfaceSlice(c.flattened)
		members, err, _ := expandSecurityGroupMembers(flattenedInterfaceSlice, models.PlatformTypeOPENSTACK)
		if err != nil {

		}
		outExpanded := members
		assert.EqualValues(t, c.expanded, outExpanded, "Error matching output and expanded: %#v vs %#v", outExpanded, c.expanded)
	}
}
