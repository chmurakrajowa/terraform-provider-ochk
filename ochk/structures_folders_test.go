package ochk

import (
	"testing"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
)

func TestFlattenFolders(t *testing.T) {
	stringPtr := func(value string) *string {
		return &value
	}

	cases := []struct {
		name      string
		expanded  []openapi.FolderInstance
		flattened []map[strfmt.UUID]interface{}
	}{
		{
			name:      "nil input",
			expanded:  nil,
			flattened: nil,
		},
		{
			name: "two folders",
			expanded: []openapi.FolderInstance{
				{
					Id:         stringPtr("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					Name:       NewNullableString("test1"),
					FolderPath: NewNullableString("/test1"),
				},
				{
					Id:         stringPtr("791bf702-22fb-4c76-bebb-1fee7ee75607"),
					Name:       NewNullableString("test2"),
					FolderPath: NewNullableString("/test1/test2"),
				},
			},
			flattened: []map[strfmt.UUID]interface{}{
				{
					"folder_id":   stringPtr("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					"folder_name": stringPtr("test1"),
					"folder_path": stringPtr("/test1"),
				},
				{
					"folder_id":   stringPtr("791bf702-22fb-4c76-bebb-1fee7ee75607"),
					"folder_name": stringPtr("test2"),
					"folder_path": stringPtr("/test1/test2"),
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			expected := mapSliceToInterfaceSlice(c.flattened)
			actual := mapSliceToInterfaceSlice(flattenFolders(c.expanded))

			assert.EqualValues(t, expected, actual)
		})
	}
}
