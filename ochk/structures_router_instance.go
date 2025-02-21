package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func flattenRouterInstancesFromScope(m []*models.RouterInstance) string {
	for _, v := range m {
		return v.RouterID.String()
	}
	return ""
}

func flattenRouterInstancesFromIDs(m []*models.RouterInstance) *schema.Set {
	s := &schema.Set{
		F: schema.HashString,
	}

	for _, v := range m {
		s.Add(v.RouterID.String())
	}
	return s
}

func expandRouterInstancesFromIDs(in []interface{}) []*models.RouterInstance {
	if len(in) == 0 {
		return nil
	}

	var out = make([]*models.RouterInstance, len(in))

	for i, v := range in {
		securityGroup := &models.RouterInstance{
			RouterID: v.(strfmt.UUID),
		}

		out[i] = securityGroup
	}

	return out
}
