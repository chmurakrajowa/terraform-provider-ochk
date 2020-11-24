package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenExpandCustomServicePorts(t *testing.T) {
	cases := []struct {
		expanded  []*models.L4PortSetEntry
		flattened []map[string]interface{}
	}{

		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},

		{
			expanded: []*models.L4PortSetEntry{
				{
					L4PortSetEntryID: "id",
					L4Protocol:       "protocol",
					SourcePorts:      transformInterfaceSliceToStringSlice(flattenStringSlice([]string{"1", "2", "3"}).List()),
					DestinationPorts: transformInterfaceSliceToStringSlice(flattenStringSlice([]string{"3", "4", "5"}).List()),
				},
				{
					L4PortSetEntryID: "id2",
					L4Protocol:       "protocol2",
					SourcePorts:      transformInterfaceSliceToStringSlice(flattenStringSlice([]string{"11", "22", "33"}).List()),
					DestinationPorts: transformInterfaceSliceToStringSlice(flattenStringSlice([]string{"33", "44", "55"}).List()),
				},
			},
			flattened: []map[string]interface{}{
				{
					"id":          "id",
					"protocol":    "protocol",
					"source":      flattenStringSlice([]string{"1", "2", "3"}),
					"destination": flattenStringSlice([]string{"3", "4", "5"}),
				},
				{
					"id":          "id2",
					"protocol":    "protocol2",
					"source":      flattenStringSlice([]string{"11", "22", "33"}),
					"destination": flattenStringSlice([]string{"33", "44", "55"}),
				},
			},
		},
	}

	for _, c := range cases {
		//FIXME comparing nested *schema.Set's gives false result because reflect.DeepEquals cannot compare type.Func
		//flattenedList := mapSliceToInterfaceSlice(c.flattened)
		//outFlattened := flattenCustomServicePorts(c.expanded)
		//assert.EqualValues(t, flattenedList, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, flattenedList)

		flattenedInterfaceSlice := mapSliceToInterfaceSlice(c.flattened)
		outExpanded := expandCustomServicePorts(flattenedInterfaceSlice)
		assert.EqualValues(t, c.expanded, outExpanded, "Error matching output and expanded: %#v vs %#v", outExpanded, c.expanded)
	}
}
