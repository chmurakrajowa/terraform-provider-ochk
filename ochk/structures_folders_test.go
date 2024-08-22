package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenFolders(t *testing.T) {
	cases := []struct {
		expanded  []*models.FolderInstance
		flattened []map[strfmt.UUID]interface{}
	}{
		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []*models.FolderInstance{
				{
					ID:         "e1675817-f1a1-45c1-988b-ec2f142867e0",
					Name:       "test1",
					FolderPath: "/test1",
				},
				{
					ID:         "791bf702-22fb-4c76-bebb-1fee7ee75607",
					Name:       "test2",
					FolderPath: "/test1/test2",
				},
			},
			flattened: []map[strfmt.UUID]interface{}{
				{
					"folder_id":   strfmt.UUID("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					"folder_name": "test1",
					"folder_path": "/test1",
				},
				{
					"folder_id":   strfmt.UUID("791bf702-22fb-4c76-bebb-1fee7ee75607"),
					"folder_name": "test2",
					"folder_path": "/test1/test2",
				},
			},
		},
	}
	for _, c := range cases {
		flattenedType := mapSliceToInterfaceSlice(c.flattened)
		outFlattened := mapSliceToInterfaceSlice(flattenFolders(c.expanded))
		assert.EqualValues(t, flattenedType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, flattenedType)
	}
}
