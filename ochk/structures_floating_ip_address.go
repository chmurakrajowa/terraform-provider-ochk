package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
)

func flattenFloatingIPAddressList(in []openapi.FloatingIp) []map[strfmt.UUID]interface{} {

	var out []map[strfmt.UUID]interface{}

	for _, v := range in {
		m := make(map[strfmt.UUID]interface{})
		m["floating_ip_id"] = v.GetFloatingIpId()
		m["display_name"] = v.GetName()
		m["public_address"] = v.GetPublicAddress()
		out = append(out, m)
	}
	return out
}
