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

// UpdateRouterResponse UpdateRouterResponse
//
// swagger:model UpdateRouterResponse
type UpdateRouterResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// request instance
	RequestInstance *RequestInstance `json:"requestInstance,omitempty"`

	// router instance
	RouterInstance *RouterInstance `json:"routerInstance,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this update router response
func (m *UpdateRouterResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequestInstance(formats); err != nil {
		res = append(res, err)
	}

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

func (m *UpdateRouterResponse) validateRequestInstance(formats strfmt.Registry) error {
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

func (m *UpdateRouterResponse) validateRouterInstance(formats strfmt.Registry) error {
	if swag.IsZero(m.RouterInstance) { // not required
		return nil
	}

	if m.RouterInstance != nil {
		if err := m.RouterInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("routerInstance")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateRouterResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this update router response based on the context it is used
func (m *UpdateRouterResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRequestInstance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRouterInstance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateRouterResponse) contextValidateRequestInstance(ctx context.Context, formats strfmt.Registry) error {

	if m.RequestInstance != nil {
		if err := m.RequestInstance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requestInstance")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateRouterResponse) contextValidateRouterInstance(ctx context.Context, formats strfmt.Registry) error {

	if m.RouterInstance != nil {
		if err := m.RouterInstance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("routerInstance")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateRouterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateRouterResponse) UnmarshalBinary(b []byte) error {
	var res UpdateRouterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
