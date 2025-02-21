package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
)

func flattenAutoNats(in []*models.NATRuleInstance) []map[strfmt.UUID]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[strfmt.UUID]interface{}

	for _, v := range in {
		if v.NatType == "AUTO" {
			m := make(map[strfmt.UUID]interface{})
			m["auto_nat_id"] = v.RuleID
			m["display_name"] = v.DisplayName
			m["virtual_network_id"] = v.VirtualNetworkInstance.VirtualNetworkID
			m["enabled"] = v.Enabled
			m["vrf_id"] = v.TierZeroRouter.RouterID
			out = append(out, m)
		}
	}
	return out
}

func flattenManualNats(in []*models.NATRuleInstance) []map[strfmt.UUID]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[strfmt.UUID]interface{}

	for _, v := range in {
		if v.NatType == "MANUAL" {
			m := make(map[strfmt.UUID]interface{})
			m["manual_nat_id"] = v.RuleID
			m["display_name"] = v.DisplayName
			m["action"] = v.Action
			m["enabled"] = v.Enabled
			m["vrf_id"] = v.TierZeroRouter.RouterID
			m["source_network"] = v.SourceNetwork
			m["destination_network"] = v.DestinationNetwork
			out = append(out, m)
		}
	}
	return out
}
