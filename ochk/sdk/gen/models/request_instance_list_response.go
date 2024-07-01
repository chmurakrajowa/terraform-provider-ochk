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

// RequestInstanceListResponse RequestInstanceListResponse
//
// swagger:model RequestInstanceListResponse
type RequestInstanceListResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// request instance list
	RequestInstanceList []*RequestInstance `json:"requestInstanceList"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this request instance list response
func (m *RequestInstanceListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequestInstanceList(formats); err != nil {
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

func (m *RequestInstanceListResponse) validateRequestInstanceList(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestInstanceList) { // not required
		return nil
	}

	for i := 0; i < len(m.RequestInstanceList); i++ {
		if swag.IsZero(m.RequestInstanceList[i]) { // not required
			continue
		}

		if m.RequestInstanceList[i] != nil {
			if err := m.RequestInstanceList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("requestInstanceList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("requestInstanceList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RequestInstanceListResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this request instance list response based on the context it is used
func (m *RequestInstanceListResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRequestInstanceList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RequestInstanceListResponse) contextValidateRequestInstanceList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RequestInstanceList); i++ {

		if m.RequestInstanceList[i] != nil {

			if swag.IsZero(m.RequestInstanceList[i]) { // not required
				return nil
			}

			if err := m.RequestInstanceList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("requestInstanceList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("requestInstanceList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RequestInstanceListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RequestInstanceListResponse) UnmarshalBinary(b []byte) error {
	var res RequestInstanceListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
