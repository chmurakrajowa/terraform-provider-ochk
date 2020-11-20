package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenExpandVirtualDisks(t *testing.T) {
	cases := []struct {
		expanded  []*models.VirtualDiskDevice
		flattened []map[string]interface{}
	}{
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []*models.VirtualDiskDevice{
				{
					ControllerID:          1,
					LunID:                 2,
					SizeMB:                3,
					VirtualDiskDeviceType: "IDE",
				},
			},
			flattened: []map[string]interface{}{
				{
					"controller_id": 1,
					"lun_id":        2,
					"size_mb":       3,
					"device_type":   "IDE",
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
			flattened: []map[string]interface{}{
				{
					"controller_id": 1,
					"lun_id":        2,
					"size_mb":       3,
					"device_type":   "IDE",
				},
				{
					"controller_id": 11,
					"lun_id":        22,
					"size_mb":       33,
					"device_type":   "IDE2",
				},
			},
		},
	}

	for _, c := range cases {
		flattenedSetType := schema.NewSet(virtualDiskHash, mapSliceToInterfaceSlice(c.flattened)).List()
		outFlattened := flattenVirtualDisks(c.expanded).List()
		assert.EqualValues(t, flattenedSetType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, c.flattened)

		flattenedInterfaceSlice := mapSliceToInterfaceSlice(c.flattened)
		outExpanded := expandVirtualDisks(flattenedInterfaceSlice)
		assert.EqualValues(t, c.expanded, outExpanded, "Error matching output and expanded: %#v vs %#v", outExpanded, c.expanded)
	}
}

func TestFlattenExpandVirtualNetworkDevices(t *testing.T) {
	cases := []struct {
		expanded  []*models.VirtualNetworkDevice
		flattened []map[string]interface{}
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
			flattened: []map[string]interface{}{
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
			flattened: []map[string]interface{}{
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
		flattenedType := mapSliceToInterfaceSlice(c.flattened)
		outFlattened := mapSliceToInterfaceSlice(flattenVirtualNetworkDevice(c.expanded))
		assert.EqualValues(t, flattenedType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, flattenedType)

		flattenedInterfaceSlice := mapSliceToInterfaceSlice(c.flattened)
		outExpanded := expandVirtualNetworkDevices(flattenedInterfaceSlice)
		assert.EqualValues(t, c.expanded, outExpanded, "Error matching output and expanded: %#v vs %#v", outExpanded, c.expanded)
	}
}
