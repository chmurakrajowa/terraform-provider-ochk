package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
)

func flattenPortsForwardingLists(in []*models.PortForwarding) []map[strfmt.UUID]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[strfmt.UUID]interface{}

	for _, v := range in {
		m := make(map[strfmt.UUID]interface{})
		m["port_forwarding_id"] = v.PortForwardingID
		m["display_name"] = v.Name
		m["floating_ip_id"] = v.FloatingIPID
		m["internal_port_id"] = v.InternalPortID
		out = append(out, m)
	}
	return out
}
