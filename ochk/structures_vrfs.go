package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/openapi/v3"
	"github.com/go-openapi/strfmt"
)

func flattenVrfs(in []openapi.RouterInstance) []map[strfmt.UUID]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[strfmt.UUID]interface{}

	for _, v := range in {
		if *v.RouterType == "TIER0" {
			m := make(map[strfmt.UUID]interface{})
			m["vrf_id"] = v.RouterId
			m["display_name"] = v.DisplayName
			out = append(out, m)
		}
	}
	return out
}
