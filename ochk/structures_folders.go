package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
)

func flattenFolders(in []*models.FolderInstance) []map[strfmt.UUID]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[strfmt.UUID]interface{}

	for _, v := range in {
		m := make(map[strfmt.UUID]interface{})
		m["folder_id"] = v.ID
		m["folder_name"] = v.Name
		m["folder_path"] = v.FolderPath
		out = append(out, m)
	}
	return out
}
