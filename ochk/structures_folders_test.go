package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenFolders(t *testing.T) {
	IdValue := "e1675817-f1a1-45c1-988b-ec2f142867e0"
	IdValue2 := "791bf702-22fb-4c76-bebb-1fee7ee75607"
	cases := []struct {
		expanded  []openapi.FolderInstance
		flattened []map[strfmt.UUID]interface{}
	}{
		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []openapi.FolderInstance{
				{
					Id:         &IdValue,
					Name:       NewNullableString("test1"),
					FolderPath: NewNullableString("/test1"),
				},
				{
					Id:         &IdValue2,
					Name:       NewNullableString("test2"),
					FolderPath: NewNullableString("/test1/test2"),
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
