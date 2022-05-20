// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdatePublicIPAllocationResponse UpdatePublicIPAllocationResponse
//
// swagger:model UpdatePublicIPAllocationResponse
type UpdatePublicIPAllocationResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// public Ip allocation
	PublicIPAllocation *PublicIPAllocation `json:"publicIpAllocation,omitempty"`

	// request instance
	RequestInstance *RequestInstance `json:"requestInstance,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this update public IP allocation response
func (m *UpdatePublicIPAllocationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePublicIPAllocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestInstance(formats); err != nil {
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

func (m *UpdatePublicIPAllocationResponse) validatePublicIPAllocation(formats strfmt.Registry) error {

	if swag.IsZero(m.PublicIPAllocation) { // not required
		return nil
	}

	if m.PublicIPAllocation != nil {
		if err := m.PublicIPAllocation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("publicIpAllocation")
			}
			return err
		}
	}

	return nil
}

func (m *UpdatePublicIPAllocationResponse) validateRequestInstance(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestInstance) { // not required
		return nil
	}

	if m.RequestInstance != nil {
		if err := m.RequestInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requestInstance")
			}
			return err
		}
	}

	return nil
}

func (m *UpdatePublicIPAllocationResponse) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdatePublicIPAllocationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdatePublicIPAllocationResponse) UnmarshalBinary(b []byte) error {
	var res UpdatePublicIPAllocationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
