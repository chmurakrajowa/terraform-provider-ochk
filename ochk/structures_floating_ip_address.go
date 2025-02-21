package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
)

func flattenFloatingIPAddressList(in []*models.FloatingIP) []map[strfmt.UUID]interface{} {

	var out []map[strfmt.UUID]interface{}

	for _, v := range in {
		m := make(map[strfmt.UUID]interface{})
		m["floating_ip_id"] = v.FloatingIPID
		m["display_name"] = v.Name
		m["public_address"] = v.PublicAddress
		out = append(out, m)
	}
	return out
}
