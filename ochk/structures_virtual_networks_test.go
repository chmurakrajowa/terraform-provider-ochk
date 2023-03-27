package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenVirtualNeworks(t *testing.T) {
	cases := []struct {
		expanded  []*models.VirtualNetworkInstance
		flattened []map[string]interface{}
	}{
		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []*models.VirtualNetworkInstance{
				{
					VirtualNetworkID: "e1675817-f1a1-45c1-988b-ec2f142867e0",
					DisplayName:      "VRF1",
					ProjectID:        "e2655817-f1a1-4c76-bebb-1fee7ee75607",
					RouterRefID:      "7364ff23-0513-48b6-97cb-637613e29424",
					FolderPath:       "/test1",
					IpamEnabled:      true,
				},
				{
					VirtualNetworkID: "791bf702-22fb-4c76-bebb-1fee7ee75607",
					DisplayName:      "VRF2",
					ProjectID:        "e2655817-f1a1-4c76-bebb-1fee7ee75607",
					RouterRefID:      "547948e9-b67d-44d1-ad69-ae9b711e289c",
					FolderPath:       "/test2",
					IpamEnabled:      true,
				},
			},
			flattened: []map[string]interface{}{
				{
					"virtual_network_id": "e1675817-f1a1-45c1-988b-ec2f142867e0",
					"display_name":       "VRF1",
					"project_id":         "e2655817-f1a1-4c76-bebb-1fee7ee75607",
					"vpc_id":             "7364ff23-0513-48b6-97cb-637613e29424",
					"folder_path":        "/test1",
					"ipam_enabled":       true,
				},
				{
					"virtual_network_id": "791bf702-22fb-4c76-bebb-1fee7ee75607",
					"display_name":       "VRF2",
					"project_id":         "e2655817-f1a1-4c76-bebb-1fee7ee75607",
					"vpc_id":             "547948e9-b67d-44d1-ad69-ae9b711e289c",
					"folder_path":        "/test2",
					"ipam_enabled":       true,
				},
			},
		},
	}
	for _, c := range cases {
		flattenedType := mapSliceToInterfaceSlice(c.flattened)
		outFlattened := mapSliceToInterfaceSlice(flattenVirtualNeworks(c.expanded))
		assert.EqualValues(t, flattenedType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, flattenedType)
	}
}
