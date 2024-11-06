package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
)

func flattenFloatingIPVmsList(in []*models.PortFwdVM) []map[strfmt.UUID]interface{} {

	var out []map[strfmt.UUID]interface{}

	for _, v := range in {
		m := make(map[strfmt.UUID]interface{})
		m["virtual_machine_id"] = v.VirtualMachineID
		m["virtual_machine_name"] = v.VirtualMachineName
		m["osc_port_id"] = v.OscPortID
		m["ip_address"] = v.IPAddress
		m["mac_address"] = v.MacAddress
		m["network_name"] = v.NetworkName
		out = append(out, m)
	}
	return out
}
