package ochk

import "github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"

func flattenBackupLists(in []*models.BackupList) []map[string]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[string]interface{}

	for _, v := range in {
		m := make(map[string]interface{})
		m["backup_list_id"] = v.BackupListID
		m["display_name"] = v.BackupListName
		out = append(out, m)
	}
	return out
}
