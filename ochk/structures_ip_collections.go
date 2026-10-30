package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
)

func flattenIPCollections(in []openapi.IpCollection) []map[string]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[string]interface{}

	for _, v := range in {
		m := make(map[string]interface{})
		m["ip_collection_id"] = v.GetId()
		m["display_name"] = v.GetDisplayName()
		m["ip_addresses"] = flattenStringSlice(v.IpCollectionAddresses)
		m["project_id"] = v.GetProjectId()

		out = append(out, m)
	}
	return out
}

func expandIPCollections(in []interface{}) []openapi.IpCollection {
	if len(in) == 0 {
		return nil
	}

	var out = make([]openapi.IpCollection, len(in))
	for i, v := range in {
		m := v.(map[strfmt.UUID]interface{})

		member := openapi.IpCollection{}

		if paramName, ok := m["ip_collection_id"].(string); ok {
			member.Id = NewNullableString(paramName)
		}

		if paramType, ok := m["display_name"].(string); ok {
			member.DisplayName = NewNullableString(paramType)
		}

		if paramValue, ok := m["ip_addresses"]; ok {
			n := paramValue.([]string)
			if len(n) == 0 {
				member.IpCollectionAddresses = nil
			} else {
				for _, x := range n {
					member.IpCollectionAddresses = append(member.IpCollectionAddresses, x)
				}
			}
		}

		if paramValue, ok := m["project_id"].(string); ok {
			member.ProjectId = NewNullableString(paramValue)
		}
		out[i] = member
	}
	return out
}
