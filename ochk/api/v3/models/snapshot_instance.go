// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SnapshotInstance snapshot instance
//
// swagger:model SnapshotInstance
type SnapshotInstance struct {

	// child snapshots
	ChildSnapshots []*SnapshotInstance `json:"childSnapshots"`

	// create time
	// Format: date-time
	CreateTime strfmt.DateTime `json:"createTime,omitempty"`

	// current
	Current bool `json:"current,omitempty"`

	// inner ID
	InnerID int32 `json:"innerID,omitempty"`

	// parent snapshot Id
	// Format: uuid
	ParentSnapshotID strfmt.UUID `json:"parentSnapshotId,omitempty"`

	// power state
	PowerState PowerState `json:"powerState,omitempty"`

	// quiesced
	Quiesced bool `json:"quiesced,omitempty"`

	// snapshot description
	SnapshotDescription string `json:"snapshotDescription,omitempty"`

	// snapshot Id
	// Format: uuid
	SnapshotID strfmt.UUID `json:"snapshotId,omitempty"`

	// snapshot mo ref
	SnapshotMoRef string `json:"snapshotMoRef,omitempty"`

	// snapshot name
	SnapshotName string `json:"snapshotName,omitempty"`

	// virtual machine Id
	// Format: uuid
	VirtualMachineID strfmt.UUID `json:"virtualMachineId,omitempty"`

	// virtual machine mo ref
	VirtualMachineMoRef string `json:"virtualMachineMoRef,omitempty"`

	// virtual machine name
	VirtualMachineName string `json:"virtualMachineName,omitempty"`
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

	if err := m.validateParentSnapshotID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePowerState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualMachineID(formats); err != nil {
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("childSnapshots" + "." + strconv.Itoa(i))
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

func (m *SnapshotInstance) validateParentSnapshotID(formats strfmt.Registry) error {
	if swag.IsZero(m.ParentSnapshotID) { // not required
		return nil
	}

	if err := validate.FormatOf("parentSnapshotId", "body", "uuid", m.ParentSnapshotID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotInstance) validatePowerState(formats strfmt.Registry) error {
	if swag.IsZero(m.PowerState) { // not required
		return nil
	}

	if err := m.PowerState.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("powerState")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("powerState")
		}
		return err
	}

	return nil
}

func (m *SnapshotInstance) validateSnapshotID(formats strfmt.Registry) error {
	if swag.IsZero(m.SnapshotID) { // not required
		return nil
	}

	if err := validate.FormatOf("snapshotId", "body", "uuid", m.SnapshotID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotInstance) validateVirtualMachineID(formats strfmt.Registry) error {
	if swag.IsZero(m.VirtualMachineID) { // not required
		return nil
	}

	if err := validate.FormatOf("virtualMachineId", "body", "uuid", m.VirtualMachineID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this snapshot instance based on the context it is used
func (m *SnapshotInstance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChildSnapshots(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePowerState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotInstance) contextValidateChildSnapshots(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ChildSnapshots); i++ {

		if m.ChildSnapshots[i] != nil {

			if swag.IsZero(m.ChildSnapshots[i]) { // not required
				return nil
			}

			if err := m.ChildSnapshots[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("childSnapshots" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("childSnapshots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnapshotInstance) contextValidatePowerState(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.PowerState) { // not required
		return nil
	}

	if err := m.PowerState.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("powerState")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("powerState")
		}
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
