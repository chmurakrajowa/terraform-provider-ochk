package ochk

import "github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"

func flattenProjects(in []*models.ProjectInstance) []map[string]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[string]interface{}

	for _, v := range in {
		m := make(map[string]interface{})
		m["project_id"] = v.ProjectID
		m["display_name"] = v.Name
		out = append(out, m)
	}
	return out
}
