package ochk

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/models"
)

func flattenSecurityGroupFromIDs(m []*models.SecurityGroup) *schema.Set {
	s := &schema.Set{
		F: schema.HashString,
	}

	for _, v := range m {
		s.Add(v.ID)
	}
	return s
}

func expandSecurityGroupFromIDs(in []interface{}) []*models.SecurityGroup {
	if len(in) == 0 {
		return nil
	}

	var out = make([]*models.SecurityGroup, len(in))

	for i, v := range in {
		securityGroup := &models.SecurityGroup{
			ID: v.(string),
		}

		out[i] = securityGroup
	}

	return out
}

func flattenSecurityGroupMembers(in []*models.SecurityGroupMember) *schema.Set {
	out := &schema.Set{
		F: securityGroupMembersHash,
	}

	for _, v := range in {
		m := make(map[string]interface{})
		m["id"] = v.ID
		m["type"] = v.MemberType

		if v.DisplayName != "" {
			m["display_name"] = v.DisplayName
		}

		out.Add(m)
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

func securityGroupMembersHash(v interface{}) int {
	m := v.(map[string]interface{})

	return schema.HashString(m["id"])
}
