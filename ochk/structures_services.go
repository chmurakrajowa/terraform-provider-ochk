package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func flattenServicesFromIDs(in []*models.ServiceInstance) *schema.Set {
	out := &schema.Set{
		F: schema.HashString,
	}

	for _, v := range in {
		out.Add(v.ServiceID)
	}
	return out
}

func expandServicesFromIDs(in []interface{}) []*models.ServiceInstance {
	if len(in) == 0 {
		return nil
	}

	var out = make([]*models.ServiceInstance, len(in))

	for i, v := range in {
		service := &models.ServiceInstance{
			ServiceID: v.(string),
		}

		out[i] = service
	}

	return out
}

func flattenServices(in []*models.ServiceInstance) []map[string]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[string]interface{}

	for _, v := range in {
		m := make(map[string]interface{})
		m["service_id"] = v.ServiceID
		m["display_name"] = v.DisplayName
		out = append(out, m)
	}
	return out
}
