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

// VcsResourcePoolsGetResponse VcsResourcePoolsGetResponse
//
// swagger:model VcsResourcePoolsGetResponse
type VcsResourcePoolsGetResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`

	// vcs resource pool instance
	VcsResourcePoolInstance *VcsResourcePools `json:"vcsResourcePoolInstance,omitempty"`
}

// Validate validates this vcs resource pools get response
func (m *VcsResourcePoolsGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcsResourcePoolInstance(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VcsResourcePoolsGetResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VcsResourcePoolsGetResponse) validateVcsResourcePoolInstance(formats strfmt.Registry) error {
	if swag.IsZero(m.VcsResourcePoolInstance) { // not required
		return nil
	}

	if m.VcsResourcePoolInstance != nil {
		if err := m.VcsResourcePoolInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vcsResourcePoolInstance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vcsResourcePoolInstance")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this vcs resource pools get response based on the context it is used
func (m *VcsResourcePoolsGetResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVcsResourcePoolInstance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VcsResourcePoolsGetResponse) contextValidateVcsResourcePoolInstance(ctx context.Context, formats strfmt.Registry) error {

	if m.VcsResourcePoolInstance != nil {
		if err := m.VcsResourcePoolInstance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vcsResourcePoolInstance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vcsResourcePoolInstance")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VcsResourcePoolsGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VcsResourcePoolsGetResponse) UnmarshalBinary(b []byte) error {
	var res VcsResourcePoolsGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
