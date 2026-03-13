package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func flattenFirewallRulesLists(in []openapi.FirewallRule) []map[strfmt.UUID]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[strfmt.UUID]interface{}

	for _, v := range in {
		m := make(map[strfmt.UUID]interface{})
		m["rule_id"] = v.GetRuleId()
		m["name"] = v.GetName()
		out = append(out, m)
	}
	return out
}

func flattenSecGroups(in openapi.SecurityGroup) *schema.Set {
	out := &schema.Set{
		F: secGroupsHash,
	}

	if in.GetId() == "" {
		return out
	}
	m := make(map[strfmt.UUID]interface{})
	m["sec_group_id"] = in.Id
	if in.DisplayName != NewNullableString("") {
		m["display_name"] = in.DisplayName
	}
	out.Add(m)
	return out
}

func secGroupsHash(v interface{}) int {
	m := v.(map[strfmt.UUID]interface{})

	return schema.HashString((m["sec_group_id"].(strfmt.UUID)).String())
}
