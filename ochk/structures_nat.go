package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
)

func flattenAutoNats(in []openapi.NATRuleInstance) []map[strfmt.UUID]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[strfmt.UUID]interface{}

	for _, v := range in {
		if *v.NatType == "AUTO" {
			m := make(map[strfmt.UUID]interface{})
			m["auto_nat_id"] = v.GetRuleId()
			m["display_name"] = v.GetDisplayName()
			m["virtual_network_id"] = v.GetVirtualNetworkId()
			m["enabled"] = v.GetEnabled()
			m["vrf_id"] = v.GetTierZeroRouterId()
			out = append(out, m)
		}
	}
	return out
}

func flattenManualNats(in []openapi.NATRuleInstance) []map[strfmt.UUID]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[strfmt.UUID]interface{}

	for _, v := range in {
		if *v.NatType == "MANUAL" {
			m := make(map[strfmt.UUID]interface{})
			m["manual_nat_id"] = v.GetRuleId()
			m["display_name"] = v.GetDisplayName()
			m["action"] = v.GetAction()
			m["enabled"] = v.GetEnabled()
			m["vrf_id"] = v.GetTierZeroRouterId()
			m["source_network"] = v.GetSourceNetwork()
			m["destination_network"] = v.GetDestinationNetwork()
			out = append(out, m)
		}
	}
	return out
}
