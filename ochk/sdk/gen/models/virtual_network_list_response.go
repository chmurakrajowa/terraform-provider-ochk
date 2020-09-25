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

// VirtualNetworkListResponse VirtualNetworkListResponse
//
// swagger:model VirtualNetworkListResponse
type VirtualNetworkListResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`

	// virtual network instance collection
	VirtualNetworkInstanceCollection []*VirtualNetworkInstance `json:"virtualNetworkInstanceCollection"`
}

// Validate validates this virtual network list response
func (m *VirtualNetworkListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualNetworkInstanceCollection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualNetworkListResponse) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VirtualNetworkListResponse) validateVirtualNetworkInstanceCollection(formats strfmt.Registry) error {

	if swag.IsZero(m.VirtualNetworkInstanceCollection) { // not required
		return nil
	}

	for i := 0; i < len(m.VirtualNetworkInstanceCollection); i++ {
		if swag.IsZero(m.VirtualNetworkInstanceCollection[i]) { // not required
			continue
		}

		if m.VirtualNetworkInstanceCollection[i] != nil {
			if err := m.VirtualNetworkInstanceCollection[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("virtualNetworkInstanceCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VirtualNetworkListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualNetworkListResponse) UnmarshalBinary(b []byte) error {
	var res VirtualNetworkListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}