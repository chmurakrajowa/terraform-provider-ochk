package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/openapi/v3"
	"github.com/go-openapi/strfmt"
)

func flattenPortsForwardingLists(in []openapi.PortForwarding) []map[strfmt.UUID]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[strfmt.UUID]interface{}

	for _, v := range in {
		m := make(map[strfmt.UUID]interface{})
		m["port_forwarding_id"] = v.PortForwardingId
		m["display_name"] = v.Name
		m["floating_ip_id"] = v.FloatingIpId
		m["internal_port_id"] = v.InternalPortId
		out = append(out, m)
	}
	return out
}
