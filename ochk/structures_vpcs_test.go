package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenVpcs(t *testing.T) {
	cases := []struct {
		expanded  []*models.RouterInstance
		flattened []map[strfmt.UUID]interface{}
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
					ProjectID:   "e2655817-f1a1-4c76-bebb-1fee7ee75607",
					ParentT0ID:  "7364ff23-0513-48b6-97cb-637613e29424",
					FolderPath:  "/test1",
					RouterType:  "TIER1",
				},
				{
					RouterID:    "791bf702-22fb-4c76-bebb-1fee7ee75607",
					DisplayName: "VRF2",
					ProjectID:   "e2655817-f1a1-4c76-bebb-1fee7ee75607",
					ParentT0ID:  "547948e9-b67d-44d1-ad69-ae9b711e289c",
					FolderPath:  "/test2",
					RouterType:  "TIER1",
				},
			},
			flattened: []map[strfmt.UUID]interface{}{
				{
					"vpc_id":       strfmt.UUID("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					"display_name": "VRF1",
					"project_id":   strfmt.UUID("e2655817-f1a1-4c76-bebb-1fee7ee75607"),
					"vrf_id":       strfmt.UUID("7364ff23-0513-48b6-97cb-637613e29424"),
					"folder_path":  "/test1",
				},
				{
					"vpc_id":       strfmt.UUID("791bf702-22fb-4c76-bebb-1fee7ee75607"),
					"display_name": "VRF2",
					"project_id":   strfmt.UUID("e2655817-f1a1-4c76-bebb-1fee7ee75607"),
					"vrf_id":       strfmt.UUID("547948e9-b67d-44d1-ad69-ae9b711e289c"),
					"folder_path":  "/test2",
				},
			},
		},
	}
	for _, c := range cases {
		flattenedType := mapSliceToInterfaceSlice(c.flattened)
		outFlattened := mapSliceToInterfaceSlice(flattenVpcs(c.expanded))
		assert.EqualValues(t, flattenedType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, flattenedType)
	}
}
