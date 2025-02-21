package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
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
					L4PortSetEntryID: "e1675817-f1a1-45c1-988b-ec2f142867aa",
					L4Protocol:       "protocol",
					SourcePorts:      transformInterfaceSliceToStringSlice(flattenStringSlice([]string{"1", "2", "3"}).List()),
					DestinationPorts: transformInterfaceSliceToStringSlice(flattenStringSlice([]string{"3", "4", "5"}).List()),
				},
				{
					L4PortSetEntryID: "21675817-f1a1-45c1-988b-ec2f142867aa",
					L4Protocol:       "protocol2",
					SourcePorts:      transformInterfaceSliceToStringSlice(flattenStringSlice([]string{"11", "22", "33"}).List()),
					DestinationPorts: transformInterfaceSliceToStringSlice(flattenStringSlice([]string{"33", "44", "55"}).List()),
				},
			},
			flattened: []map[string]interface{}{
				{
					"id":          "e1675817-f1a1-45c1-988b-ec2f142867aa",
					"protocol":    "protocol",
					"source":      flattenStringSlice([]string{"1", "2", "3"}),
					"destination": flattenStringSlice([]string{"3", "4", "5"}),
				},
				{
					"id":          "21675817-f1a1-45c1-988b-ec2f142867aa",
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

		flattenedInterfaceSlice := mapSliceToInterfaceSliceStr(c.flattened)
		outExpanded := expandCustomServicePorts(flattenedInterfaceSlice)
		assert.EqualValues(t, c.expanded, outExpanded, "Error matching output and expanded: %#v vs %#v", outExpanded, c.expanded)
	}
}

func TestFlattenCustomServices(t *testing.T) {
	cases := []struct {
		expanded  []*models.CustomServiceInstance
		flattened []map[strfmt.UUID]interface{}
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
			flattened: []map[strfmt.UUID]interface{}{
				{
					"custom_service_id": strfmt.UUID("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					"display_name":      "test1",
					"project_id":        strfmt.UUID("aa675817-f1a1-45c1-988b-ec2f142867e1"),
				},
				{
					"custom_service_id": strfmt.UUID("791bf702-22fb-4c76-bebb-1fee7ee75607"),
					"display_name":      "test2",
					"project_id":        strfmt.UUID("aa675817-f1a1-45c1-988b-ec2f142867e1"),
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
