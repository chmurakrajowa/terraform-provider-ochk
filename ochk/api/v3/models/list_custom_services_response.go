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

// ListCustomServicesResponse list custom services response
//
// swagger:model ListCustomServicesResponse
type ListCustomServicesResponse struct {

	// custom service instance collection
	CustomServiceInstanceCollection []*CustomServiceInstance `json:"customServiceInstanceCollection"`

	// messages
	Messages string `json:"messages,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this list custom services response
func (m *ListCustomServicesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomServiceInstanceCollection(formats); err != nil {
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

func (m *ListCustomServicesResponse) validateCustomServiceInstanceCollection(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomServiceInstanceCollection) { // not required
		return nil
	}

	for i := 0; i < len(m.CustomServiceInstanceCollection); i++ {
		if swag.IsZero(m.CustomServiceInstanceCollection[i]) { // not required
			continue
		}

		if m.CustomServiceInstanceCollection[i] != nil {
			if err := m.CustomServiceInstanceCollection[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customServiceInstanceCollection" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customServiceInstanceCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ListCustomServicesResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this list custom services response based on the context it is used
func (m *ListCustomServicesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomServiceInstanceCollection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListCustomServicesResponse) contextValidateCustomServiceInstanceCollection(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CustomServiceInstanceCollection); i++ {

		if m.CustomServiceInstanceCollection[i] != nil {

			if swag.IsZero(m.CustomServiceInstanceCollection[i]) { // not required
				return nil
			}

			if err := m.CustomServiceInstanceCollection[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customServiceInstanceCollection" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customServiceInstanceCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListCustomServicesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListCustomServicesResponse) UnmarshalBinary(b []byte) error {
	var res ListCustomServicesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
