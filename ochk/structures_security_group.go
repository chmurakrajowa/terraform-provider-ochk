package ochk

import (
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/models"
)

func flattenSecurityGroupMembers(in []*models.SecurityGroupMember) []map[string]interface{} {
	var out = make([]map[string]interface{}, len(in))

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

func expandSecurityGroupMembers(in []interface{}) []*models.SecurityGroupMember {
	if len(in) == 0 {
		return nil
	}

	var out = make([]*models.SecurityGroupMember, len(in))
	for i, v := range in {
		m := v.(map[string]interface{})

		member := &models.SecurityGroupMember{
			ID:         m["id"].(string),
			MemberType: m["type"].(string),
		}

		if displayName, ok := m["display_name"].(string); ok && displayName != "" {
			member.DisplayName = displayName
		}

		out[i] = member
	}
	return out
}
