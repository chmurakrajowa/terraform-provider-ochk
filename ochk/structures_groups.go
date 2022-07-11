package ochk

//func flattenGroupInstancesFromIDs(m []*models.GroupInstance) *schema.Set {
//	s := &schema.Set{
//		F: schema.HashString,
//	}
//
//	for _, v := range m {
//		s.Add(v.GroupID)
//	}
//
//	return s
//}

//func expandGroupsFromIDs(in []interface{}) []*models.GroupInstance {
//	var out = make([]*models.GroupInstance, len(in))
//	for i, v := range in {
//		group := &models.GroupInstance{
//			GroupID: v.(string),
//		}
//		out[i] = group
//	}
//	return out
//}
