package ochk

import (
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func flattenCustomServicePorts(in []openapi.L4PortSetEntry) []interface{} {
	if in == nil {
		return nil
	}

	out := make([]interface{}, 0)

	for _, v := range in {
		m := make(map[string]interface{})
		m["id"] = v.GetL4PortSetEntryId()
		m["protocol"] = v.L4Protocol
		m["source"] = flattenStringSlice(v.SourcePorts)
		m["destination"] = flattenStringSlice(v.DestinationPorts)
		out = append(out, m)
	}

	return out
}

func expandCustomServicePorts(in []interface{}) []openapi.L4PortSetEntry {
	if len(in) == 0 {
		return nil
	}
	var out = make([]openapi.L4PortSetEntry, len(in))
	for i, v := range in {
		m := v.(map[string]interface{})

		member := openapi.L4PortSetEntry{
			L4PortSetEntryId: openapi.NullableString{},
			L4Protocol:       castStringToL4ProtocolEnum(m["protocol"].(string)).Ptr(),
			SourcePorts:      transformSetToStringSlice(m["source"].(*schema.Set)),
			DestinationPorts: transformSetToStringSlice(m["destination"].(*schema.Set)),
		}

		out[i] = member
	}

	return out
}

func flattenCustomServicesFromIDs(in []openapi.CustomServiceInstance) *schema.Set {
	out := &schema.Set{
		F: schema.HashString,
	}

	for _, v := range in {
		out.Add(fmt.Sprint(v.GetServiceId()))
	}
	return out
}

func expandCustomServicesFromIDs(in []interface{}) []openapi.CustomServiceInstance {
	if len(in) == 0 {
		return nil
	}

	var out = make([]openapi.CustomServiceInstance, len(in))

	for i, v := range in {
		idValue := strfmt.UUID.String(strfmt.UUID(v.(string)))
		service := openapi.CustomServiceInstance{
			ServiceId: NewNullableString(idValue),
		}

		out[i] = service
	}

	return out
}

func flattenCustomServices(in []openapi.CustomServiceInstance) []map[strfmt.UUID]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[strfmt.UUID]interface{}

	for _, v := range in {
		m := make(map[strfmt.UUID]interface{})
		m["custom_service_id"] = v.GetServiceId()
		m["display_name"] = v.GetDisplayName()
		m["project_id"] = v.GetProjectId()
		out = append(out, m)
	}
	return out
}

type L4ProtocolType = openapi.L4Protocol

const (
	TCP L4ProtocolType = "TCP"
	UDP L4ProtocolType = "UDP"
)

func castStringToL4ProtocolEnum(e string) openapi.L4Protocol {
	switch e {
	case "TCP":
		return TCP
	case "UDP":
		return UDP
	default:
		return ""
	}
}
