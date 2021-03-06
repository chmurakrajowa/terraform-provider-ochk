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

// EdgeClusterListResponse EdgeClusterListResponse
//
// swagger:model EdgeClusterListResponse
type EdgeClusterListResponse struct {

	// edge cluster collection
	EdgeClusterCollection []*EdgeClusterInstance `json:"edgeClusterCollection"`

	// messages
	Messages string `json:"messages,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this edge cluster list response
func (m *EdgeClusterListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEdgeClusterCollection(formats); err != nil {
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

func (m *EdgeClusterListResponse) validateEdgeClusterCollection(formats strfmt.Registry) error {

	if swag.IsZero(m.EdgeClusterCollection) { // not required
		return nil
	}

	for i := 0; i < len(m.EdgeClusterCollection); i++ {
		if swag.IsZero(m.EdgeClusterCollection[i]) { // not required
			continue
		}

		if m.EdgeClusterCollection[i] != nil {
			if err := m.EdgeClusterCollection[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("edgeClusterCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EdgeClusterListResponse) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EdgeClusterListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgeClusterListResponse) UnmarshalBinary(b []byte) error {
	var res EdgeClusterListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
