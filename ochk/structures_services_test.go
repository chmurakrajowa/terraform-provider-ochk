package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenExpandServicesFromIDs(t *testing.T) {
	cases := []struct {
		expanded        []openapi.ServiceInstance
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
			expanded: []openapi.ServiceInstance{
				{
					ServiceId: NewNullableString("afdb07d8-d0d2-11ea-87d0-0242ac130003"),
				},
			},
			flattened: []interface{}{
				"afdb07d8-d0d2-11ea-87d0-0242ac130003",
			},
		},

		// more fields then necessary (test only flatten)
		{
			expanded: []openapi.ServiceInstance{
				{
					DisplayName: NewNullableString("http"),
					ServiceId:   NewNullableString("afdb07d8-d0d2-11ea-87d0-0242ac130003"),
				},
			},
			flattened: []interface{}{
				"afdb07d8-d0d2-11ea-87d0-0242ac130003",
			},
			onlyTestFlatten: true,
		},

		// multile routers
		{
			expanded: []openapi.ServiceInstance{
				{
					ServiceId: NewNullableString("afdb07d8-d0d2-11ea-87d0-0242ac130003"),
				},
				{
					ServiceId: NewNullableString("114d82f0-79cc-4501-a574-dd920b6b6e7e"),
				},
			},
			flattened: []interface{}{
				"afdb07d8-d0d2-11ea-87d0-0242ac130003",
				"114d82f0-79cc-4501-a574-dd920b6b6e7e",
			},
		},
	}

	for _, c := range cases {
		//flattenedSetType := schema.NewSet(schema.HashString, c.flattened).List()
		//outFlattened := flattenServicesFromIDs(c.expanded).List()
		//assert.EqualValues(t, flattenedSetType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, c.flattened)

		if !c.onlyTestFlatten {
			outExpanded := expandServicesFromIDs(c.flattened)
			assert.EqualValues(t, c.expanded, outExpanded, "Error matching output and expanded: %#v vs %#v", outExpanded, c.expanded)
		}

	}
}

func TestFlattenServices(t *testing.T) {
	cases := []struct {
		expanded  []openapi.ServiceInstance
		flattened []map[strfmt.UUID]interface{}
	}{
		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []openapi.ServiceInstance{
				{
					ServiceId:   NewNullableString("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					DisplayName: NewNullableString("test1"),
				},
				{
					ServiceId:   NewNullableString("791bf702-22fb-4c76-bebb-1fee7ee75607"),
					DisplayName: NewNullableString("test2"),
				},
			},
			flattened: []map[strfmt.UUID]interface{}{
				{
					"service_id":   strfmt.UUID("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					"display_name": "test1",
				},
				{
					"service_id":   strfmt.UUID("791bf702-22fb-4c76-bebb-1fee7ee75607"),
					"display_name": "test2",
				},
			},
		},
	}
	for _, c := range cases {
		flattenedType := mapSliceToInterfaceSlice(c.flattened)
		outFlattened := mapSliceToInterfaceSlice(flattenServices(c.expanded))
		assert.EqualValues(t, flattenedType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, flattenedType)
	}
}
