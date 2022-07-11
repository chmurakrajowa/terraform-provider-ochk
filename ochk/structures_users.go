package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func flattenUserInstancesFromIDs(m []*models.UserInstance) *schema.Set {
	s := &schema.Set{
		F: schema.HashString,
	}

	for _, v := range m {
		s.Add(v.UserID)
	}

	return s
}

//func expandUsersFromIDs(in []interface{}) []*models.UserInstance {
//	var out = make([]*models.UserInstance, len(in))
//	for i, v := range in {
//		user := &models.UserInstance{
//			UserID: v.(string),
//		}
//		out[i] = user
//	}
//	return out
//}
