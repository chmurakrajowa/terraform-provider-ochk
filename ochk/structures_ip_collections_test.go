package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenIPCollections(t *testing.T) {
	cases := []struct {
		expanded  []*models.IPCollection
		flattened []map[strfmt.UUID]interface{}
	}{
		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []*models.IPCollection{
				{
					ID:          "e1675817-f1a1-45c1-988b-ec2f142867e0",
					DisplayName: "test1",
					//FIXME comparing nested *schema.Set's gives false result
					//IPCollectionAddresses: transformInterfaceSliceToStringSlice(flattenStringSlice([]string{"192.168.0.1"}).List()),
					ProjectID: "791bf702-22fb-4c76-bebb-1fee7ee75607",
				},
				{
					ID:          "1c1f244a-6ec0-47c6-961a-f255d64e954b",
					DisplayName: "test2",
					//FIXME comparing nested *schema.Set's gives false result
					//IPCollectionAddresses: transformInterfaceSliceToStringSlice(flattenStringSlice([]string{"192.168.1.1"}).List()),
					ProjectID: "0d3f34d8-4364-45b7-ae69-0635fa438cab",
				},
			},
			flattened: []map[strfmt.UUID]interface{}{
				{
					"ip_collection_id": "e1675817-f1a1-45c1-988b-ec2f142867e0",
					"display_name":     "test1",
					//FIXME comparing nested *schema.Set's gives false result
					//"ip_addresses": transformInterfaceSliceToStringSlice(flattenStringSlice([]string{"192.168.0.1"}).List()),
					"project_id": "791bf702-22fb-4c76-bebb-1fee7ee75607",
				},
				{
					"ip_collection_id": "1c1f244a-6ec0-47c6-961a-f255d64e954b",
					"display_name":     "test2",
					//FIXME comparing nested *schema.Set's gives false result
					//"ip_addresses": transformInterfaceSliceToStringSlice(flattenStringSlice([]string{"192.168.1.1"}).List()),
					"project_id": "0d3f34d8-4364-45b7-ae69-0635fa438cab",
				},
			},
		},
	}
	for _, c := range cases {

		outExpanded := flattenIPCollections(c.expanded)
		flattenedType := mapSliceToInterfaceSlice(c.flattened)
		flattenedTypeExpanded := expandIPCollections(flattenedType)
		flattenedExpanded := flattenIPCollections(flattenedTypeExpanded)

		assert.EqualValues(t, outExpanded, flattenedExpanded, "Error matching output and flattened: %#v vs %#v", flattenedExpanded, outExpanded)
	}
}
