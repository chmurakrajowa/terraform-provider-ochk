package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
)

func flattenKMSKeys(in []*models.KeyInstance) []map[string]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[string]interface{}

	for _, v := range in {
		m := make(map[string]interface{})
		m["kms_key_id"] = v.ID
		m["display_name"] = v.Name
		m["key_usage"] = flattenStringSlice(v.KeyUsageList)
		m["state"] = v.State

		out = append(out, m)
	}
	return out
}

func expandKMSKeys(in []interface{}) []*models.KeyInstance {
	if len(in) == 0 {
		return nil
	}

	var out = make([]*models.KeyInstance, len(in))
	for i, v := range in {
		m := v.(map[strfmt.UUID]interface{})

		member := &models.KeyInstance{}

		if paramName, ok := m["kms_key_id"].(string); ok {
			member.ID = paramName
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
