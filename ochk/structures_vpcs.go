package ochk

import "github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"

func flattenVpcs(in []*models.RouterInstance) []map[string]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[string]interface{}

	for _, v := range in {
		if v.RouterType == "TIER1" {
			m := make(map[string]interface{})
			m["vpc_id"] = v.RouterID
			m["vrf_id"] = v.ParentT0ID
			m["display_name"] = v.DisplayName
			m["project_id"] = v.ProjectID
			m["folder_path"] = v.FolderPath
			out = append(out, m)
		}
	}
	return out
}
