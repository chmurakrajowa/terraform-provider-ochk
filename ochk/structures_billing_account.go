package ochk

import (
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func flattenAccount(in []*models.AccountInstance) []map[string]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[string]interface{}

	for _, v := range in {
		m := make(map[string]interface{})
		m["account_id"] = fmt.Sprint(v.AccountID)
		m["display_name"] = v.AccountName

		out = append(out, m)
	}
	return out
}

func flattenAccProjects(in []*models.AccountProjectInstance) *schema.Set {
	out := &schema.Set{
		F: projectsHash,
	}

	for _, v := range in {
		m := make(map[strfmt.UUID]interface{})
		m["project_id"] = v.ProjectID
		if v.Name != "" {
			m["display_name"] = v.Name
		}
		out.Add(m)
	}
	return out
}

func projectsHash(v interface{}) int {
	m := v.(map[string]interface{})

	return schema.HashString(m["project_id"])
}

func expandAcctProjects(in []interface{}) []*models.AccountProjectInstance {
	if len(in) == 0 {
		return nil
	}

	var out = make([]*models.AccountProjectInstance, len(in))
	for i, v := range in {
		m := v.(map[strfmt.UUID]interface{})

		member := &models.AccountProjectInstance{
			ProjectID: m["project_id"].(strfmt.UUID),
		}

		if displayName, ok := m["display_name"].(string); ok && displayName != "" {
			member.Name = displayName
		}

		out[i] = member
	}
	return out
}
