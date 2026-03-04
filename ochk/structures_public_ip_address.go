package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
)

func flattenPublicIPAddress(in []openapi.PublicIpAllocation) []map[strfmt.UUID]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[strfmt.UUID]interface{}

	for _, v := range in {
		m := make(map[strfmt.UUID]interface{})
		m["public_ip_address_id"] = v.PublicIpAddress.GetIpAddressId()
		m["display_name"] = v.GetName()
		m["ip_address"] = v.PublicIpAddress.GetIpAddress()
		out = append(out, m)
	}
	return out
}
