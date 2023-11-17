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

// FolderListResponse FolderListResponse
//
// swagger:model FolderListResponse
type FolderListResponse struct {

	// folder instance collection
	FolderInstanceCollection []*FolderInstance `json:"folderInstanceCollection"`

	// messages
	Messages string `json:"messages,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this folder list response
func (m *FolderListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFolderInstanceCollection(formats); err != nil {
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

func (m *FolderListResponse) validateFolderInstanceCollection(formats strfmt.Registry) error {
	if swag.IsZero(m.FolderInstanceCollection) { // not required
		return nil
	}

	for i := 0; i < len(m.FolderInstanceCollection); i++ {
		if swag.IsZero(m.FolderInstanceCollection[i]) { // not required
			continue
		}

		if m.FolderInstanceCollection[i] != nil {
			if err := m.FolderInstanceCollection[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("folderInstanceCollection" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("folderInstanceCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FolderListResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this folder list response based on the context it is used
func (m *FolderListResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFolderInstanceCollection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FolderListResponse) contextValidateFolderInstanceCollection(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FolderInstanceCollection); i++ {

		if m.FolderInstanceCollection[i] != nil {

			if swag.IsZero(m.FolderInstanceCollection[i]) { // not required
				return nil
			}

			if err := m.FolderInstanceCollection[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("folderInstanceCollection" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("folderInstanceCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FolderListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FolderListResponse) UnmarshalBinary(b []byte) error {
	var res FolderListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
