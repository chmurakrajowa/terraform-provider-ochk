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

// NetworkListResponse NetworkListResponse
//
// swagger:model NetworkListResponse
type NetworkListResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`

	// vcs network instance collection
	VcsNetworkInstanceCollection []*VCSNetworkInstance `json:"vcsNetworkInstanceCollection"`
}

// Validate validates this network list response
func (m *NetworkListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcsNetworkInstanceCollection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkListResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NetworkListResponse) validateVcsNetworkInstanceCollection(formats strfmt.Registry) error {
	if swag.IsZero(m.VcsNetworkInstanceCollection) { // not required
		return nil
	}

	for i := 0; i < len(m.VcsNetworkInstanceCollection); i++ {
		if swag.IsZero(m.VcsNetworkInstanceCollection[i]) { // not required
			continue
		}

		if m.VcsNetworkInstanceCollection[i] != nil {
			if err := m.VcsNetworkInstanceCollection[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vcsNetworkInstanceCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this network list response based on the context it is used
func (m *NetworkListResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVcsNetworkInstanceCollection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkListResponse) contextValidateVcsNetworkInstanceCollection(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VcsNetworkInstanceCollection); i++ {

		if m.VcsNetworkInstanceCollection[i] != nil {
			if err := m.VcsNetworkInstanceCollection[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vcsNetworkInstanceCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkListResponse) UnmarshalBinary(b []byte) error {
	var res NetworkListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
