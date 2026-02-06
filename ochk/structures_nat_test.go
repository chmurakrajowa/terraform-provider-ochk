package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/openapi/v3"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//type NetworkNATType = openapi.NATType
//
//const (
//	NetworkManual NetworkNATType = "MANUAL"
//	NetworkAuto   NetworkNATType = "AUTO"
//)

func TestFlattenAutoNat(t *testing.T) {
	cases := []struct {
		expanded  []openapi.NATRuleInstance
		flattened []map[strfmt.UUID]interface{}
	}{
		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []openapi.NATRuleInstance{
				{
					NatType:     openapi.AUTO.Ptr(),
					RuleId:      NewNullableString("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					DisplayName: NewNullableString("nat-test1"),
					VirtualNetworkInstance: &openapi.VirtualNetworkInstance{
						VirtualNetworkId: NewNullableString("2ae951f1-5285-496c-b598-aabe1a792319"),
					},
					Enabled: NewNullableBool(true),
					TierZeroRouter: &openapi.RouterInstance{
						RouterId: NewNullableString("547948e9-b67d-44d1-ad69-ae9b711e289c"),
					},
				},
			},
			flattened: []map[strfmt.UUID]interface{}{
				{
					"auto_nat_id":        strfmt.UUID("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					"display_name":       "nat-test1",
					"virtual_network_id": strfmt.UUID("2ae951f1-5285-496c-b598-aabe1a792319"),
					"enabled":            true,
					"vrf_id":             strfmt.UUID("547948e9-b67d-44d1-ad69-ae9b711e289c"),
				},
			},
		},
		{
			expanded: []openapi.NATRuleInstance{
				{
					NatType:     openapi.AUTO.Ptr(),
					RuleId:      NewNullableString("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					DisplayName: NewNullableString("nat-test1"),
					VirtualNetworkInstance: &openapi.VirtualNetworkInstance{
						VirtualNetworkId: NewNullableString("2ae951f1-5285-496c-b598-aabe1a792319"),
					},
					Enabled: NewNullableBool(false),
					TierZeroRouter: &openapi.RouterInstance{
						RouterId: NewNullableString("547948e9-b67d-44d1-ad69-ae9b711e289c"),
					},
				},
				{
					NatType:     openapi.AUTO.Ptr(),
					RuleId:      NewNullableString("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					DisplayName: NewNullableString("nat-test2"),
					VirtualNetworkInstance: &openapi.VirtualNetworkInstance{
						VirtualNetworkId: NewNullableString("2ae951f1-5285-496c-b598-aabe1a792319"),
					},
					Enabled: NewNullableBool(false),
					TierZeroRouter: &openapi.RouterInstance{
						RouterId: NewNullableString("b0908315-4c61-4326-b18d-2d145e6937a3"),
					},
				},
			},
			flattened: []map[strfmt.UUID]interface{}{
				{
					"auto_nat_id":        strfmt.UUID("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					"display_name":       "nat-test1",
					"virtual_network_id": strfmt.UUID("2ae951f1-5285-496c-b598-aabe1a792319"),
					"enabled":            false,
					"vrf_id":             strfmt.UUID("547948e9-b67d-44d1-ad69-ae9b711e289c"),
				},
				{
					"auto_nat_id":        strfmt.UUID("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					"display_name":       "nat-test2",
					"virtual_network_id": strfmt.UUID("2ae951f1-5285-496c-b598-aabe1a792319"),
					"enabled":            false,
					"vrf_id":             strfmt.UUID("b0908315-4c61-4326-b18d-2d145e6937a3"),
				},
			},
		},
	}
	for _, c := range cases {
		flattenedType := mapSliceToInterfaceSlice(c.flattened)
		outFlattened := mapSliceToInterfaceSlice(flattenAutoNats(c.expanded))

		assert.EqualValues(t, flattenedType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, flattenedType)
	}
}

func TestFlattenManualNat(t *testing.T) {
	cases := []struct {
		expanded  []openapi.NATRuleInstance
		flattened []map[strfmt.UUID]interface{}
	}{
		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []openapi.NATRuleInstance{
				{
					NatType:     NetworkManual.Ptr(),
					RuleId:      NewNullableString("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					DisplayName: NewNullableString("nat-test1"),
					Action:      openapi.DNAT.Ptr(),
					Enabled:     NewNullableBool(true),
					TierZeroRouter: &openapi.RouterInstance{
						RouterId: NewNullableString("547948e9-b67d-44d1-ad69-ae9b711e289c"),
					},
					SourceNetwork:      NewNullableString("192.168.15.0/24"),
					DestinationNetwork: NewNullableString("192.168.0.0/24"),
				},
			},
			flattened: []map[strfmt.UUID]interface{}{
				{
					"manual_nat_id":       strfmt.UUID("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					"display_name":        "nat-test1",
					"action":              openapi.NATRuleAction("DNAT"),
					"enabled":             true,
					"vrf_id":              strfmt.UUID("547948e9-b67d-44d1-ad69-ae9b711e289c"),
					"source_network":      "192.168.15.0/24",
					"destination_network": "192.168.0.0/24",
				},
			},
		},
		{
			expanded: []openapi.NATRuleInstance{
				{
					NatType:     NetworkManual.Ptr(),
					RuleId:      NewNullableString("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					DisplayName: NewNullableString("nat-test2"),
					Action:      openapi.DNAT.Ptr(),
					Enabled:     NewNullableBool(false),
					TierZeroRouter: &openapi.RouterInstance{
						RouterId: NewNullableString("b0908315-4c61-4326-b18d-2d145e6937a3"),
					},
					SourceNetwork:      NewNullableString("192.168.0.0/24"),
					DestinationNetwork: NewNullableString("192.168.1.0/24"),
				},
				{
					NatType:     NetworkManual.Ptr(),
					RuleId:      NewNullableString("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					DisplayName: NewNullableString("nat-test1"),
					Action:      openapi.DNAT.Ptr(),
					Enabled:     NewNullableBool(true),
					TierZeroRouter: &openapi.RouterInstance{
						RouterId: NewNullableString("547948e9-b67d-44d1-ad69-ae9b711e289c"),
					},
					SourceNetwork:      NewNullableString("192.168.15.0/24"),
					DestinationNetwork: NewNullableString("192.168.0.0/24"),
				},
			},
			flattened: []map[strfmt.UUID]interface{}{
				{
					"manual_nat_id":       strfmt.UUID("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					"display_name":        "nat-test2",
					"action":              openapi.NATRuleAction("DNAT"),
					"enabled":             false,
					"vrf_id":              strfmt.UUID("b0908315-4c61-4326-b18d-2d145e6937a3"),
					"source_network":      "192.168.0.0/24",
					"destination_network": "192.168.1.0/24",
				},
				{
					"manual_nat_id":       strfmt.UUID("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					"display_name":        "nat-test1",
					"action":              openapi.NATRuleAction("DNAT"),
					"enabled":             true,
					"vrf_id":              strfmt.UUID("547948e9-b67d-44d1-ad69-ae9b711e289c"),
					"source_network":      "192.168.15.0/24",
					"destination_network": "192.168.0.0/24",
				},
			},
		},
	}
	for _, c := range cases {
		flattenedType := mapSliceToInterfaceSlice(c.flattened)
		outFlattened := mapSliceToInterfaceSlice(flattenManualNats(c.expanded))

		assert.EqualValues(t, flattenedType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, flattenedType)
	}
}
