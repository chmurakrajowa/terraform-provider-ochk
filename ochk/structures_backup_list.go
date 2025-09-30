package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/openapi/v3"
	"github.com/go-openapi/strfmt"
)

func flattenBackupLists(in []openapi.BackupList) []map[strfmt.UUID]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[strfmt.UUID]interface{}

	for _, v := range in {
		m := make(map[strfmt.UUID]interface{})
		m["backup_list_id"] = v.BackupListId
		m["display_name"] = v.BackupListName
		out = append(out, m)
	}
	return out
}
