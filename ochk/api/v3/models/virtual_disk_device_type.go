// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// VirtualDiskDeviceType virtual disk device type
//
// swagger:model VirtualDiskDeviceType
type VirtualDiskDeviceType string

func NewVirtualDiskDeviceType(value VirtualDiskDeviceType) *VirtualDiskDeviceType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VirtualDiskDeviceType.
func (m VirtualDiskDeviceType) Pointer() *VirtualDiskDeviceType {
	return &m
}

const (

	// VirtualDiskDeviceTypeSCSI captures enum value "SCSI"
	VirtualDiskDeviceTypeSCSI VirtualDiskDeviceType = "SCSI"

	// VirtualDiskDeviceTypeIDE captures enum value "IDE"
	VirtualDiskDeviceTypeIDE VirtualDiskDeviceType = "IDE"

	// VirtualDiskDeviceTypeSATA captures enum value "SATA"
	VirtualDiskDeviceTypeSATA VirtualDiskDeviceType = "SATA"
)

// for schema
var virtualDiskDeviceTypeEnum []interface{}

func init() {
	var res []VirtualDiskDeviceType
	if err := json.Unmarshal([]byte(`["SCSI","IDE","SATA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		virtualDiskDeviceTypeEnum = append(virtualDiskDeviceTypeEnum, v)
	}
}

func (m VirtualDiskDeviceType) validateVirtualDiskDeviceTypeEnum(path, location string, value VirtualDiskDeviceType) error {
	if err := validate.EnumCase(path, location, value, virtualDiskDeviceTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this virtual disk device type
func (m VirtualDiskDeviceType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVirtualDiskDeviceTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this virtual disk device type based on context it is used
func (m VirtualDiskDeviceType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
