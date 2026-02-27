package ochk

import (
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func flattenAccount(in []openapi.AccountInstance) []map[string]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[string]interface{}

	for _, v := range in {
		m := make(map[string]interface{})
		m["account_id"] = fmt.Sprint(v.GetAccountId())
		m["display_name"] = v.AccountName.Get()

		out = append(out, m)
	}
	return out
}

func flattenAccProjects(in []openapi.AccountProjectInstance) *schema.Set {
	out := &schema.Set{
		F: projectsHash,
	}

	for _, v := range in {
		m := make(map[string]interface{})
		m["project_id"] = v.GetProjectId()
		if v.Name != NewNullableString("") {
			m["display_name"] = v.Name.Get()
		}
		out.Add(m)
	}
	return out
}

//func projectsHash(v interface{}) int {
//	m := v.(map[strfmt.UUID]interface{})
//
//	return schema.HashString((m["project_id"].(strfmt.UUID)).String())
//}

func projectsHash(v interface{}) int {
	m := v.(map[string]interface{})

	projectID := m["project_id"].(string)

	return schema.HashString(projectID)
}

func expandAcctProjects(in []interface{}) []openapi.AccountProjectInstance {
	if len(in) == 0 {
		return nil
	}

	var out = make([]openapi.AccountProjectInstance, len(in))
	for i, v := range in {
		m := v.(map[strfmt.UUID]interface{})

		member := openapi.AccountProjectInstance{
			ProjectId: NewNullableString(m["project_id"].(string)),
		}

		if displayName, ok := m["display_name"].(string); ok && displayName != "" {
			member.Name = NewNullableString(displayName)
		}

		out[i] = member
	}
	return out
}
