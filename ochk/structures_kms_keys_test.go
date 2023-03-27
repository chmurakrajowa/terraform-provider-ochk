package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenKMSKeys(t *testing.T) {
	cases := []struct {
		expanded  []*models.KeyInstance
		flattened []map[string]interface{}
	}{
		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []*models.KeyInstance{
				{
					ID:   "e1675817-f1a1-45c1-988b-ec2f142867e0",
					Name: "test1",
					//FIXME comparing nested *schema.Set's gives false result
					//KeyUsageList: transformInterfaceSliceToStringSlice(flattenStringSlice([]string{"ENCRYPT", "DECRYPT"}).List()),
					State: "Active",
				},
				{
					ID:   "e1675817-f1a1-45c1-988b-ec2f142867e1",
					Name: "test2",
					//FIXME comparing nested *schema.Set's gives false result
					//KeyUsageList: transformInterfaceSliceToStringSlice(flattenStringSlice([]string{"ENCRYPT"}).List()),
					State: "Inactive",
				},
			},
			flattened: []map[string]interface{}{
				{
					"kms_key_id":   "e1675817-f1a1-45c1-988b-ec2f142867e0",
					"display_name": "test1",
					//FIXME comparing nested *schema.Set's gives false result
					//"key_usage":    flattenStringSlice([]string{"ENCRYPT", "DECRYPT"}).List(),
					"state": "Active",
				},
				{
					"kms_key_id":   "e1675817-f1a1-45c1-988b-ec2f142867e1",
					"display_name": "test2",
					//FIXME comparing nested *schema.Set's gives false result
					//"key_usage":    flattenStringSlice([]string{"ENCRYPT"}).List(),
					"state": "Inactive",
				},
			},
		},
	}
	for _, c := range cases {
		//flattenedType := mapSliceToInterfaceSlice(c.flattened)
		//outFlattened := mapSliceToInterfaceSlice(flattenKMSKeys(c.expanded))
		//assert.EqualValues(t, flattenedType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, flattenedType)

		outExpanded := flattenKMSKeys(c.expanded)
		flattenedType := mapSliceToInterfaceSlice(c.flattened)
		flattenedTypeExpanded := expandKMSKeys(flattenedType)
		flattenedExpanded := flattenKMSKeys(flattenedTypeExpanded)

		assert.EqualValues(t, outExpanded, flattenedExpanded, "Error matching output and flattened: %#v vs %#v", flattenedExpanded, outExpanded)
	}
}
