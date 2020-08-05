package ochk

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/models"
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
