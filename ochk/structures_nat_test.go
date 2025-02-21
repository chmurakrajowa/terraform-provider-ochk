package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenAutoNat(t *testing.T) {
	cases := []struct {
		expanded  []*models.NATRuleInstance
		flattened []map[strfmt.UUID]interface{}
	}{
		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []*models.NATRuleInstance{
				{
					NatType:     "AUTO",
					RuleID:      "e1675817-f1a1-45c1-988b-ec2f142867e0",
					DisplayName: "nat-test1",
					VirtualNetworkInstance: &models.VirtualNetworkInstance{
						VirtualNetworkID: "2ae951f1-5285-496c-b598-aabe1a792319",
					},
					Enabled: true,
					TierZeroRouter: &models.RouterInstance{
						RouterID: "547948e9-b67d-44d1-ad69-ae9b711e289c",
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
			expanded: []*models.NATRuleInstance{
				{
					NatType:     "AUTO",
					RuleID:      "e1675817-f1a1-45c1-988b-ec2f142867e0",
					DisplayName: "nat-test1",
					VirtualNetworkInstance: &models.VirtualNetworkInstance{
						VirtualNetworkID: "2ae951f1-5285-496c-b598-aabe1a792319",
					},
					Enabled: false,
					TierZeroRouter: &models.RouterInstance{
						RouterID: "547948e9-b67d-44d1-ad69-ae9b711e289c",
					},
				},
				{
					NatType:     "AUTO",
					RuleID:      "e1675817-f1a1-45c1-988b-ec2f142867e0",
					DisplayName: "nat-test2",
					VirtualNetworkInstance: &models.VirtualNetworkInstance{
						VirtualNetworkID: "2ae951f1-5285-496c-b598-aabe1a792319",
					},
					Enabled: false,
					TierZeroRouter: &models.RouterInstance{
						RouterID: "b0908315-4c61-4326-b18d-2d145e6937a3",
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
		expanded  []*models.NATRuleInstance
		flattened []map[strfmt.UUID]interface{}
	}{
		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []*models.NATRuleInstance{
				{
					NatType:     "MANUAL",
					RuleID:      strfmt.UUID("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					DisplayName: "nat-test1",
					Action:      models.NATRuleAction("DNAT"),
					Enabled:     true,
					TierZeroRouter: &models.RouterInstance{
						RouterID: strfmt.UUID("547948e9-b67d-44d1-ad69-ae9b711e289c"),
					},
					SourceNetwork:      "192.168.15.0/24",
					DestinationNetwork: "192.168.0.0/24",
				},
			},
			flattened: []map[strfmt.UUID]interface{}{
				{
					"manual_nat_id":       strfmt.UUID("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					"display_name":        "nat-test1",
					"action":              models.NATRuleAction("DNAT"),
					"enabled":             true,
					"vrf_id":              strfmt.UUID("547948e9-b67d-44d1-ad69-ae9b711e289c"),
					"source_network":      "192.168.15.0/24",
					"destination_network": "192.168.0.0/24",
				},
			},
		},
		{
			expanded: []*models.NATRuleInstance{
				{
					NatType:     "MANUAL",
					RuleID:      "e1675817-f1a1-45c1-988b-ec2f142867e0",
					DisplayName: "nat-test2",
					Action:      "DNAT",
					Enabled:     false,
					TierZeroRouter: &models.RouterInstance{
						RouterID: "b0908315-4c61-4326-b18d-2d145e6937a3",
					},
					SourceNetwork:      "192.168.0.0/24",
					DestinationNetwork: "192.168.1.0/24",
				},
				{
					NatType:     "MANUAL",
					RuleID:      "e1675817-f1a1-45c1-988b-ec2f142867e0",
					DisplayName: "nat-test1",
					Action:      "DNAT",
					Enabled:     true,
					TierZeroRouter: &models.RouterInstance{
						RouterID: "547948e9-b67d-44d1-ad69-ae9b711e289c",
					},
					SourceNetwork:      "192.168.15.0/24",
					DestinationNetwork: "192.168.0.0/24",
				},
			},
			flattened: []map[strfmt.UUID]interface{}{
				{
					"manual_nat_id":       strfmt.UUID("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					"display_name":        "nat-test2",
					"action":              models.NATRuleAction("DNAT"),
					"enabled":             false,
					"vrf_id":              strfmt.UUID("b0908315-4c61-4326-b18d-2d145e6937a3"),
					"source_network":      "192.168.0.0/24",
					"destination_network": "192.168.1.0/24",
				},
				{
					"manual_nat_id":       strfmt.UUID("e1675817-f1a1-45c1-988b-ec2f142867e0"),
					"display_name":        "nat-test1",
					"action":              models.NATRuleAction("DNAT"),
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
