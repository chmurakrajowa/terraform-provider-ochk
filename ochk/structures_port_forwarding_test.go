package ochk

import (
	"testing"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
)

func TestFlattenPortForwarding(t *testing.T) {
	floatingIPID := "f1afcad5-f63e-4d99-9669-6081135bffc2"
	internalPortID := "11752b30-9749-4534-8ec1-2b0e9670b92b"
	portForwardingID := "e1675817-f1a1-45c1-988b-ec2f142867e0"

	cases := []struct {
		name      string
		expanded  []openapi.PortForwarding
		flattened []map[strfmt.UUID]interface{}
	}{
		{
			name:      "nil input",
			expanded:  nil,
			flattened: nil,
		},
		{
			name: "port forwarding with default values",
			expanded: []openapi.PortForwarding{
				{
					FloatingIpId:     &floatingIPID,
					Name:             NewNullableString("port-fwd-test1"),
					InternalPortId:   &internalPortID,
					PortForwardingId: &portForwardingID,
				},
			},
			flattened: []map[strfmt.UUID]interface{}{
				{
					"floating_ip_id":      "f1afcad5-f63e-4d99-9669-6081135bffc2",
					"display_name":        "port-fwd-test1",
					"internal_port_id":    "11752b30-9749-4534-8ec1-2b0e9670b92b",
					"port_forwarding_id":  "e1675817-f1a1-45c1-988b-ec2f142867e0",
					"description":         "",
					"external_port":       int32(0),
					"external_port_range": "",
					"internal_ip_address": "",
					"internal_port":       int32(0),
					"internal_port_range": "",
					"protocol":            "",
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			expected := mapSliceToInterfaceSlice(c.flattened)
			actual := mapSliceToInterfaceSlice(
				flattenPortsForwardingLists(c.expanded),
			)

			assert.EqualValues(t, expected, actual)
		})
	}
}
