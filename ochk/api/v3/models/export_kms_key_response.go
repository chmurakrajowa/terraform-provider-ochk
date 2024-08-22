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

// ExportKmsKeyResponse export kms key response
//
// swagger:model ExportKmsKeyResponse
type ExportKmsKeyResponse struct {

	// key instance
	KeyInstance *KeyInstance `json:"keyInstance,omitempty"`

	// messages
	Messages string `json:"messages,omitempty"`

	// request instance
	RequestInstance *RequestInstance `json:"requestInstance,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this export kms key response
func (m *ExportKmsKeyResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKeyInstance(formats); err != nil {
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

func (m *ExportKmsKeyResponse) validateKeyInstance(formats strfmt.Registry) error {
	if swag.IsZero(m.KeyInstance) { // not required
		return nil
	}

	if m.KeyInstance != nil {
		if err := m.KeyInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keyInstance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keyInstance")
			}
			return err
		}
	}

	return nil
}

func (m *ExportKmsKeyResponse) validateRequestInstance(formats strfmt.Registry) error {
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

func (m *ExportKmsKeyResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this export kms key response based on the context it is used
func (m *ExportKmsKeyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKeyInstance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequestInstance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExportKmsKeyResponse) contextValidateKeyInstance(ctx context.Context, formats strfmt.Registry) error {

	if m.KeyInstance != nil {

		if swag.IsZero(m.KeyInstance) { // not required
			return nil
		}

		if err := m.KeyInstance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keyInstance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keyInstance")
			}
			return err
		}
	}

	return nil
}

func (m *ExportKmsKeyResponse) contextValidateRequestInstance(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *ExportKmsKeyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExportKmsKeyResponse) UnmarshalBinary(b []byte) error {
	var res ExportKmsKeyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
