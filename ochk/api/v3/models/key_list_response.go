// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// KeyListResponse key list response
//
// swagger:model KeyListResponse
type KeyListResponse struct {

	// key instance collection
	KeyInstanceCollection []*KeyInstance `json:"keyInstanceCollection"`

	// messages
	Messages string `json:"messages,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this key list response
func (m *KeyListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKeyInstanceCollection(formats); err != nil {
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

func (m *KeyListResponse) validateKeyInstanceCollection(formats strfmt.Registry) error {
	if swag.IsZero(m.KeyInstanceCollection) { // not required
		return nil
	}

	for i := 0; i < len(m.KeyInstanceCollection); i++ {
		if swag.IsZero(m.KeyInstanceCollection[i]) { // not required
			continue
		}

		if m.KeyInstanceCollection[i] != nil {
			if err := m.KeyInstanceCollection[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("keyInstanceCollection" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("keyInstanceCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *KeyListResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this key list response based on the context it is used
func (m *KeyListResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKeyInstanceCollection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeyListResponse) contextValidateKeyInstanceCollection(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.KeyInstanceCollection); i++ {

		if m.KeyInstanceCollection[i] != nil {

			if swag.IsZero(m.KeyInstanceCollection[i]) { // not required
				return nil
			}

			if err := m.KeyInstanceCollection[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("keyInstanceCollection" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("keyInstanceCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *KeyListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeyListResponse) UnmarshalBinary(b []byte) error {
	var res KeyListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
