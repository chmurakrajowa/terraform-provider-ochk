package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenPublicIPAddress(t *testing.T) {
	cases := []struct {
		expanded  []*models.PublicIPAllocation
		flattened []map[string]interface{}
	}{
		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []*models.PublicIPAllocation{
				{
					Name: "test1",
					PublicIPAddress: &models.PublicIPAddress{
						IPAddress:   "212.86.15.4",
						IPAddressID: "10",
					},
				},
				{
					Name: "test2",
					PublicIPAddress: &models.PublicIPAddress{
						IPAddress:   "212.88.15.4",
						IPAddressID: "20",
					},
				},
			},
			flattened: []map[string]interface{}{
				{
					"public_ip_address_id": "10",
					"display_name":         "test1",
					"ip_address":           "212.86.15.4",
				},
				{
					"public_ip_address_id": "20",
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
