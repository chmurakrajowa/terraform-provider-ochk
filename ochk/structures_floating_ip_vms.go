package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
)

func flattenFloatingIPVmsList(in []openapi.PortFwdVm) []map[strfmt.UUID]interface{} {

	var out []map[strfmt.UUID]interface{}

	for _, v := range in {
		m := make(map[strfmt.UUID]interface{})
		m["virtual_machine_id"] = v.GetVirtualMachineId()
		m["virtual_machine_name"] = v.GetVirtualMachineName()
		m["osc_port_id"] = v.GetOscPortId()
		m["ip_address"] = v.GetIpAddress()
		m["mac_address"] = v.GetMacAddress()
		m["network_name"] = v.GetNetworkName()
		out = append(out, m)
	}
	return out
}
