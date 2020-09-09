package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenExpandUserInstancesFromIDs(t *testing.T) {
	cases := []struct {
		expanded  []*models.UserInstance
		flattened []interface{}
	}{
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []*models.UserInstance{
				{
					UserID: "afdb07d8-d0d2-11ea-87d0-0242ac130003",
				},
			},
			flattened: []interface{}{
				"afdb07d8-d0d2-11ea-87d0-0242ac130003",
			},
		},
		{
			expanded: []*models.UserInstance{
				{
					UserID: "afdb07d8-d0d2-11ea-87d0-0242ac130003",
				},
				{
					UserID: "114d82f0-79cc-4501-a574-dd920b6b6e7e",
				},
			},
			flattened: []interface{}{
				"afdb07d8-d0d2-11ea-87d0-0242ac130003",
				"114d82f0-79cc-4501-a574-dd920b6b6e7e",
			},
		},
	}

	for _, c := range cases {
		flattenedSetType := schema.NewSet(schema.HashString, c.flattened).List()
		outFlattened := flattenUserInstancesFromIDs(c.expanded).List()
		assert.EqualValues(t, flattenedSetType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, c.flattened)
		outExpanded := expandUserInstancesFromIDs(c.flattened)
		assert.EqualValues(t, c.expanded, outExpanded, "Error matching output and expanded: %#v vs %#v", outExpanded, c.expanded)
	}
}

func TestFlattenExpandVCSNetworkInstancesFromIDs(t *testing.T) {
	cases := []struct {
		expanded  []*models.VCSNetworkInstance
		flattened []interface{}
	}{
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []*models.VCSNetworkInstance{
				{
					NetworkID: "afdb07d8-d0d2-11ea-87d0-0242ac130003",
				},
			},
			flattened: []interface{}{
				"afdb07d8-d0d2-11ea-87d0-0242ac130003",
			},
		},
		{
			expanded: []*models.VCSNetworkInstance{
				{
					NetworkID: "afdb07d8-d0d2-11ea-87d0-0242ac130003",
				},
				{
					NetworkID: "114d82f0-79cc-4501-a574-dd920b6b6e7e",
				},
			},
			flattened: []interface{}{
				"afdb07d8-d0d2-11ea-87d0-0242ac130003",
				"114d82f0-79cc-4501-a574-dd920b6b6e7e",
			},
		},
	}

	for _, c := range cases {
		flattenedSetType := schema.NewSet(schema.HashString, c.flattened).List()
		outFlattened := flattenVCSNetworkInstancesFromIDs(c.expanded).List()
		assert.EqualValues(t, flattenedSetType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, c.flattened)
		outExpanded := expandVCSNetworkInstancesFromIDs(c.flattened)
		assert.EqualValues(t, c.expanded, outExpanded, "Error matching output and expanded: %#v vs %#v", outExpanded, c.expanded)
	}
}
