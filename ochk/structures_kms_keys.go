package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/openapi/v3"
	"github.com/go-openapi/strfmt"
)

func flattenKMSKeys(in []openapi.KeyInstance) []map[string]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[string]interface{}

	for _, v := range in {
		m := make(map[string]interface{})
		m["kms_key_id"] = v.Id
		m["display_name"] = v.Name
		m["key_usage"] = flattenStringSlice(v.KeyUsageList)
		m["state"] = v.State

		out = append(out, m)
	}
	return out
}

func expandKMSKeys(in []interface{}) []openapi.KeyInstance {
	if len(in) == 0 {
		return nil
	}

	var out = make([]openapi.KeyInstance, len(in))
	for i, v := range in {
		m := v.(map[strfmt.UUID]interface{})

		member := openapi.KeyInstance{}

		if paramName, ok := m["kms_key_id"].(string); ok {
			member.Id = paramName
		}

		if paramType, ok := m["display_name"].(string); ok {
			member.Name = paramType
		}

		if paramValue, ok := m["key_usage"]; ok {
			n := paramValue.([]string)
			if len(n) == 0 {
				member.KeyUsageList = nil
			} else {
				for _, x := range n {
					member.KeyUsageList = append(member.KeyUsageList, x)
				}
			}
		}

		if paramValue, ok := m["state"].(string); ok {
			member.State = paramValue
		}
		out[i] = member
	}
	return out
}
