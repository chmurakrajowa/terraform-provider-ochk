package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirewallOscRule(t *testing.T) {
	cases := []struct {
		expanded  []*models.FirewallRule
		flattened []map[strfmt.UUID]interface{}
	}{
		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []*models.FirewallRule{
				{
					RuleID: "c3514ee8-ee12-42b8-8dae-6f97c5271721",
					Name:   "test1",
				},
			},
			flattened: []map[strfmt.UUID]interface{}{
				{
					"rule_id": strfmt.UUID("c3514ee8-ee12-42b8-8dae-6f97c5271721"),
					"name":    "test1",
				},
			},
		},
	}
	for _, c := range cases {
		flattenedType := mapSliceToInterfaceSlice(c.flattened)
		outFlattened := mapSliceToInterfaceSlice(flattenFirewallRulesLists(c.expanded))
		assert.EqualValues(t, flattenedType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, flattenedType)
	}
}
