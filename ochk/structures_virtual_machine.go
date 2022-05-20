package ochk

import (
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func virtualDiskHash(v interface{}) int {
	m := v.(map[string]interface{})

	return schema.HashString(fmt.Sprintf("%d%d", m["controller_id"], m["lun_id"]))
}

func flattenDeploymentParams(in []*models.DeploymentParam) []map[string]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[string]interface{}

	for _, v := range in {
		m := make(map[string]interface{})
		m["param_name"] = v.ParamName
		m["param_type"] = v.ParamType
		m["param_value"] = v.ParamValue

		out = append(out, m)
	}
	return out
}

func expandVDeploymentParams(in []interface{}) []*models.DeploymentParam {
	var out = make([]*models.DeploymentParam, len(in))
	for i, v := range in {
		m := v.(map[string]interface{})

		member := &models.DeploymentParam{}

		if paramName, ok := m["param_name"].(string); ok {
			member.ParamName = paramName
		}
		if paramType, ok := m["param_type"].(string); ok {
			member.ParamType = paramType
		}
		if paramValue, ok := m["param_value"].(string); ok {
			member.ParamValue = paramValue
		}
		out[i] = member
	}
	return out
}

func flattenVirtualDisks(in []*models.VirtualDiskDevice) *schema.Set {
	if len(in) == 0 {
		return nil
	}

	out := &schema.Set{
		F: virtualDiskHash,
	}

	for _, v := range in {
		m := make(map[string]interface{})
		m["controller_id"] = int(v.ControllerID)
		m["lun_id"] = int(v.LunID)
		m["size_mb"] = int(v.SizeMB)
		m["device_type"] = v.VirtualDiskDeviceType

		out.Add(m)
	}
	return out
}

func expandVirtualDisks(in []interface{}) []*models.VirtualDiskDevice {
	if len(in) == 0 {
		return nil
	}

	var out = make([]*models.VirtualDiskDevice, len(in))
	for i, v := range in {
		m := v.(map[string]interface{})

		member := &models.VirtualDiskDevice{}

		if controllerID, ok := m["controller_id"].(int); ok {
			member.ControllerID = int32(controllerID)
		}

		if lunID, ok := m["lun_id"].(int); ok {
			member.LunID = int32(lunID)
		}

		if sizeMB, ok := m["size_mb"].(int); ok {
			member.SizeMB = int64(sizeMB)
		}

		if deviceType, ok := m["device_type"].(string); ok {
			member.VirtualDiskDeviceType = deviceType
		}

		out[i] = member
	}

	return out
}

func flattenVirtualNetworkDevice(in []*models.VirtualNetworkDevice) []map[string]interface{} {
	var out []map[string]interface{}
	for _, v := range in {
		m := make(map[string]interface{})
		m["device_id"] = v.DeviceID
		if v.VirtualNetworkInstance != nil {
			m["virtual_network_id"] = v.VirtualNetworkInstance.VirtualNetworkID
		}

		out = append(out, m)
	}

	return out
}

func expandVirtualNetworkDevices(in []interface{}) []*models.VirtualNetworkDevice {
	if len(in) == 0 {
		return nil
	}

	var out = make([]*models.VirtualNetworkDevice, len(in))
	for i, v := range in {
		m := v.(map[string]interface{})

		member := &models.VirtualNetworkDevice{}

		if virtualNetworkID, ok := m["virtual_network_id"].(string); ok && virtualNetworkID != "" {
			member.VirtualNetworkInstance = &models.VirtualNetworkInstance{
				VirtualNetworkID: virtualNetworkID,
			}
		}

		if deviceID, ok := m["device_id"].(string); ok && deviceID != "" {
			member.DeviceID = deviceID
		}

		out[i] = member
	}

	return out
}

func flattenBackupListsFromIDs(m []*models.BackupList) *schema.Set {
	s := &schema.Set{
		F: schema.HashString,
	}

	for _, v := range m {
		s.Add(v.BackupListID)
	}

	return s
}

func expandBackupListsFromIDs(in []interface{}) []*models.BackupList {
	if len(in) == 0 {
		return nil
	}

	var out = make([]*models.BackupList, len(in))

	for i, v := range in {
		BackupListInstance := &models.BackupList{
			BackupListID: v.(string),
		}

		out[i] = BackupListInstance
	}
	return out
}

func flattenSystemTagsListsFromIDs(m []*models.SystemTag) *schema.Set {
	s := &schema.Set{
		F: schema.HashString,
	}

	for _, v := range m {
		s.Add(fmt.Sprint(v.SystemTagID))
	}

	return s
}

func expandSystemTagsListsFromIDs(in []interface{}) []*models.SystemTag {
	if len(in) == 0 {
		return nil
	}

	var out = make([]*models.SystemTag, len(in))

	for i, v := range in {

		var systemTagIDInt32 int32
		fmt.Sscan(v.(string), &systemTagIDInt32)

		SystemTagInstance := &models.SystemTag{
			SystemTagID: systemTagIDInt32,
		}

		out[i] = SystemTagInstance
	}
	return out
}

func flattenBillingTagsListsFromIDs(m []*models.BillingTag) *schema.Set {
	s := &schema.Set{
		F: schema.HashString,
	}

	for _, v := range m {
		s.Add(fmt.Sprint(v.BillingTagID))
	}

	return s
}

func expandBillingTagsListsFromIDs(in []interface{}) []*models.BillingTag {
	if len(in) == 0 {
		return nil
	}

	var out = make([]*models.BillingTag, len(in))

	for i, v := range in {

		var billingTagIDInt32 int32
		fmt.Sscan(v.(string), &billingTagIDInt32)

		BillingTagInstance := &models.BillingTag{
			BillingTagID: billingTagIDInt32,
		}

		out[i] = BillingTagInstance
	}
	return out
}
