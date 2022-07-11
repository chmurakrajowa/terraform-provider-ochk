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

// TenantInfoGetResponse TenantInfoGetResponse
//
// swagger:model TenantInfoGetResponse
type TenantInfoGetResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// tenant info
	TenantInfo *TenantInfo `json:"tenantInfo,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this tenant info get response
func (m *TenantInfoGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTenantInfo(formats); err != nil {
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

func (m *TenantInfoGetResponse) validateTenantInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.TenantInfo) { // not required
		return nil
	}

	if m.TenantInfo != nil {
		if err := m.TenantInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenantInfo")
			}
			return err
		}
	}

	return nil
}

func (m *TenantInfoGetResponse) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TenantInfoGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TenantInfoGetResponse) UnmarshalBinary(b []byte) error {
	var res TenantInfoGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}