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

// SnapshotListResponse SnapshotListResponse
//
// swagger:model SnapshotListResponse
type SnapshotListResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// snapshot instance collection
	SnapshotInstanceCollection []*SnapshotInstance `json:"snapshotInstanceCollection"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this snapshot list response
func (m *SnapshotListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSnapshotInstanceCollection(formats); err != nil {
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

func (m *SnapshotListResponse) validateSnapshotInstanceCollection(formats strfmt.Registry) error {

	if swag.IsZero(m.SnapshotInstanceCollection) { // not required
		return nil
	}

	for i := 0; i < len(m.SnapshotInstanceCollection); i++ {
		if swag.IsZero(m.SnapshotInstanceCollection[i]) { // not required
			continue
		}

		if m.SnapshotInstanceCollection[i] != nil {
			if err := m.SnapshotInstanceCollection[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("snapshotInstanceCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnapshotListResponse) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotListResponse) UnmarshalBinary(b []byte) error {
	var res SnapshotListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
