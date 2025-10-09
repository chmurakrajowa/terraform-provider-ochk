package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/openapi/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenExpandVirtualDisks(t *testing.T) {
	cases := []struct {
		expanded  []openapi.VirtualDiskDevice
		flattened []map[string]interface{}
	}{
		{
			expanded: []openapi.VirtualDiskDevice{
				{
					ControllerId:          1,
					LunId:                 2,
					SizeMB:                3,
					VirtualDiskDeviceType: openapi.VirtualDiskDeviceType("IDE"),
				},
			},
			flattened: []map[string]interface{}{
				{
					"controller_id": 1,
					"lun_id":        2,
					"size_mb":       3,
					"device_type":   openapi.VirtualDiskDeviceType("IDE"),
				},
			},
		},
		{
			expanded: []openapi.VirtualDiskDevice{
				{
					ControllerId:          1,
					LunId:                 2,
					SizeMB:                3,
					VirtualDiskDeviceType: "IDE",
				},
				{
					ControllerId:          11,
					LunId:                 22,
					SizeMB:                33,
					VirtualDiskDeviceType: openapi.VirtualDiskDeviceType("IDE2"),
				},
			},
			flattened: []map[string]interface{}{
				{
					"controller_id": 1,
					"lun_id":        2,
					"size_mb":       3,
					"device_type":   openapi.VirtualDiskDeviceType("IDE"),
				},
				{
					"controller_id": 11,
					"lun_id":        22,
					"size_mb":       33,
					"device_type":   openapi.VirtualDiskDeviceType("IDE2"),
				},
			},
		},
	}

	for _, c := range cases {
		//flattenedSetType := schema.NewSet(virtualDiskHash, mapSliceToInterfaceSlice(c.flattened)).List()
		//outFlattened := flattenVirtualDisks(c.expanded).List()
		//assert.EqualValues(t, flattenedSetType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, c.flattened)

		flattenedInterfaceSlice := mapSliceToInterfaceSliceStr(c.flattened)
		outExpanded := expandVirtualDisks(flattenedInterfaceSlice)
		assert.EqualValues(t, c.expanded, outExpanded, "Error matching output and expanded: %#v vs %#v", outExpanded, c.expanded)
	}

}

func TestFlattenExpandVirtualNetworkDevices(t *testing.T) {
	cases := []struct {
		expanded  []openapi.VirtualNetworkDevice
		flattened []map[string]interface{}
	}{
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []openapi.VirtualNetworkDevice{
				{
					DeviceId:               NewNullableString("123"),
					VirtualNetworkInstance: openapi.VirtualNetworkInstance{VirtualNetworkId: NewNullableString("vnet-id")},
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
			expanded: []openapi.VirtualNetworkDevice{
				{
					DeviceId:               NewNullableString("123"),
					VirtualNetworkInstance: openapi.VirtualNetworkInstance{VirtualNetworkId: NewNullableString("vnet-id")},
				},
				{
					DeviceId:               NewNullableString("1234"),
					VirtualNetworkInstance: openapi.VirtualNetworkInstance{VirtualNetworkId: NewNullableString("vnet-id2")},
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

		flattenedInterfaceSlice := mapSliceToInterfaceSliceStr(c.flattened)
		outExpanded := expandVirtualNetworkDevices(flattenedInterfaceSlice)
		assert.EqualValues(t, c.expanded, outExpanded, "Error matching output and expanded: %#v vs %#v", outExpanded, c.expanded)
	}
}
