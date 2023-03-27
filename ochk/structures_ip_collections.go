package ochk

import "github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"

func flattenIPCollections(in []*models.IPCollection) []map[string]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[string]interface{}

	for _, v := range in {
		m := make(map[string]interface{})
		m["ip_collection_id"] = v.ID
		m["display_name"] = v.DisplayName
		m["ip_addresses"] = flattenStringSlice(v.IPCollectionAddresses)
		m["project_id"] = v.ProjectID

		out = append(out, m)
	}
	return out
}

func expandIPCollections(in []interface{}) []*models.IPCollection {
	if len(in) == 0 {
		return nil
	}

	var out = make([]*models.IPCollection, len(in))
	for i, v := range in {
		m := v.(map[string]interface{})

		member := &models.IPCollection{}

		if paramName, ok := m["ip_collection_id"].(string); ok {
			member.ID = paramName
		}

		if paramType, ok := m["display_name"].(string); ok {
			member.DisplayName = paramType
		}

		if paramValue, ok := m["ip_addresses"]; ok {
			n := paramValue.([]string)
			if len(n) == 0 {
				member.IPCollectionAddresses = nil
			} else {
				for _, x := range n {
					member.IPCollectionAddresses = append(member.IPCollectionAddresses, x)
				}
			}
		}

		if paramValue, ok := m["project_id"].(string); ok {
			member.ProjectID = paramValue
		}
		out[i] = member
	}
	return out
}
