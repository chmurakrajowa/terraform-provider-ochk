package ochk

import (
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenFloatingIpAddress(t *testing.T) {
	cases := []struct {
		expanded  []*models.FloatingIP
		flattened []map[strfmt.UUID]interface{}
	}{
		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []*models.FloatingIP{
				{
					FloatingIPID:  "92457621-e5af-49d6-921e-184566018fa9",
					Name:          "publicIp1",
					PublicAddress: "203.0.113.15",
				},
			},
			flattened: []map[strfmt.UUID]interface{}{
				{
					"floating_ip_id": strfmt.UUID("92457621-e5af-49d6-921e-184566018fa9"),
					"display_name":   "publicIp1",
					"public_address": "203.0.113.15",
				},
			},
		},
	}
	for i, c := range cases {
		fmt.Printf("index === %d\n", i)
		flattenedType := mapSliceToInterfaceSlice(c.flattened)
		fmt.Println("flattenedType >>>>> %w", flattenedType)
		outFlattened := mapSliceToInterfaceSlice(flattenFloatingIPAddressList(c.expanded))
		fmt.Println("outFlattened >>>>> %w", outFlattened)
		assert.EqualValues(t, flattenedType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, flattenedType)
	}
}
