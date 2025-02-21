package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenPortForwarding(t *testing.T) {
	cases := []struct {
		expanded  []*models.PortForwarding
		flattened []map[strfmt.UUID]interface{}
	}{
		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []*models.PortForwarding{
				{
					FloatingIPID:     "f1afcad5-f63e-4d99-9669-6081135bffc2",
					Name:             "port-fwd-test1",
					InternalPortID:   "11752b30-9749-4534-8ec1-2b0e9670b92b",
					PortForwardingID: "e1675817-f1a1-45c1-988b-ec2f142867e0",
				},
			},
			flattened: []map[strfmt.UUID]interface{}{
				{
					"floating_ip_id":     strfmt.UUID("f1afcad5-f63e-4d99-9669-6081135bffc2"),
					"display_name":       "port-fwd-test1",
					"internal_port_id":   strfmt.UUID("11752b30-9749-4534-8ec1-2b0e9670b92b"),
					"port_forwarding_id": strfmt.UUID("e1675817-f1a1-45c1-988b-ec2f142867e0"),
				},
			},
		},
	}
	for _, c := range cases {
		flattenedType := mapSliceToInterfaceSlice(c.flattened)
		outFlattened := mapSliceToInterfaceSlice(flattenPortsForwardingLists(c.expanded))

		assert.EqualValues(t, flattenedType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, flattenedType)
	}
}
