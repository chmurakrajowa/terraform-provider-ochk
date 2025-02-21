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

// ListRequestsResponse list requests response
//
// swagger:model ListRequestsResponse
type ListRequestsResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// request instance collection
	RequestInstanceCollection []*RequestInstance `json:"requestInstanceCollection"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this list requests response
func (m *ListRequestsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequestInstanceCollection(formats); err != nil {
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

func (m *ListRequestsResponse) validateRequestInstanceCollection(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestInstanceCollection) { // not required
		return nil
	}

	for i := 0; i < len(m.RequestInstanceCollection); i++ {
		if swag.IsZero(m.RequestInstanceCollection[i]) { // not required
			continue
		}

		if m.RequestInstanceCollection[i] != nil {
			if err := m.RequestInstanceCollection[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("requestInstanceCollection" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("requestInstanceCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ListRequestsResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this list requests response based on the context it is used
func (m *ListRequestsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRequestInstanceCollection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListRequestsResponse) contextValidateRequestInstanceCollection(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RequestInstanceCollection); i++ {

		if m.RequestInstanceCollection[i] != nil {

			if swag.IsZero(m.RequestInstanceCollection[i]) { // not required
				return nil
			}

			if err := m.RequestInstanceCollection[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("requestInstanceCollection" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("requestInstanceCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListRequestsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListRequestsResponse) UnmarshalBinary(b []byte) error {
	var res ListRequestsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
