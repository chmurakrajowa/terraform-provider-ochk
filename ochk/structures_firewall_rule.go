package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

func flattenFirewallRulePosition(in *models.Position) []map[string]interface{} {
	if in == nil {
		return nil
	}

	out := make([]map[string]interface{}, 1)
	position := make(map[string]interface{})

	position["rule_id"] = in.RuleID
	position["revise_operation"] = in.ReviseOperation

	out = append(out, position)

	return out
}

func expandFirewallRulePosition(in []interface{}) *models.Position {
	if len(in) == 0 {
		return nil
	}

	position := in[0].(map[string]interface{})

	return &models.Position{
		RuleID:          position["rule_id"].(string),
		ReviseOperation: position["revise_operation"].(string),
	}
}
