package ochk

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/models"
)

func flattenSecurityGroupMembers(in []*models.SecurityGroupMember) []map[string]interface{} {
	var out = make([]map[string]interface{}, len(in), len(in))

	for i, v := range in {
		m := make(map[string]interface{})
		m["id"] = v.ID
		m["type"] = v.MemberType

		if v.DisplayName != "" {
			m["display_name"] = v.DisplayName
		}

		out[i] = m
	}
	return out
}

func expandSecurityGroupMembers(list *schema.Set) []*models.SecurityGroupMember {
	in := list.List()
	if len(in) == 0 {
		return nil
	}

	var out = make([]*models.SecurityGroupMember, len(in), len(in))
	for i, v := range in {
		m := v.(map[string]interface{})

		member := &models.SecurityGroupMember{
			ID:         m["id"].(string),
			MemberType: m["type"].(string),
		}

		if displayName := m["display_name"].(string); displayName != "" {
			member.DisplayName = displayName
		}

		out[i] = member
	}
	return out
}
