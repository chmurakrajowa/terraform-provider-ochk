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

// RouterGetResponse RouterGetResponse
//
// swagger:model RouterGetResponse
type RouterGetResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// router instance
	RouterInstance *RouterInstance `json:"routerInstance,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this router get response
func (m *RouterGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRouterInstance(formats); err != nil {
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

func (m *RouterGetResponse) validateRouterInstance(formats strfmt.Registry) error {
	if swag.IsZero(m.RouterInstance) { // not required
		return nil
	}

	if m.RouterInstance != nil {
		if err := m.RouterInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("routerInstance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("routerInstance")
			}
			return err
		}
	}

	return nil
}

func (m *RouterGetResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this router get response based on the context it is used
func (m *RouterGetResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRouterInstance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RouterGetResponse) contextValidateRouterInstance(ctx context.Context, formats strfmt.Registry) error {

	if m.RouterInstance != nil {

		if swag.IsZero(m.RouterInstance) { // not required
			return nil
		}

		if err := m.RouterInstance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("routerInstance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("routerInstance")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RouterGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RouterGetResponse) UnmarshalBinary(b []byte) error {
	var res RouterGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
