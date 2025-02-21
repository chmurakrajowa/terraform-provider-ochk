package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
)

//todo check what with posistion?
//
//func flattenFirewallRulePosition(in *models) []map[string]interface{} {
//	if in == nil {
//		return nil
//	}
//
//	out := make([]map[string]interface{}, 0)
//	position := make(map[string]interface{})
//
//	position["rule_id"] = in.RuleID
//	position["revise_operation"] = in.ReviseOperation
//
//	out = append(out, position)
//
//	return out
//}
//
//func expandFirewallRulePosition(in []interface{}) *models.Position {
//	if len(in) == 0 {
//		return nil
//	}
//
//	position := in[0].(map[string]interface{})
//
//	return &models.Position{
//		RuleID:          position["rule_id"].(string),
//		ReviseOperation: position["revise_operation"].(string),
//	}
//}

func flattenFirewallEWRulesLists(in []*models.DfwRule) []map[string]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[string]interface{}

	for _, v := range in {
		m := make(map[string]interface{})
		m["firewall_ew_rule_id"] = v.RuleID
		m["display_name"] = v.DisplayName
		m["project_id"] = v.ProjectID
		out = append(out, m)
	}
	return out
}

func flattenFirewallSNRulesLists(in []*models.GfwRule) []map[string]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[string]interface{}

	for _, v := range in {
		m := make(map[string]interface{})
		m["firewall_sn_rule_id"] = v.RuleID
		m["display_name"] = v.DisplayName
		m["project_id"] = v.ProjectID
		out = append(out, m)
	}
	return out
}
