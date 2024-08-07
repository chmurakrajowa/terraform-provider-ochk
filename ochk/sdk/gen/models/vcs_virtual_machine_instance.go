// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VcsVirtualMachineInstance VcsVirtualMachineInstance
//
// swagger:model VcsVirtualMachineInstance
type VcsVirtualMachineInstance struct {

	// additional virtual disk device collection
	AdditionalVirtualDiskDeviceCollection []*VirtualDiskDevice `json:"additionalVirtualDiskDeviceCollection"`

	// backup list collection
	BackupListCollection []*BackupList `json:"backupListCollection"`

	// cpu count
	CPUCount int32 `json:"cpuCount,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// creation date
	// Format: date-time
	CreationDate *strfmt.DateTime `json:"creationDate,omitempty"`

	// deployment instance
	DeploymentInstance *DeploymentInstance `json:"deploymentInstance,omitempty"`

	// deployment params
	DeploymentParams []*DeploymentParam `json:"deploymentParams"`

	// dns search suffix
	DNSSearchSuffix string `json:"dnsSearchSuffix,omitempty"`

	// dns suffix
	DNSSuffix string `json:"dnsSuffix,omitempty"`

	// encryption instance
	EncryptionInstance *EncryptionInstance `json:"encryptionInstance,omitempty"`

	// folder path
	FolderPath string `json:"folderPath,omitempty"`

	// initial password
	InitialPassword string `json:"initialPassword,omitempty"`

	// initial user name
	InitialUserName string `json:"initialUserName,omitempty"`

	// ip address
	IPAddress string `json:"ipAddress,omitempty"`

	// iso instance
	IsoInstance *IsoInstance `json:"isoInstance,omitempty"`

	// lic settings
	LicSettings *LicSettings `json:"licSettings,omitempty"`

	// memory size m b
	MemorySizeMB int32 `json:"memorySizeMB,omitempty"`

	// modification date
	// Format: date-time
	ModificationDate *strfmt.DateTime `json:"modificationDate,omitempty"`

	// modified by
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// os type
	// Enum: [LINUX WINDOWS]
	OsType string `json:"osType,omitempty"`

	// os virtual disk device
	OsVirtualDiskDevice *VirtualDiskDevice `json:"osVirtualDiskDevice,omitempty"`

	// ovf Ip configuration
	OvfIPConfiguration bool `json:"ovfIpConfiguration,omitempty"`

	// power state
	// Enum: [poweredOff poweredOn suspended]
	PowerState string `json:"powerState,omitempty"`

	// primary Dns address
	PrimaryDNSAddress string `json:"primaryDnsAddress,omitempty"`

	// primary wins address
	PrimaryWinsAddress string `json:"primaryWinsAddress,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`

	// read only
	ReadOnly bool `json:"readOnly,omitempty"`

	// secondary Dns address
	SecondaryDNSAddress string `json:"secondaryDnsAddress,omitempty"`

	// secondary wins address
	SecondaryWinsAddress string `json:"secondaryWinsAddress,omitempty"`

	// ssh key
	SSHKey string `json:"sshKey,omitempty"`

	// storage policy
	// Enum: [ENTERPRISE STANDARD_W1 STANDARD_W2]
	StoragePolicy string `json:"storagePolicy,omitempty"`

	// tags
	Tags []*Tag `json:"tags"`

	// virtual machine Id
	VirtualMachineID string `json:"virtualMachineId,omitempty"`

	// virtual machine name
	VirtualMachineName string `json:"virtualMachineName,omitempty"`

	// virtual network devices
	VirtualNetworkDevices []*VirtualNetworkDevice `json:"virtualNetworkDevices"`
}

// Validate validates this vcs virtual machine instance
func (m *VcsVirtualMachineInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalVirtualDiskDeviceCollection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackupListCollection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeploymentInstance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeploymentParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncryptionInstance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsoInstance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLicSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModificationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOsType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOsVirtualDiskDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePowerState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoragePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualNetworkDevices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VcsVirtualMachineInstance) validateAdditionalVirtualDiskDeviceCollection(formats strfmt.Registry) error {
	if swag.IsZero(m.AdditionalVirtualDiskDeviceCollection) { // not required
		return nil
	}

	for i := 0; i < len(m.AdditionalVirtualDiskDeviceCollection); i++ {
		if swag.IsZero(m.AdditionalVirtualDiskDeviceCollection[i]) { // not required
			continue
		}

		if m.AdditionalVirtualDiskDeviceCollection[i] != nil {
			if err := m.AdditionalVirtualDiskDeviceCollection[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additionalVirtualDiskDeviceCollection" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("additionalVirtualDiskDeviceCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VcsVirtualMachineInstance) validateBackupListCollection(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupListCollection) { // not required
		return nil
	}

	for i := 0; i < len(m.BackupListCollection); i++ {
		if swag.IsZero(m.BackupListCollection[i]) { // not required
			continue
		}

		if m.BackupListCollection[i] != nil {
			if err := m.BackupListCollection[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backupListCollection" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("backupListCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VcsVirtualMachineInstance) validateCreationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VcsVirtualMachineInstance) validateDeploymentInstance(formats strfmt.Registry) error {
	if swag.IsZero(m.DeploymentInstance) { // not required
		return nil
	}

	if m.DeploymentInstance != nil {
		if err := m.DeploymentInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deploymentInstance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deploymentInstance")
			}
			return err
		}
	}

	return nil
}

func (m *VcsVirtualMachineInstance) validateDeploymentParams(formats strfmt.Registry) error {
	if swag.IsZero(m.DeploymentParams) { // not required
		return nil
	}

	for i := 0; i < len(m.DeploymentParams); i++ {
		if swag.IsZero(m.DeploymentParams[i]) { // not required
			continue
		}

		if m.DeploymentParams[i] != nil {
			if err := m.DeploymentParams[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deploymentParams" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("deploymentParams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VcsVirtualMachineInstance) validateEncryptionInstance(formats strfmt.Registry) error {
	if swag.IsZero(m.EncryptionInstance) { // not required
		return nil
	}

	if m.EncryptionInstance != nil {
		if err := m.EncryptionInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encryptionInstance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("encryptionInstance")
			}
			return err
		}
	}

	return nil
}

func (m *VcsVirtualMachineInstance) validateIsoInstance(formats strfmt.Registry) error {
	if swag.IsZero(m.IsoInstance) { // not required
		return nil
	}

	if m.IsoInstance != nil {
		if err := m.IsoInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("isoInstance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("isoInstance")
			}
			return err
		}
	}

	return nil
}

func (m *VcsVirtualMachineInstance) validateLicSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.LicSettings) { // not required
		return nil
	}

	if m.LicSettings != nil {
		if err := m.LicSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("licSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("licSettings")
			}
			return err
		}
	}

	return nil
}

func (m *VcsVirtualMachineInstance) validateModificationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ModificationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("modificationDate", "body", "date-time", m.ModificationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var vcsVirtualMachineInstanceTypeOsTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LINUX","WINDOWS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vcsVirtualMachineInstanceTypeOsTypePropEnum = append(vcsVirtualMachineInstanceTypeOsTypePropEnum, v)
	}
}

const (

	// VcsVirtualMachineInstanceOsTypeLINUX captures enum value "LINUX"
	VcsVirtualMachineInstanceOsTypeLINUX string = "LINUX"

	// VcsVirtualMachineInstanceOsTypeWINDOWS captures enum value "WINDOWS"
	VcsVirtualMachineInstanceOsTypeWINDOWS string = "WINDOWS"
)

// prop value enum
func (m *VcsVirtualMachineInstance) validateOsTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, vcsVirtualMachineInstanceTypeOsTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VcsVirtualMachineInstance) validateOsType(formats strfmt.Registry) error {
	if swag.IsZero(m.OsType) { // not required
		return nil
	}

	// value enum
	if err := m.validateOsTypeEnum("osType", "body", m.OsType); err != nil {
		return err
	}

	return nil
}

func (m *VcsVirtualMachineInstance) validateOsVirtualDiskDevice(formats strfmt.Registry) error {
	if swag.IsZero(m.OsVirtualDiskDevice) { // not required
		return nil
	}

	if m.OsVirtualDiskDevice != nil {
		if err := m.OsVirtualDiskDevice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("osVirtualDiskDevice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("osVirtualDiskDevice")
			}
			return err
		}
	}

	return nil
}

var vcsVirtualMachineInstanceTypePowerStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["poweredOff","poweredOn","suspended"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vcsVirtualMachineInstanceTypePowerStatePropEnum = append(vcsVirtualMachineInstanceTypePowerStatePropEnum, v)
	}
}

const (

	// VcsVirtualMachineInstancePowerStatePoweredOff captures enum value "poweredOff"
	VcsVirtualMachineInstancePowerStatePoweredOff string = "poweredOff"

	// VcsVirtualMachineInstancePowerStatePoweredOn captures enum value "poweredOn"
	VcsVirtualMachineInstancePowerStatePoweredOn string = "poweredOn"

	// VcsVirtualMachineInstancePowerStateSuspended captures enum value "suspended"
	VcsVirtualMachineInstancePowerStateSuspended string = "suspended"
)

// prop value enum
func (m *VcsVirtualMachineInstance) validatePowerStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, vcsVirtualMachineInstanceTypePowerStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VcsVirtualMachineInstance) validatePowerState(formats strfmt.Registry) error {
	if swag.IsZero(m.PowerState) { // not required
		return nil
	}

	// value enum
	if err := m.validatePowerStateEnum("powerState", "body", m.PowerState); err != nil {
		return err
	}

	return nil
}

var vcsVirtualMachineInstanceTypeStoragePolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENTERPRISE","STANDARD_W1","STANDARD_W2"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vcsVirtualMachineInstanceTypeStoragePolicyPropEnum = append(vcsVirtualMachineInstanceTypeStoragePolicyPropEnum, v)
	}
}

const (

	// VcsVirtualMachineInstanceStoragePolicyENTERPRISE captures enum value "ENTERPRISE"
	VcsVirtualMachineInstanceStoragePolicyENTERPRISE string = "ENTERPRISE"

	// VcsVirtualMachineInstanceStoragePolicySTANDARDW1 captures enum value "STANDARD_W1"
	VcsVirtualMachineInstanceStoragePolicySTANDARDW1 string = "STANDARD_W1"

	// VcsVirtualMachineInstanceStoragePolicySTANDARDW2 captures enum value "STANDARD_W2"
	VcsVirtualMachineInstanceStoragePolicySTANDARDW2 string = "STANDARD_W2"
)

// prop value enum
func (m *VcsVirtualMachineInstance) validateStoragePolicyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, vcsVirtualMachineInstanceTypeStoragePolicyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VcsVirtualMachineInstance) validateStoragePolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.StoragePolicy) { // not required
		return nil
	}

	// value enum
	if err := m.validateStoragePolicyEnum("storagePolicy", "body", m.StoragePolicy); err != nil {
		return err
	}

	return nil
}

func (m *VcsVirtualMachineInstance) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VcsVirtualMachineInstance) validateVirtualNetworkDevices(formats strfmt.Registry) error {
	if swag.IsZero(m.VirtualNetworkDevices) { // not required
		return nil
	}

	for i := 0; i < len(m.VirtualNetworkDevices); i++ {
		if swag.IsZero(m.VirtualNetworkDevices[i]) { // not required
			continue
		}

		if m.VirtualNetworkDevices[i] != nil {
			if err := m.VirtualNetworkDevices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("virtualNetworkDevices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("virtualNetworkDevices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this vcs virtual machine instance based on the context it is used
func (m *VcsVirtualMachineInstance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdditionalVirtualDiskDeviceCollection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBackupListCollection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeploymentInstance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeploymentParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEncryptionInstance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsoInstance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLicSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOsVirtualDiskDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVirtualNetworkDevices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VcsVirtualMachineInstance) contextValidateAdditionalVirtualDiskDeviceCollection(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AdditionalVirtualDiskDeviceCollection); i++ {

		if m.AdditionalVirtualDiskDeviceCollection[i] != nil {

			if swag.IsZero(m.AdditionalVirtualDiskDeviceCollection[i]) { // not required
				return nil
			}

			if err := m.AdditionalVirtualDiskDeviceCollection[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additionalVirtualDiskDeviceCollection" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("additionalVirtualDiskDeviceCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VcsVirtualMachineInstance) contextValidateBackupListCollection(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BackupListCollection); i++ {

		if m.BackupListCollection[i] != nil {

			if swag.IsZero(m.BackupListCollection[i]) { // not required
				return nil
			}

			if err := m.BackupListCollection[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backupListCollection" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("backupListCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VcsVirtualMachineInstance) contextValidateDeploymentInstance(ctx context.Context, formats strfmt.Registry) error {

	if m.DeploymentInstance != nil {

		if swag.IsZero(m.DeploymentInstance) { // not required
			return nil
		}

		if err := m.DeploymentInstance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deploymentInstance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deploymentInstance")
			}
			return err
		}
	}

	return nil
}

func (m *VcsVirtualMachineInstance) contextValidateDeploymentParams(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DeploymentParams); i++ {

		if m.DeploymentParams[i] != nil {

			if swag.IsZero(m.DeploymentParams[i]) { // not required
				return nil
			}

			if err := m.DeploymentParams[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deploymentParams" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("deploymentParams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VcsVirtualMachineInstance) contextValidateEncryptionInstance(ctx context.Context, formats strfmt.Registry) error {

	if m.EncryptionInstance != nil {

		if swag.IsZero(m.EncryptionInstance) { // not required
			return nil
		}

		if err := m.EncryptionInstance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encryptionInstance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("encryptionInstance")
			}
			return err
		}
	}

	return nil
}

func (m *VcsVirtualMachineInstance) contextValidateIsoInstance(ctx context.Context, formats strfmt.Registry) error {

	if m.IsoInstance != nil {

		if swag.IsZero(m.IsoInstance) { // not required
			return nil
		}

		if err := m.IsoInstance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("isoInstance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("isoInstance")
			}
			return err
		}
	}

	return nil
}

func (m *VcsVirtualMachineInstance) contextValidateLicSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.LicSettings != nil {

		if swag.IsZero(m.LicSettings) { // not required
			return nil
		}

		if err := m.LicSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("licSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("licSettings")
			}
			return err
		}
	}

	return nil
}

func (m *VcsVirtualMachineInstance) contextValidateOsVirtualDiskDevice(ctx context.Context, formats strfmt.Registry) error {

	if m.OsVirtualDiskDevice != nil {

		if swag.IsZero(m.OsVirtualDiskDevice) { // not required
			return nil
		}

		if err := m.OsVirtualDiskDevice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("osVirtualDiskDevice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("osVirtualDiskDevice")
			}
			return err
		}
	}

	return nil
}

func (m *VcsVirtualMachineInstance) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {

			if swag.IsZero(m.Tags[i]) { // not required
				return nil
			}

			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VcsVirtualMachineInstance) contextValidateVirtualNetworkDevices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VirtualNetworkDevices); i++ {

		if m.VirtualNetworkDevices[i] != nil {

			if swag.IsZero(m.VirtualNetworkDevices[i]) { // not required
				return nil
			}

			if err := m.VirtualNetworkDevices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("virtualNetworkDevices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("virtualNetworkDevices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VcsVirtualMachineInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VcsVirtualMachineInstance) UnmarshalBinary(b []byte) error {
	var res VcsVirtualMachineInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
