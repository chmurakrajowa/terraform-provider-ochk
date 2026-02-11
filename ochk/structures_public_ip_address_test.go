package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenPublicIPAddress(t *testing.T) {
	cases := []struct {
		expanded  []openapi.PublicIpAllocation
		flattened []map[strfmt.UUID]interface{}
	}{
		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []openapi.PublicIpAllocation{
				{
					Name: NewNullableString("test1"),
					PublicIpAddress: &openapi.PublicIpAddress{
						IpAddress:   NewNullableString("212.86.15.4"),
						IpAddressId: NewNullableString("f3f8c5c3-d3ca-4ec7-b553-81208cbc016b"),
					},
				},
				{
					Name: NewNullableString("test2"),
					PublicIpAddress: &openapi.PublicIpAddress{
						IpAddress:   NewNullableString("212.88.15.4"),
						IpAddressId: NewNullableString("a8d33488-f662-43af-9755-2509ce09d24b"),
					},
				},
			},
			flattened: []map[strfmt.UUID]interface{}{
				{
					"public_ip_address_id": strfmt.UUID("f3f8c5c3-d3ca-4ec7-b553-81208cbc016b"),
					"display_name":         "test1",
					"ip_address":           "212.86.15.4",
				},
				{
					"public_ip_address_id": strfmt.UUID("a8d33488-f662-43af-9755-2509ce09d24b"),
					"display_name":         "test2",
					"ip_address":           "212.88.15.4",
				},
			},
		},
	}
	for _, c := range cases {
		flattenedType := mapSliceToInterfaceSlice(c.flattened)
		outFlattened := mapSliceToInterfaceSlice(flattenPublicIPAddress(c.expanded))
		assert.EqualValues(t, flattenedType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, flattenedType)
	}
}
