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

// ProjectListResponse ProjectListResponse
//
// swagger:model ProjectListResponse
type ProjectListResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// project instance collection
	ProjectInstanceCollection []*ProjectInstance `json:"projectInstanceCollection"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this project list response
func (m *ProjectListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProjectInstanceCollection(formats); err != nil {
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

func (m *ProjectListResponse) validateProjectInstanceCollection(formats strfmt.Registry) error {
	if swag.IsZero(m.ProjectInstanceCollection) { // not required
		return nil
	}

	for i := 0; i < len(m.ProjectInstanceCollection); i++ {
		if swag.IsZero(m.ProjectInstanceCollection[i]) { // not required
			continue
		}

		if m.ProjectInstanceCollection[i] != nil {
			if err := m.ProjectInstanceCollection[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("projectInstanceCollection" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("projectInstanceCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProjectListResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this project list response based on the context it is used
func (m *ProjectListResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProjectInstanceCollection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectListResponse) contextValidateProjectInstanceCollection(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProjectInstanceCollection); i++ {

		if m.ProjectInstanceCollection[i] != nil {

			if swag.IsZero(m.ProjectInstanceCollection[i]) { // not required
				return nil
			}

			if err := m.ProjectInstanceCollection[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("projectInstanceCollection" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("projectInstanceCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectListResponse) UnmarshalBinary(b []byte) error {
	var res ProjectListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
