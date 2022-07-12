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

// VcsVirtualMachineUpdateResponse VcsVirtualMachineUpdateResponse
//
// swagger:model VcsVirtualMachineUpdateResponse
type VcsVirtualMachineUpdateResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// request instance
	RequestInstance *RequestInstance `json:"requestInstance,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`

	// vcs virtual machine instance
	VcsVirtualMachineInstance *VcsVirtualMachineInstance `json:"vcsVirtualMachineInstance,omitempty"`
}

// Validate validates this vcs virtual machine update response
func (m *VcsVirtualMachineUpdateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequestInstance(formats); err != nil {
		res = append(res, err)
	}

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

func (m *VcsVirtualMachineUpdateResponse) validateRequestInstance(formats strfmt.Registry) error {
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

func (m *VcsVirtualMachineUpdateResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VcsVirtualMachineUpdateResponse) validateVcsVirtualMachineInstance(formats strfmt.Registry) error {
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

// ContextValidate validate this vcs virtual machine update response based on the context it is used
func (m *VcsVirtualMachineUpdateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRequestInstance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVcsVirtualMachineInstance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VcsVirtualMachineUpdateResponse) contextValidateRequestInstance(ctx context.Context, formats strfmt.Registry) error {

	if m.RequestInstance != nil {
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

func (m *VcsVirtualMachineUpdateResponse) contextValidateVcsVirtualMachineInstance(ctx context.Context, formats strfmt.Registry) error {

	if m.VcsVirtualMachineInstance != nil {
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
func (m *VcsVirtualMachineUpdateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VcsVirtualMachineUpdateResponse) UnmarshalBinary(b []byte) error {
	var res VcsVirtualMachineUpdateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
