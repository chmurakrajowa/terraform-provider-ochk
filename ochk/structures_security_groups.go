package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
)

func flattenSecurityGroups(in []openapi.SecurityGroup) []map[strfmt.UUID]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[strfmt.UUID]interface{}

	for _, v := range in {
		m := make(map[strfmt.UUID]interface{})
		m["security_group_id"] = v.Id
		m["display_name"] = v.DisplayName
		m["project_id"] = v.ProjectId
		out = append(out, m)
	}
	return out
}
