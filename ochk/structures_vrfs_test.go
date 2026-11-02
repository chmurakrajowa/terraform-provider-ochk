package ochk

import (
	"testing"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
)

func TestFlattenVrfs(t *testing.T) {
	stringPtr := func(value string) *string {
		return &value
	}
	cases := []struct {
		expanded  []openapi.RouterInstance
		flattened []map[strfmt.UUID]interface{}
	}{
		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []openapi.RouterInstance{
				{
					RouterId:    NewNullableString("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					DisplayName: NewNullableString("VRF1"),
					RouterType:  openapi.ROUTERTYPE_TIER0.Ptr(),
				},
				{
					RouterId:    NewNullableString("791bf702-22fb-4c76-bebb-1fee7ee75607"),
					DisplayName: NewNullableString("VRF2"),
					RouterType:  openapi.ROUTERTYPE_TIER0.Ptr(),
				},
			},
			flattened: []map[strfmt.UUID]interface{}{
				{
					"vrf_id":       stringPtr("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					"display_name": stringPtr("VRF1"),
				},
				{
					"vrf_id":       stringPtr("791bf702-22fb-4c76-bebb-1fee7ee75607"),
					"display_name": stringPtr("VRF2"),
				},
			},
		},
	}
	for _, c := range cases {
		flattenedType := mapSliceToInterfaceSlice(c.flattened)
		outFlattened := mapSliceToInterfaceSlice(flattenVrfs(c.expanded))
		assert.EqualValues(t, flattenedType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, flattenedType)
	}
}
