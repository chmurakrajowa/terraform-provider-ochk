// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VirtualNetworkDevice VirtualNetworkDevice
//
// swagger:model VirtualNetworkDevice
type VirtualNetworkDevice struct {

	// device Id
	DeviceID string `json:"deviceId,omitempty"`

	// virtual network instance
	VirtualNetworkInstance *VirtualNetworkInstance `json:"virtualNetworkInstance,omitempty"`
}

// Validate validates this virtual network device
func (m *VirtualNetworkDevice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVirtualNetworkInstance(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualNetworkDevice) validateVirtualNetworkInstance(formats strfmt.Registry) error {
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

// ContextValidate validate this virtual network device based on the context it is used
func (m *VirtualNetworkDevice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVirtualNetworkInstance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualNetworkDevice) contextValidateVirtualNetworkInstance(ctx context.Context, formats strfmt.Registry) error {

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
func (m *VirtualNetworkDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualNetworkDevice) UnmarshalBinary(b []byte) error {
	var res VirtualNetworkDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
