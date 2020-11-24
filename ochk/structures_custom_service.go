package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
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
			L4PortSetEntryID: m["id"].(string),
			L4Protocol:       m["protocol"].(string),
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
			ServiceID: v.(string),
		}

		out[i] = service
	}

	return out
}
