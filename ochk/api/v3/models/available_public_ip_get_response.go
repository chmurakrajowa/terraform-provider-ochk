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

// AvailablePublicIPGetResponse available public Ip get response
//
// swagger:model AvailablePublicIpGetResponse
type AvailablePublicIPGetResponse struct {

	// public Ip address
	PublicIPAddress *PublicIPAddress `json:"publicIpAddress,omitempty"`
}

// Validate validates this available public Ip get response
func (m *AvailablePublicIPGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePublicIPAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AvailablePublicIPGetResponse) validatePublicIPAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.PublicIPAddress) { // not required
		return nil
	}

	if m.PublicIPAddress != nil {
		if err := m.PublicIPAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("publicIpAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("publicIpAddress")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this available public Ip get response based on the context it is used
func (m *AvailablePublicIPGetResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePublicIPAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AvailablePublicIPGetResponse) contextValidatePublicIPAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.PublicIPAddress != nil {

		if swag.IsZero(m.PublicIPAddress) { // not required
			return nil
		}

		if err := m.PublicIPAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("publicIpAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("publicIpAddress")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AvailablePublicIPGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AvailablePublicIPGetResponse) UnmarshalBinary(b []byte) error {
	var res AvailablePublicIPGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
