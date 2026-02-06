package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/openapi/v3"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func flattenRouterInstancesFromScope(m []openapi.RouterInstance) string {
	for _, v := range m {
		return v.GetRouterId()
	}
	return ""
}

func flattenRouterInstancesFromIDs(m []openapi.RouterInstance) *schema.Set {
	s := &schema.Set{
		F: schema.HashString,
	}

	for _, v := range m {
		s.Add(v.GetRouterId())
	}
	return s
}

func expandRouterInstancesFromIDs(in []interface{}) []openapi.RouterInstance {
	if len(in) == 0 {
		return nil
	}

	var out = make([]openapi.RouterInstance, len(in))

	for i, v := range in {
		securityGroup := openapi.RouterInstance{
			RouterId: NewNullableString(v.(string)),
		}

		out[i] = securityGroup
	}

	return out
}
