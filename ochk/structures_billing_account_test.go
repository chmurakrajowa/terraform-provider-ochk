package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenExpandAccountProjects(t *testing.T) {
	cases := []struct {
		expanded  []*models.AccountProjectInstance
		flattened []map[string]interface{}
	}{
		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []*models.AccountProjectInstance{
				{
					ProjectID: "e1675817-f1a1-45c1-988b-ec2f142867aa",
				},
			},
			flattened: []map[string]interface{}{
				{
					"project_id": "e1675817-f1a1-45c1-988b-ec2f142867aa",
				},
			},
		},
	}
	for _, c := range cases {
		flattenedSetType := schema.NewSet(projectsHash, mapSliceToInterfaceSlice(c.flattened)).List()
		outFlattened := flattenAccProjects(c.expanded).List()
		assert.EqualValues(t, flattenedSetType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, c.flattened)

		flattenedInterfaceSlice := mapSliceToInterfaceSlice(c.flattened)
		outExpanded := expandAcctProjects(flattenedInterfaceSlice)
		assert.EqualValues(t, c.expanded, outExpanded, "Error matching output and expanded: %#v vs %#v", outExpanded, c.expanded)
	}
}
