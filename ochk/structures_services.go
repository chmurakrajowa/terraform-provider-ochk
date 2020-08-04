package ochk

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/models"
)

func flattenServicesFromIDs(m []*models.ServiceInstance) *schema.Set {
	s := &schema.Set{}

	for _, v := range m {
		s.Add(v.ServiceID)
	}
	return s
}

func expandServicesFromIDs(l []interface{}) []*models.ServiceInstance {
	if len(l) == 0 {
		return nil
	}

	var m = make([]*models.ServiceInstance, len(l))

	for i, v := range l {
		service := &models.ServiceInstance{
			ServiceID: v.(string),
		}

		m[i] = service
	}

	return m
}
