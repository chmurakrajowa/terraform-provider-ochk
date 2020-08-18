package ochk

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenExpandIPSetAddresses(t *testing.T) {
	cases := []struct {
		expanded  []*models.IPSetAddress
		flattened []map[string]interface{}
	}{

		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},

		// single address
		{
			expanded: []*models.IPSetAddress{
				{
					IPSetAddressID: 60,
					IPAddress:      "8.8.8.8/24",
				},
			},
			flattened: []map[string]interface{}{
				{
					"id":      60,
					"address": "8.8.8.8/24",
				},
			},
		},

		// multiple addresses
		{
			expanded: []*models.IPSetAddress{
				{
					IPSetAddressID: 60,
					IPAddress:      "8.8.8.8/24",
				},
				{
					IPSetAddressID: 62,
					IPAddress:      "11.11.11.11/24",
				},
			},
			flattened: []map[string]interface{}{
				{
					"id":      60,
					"address": "8.8.8.8/24",
				},
				{
					"id":      62,
					"address": "11.11.11.11/24",
				},
			},
		},
	}

	for _, c := range cases {
		flattenedSetType := schema.NewSet(ipSetAddressHash, mapSliceToInterfaceSlice(c.flattened)).List()
		outFlattened := flattenIPSetAddresses(c.expanded).List()
		assert.EqualValues(t, flattenedSetType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, c.flattened)

		flattenedInterfaceSlice := mapSliceToInterfaceSlice(c.flattened)
		outExpanded := expandIPSetAddresses(flattenedInterfaceSlice)
		assert.EqualValues(t, c.expanded, outExpanded, "Error matching output and expanded: %#v vs %#v", outExpanded, c.expanded)
	}
}
