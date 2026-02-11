package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func flattenServicesFromIDs(in []openapi.ServiceInstance) *schema.Set {
	out := &schema.Set{
		F: schema.HashString,
	}

	for _, v := range in {
		out.Add(v.GetServiceId())
	}
	return out
}

func expandServicesFromIDs(in []interface{}) []openapi.ServiceInstance {
	if len(in) == 0 {
		return nil
	}

	var out = make([]openapi.ServiceInstance, len(in))

	for i, v := range in {
		idValue := strfmt.UUID.String(strfmt.UUID(v.(string)))
		service := openapi.ServiceInstance{
			ServiceId: NewNullableString(idValue),
		}

		out[i] = service
	}

	return out
}

func flattenServices(in []openapi.ServiceInstance) []map[strfmt.UUID]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[strfmt.UUID]interface{}

	for _, v := range in {
		m := make(map[strfmt.UUID]interface{})
		m["service_id"] = v.ServiceId
		m["display_name"] = v.DisplayName
		out = append(out, m)
	}
	return out
}
