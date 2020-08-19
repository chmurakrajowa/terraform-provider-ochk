package ochk

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenExpandServicesFromIDs(t *testing.T) {
	cases := []struct {
		expanded        []*models.ServiceInstance
		flattened       []interface{}
		onlyTestFlatten bool
	}{
		// nil
		{
			expanded:  nil,
			flattened: nil,
		},

		// single router
		{
			expanded: []*models.ServiceInstance{
				{
					ServiceID: "afdb07d8-d0d2-11ea-87d0-0242ac130003",
				},
			},
			flattened: []interface{}{
				"afdb07d8-d0d2-11ea-87d0-0242ac130003",
			},
		},

		// more fields then necessary (test only flatten)
		{
			expanded: []*models.ServiceInstance{
				{
					DisplayName: "http",
					ServiceID:   "afdb07d8-d0d2-11ea-87d0-0242ac130003",
				},
			},
			flattened: []interface{}{
				"afdb07d8-d0d2-11ea-87d0-0242ac130003",
			},
			onlyTestFlatten: true,
		},

		// multile routers
		{
			expanded: []*models.ServiceInstance{
				{
					ServiceID: "afdb07d8-d0d2-11ea-87d0-0242ac130003",
				},
				{
					ServiceID: "114d82f0-79cc-4501-a574-dd920b6b6e7e",
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
		outFlattened := flattenServicesFromIDs(c.expanded).List()
		assert.EqualValues(t, flattenedSetType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, c.flattened)

		if !c.onlyTestFlatten {
			outExpanded := expandServicesFromIDs(c.flattened)
			assert.EqualValues(t, c.expanded, outExpanded, "Error matching output and expanded: %#v vs %#v", outExpanded, c.expanded)
		}
	}
}
