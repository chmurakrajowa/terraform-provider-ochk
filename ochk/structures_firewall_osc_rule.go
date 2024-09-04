package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
)

func flattenFirewallRulesLists(in []*models.FirewallRule) []map[string]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[string]interface{}

	for _, v := range in {
		m := make(map[string]interface{})
		m["rule_id"] = v.RuleID
		m["name"] = v.Name
		m["project_id"] = v.ProjectExternalID
		m["description"] = v.Description
		m["direction"] = v.Direction
		m["etherType"] = v.EtherType
		m["port_range_min"] = v.PortRangeMin
		m["port_range_max"] = v.PortRangeMax
		m["protocol"] = v.Protocol
		m["security_group_id"] = v.SecurityGroup.ID
		m["remote_ip_prefix"] = v.RemoteIPPrefix
		out = append(out, m)
	}
	return out
}
