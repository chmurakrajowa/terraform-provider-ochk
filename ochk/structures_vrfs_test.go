package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenVrfs(t *testing.T) {
	cases := []struct {
		expanded  []*models.RouterInstance
		flattened []map[string]interface{}
	}{
		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []*models.RouterInstance{
				{
					RouterID:    "e1675817-f1a1-45c1-988b-ec2f142867e0",
					DisplayName: "VRF1",
					RouterType:  "TIER0",
				},
				{
					RouterID:    "791bf702-22fb-4c76-bebb-1fee7ee75607",
					DisplayName: "VRF2",
					RouterType:  "TIER0",
				},
			},
			flattened: []map[string]interface{}{
				{
					"vrf_id":       "e1675817-f1a1-45c1-988b-ec2f142867e0",
					"display_name": "VRF1",
				},
				{
					"vrf_id":       "791bf702-22fb-4c76-bebb-1fee7ee75607",
					"display_name": "VRF2",
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
