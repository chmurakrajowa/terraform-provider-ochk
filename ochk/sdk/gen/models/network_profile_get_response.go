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

// NetworkProfileGetResponse NetworkProfileGetResponse
//
// swagger:model NetworkProfileGetResponse
type NetworkProfileGetResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// network profile
	NetworkProfile *NetworkProfileInstance `json:"networkProfile,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this network profile get response
func (m *NetworkProfileGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworkProfile(formats); err != nil {
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

func (m *NetworkProfileGetResponse) validateNetworkProfile(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkProfile) { // not required
		return nil
	}

	if m.NetworkProfile != nil {
		if err := m.NetworkProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkProfile")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkProfileGetResponse) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkProfileGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkProfileGetResponse) UnmarshalBinary(b []byte) error {
	var res NetworkProfileGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}