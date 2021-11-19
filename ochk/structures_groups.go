package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func flattenGrouoInstancesFromIDs(m []*models.GroupInstance) *schema.Set {
	s := &schema.Set{
		F: schema.HashString,
	}

	for _, v := range m {
		s.Add(v.GroupID)
	}

	return s
}

//func expandGroupInstancesFromIDs(in []interface{}) []*models.GroupInstance {
//	if len(in) == 0 {
//		return nil
//	}
//
//	var out = make([]*models.GroupInstance, len(in))
//
//	for i, v := range in {
//		userInstance := &models.GroupInstance{
//			GroupID: v.(string),
//		}
//
//		out[i] = userInstance
//	}
//
//	return out
//}
