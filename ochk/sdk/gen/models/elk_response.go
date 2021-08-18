// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ElkResponse ElkResponse
//
// swagger:model ElkResponse
type ElkResponse struct {

	// hits
	Hits *Hits `json:"hits,omitempty"`

	// timed out
	TimedOut bool `json:"timed_out,omitempty"`
}

// Validate validates this elk response
func (m *ElkResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHits(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElkResponse) validateHits(formats strfmt.Registry) error {

	if swag.IsZero(m.Hits) { // not required
		return nil
	}

	if m.Hits != nil {
		if err := m.Hits.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hits")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ElkResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElkResponse) UnmarshalBinary(b []byte) error {
	var res ElkResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
