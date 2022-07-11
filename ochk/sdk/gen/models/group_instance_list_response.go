// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GroupInstanceListResponse GroupInstanceListResponse
//
// swagger:model GroupInstanceListResponse
type GroupInstanceListResponse struct {

	// group instance collection
	GroupInstanceCollection []*ADGroupInstance `json:"groupInstanceCollection"`

	// messages
	Messages string `json:"messages,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this group instance list response
func (m *GroupInstanceListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupInstanceCollection(formats); err != nil {
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

func (m *GroupInstanceListResponse) validateGroupInstanceCollection(formats strfmt.Registry) error {

	if swag.IsZero(m.GroupInstanceCollection) { // not required
		return nil
	}

	for i := 0; i < len(m.GroupInstanceCollection); i++ {
		if swag.IsZero(m.GroupInstanceCollection[i]) { // not required
			continue
		}

		if m.GroupInstanceCollection[i] != nil {
			if err := m.GroupInstanceCollection[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groupInstanceCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GroupInstanceListResponse) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GroupInstanceListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupInstanceListResponse) UnmarshalBinary(b []byte) error {
	var res GroupInstanceListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
