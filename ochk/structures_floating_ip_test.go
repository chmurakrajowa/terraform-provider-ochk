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
				//{
				//	FloatingIPID:  "791bf702-22fb-4c76-bebb-1fee7ee75607",
				//	Name:          "test2",
				//	PublicAddress: "/test1/test2",
				//},
			},
			flattened: []map[strfmt.UUID]interface{}{
				{
					"floating_ip_id": strfmt.UUID("92457621-e5af-49d6-921e-184566018fa9"),
					"display_name":   "publicIp1",
					"public_address": "203.0.113.15",
				},
				//{
				//	"id":           strfmt.UUID("791bf702-22fb-4c76-bebb-1fee7ee75607"),
				//	"display_name": "test2",
				//	"description":  "/test1/test2",
				//},
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
