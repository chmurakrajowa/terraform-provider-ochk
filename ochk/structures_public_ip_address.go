package ochk

import "github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"

func flattenPublicIPAddress(in []*models.PublicIPAllocation) []map[string]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[string]interface{}

	for _, v := range in {
		m := make(map[string]interface{})
		m["public_ip_address_id"] = v.PublicIPAddress.IPAddressID
		m["display_name"] = v.Name
		m["ip_address"] = v.PublicIPAddress.IPAddress
		out = append(out, m)
	}
	return out
}
