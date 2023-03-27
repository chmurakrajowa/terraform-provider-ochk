package ochk

import "github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"

func flattenVrfs(in []*models.RouterInstance) []map[string]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[string]interface{}

	for _, v := range in {
		if v.RouterType == "TIER0" {
			m := make(map[string]interface{})
			m["vrf_id"] = v.RouterID
			m["display_name"] = v.DisplayName
			out = append(out, m)
		}
	}
	return out
}
