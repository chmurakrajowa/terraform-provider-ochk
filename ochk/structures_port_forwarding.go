package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
)

func flattenPortsForwardingLists(in []openapi.PortForwarding) []map[strfmt.UUID]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[strfmt.UUID]interface{}

	for _, v := range in {
		m := make(map[strfmt.UUID]interface{})
		m["port_forwarding_id"] = v.GetPortForwardingId()
		m["display_name"] = v.GetName()
		m["floating_ip_id"] = v.GetFloatingIpId()
		m["internal_port_id"] = v.GetInternalPortId()
		out = append(out, m)
	}
	return out
}
