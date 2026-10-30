package ochk

import (
	"testing"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
)

func TestFlattenExpandCustomServicePorts(t *testing.T) {
	cases := []struct {
		name      string
		expanded  []openapi.L4PortSetEntry
		flattened []map[string]interface{}
	}{
		{
			name:      "nil input",
			expanded:  nil,
			flattened: nil,
		},
		{
			name: "TCP and UDP ports",
			expanded: []openapi.L4PortSetEntry{
				{
					L4Protocol: openapi.L4Protocol("TCP").Ptr(),
					SourcePorts: transformInterfaceSliceToStringSlice(
						flattenStringSlice([]string{"1", "2", "3"}).List(),
					),
					DestinationPorts: transformInterfaceSliceToStringSlice(
						flattenStringSlice([]string{"3", "4", "5"}).List(),
					),
				},
				{
					L4Protocol: openapi.L4Protocol("UDP").Ptr(),
					SourcePorts: transformInterfaceSliceToStringSlice(
						flattenStringSlice([]string{"11", "22", "33"}).List(),
					),
					DestinationPorts: transformInterfaceSliceToStringSlice(
						flattenStringSlice([]string{"33", "44", "55"}).List(),
					),
				},
			},
			flattened: []map[string]interface{}{
				{
					"id":          "e1675817-f1a1-45c1-988b-ec2f142867aa",
					"protocol":    "TCP",
					"source":      flattenStringSlice([]string{"1", "2", "3"}),
					"destination": flattenStringSlice([]string{"3", "4", "5"}),
				},
				{
					"id":          "21675817-f1a1-45c1-988b-ec2f142867aa",
					"protocol":    "UDP",
					"source":      flattenStringSlice([]string{"11", "22", "33"}),
					"destination": flattenStringSlice([]string{"33", "44", "55"}),
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := mapSliceToInterfaceSliceStr(c.flattened)
			actual := expandCustomServicePorts(input)

			assert.EqualValues(t, c.expanded, actual)
		})
	}
}

func TestFlattenCustomServices(t *testing.T) {
	cases := []struct {
		name      string
		expanded  []openapi.CustomServiceInstance
		flattened []map[strfmt.UUID]interface{}
	}{
		{
			name:      "nil input",
			expanded:  nil,
			flattened: nil,
		},
		{
			name: "two custom services",
			expanded: []openapi.CustomServiceInstance{
				{
					ServiceId:   NewNullableString("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					DisplayName: NewNullableString("test1"),
					ProjectId:   NewNullableString("aa675817-f1a1-45c1-988b-ec2f142867e1"),
				},
				{
					ServiceId:   NewNullableString("791bf702-22fb-4c76-bebb-1fee7ee75607"),
					DisplayName: NewNullableString("test2"),
					ProjectId:   NewNullableString("aa675817-f1a1-45c1-988b-ec2f142867e1"),
				},
			},
			flattened: []map[strfmt.UUID]interface{}{
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
		t.Run(c.name, func(t *testing.T) {
			expected := mapSliceToInterfaceSlice(c.flattened)
			actual := mapSliceToInterfaceSlice(flattenCustomServices(c.expanded))

			assert.EqualValues(t, expected, actual)
		})
	}
}
