package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func flattenCustomServicePorts(in []*models.L4PortSetEntry) []interface{} {
	if in == nil {
		return nil
	}

	out := make([]interface{}, 0)

	for _, v := range in {
		m := make(map[string]interface{})
		m["id"] = v.L4PortSetEntryID
		m["protocol"] = v.L4Protocol
		m["source"] = flattenStringSlice(v.SourcePorts)
		m["destination"] = flattenStringSlice(v.DestinationPorts)
		out = append(out, m)
	}

	return out
}

func expandCustomServicePorts(in []interface{}) []*models.L4PortSetEntry {
	if len(in) == 0 {
		return nil
	}

	var out = make([]*models.L4PortSetEntry, len(in))
	for i, v := range in {
		m := v.(map[string]interface{})

		member := &models.L4PortSetEntry{
			L4PortSetEntryID: strfmt.UUID(m["id"].(string)),
			L4Protocol:       models.L4Protocol(m["protocol"].(string)),
			SourcePorts:      transformSetToStringSlice(m["source"].(*schema.Set)),
			DestinationPorts: transformSetToStringSlice(m["destination"].(*schema.Set)),
		}

		out[i] = member
	}

	return out
}

func flattenCustomServicesFromIDs(in []*models.CustomServiceInstance) *schema.Set {
	out := &schema.Set{
		F: schema.HashString,
	}

	for _, v := range in {
		out.Add(v.ServiceID)
	}
	return out
}

func expandCustomServicesFromIDs(in []interface{}) []*models.CustomServiceInstance {
	if len(in) == 0 {
		return nil
	}

	var out = make([]*models.CustomServiceInstance, len(in))

	for i, v := range in {
		service := &models.CustomServiceInstance{
			ServiceID: v.(strfmt.UUID),
		}

		out[i] = service
	}

	return out
}

func flattenCustomServices(in []*models.CustomServiceInstance) []map[strfmt.UUID]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[strfmt.UUID]interface{}

	for _, v := range in {
		m := make(map[strfmt.UUID]interface{})
		m["custom_service_id"] = v.ServiceID
		m["display_name"] = v.DisplayName
		m["project_id"] = v.ProjectID
		out = append(out, m)
	}
	return out
}
