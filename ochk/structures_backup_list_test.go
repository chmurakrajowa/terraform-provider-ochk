package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenBackupLists(t *testing.T) {
	cases := []struct {
		expanded  []openapi.BackupList
		flattened []map[strfmt.UUID]interface{}
	}{
		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []openapi.BackupList{
				{
					BackupListId:   NewNullableString("5403439c-38a5-4f98-a58b-134072260bfb"),
					BackupListName: NewNullableString("Platinium (1 week / 24h)"),
				},
				{
					BackupListId:   NewNullableString("791bf702-22fb-4c76-bebb-1fee7ee75607"),
					BackupListName: NewNullableString("Platinium (1 day / 8h)"),
				},
			},
			flattened: []map[strfmt.UUID]interface{}{
				{
					"backup_list_id": strfmt.UUID("5403439c-38a5-4f98-a58b-134072260bfb"),
					"display_name":   "Platinium (1 week / 24h)",
				},
				{
					"backup_list_id": strfmt.UUID("791bf702-22fb-4c76-bebb-1fee7ee75607"),
					"display_name":   "Platinium (1 day / 8h)",
				},
			},
		},
	}
	for _, c := range cases {
		flattenedType := mapSliceToInterfaceSlice(c.flattened)
		outFlattened := mapSliceToInterfaceSlice(flattenBackupLists(c.expanded))
		assert.EqualValues(t, flattenedType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, flattenedType)
	}
}
