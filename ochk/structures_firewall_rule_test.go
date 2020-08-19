package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenExpandFirewallRulePosition(t *testing.T) {
	cases := []struct {
		expanded  *models.Position
		flattened map[string]interface{}
	}{
		// nil
		{
			expanded:  nil,
			flattened: nil,
		},

		// full object
		{
			expanded: &models.Position{
				RuleID:          "afdb07d8-d0d2-11ea-87d0-0242ac130003",
				ReviseOperation: "TOP",
			},
			flattened: map[string]interface{}{
				"rule_id":          "afdb07d8-d0d2-11ea-87d0-0242ac130003",
				"revise_operation": "TOP",
			},
		},
	}

	for _, c := range cases {
		flattenedSliceMap := mapToMapSlice(c.flattened)
		outFlattened := flattenFirewallRulePosition(c.expanded)
		assert.EqualValues(t, flattenedSliceMap, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, c.flattened)

		flattenedSlice := mapToInterfaceSlice(c.flattened)
		outExpanded := expandFirewallRulePosition(flattenedSlice)
		assert.EqualValues(t, c.expanded, outExpanded, "Error matching output and expanded: %#v vs %#v", outExpanded, c.expanded)
	}
}
