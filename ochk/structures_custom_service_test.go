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

func TestFlattenCustomServices(t *testing.T) {
	cases := []struct {
		expanded  []*models.CustomServiceInstance
		flattened []map[string]interface{}
	}{
		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []*models.CustomServiceInstance{
				{
					ServiceID:   "e1675817-f1a1-45c1-988b-ec2f142867e0",
					DisplayName: "test1",
					ProjectID:   "aa675817-f1a1-45c1-988b-ec2f142867e1",
				},
				{
					ServiceID:   "791bf702-22fb-4c76-bebb-1fee7ee75607",
					DisplayName: "test2",
					ProjectID:   "aa675817-f1a1-45c1-988b-ec2f142867e1",
				},
			},
			flattened: []map[string]interface{}{
				{
					"custom_service_id": "e1675817-f1a1-45c1-988b-ec2f142867e0",
					"display_name":      "test1",
					"project_id":        "aa675817-f1a1-45c1-988b-ec2f142867e1",
				},
				{
					"custom_service_id": "791bf702-22fb-4c76-bebb-1fee7ee75607",
					"display_name":      "test2",
					"project_id":        "aa675817-f1a1-45c1-988b-ec2f142867e1",
				},
			},
		},
	}
	for _, c := range cases {
		flattenedType := mapSliceToInterfaceSlice(c.flattened)
		outFlattened := mapSliceToInterfaceSlice(flattenCustomServices(c.expanded))
		assert.EqualValues(t, flattenedType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, flattenedType)
	}
}
