package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
)

func flattenVpcs(in []openapi.RouterInstance) []map[strfmt.UUID]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[strfmt.UUID]interface{}

	for _, v := range in {
		if *v.RouterType == "TIER1" {
			m := make(map[strfmt.UUID]interface{})
			m["vpc_id"] = v.GetRouterId()
			m["vrf_id"] = v.ParentT0Id.Get()
			m["display_name"] = v.GetDisplayName()
			m["project_id"] = v.GetProjectId()
			m["folder_path"] = v.GetFolderPath()
			out = append(out, m)
		}
	}
	return out
}
