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

// VirtualNetworkUpdateResponse VirtualNetworkUpdateResponse
//
// swagger:model VirtualNetworkUpdateResponse
type VirtualNetworkUpdateResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// request instance
	RequestInstance *RequestInstance `json:"requestInstance,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`

	// virtual network instance
	VirtualNetworkInstance *VirtualNetworkInstance `json:"virtualNetworkInstance,omitempty"`
}

// Validate validates this virtual network update response
func (m *VirtualNetworkUpdateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequestInstance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualNetworkInstance(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualNetworkUpdateResponse) validateRequestInstance(formats strfmt.Registry) error {
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

func (m *VirtualNetworkUpdateResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VirtualNetworkUpdateResponse) validateVirtualNetworkInstance(formats strfmt.Registry) error {
	if swag.IsZero(m.VirtualNetworkInstance) { // not required
		return nil
	}

	if m.VirtualNetworkInstance != nil {
		if err := m.VirtualNetworkInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtualNetworkInstance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("virtualNetworkInstance")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this virtual network update response based on the context it is used
func (m *VirtualNetworkUpdateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRequestInstance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVirtualNetworkInstance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualNetworkUpdateResponse) contextValidateRequestInstance(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VirtualNetworkUpdateResponse) contextValidateVirtualNetworkInstance(ctx context.Context, formats strfmt.Registry) error {

	if m.VirtualNetworkInstance != nil {

		if swag.IsZero(m.VirtualNetworkInstance) { // not required
			return nil
		}

		if err := m.VirtualNetworkInstance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtualNetworkInstance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("virtualNetworkInstance")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VirtualNetworkUpdateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualNetworkUpdateResponse) UnmarshalBinary(b []byte) error {
	var res VirtualNetworkUpdateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
