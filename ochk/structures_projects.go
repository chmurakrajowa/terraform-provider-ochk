package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
)

func flattenProjects(in []*models.ProjectInstance) []map[strfmt.UUID]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[strfmt.UUID]interface{}

	for _, v := range in {
		m := make(map[strfmt.UUID]interface{})
		m["project_id"] = v.ProjectID
		m["display_name"] = v.Name
		out = append(out, m)
	}
	return out
}
