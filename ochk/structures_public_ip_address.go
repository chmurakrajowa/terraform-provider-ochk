package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
)

func flattenPublicIPAddress(in []*models.PublicIPAllocation) []map[strfmt.UUID]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[strfmt.UUID]interface{}

	for _, v := range in {
		m := make(map[strfmt.UUID]interface{})
		m["public_ip_address_id"] = v.PublicIPAddress.IPAddressID
		m["display_name"] = v.Name
		m["ip_address"] = v.PublicIPAddress.IPAddress
		out = append(out, m)
	}
	return out
}
