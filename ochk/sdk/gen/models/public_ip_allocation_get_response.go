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

// PublicIPAllocationGetResponse PublicIpAllocationGetResponse
//
// swagger:model PublicIpAllocationGetResponse
type PublicIPAllocationGetResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// public Ip allocation
	PublicIPAllocation *PublicIPAllocation `json:"publicIpAllocation,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this public Ip allocation get response
func (m *PublicIPAllocationGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePublicIPAllocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicIPAllocationGetResponse) validatePublicIPAllocation(formats strfmt.Registry) error {
	if swag.IsZero(m.PublicIPAllocation) { // not required
		return nil
	}

	if m.PublicIPAllocation != nil {
		if err := m.PublicIPAllocation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("publicIpAllocation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("publicIpAllocation")
			}
			return err
		}
	}

	return nil
}

func (m *PublicIPAllocationGetResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this public Ip allocation get response based on the context it is used
func (m *PublicIPAllocationGetResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePublicIPAllocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicIPAllocationGetResponse) contextValidatePublicIPAllocation(ctx context.Context, formats strfmt.Registry) error {

	if m.PublicIPAllocation != nil {

		if swag.IsZero(m.PublicIPAllocation) { // not required
			return nil
		}

		if err := m.PublicIPAllocation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("publicIpAllocation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("publicIpAllocation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublicIPAllocationGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicIPAllocationGetResponse) UnmarshalBinary(b []byte) error {
	var res PublicIPAllocationGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
