package ochk

import "github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/models"

func flattenSecurityGroupMember(in []*models.SecurityGroupMember) []interface{} {
	var out = make([]interface{}, len(in), len(in))

	for i, v := range in {
		m := make(map[string]interface{})
		m["id"] = v.ExternalID
		m["type"] = v.MemberType
		m["displayName"] = v.DisplayName

		out[i] = m
	}
	return out
}

func expandSecurityGroupMember(in []interface{}) []map[string]interface{} {
	var out = make([]models.SecurityGroupMember, len(in), len(in))

	for i, v := range in {
		m := make(map[string]interface{})
		m["id"] = v.ExternalID
		m["type"] = v.MemberType
		m["displayName"] = v.DisplayName

		out[i] = m
	}
	return out
}
