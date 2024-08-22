package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenExpandVirtualDisks(t *testing.T) {
	cases := []struct {
		expanded  []*models.VirtualDiskDevice
		flattened []map[strfmt.UUID]interface{}
	}{
		{
			expanded: []*models.VirtualDiskDevice{
				{
					ControllerID:          1,
					LunID:                 2,
					SizeMB:                3,
					VirtualDiskDeviceType: models.VirtualDiskDeviceType("IDE"),
				},
			},
			flattened: []map[strfmt.UUID]interface{}{
				{
					"controller_id": 1,
					"lun_id":        2,
					"size_mb":       3,
					"device_type":   models.VirtualDiskDeviceType("IDE"),
				},
			},
		},
		{
			expanded: []*models.VirtualDiskDevice{
				{
					ControllerID:          1,
					LunID:                 2,
					SizeMB:                3,
					VirtualDiskDeviceType: "IDE",
				},
				{
					ControllerID:          11,
					LunID:                 22,
					SizeMB:                33,
					VirtualDiskDeviceType: "IDE2",
				},
			},
			flattened: []map[strfmt.UUID]interface{}{
				{
					"controller_id": 1,
					"lun_id":        2,
					"size_mb":       3,
					"device_type":   models.VirtualDiskDeviceType("IDE"),
				},
				{
					"controller_id": 11,
					"lun_id":        22,
					"size_mb":       33,
					"device_type":   models.VirtualDiskDeviceType("IDE2"),
				},
			},
		},
	}

	for _, c := range cases {
		//flattenedSetType := schema.NewSet(virtualDiskHash, mapSliceToInterfaceSlice(c.flattened)).List()
		//outFlattened := flattenVirtualDisks(c.expanded).List()
		//assert.EqualValues(t, flattenedSetType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, c.flattened)

		flattenedInterfaceSlice := mapSliceToInterfaceSlice(c.flattened)
		outExpanded := expandVirtualDisks(flattenedInterfaceSlice)
		assert.EqualValues(t, c.expanded, outExpanded, "Error matching output and expanded: %#v vs %#v", outExpanded, c.expanded)
	}

}

func TestFlattenExpandVirtualNetworkDevices(t *testing.T) {
	cases := []struct {
		expanded  []*models.VirtualNetworkDevice
		flattened []map[strfmt.UUID]interface{}
	}{
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []*models.VirtualNetworkDevice{
				{
					DeviceID:               "123",
					VirtualNetworkInstance: &models.VirtualNetworkInstance{VirtualNetworkID: "vnet-id"},
				},
			},
			flattened: []map[strfmt.UUID]interface{}{
				{
					"device_id":          "123",
					"virtual_network_id": "vnet-id",
				},
			},
		},
		{
			expanded: []*models.VirtualNetworkDevice{
				{
					DeviceID:               "123",
					VirtualNetworkInstance: &models.VirtualNetworkInstance{VirtualNetworkID: "vnet-id"},
				},
				{
					DeviceID:               "1234",
					VirtualNetworkInstance: &models.VirtualNetworkInstance{VirtualNetworkID: "vnet-id2"},
				},
			},
			flattened: []map[strfmt.UUID]interface{}{
				{
					"device_id":          "123",
					"virtual_network_id": "vnet-id",
				},
				{
					"device_id":          "1234",
					"virtual_network_id": "vnet-id2",
				},
			},
		},
	}

	for _, c := range cases {

		flattenedInterfaceSlice := mapSliceToInterfaceSlice(c.flattened)
		outExpanded := expandVirtualNetworkDevices(flattenedInterfaceSlice)
		assert.EqualValues(t, c.expanded, outExpanded, "Error matching output and expanded: %#v vs %#v", outExpanded, c.expanded)
	}
}
