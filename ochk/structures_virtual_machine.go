package ochk

import (
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func virtualDiskHash(v interface{}) int {
	m := v.(map[strfmt.UUID]interface{})
	return schema.HashString(fmt.Sprintf("%d%d", m["controller_id"], m["lun_id"]))
}

func flattenVirtualMachines(in []openapi.VirtualMachineInstance) []map[string]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[string]interface{}

	for _, v := range in {
		m := make(map[string]interface{})
		m["virtual_machine_id"] = v.GetVirtualMachineId()
		m["display_name"] = v.GetVirtualMachineName()
		m["folder_path"] = v.GetFolderPath()
		m["project_id"] = v.GetProjectId()
		out = append(out, m)
	}
	return out
}

//TODO check what is this
//
//func flattenDeploymentParams(in []*models.DeploymentParam) []map[string]interface{} {
//	if len(in) == 0 {
//		return nil
//	}
//
//	var out []map[string]interface{}
//
//	for _, v := range in {
//		m := make(map[string]interface{})
//		m["param_name"] = v.ParamName
//		m["param_type"] = v.ParamType
//		m["param_value"] = v.ParamValue
//
//		out = append(out, m)
//	}
//	return out
//}
//
//func expandVDeploymentParams(in []interface{}) []*models.DeploymentParam {
//	var out = make([]*models.DeploymentParam, len(in))
//	for i, v := range in {
//		m := v.(map[string]interface{})
//
//		member := &models.DeploymentParam{}
//
//		if paramName, ok := m["param_name"].(string); ok {
//			member.ParamName = paramName
//		}
//		if paramType, ok := m["param_type"].(string); ok {
//			member.ParamType = paramType
//		}
//		if paramValue, ok := m["param_value"].(string); ok {
//			member.ParamValue = paramValue
//		}
//		out[i] = member
//	}
//	return out
//}

func flattenVirtualDisks(in []openapi.VirtualDiskDevice) *schema.Set {
	if len(in) == 0 {
		return nil
	}

	out := &schema.Set{
		F: virtualDiskHash,
	}

	for _, v := range in {
		m := make(map[strfmt.UUID]interface{})
		m["controller_id"] = v.GetControllerId()
		m["lun_id"] = v.GetLunId()
		m["size_mb"] = v.GetSizeMB()
		m["device_type"] = v.GetVirtualDiskDeviceType()

		out.Add(m)
	}
	return out
}

func validateVirtualMachine(d *schema.ResourceData, platformType openapi.PlatformType) string {
	for i, v := range d.Get("additional_virtual_disks").(*schema.Set).List() {
		m := v.(map[string]interface{})
		if m["size_mb"].(int) < 1024 {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf(E1002, m["size_mb"])
		}
	}

	if platformType == openapi.PLATFORMTYPE_OPENSTACK {
		if len(d.Get("backup_lists").(*schema.Set).List()) > 0 {
			return fmt.Sprintf(E1003, "backup_lists", "backup_lists")
		}

		if len(d.Get("encryption_recrypt").(string)) > 0 {
			return fmt.Sprintf(E1004, "encryption_recrypt")
		}

	}
	return ""
}

func expandVirtualDisks(in []interface{}) []openapi.VirtualDiskDevice {
	if len(in) == 0 {
		return nil
	}
	var out = make([]openapi.VirtualDiskDevice, len(in))

	for i, v := range in {
		m := v.(map[string]interface{})

		//m := v.(map[strfmt.UUID]interface{})

		member := openapi.VirtualDiskDevice{}
		if controllerID, ok := m["controller_id"].(int); ok {
			controllerIDValue := int32(controllerID)
			member.ControllerId = &controllerIDValue
		}

		if lunID, ok := m["lun_id"].(int); ok {
			lunIdIntValue := int32(lunID)
			member.LunId = &lunIdIntValue
		}
		if sizeMB, ok := m["size_mb"].(int); ok {
			sizeMBValue := int64(sizeMB)

			member.SizeMB = &sizeMBValue
		}

		if deviceType, ok := m["device_type"].(openapi.VirtualDiskDeviceType); ok {
			member.VirtualDiskDeviceType = &deviceType
		}

		out[i] = member
	}

	return out
}

func flattenVirtualNetworkDevice(in []openapi.VirtualNetworkDevice) []map[strfmt.UUID]interface{} {
	var out []map[strfmt.UUID]interface{}
	for _, v := range in {
		m := make(map[strfmt.UUID]interface{})
		m["device_id"] = v.GetDeviceId()
		if v.VirtualNetworkInstance != nil {
			m["virtual_network_id"] = v.VirtualNetworkInstance.GetVirtualNetworkId()
		}

		out = append(out, m)
	}

	return out
}

func expandVirtualNetworkDevices(in []interface{}) []openapi.VirtualNetworkDevice {
	if len(in) == 0 {
		return nil
	}

	var out = make([]openapi.VirtualNetworkDevice, len(in))
	for i, v := range in {
		m := v.(map[string]interface{})
		//m := v.(map[strfmt.UUID]interface{})
		member := openapi.VirtualNetworkDevice{}

		if virtualNetworkID, ok := m["virtual_network_id"].(string); ok && virtualNetworkID != "" {
			member.VirtualNetworkInstance = &openapi.VirtualNetworkInstance{
				VirtualNetworkId: NewNullableString(virtualNetworkID),
			}
		}

		if deviceID, ok := m["device_id"].(string); ok && deviceID != "" {
			member.DeviceId = NewNullableString(deviceID)
		}

		out[i] = member
	}

	return out
}

func flattenBackupListsFromIDs(m []openapi.BackupList) *schema.Set {
	s := &schema.Set{
		F: schema.HashString,
	}

	for _, v := range m {
		s.Add(v.GetBackupListId())
	}

	return s
}

func expandBackupListsFromIDs(in []interface{}) []openapi.BackupList {
	if len(in) == 0 {
		return nil
	}
	var out = make([]openapi.BackupList, len(in))

	for i, v := range in {
		value := strfmt.UUID.String(strfmt.UUID(v.(string)))
		BackupListInstance := openapi.BackupList{
			BackupListId: NewNullableString(value),
		}

		out[i] = BackupListInstance
	}
	return out
}

func flattenTagsListsFromIDs(m []openapi.Tag) *schema.Set {
	s := &schema.Set{
		F: schema.HashString,
	}

	for _, v := range m {
		s.Add(fmt.Sprint(v.GetTagId()))
	}

	return s
}

func expandTagsListsFromIDs(in []interface{}) []openapi.Tag {

	var out = make([]openapi.Tag, len(in))

	for i, v := range in {

		var tagIDInt32 int32
		_, err := fmt.Sscan(v.(string), &tagIDInt32)
		if err != nil {
			return nil
		}

		TagInstance := openapi.Tag{
			TagId: &tagIDInt32,
		}

		out[i] = TagInstance
	}
	return out
}
