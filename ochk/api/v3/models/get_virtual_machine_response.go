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

// GetVirtualMachineResponse get virtual machine response
//
// swagger:model GetVirtualMachineResponse
type GetVirtualMachineResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`

	// vcs virtual machine instance
	VcsVirtualMachineInstance *VirtualMachineInstance `json:"vcsVirtualMachineInstance,omitempty"`
}

// Validate validates this get virtual machine response
func (m *GetVirtualMachineResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcsVirtualMachineInstance(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetVirtualMachineResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetVirtualMachineResponse) validateVcsVirtualMachineInstance(formats strfmt.Registry) error {
	if swag.IsZero(m.VcsVirtualMachineInstance) { // not required
		return nil
	}

	if m.VcsVirtualMachineInstance != nil {
		if err := m.VcsVirtualMachineInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vcsVirtualMachineInstance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vcsVirtualMachineInstance")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get virtual machine response based on the context it is used
func (m *GetVirtualMachineResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVcsVirtualMachineInstance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetVirtualMachineResponse) contextValidateVcsVirtualMachineInstance(ctx context.Context, formats strfmt.Registry) error {

	if m.VcsVirtualMachineInstance != nil {

		if swag.IsZero(m.VcsVirtualMachineInstance) { // not required
			return nil
		}

		if err := m.VcsVirtualMachineInstance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vcsVirtualMachineInstance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vcsVirtualMachineInstance")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetVirtualMachineResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetVirtualMachineResponse) UnmarshalBinary(b []byte) error {
	var res GetVirtualMachineResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
