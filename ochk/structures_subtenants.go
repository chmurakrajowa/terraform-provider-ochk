package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func flattenVCSNetworkInstancesFromIDs(m []*models.VCSNetworkInstance) *schema.Set {
	s := &schema.Set{
		F: schema.HashString,
	}

	for _, v := range m {
		s.Add(v.NetworkID)
	}

	return s
}

func expandVCSNetworkInstancesFromIDs(in []interface{}) []*models.VCSNetworkInstance {
	if len(in) == 0 {
		return nil
	}

	var out = make([]*models.VCSNetworkInstance, len(in))

	for i, v := range in {
		userInstance := &models.VCSNetworkInstance{
			NetworkID: v.(string),
		}

		out[i] = userInstance
	}

	return out
}
