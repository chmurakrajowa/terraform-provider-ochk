package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
			ID: v.(strfmt.UUID),
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

func expandSecurityGroupMembers(in []interface{}, platformType models.PlatformType) ([]*models.SecurityGroupMember, diag.Diagnostics, string) {
	var out = make([]*models.SecurityGroupMember, len(in))
	if len(in) == 0 {
		return out, nil, ""
	}
	for i, v := range in {
		m := v.(map[string]interface{})

		member := &models.SecurityGroupMember{
			ID:         strfmt.UUID(m["id"].(string)),
			MemberType: models.SecurityGroupMemberType(m["type"].(string)),
		}
		if platformType == "OPENSTACK" {
			if member.MemberType == "IPCOLLECTION" {
				return nil, diag.Errorf("error while expand security group:'' %+v", IPCOLLECTION), "IPCOLLECTION"
			} else if member.MemberType == "LOGICAL_PORT" {
				return nil, diag.Errorf("error while expand security group:'' %+v", LOGICAL_PORT), "LOGICAL_PORT"
			} else if member.MemberType == "IPSET" {
				return nil, diag.Errorf("error while expand security group:'' %+v", LOGICAL_PORT), "IPSET"
			}
		}

		if platformType == "VMWARE" {
			if member.MemberType == "IPSET" {
				return nil, diag.Errorf("error while expand security group:'' %+v", IPSET), "IPSET"
			} else if member.MemberType == "LOGICAL_PORT" {
				return nil, diag.Errorf("error while expand security group:'' %+v", LOGICAL_PORT), "LOGICAL_PORT"
			}
		}

		if displayName, ok := m["display_name"].(string); ok && displayName != "" {
			member.DisplayName = displayName
		}
		out[i] = member
	}
	return out, nil, ""
}

func securityGroupMembersHash(v interface{}) int {
	m := v.(map[string]interface{})

	return schema.HashString(m["id"].(strfmt.UUID).String())
}
