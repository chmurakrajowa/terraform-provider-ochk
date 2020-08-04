package ochk

import (
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/models"
)

func flattenFirewallRulePosition(in *models.Position) interface{} {
	if in == nil {
		return nil
	}

	out := make(map[string]interface{})

	out["rule_id"] = in.RuleID
	out["revise_operation"] = in.ReviseOperation

	return out
}

func expandFirewallRulePosition(in interface{}) *models.Position {
	if in == nil {
		return nil
	}

	m := in.(map[string]interface{})

	return &models.Position{
		RuleID:          m["rule_id"].(string),
		ReviseOperation: m["revise_operation"].(string),
	}
}
