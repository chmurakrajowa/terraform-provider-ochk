// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateVirtualMachineResponse update virtual machine response
//
// swagger:model UpdateVirtualMachineResponse
type UpdateVirtualMachineResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// request instance
	RequestInstance *RequestInstance `json:"requestInstance,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`

	// virtual machine instance
	VirtualMachineInstance *VirtualMachineInstance `json:"virtualMachineInstance,omitempty"`
}

// Validate validates this update virtual machine response
func (m *UpdateVirtualMachineResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequestInstance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualMachineInstance(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateVirtualMachineResponse) validateRequestInstance(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestInstance) { // not required
		return nil
	}

	if m.RequestInstance != nil {
		if err := m.RequestInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requestInstance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("requestInstance")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateVirtualMachineResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpdateVirtualMachineResponse) validateVirtualMachineInstance(formats strfmt.Registry) error {
	if swag.IsZero(m.VirtualMachineInstance) { // not required
		return nil
	}

	if m.VirtualMachineInstance != nil {
		if err := m.VirtualMachineInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtualMachineInstance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("virtualMachineInstance")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update virtual machine response based on the context it is used
func (m *UpdateVirtualMachineResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRequestInstance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVirtualMachineInstance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateVirtualMachineResponse) contextValidateRequestInstance(ctx context.Context, formats strfmt.Registry) error {

	if m.RequestInstance != nil {

		if swag.IsZero(m.RequestInstance) { // not required
			return nil
		}

		if err := m.RequestInstance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requestInstance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("requestInstance")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateVirtualMachineResponse) contextValidateVirtualMachineInstance(ctx context.Context, formats strfmt.Registry) error {

	if m.VirtualMachineInstance != nil {

		if swag.IsZero(m.VirtualMachineInstance) { // not required
			return nil
		}

		if err := m.VirtualMachineInstance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtualMachineInstance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("virtualMachineInstance")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateVirtualMachineResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateVirtualMachineResponse) UnmarshalBinary(b []byte) error {
	var res UpdateVirtualMachineResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
