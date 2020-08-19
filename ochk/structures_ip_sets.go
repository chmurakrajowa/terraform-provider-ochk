package ochk

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

func flattenIPSetAddresses(in []*models.IPSetAddress) *schema.Set {
	out := &schema.Set{
		F: ipSetAddressHash,
	}

	for _, v := range in {
		m := make(map[string]interface{})
		m["id"] = int(v.IPSetAddressID)
		m["address"] = v.IPAddress

		out.Add(m)
	}
	return out
}

func expandIPSetAddresses(in []interface{}) []*models.IPSetAddress {
	if len(in) == 0 {
		return nil
	}

	var out = make([]*models.IPSetAddress, len(in))
	for i, v := range in {
		m := v.(map[string]interface{})

		address := &models.IPSetAddress{
			IPSetAddressID: int32(m["id"].(int)),
			IPAddress:      m["address"].(string),
		}

		out[i] = address
	}
	return out
}

func ipSetAddressHash(v interface{}) int {
	m := v.(map[string]interface{})

	return schema.HashInt(m["id"])
}
