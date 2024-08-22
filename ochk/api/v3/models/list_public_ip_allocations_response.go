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

// ListPublicIPAllocationsResponse list public Ip allocations response
//
// swagger:model ListPublicIpAllocationsResponse
type ListPublicIPAllocationsResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// public Ip allocation collection
	PublicIPAllocationCollection []*PublicIPAllocation `json:"publicIpAllocationCollection"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this list public Ip allocations response
func (m *ListPublicIPAllocationsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePublicIPAllocationCollection(formats); err != nil {
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

func (m *ListPublicIPAllocationsResponse) validatePublicIPAllocationCollection(formats strfmt.Registry) error {
	if swag.IsZero(m.PublicIPAllocationCollection) { // not required
		return nil
	}

	for i := 0; i < len(m.PublicIPAllocationCollection); i++ {
		if swag.IsZero(m.PublicIPAllocationCollection[i]) { // not required
			continue
		}

		if m.PublicIPAllocationCollection[i] != nil {
			if err := m.PublicIPAllocationCollection[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("publicIpAllocationCollection" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("publicIpAllocationCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ListPublicIPAllocationsResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this list public Ip allocations response based on the context it is used
func (m *ListPublicIPAllocationsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePublicIPAllocationCollection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListPublicIPAllocationsResponse) contextValidatePublicIPAllocationCollection(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PublicIPAllocationCollection); i++ {

		if m.PublicIPAllocationCollection[i] != nil {

			if swag.IsZero(m.PublicIPAllocationCollection[i]) { // not required
				return nil
			}

			if err := m.PublicIPAllocationCollection[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("publicIpAllocationCollection" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("publicIpAllocationCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListPublicIPAllocationsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListPublicIPAllocationsResponse) UnmarshalBinary(b []byte) error {
	var res ListPublicIPAllocationsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
