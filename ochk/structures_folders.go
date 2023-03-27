package ochk

import "github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"

func flattenFolders(in []*models.FolderInstance) []map[string]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[string]interface{}

	for _, v := range in {
		m := make(map[string]interface{})
		m["folder_id"] = v.ID
		m["folder_name"] = v.Name
		m["folder_path"] = v.FolderPath
		out = append(out, m)
	}
	return out
}
