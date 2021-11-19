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

// VirtualMachineListResponse VirtualMachineListResponse
//
// swagger:model VirtualMachineListResponse
type VirtualMachineListResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`

	// virtual machine collection
	VirtualMachineCollection []*VirtualMachine `json:"virtualMachineCollection"`
}

// Validate validates this virtual machine list response
func (m *VirtualMachineListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualMachineCollection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualMachineListResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VirtualMachineListResponse) validateVirtualMachineCollection(formats strfmt.Registry) error {
	if swag.IsZero(m.VirtualMachineCollection) { // not required
		return nil
	}

	for i := 0; i < len(m.VirtualMachineCollection); i++ {
		if swag.IsZero(m.VirtualMachineCollection[i]) { // not required
			continue
		}

		if m.VirtualMachineCollection[i] != nil {
			if err := m.VirtualMachineCollection[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("virtualMachineCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this virtual machine list response based on the context it is used
func (m *VirtualMachineListResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVirtualMachineCollection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualMachineListResponse) contextValidateVirtualMachineCollection(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VirtualMachineCollection); i++ {

		if m.VirtualMachineCollection[i] != nil {
			if err := m.VirtualMachineCollection[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("virtualMachineCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VirtualMachineListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualMachineListResponse) UnmarshalBinary(b []byte) error {
	var res VirtualMachineListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
