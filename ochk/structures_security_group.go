package ochk

import (
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func flattenSecurityGroupFromIDs(m []openapi.SecurityGroup) *schema.Set {
	s := &schema.Set{
		F: schema.HashString,
	}

	for _, v := range m {
		s.Add(fmt.Sprint(v.GetId()))
	}
	return s
}

func expandSecurityGroupFromIDs(in []interface{}) []openapi.SecurityGroup {
	if len(in) == 0 {
		return nil
	}

	var out = make([]openapi.SecurityGroup, len(in))

	for i, v := range in {
		idValue := strfmt.UUID.String(strfmt.UUID(v.(string)))
		securityGroup := openapi.SecurityGroup{
			Id: NewNullableString(idValue),
		}
		out[i] = securityGroup
	}

	return out
}

func flattenSecurityGroupMembers(in []openapi.SecurityGroupMember) *schema.Set {
	out := &schema.Set{
		F: securityGroupMembersHash,
	}

	for _, v := range in {
		m := make(map[string]interface{})
		m["id"] = v.GetId()
		m["type"] = v.GetMemberType()

		if v.DisplayName != NewNullableString("") {
			m["display_name"] = v.GetDisplayName()
		}

		out.Add(m)
	}
	return out
}

func expandSecurityGroupMembers(in []interface{}, platformType openapi.PlatformType) ([]openapi.SecurityGroupMember, diag.Diagnostics, string) {
	var out = make([]openapi.SecurityGroupMember, len(in))
	if len(in) == 0 {
		return out, nil, ""
	}
	for i, v := range in {
		m := v.(map[string]interface{})

		memberTypeValue := openapi.SecurityGroupMemberType(m["type"].(string))
		member := openapi.SecurityGroupMember{
			Id:         NewNullableString(m["id"].(string)),
			MemberType: &memberTypeValue,
		}
		if platformType == "OPENSTACK" {
			if *member.MemberType == "IPCOLLECTION" {
				return nil, diag.Errorf("error while expand security group:'' %+v", IPCOLLECTION), "IPCOLLECTION"
			} else if *member.MemberType == "LOGICAL_PORT" {
				return nil, diag.Errorf("error while expand security group:'' %+v", LOGICAL_PORT), "LOGICAL_PORT"
			} else if *member.MemberType == "IPSET" {
				return nil, diag.Errorf("error while expand security group:'' %+v", LOGICAL_PORT), "IPSET"
			} else if *member.MemberType == "SEGMENT" {
				return nil, diag.Errorf("error while expand security group:'' %+v", SEGMENT), "SEGMENT"
			} else if *member.MemberType == "GROUP" {
				return nil, diag.Errorf("error while expand security group:'' %+v", GROUP), "GROUP"
			}
		}

		if platformType == "VMWARE" {
			if *member.MemberType == "IPSET" {
				return nil, diag.Errorf("error while expand security group:'' %+v", IPSET), "IPSET"
			} else if *member.MemberType == "LOGICAL_PORT" {
				return nil, diag.Errorf("error while expand security group:'' %+v", LOGICAL_PORT), "LOGICAL_PORT"
			} else if *member.MemberType == "SEGMENT" {
				return nil, diag.Errorf("error while expand security group:'' %+v", SEGMENT), "SEGMENT"
			} else if *member.MemberType == "GROUP" {
				return nil, diag.Errorf("error while expand security group:'' %+v", GROUP), "GROUP"
			}

			if !contains([]string{string(VIRTUAL_MACHINE), string(IPCOLLECTION)}, string(memberTypeValue)) {
				return nil, diag.Errorf("error while expand security group:'' %+v", VIRTUAL_MACHINE, IPCOLLECTION),
					"Available MemberType: " + string(VIRTUAL_MACHINE) + " " + string(IPCOLLECTION)
			}

		}

		if displayName, ok := m["display_name"].(string); ok && displayName != "" {
			member.DisplayName = NewNullableString(displayName)
		}
		out[i] = member
	}
	return out, nil, ""
}

func securityGroupMembersHash(v interface{}) int {
	m := v.(map[string]interface{})
	return schema.HashString(m["id"].(string))
}
