// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SnapshotInstance SnapshotInstance
//
// swagger:model SnapshotInstance
type SnapshotInstance struct {

	// child snapshots
	ChildSnapshots []*SnapshotInstance `json:"childSnapshots"`

	// create time
	// Format: date-time
	CreateTime *strfmt.DateTime `json:"createTime,omitempty"`

	// power state
	// Enum: [poweredOff poweredOn suspended]
	PowerState string `json:"powerState,omitempty"`

	// quiesced
	Quiesced bool `json:"quiesced,omitempty"`

	// snapshot description
	SnapshotDescription string `json:"snapshotDescription,omitempty"`

	// snapshot Id
	SnapshotID string `json:"snapshotId,omitempty"`

	// snapshot name
	SnapshotName string `json:"snapshotName,omitempty"`

	// virtual machine Id
	VirtualMachineID string `json:"virtualMachineId,omitempty"`
}

// Validate validates this snapshot instance
func (m *SnapshotInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChildSnapshots(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePowerState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotInstance) validateChildSnapshots(formats strfmt.Registry) error {

	if swag.IsZero(m.ChildSnapshots) { // not required
		return nil
	}

	for i := 0; i < len(m.ChildSnapshots); i++ {
		if swag.IsZero(m.ChildSnapshots[i]) { // not required
			continue
		}

		if m.ChildSnapshots[i] != nil {
			if err := m.ChildSnapshots[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("childSnapshots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnapshotInstance) validateCreateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("createTime", "body", "date-time", m.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var snapshotInstanceTypePowerStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["poweredOff","poweredOn","suspended"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		snapshotInstanceTypePowerStatePropEnum = append(snapshotInstanceTypePowerStatePropEnum, v)
	}
}

const (

	// SnapshotInstancePowerStatePoweredOff captures enum value "poweredOff"
	SnapshotInstancePowerStatePoweredOff string = "poweredOff"

	// SnapshotInstancePowerStatePoweredOn captures enum value "poweredOn"
	SnapshotInstancePowerStatePoweredOn string = "poweredOn"

	// SnapshotInstancePowerStateSuspended captures enum value "suspended"
	SnapshotInstancePowerStateSuspended string = "suspended"
)

// prop value enum
func (m *SnapshotInstance) validatePowerStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, snapshotInstanceTypePowerStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SnapshotInstance) validatePowerState(formats strfmt.Registry) error {

	if swag.IsZero(m.PowerState) { // not required
		return nil
	}

	// value enum
	if err := m.validatePowerStateEnum("powerState", "body", m.PowerState); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotInstance) UnmarshalBinary(b []byte) error {
	var res SnapshotInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
