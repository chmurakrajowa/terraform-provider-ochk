package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/openapi/v3"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenExpandRouterInstanceFromIDs(t *testing.T) {
	cases := []struct {
		expanded        []openapi.RouterInstance
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
			expanded: []openapi.RouterInstance{
				{
					RouterId: NewNullableString("afdb07d8-d0d2-11ea-87d0-0242ac130003"),
				},
			},
			flattened: []interface{}{
				strfmt.UUID("afdb07d8-d0d2-11ea-87d0-0242ac130003"),
			},
		},

		// more fields then necessary (test only flatten)
		{
			expanded: []openapi.RouterInstance{
				{
					DisplayName: NewNullableString("T1"),
					RouterId:    NewNullableString("afdb07d8-d0d2-11ea-87d0-0242ac130003"),
					RouterType:  "Ingress Router",
				},
			},
			flattened: []interface{}{
				strfmt.UUID("afdb07d8-d0d2-11ea-87d0-0242ac130003"),
			},
			onlyTestFlatten: true,
		},

		// multiple routers
		{
			expanded: []openapi.RouterInstance{
				{
					RouterId: NewNullableString("afdb07d8-d0d2-11ea-87d0-0242ac130003"),
				},
				{
					RouterId: NewNullableString("114d82f0-79cc-4501-a574-dd920b6b6e7e"),
				},
			},
			flattened: []interface{}{
				strfmt.UUID("afdb07d8-d0d2-11ea-87d0-0242ac130003"),
				strfmt.UUID("114d82f0-79cc-4501-a574-dd920b6b6e7e"),
			},
		},
	}

	for _, c := range cases {
		//flattenedSetType := schema.NewSet(schema.HashString, c.flattened).List()
		//outFlattened := flattenRouterInstancesFromIDs(c.expanded).List()
		//assert.EqualValues(t, flattenedSetType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, c.flattened)

		if !c.onlyTestFlatten {
			outExpanded := expandRouterInstancesFromIDs(c.flattened)
			assert.EqualValues(t, c.expanded, outExpanded, "Error matching output and expanded: %#v vs %#v", outExpanded, c.expanded)
		}
	}
}
