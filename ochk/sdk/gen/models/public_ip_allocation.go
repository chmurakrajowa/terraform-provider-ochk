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

// PublicIPAllocation PublicIpAllocation
//
// swagger:model PublicIpAllocation
type PublicIPAllocation struct {

	// allocation Id
	AllocationID int32 `json:"allocationId,omitempty"`

	// assignment date
	// Format: date-time
	AssignmentDate *strfmt.DateTime `json:"assignmentDate,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// public Ip address
	PublicIPAddress *PublicIPAddress `json:"publicIpAddress,omitempty"`

	// service list
	ServiceList []*ServiceInstance `json:"serviceList"`

	// services
	Services string `json:"services,omitempty"`
}

// Validate validates this public Ip allocation
func (m *PublicIPAllocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssignmentDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicIPAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicIPAllocation) validateAssignmentDate(formats strfmt.Registry) error {
	if swag.IsZero(m.AssignmentDate) { // not required
		return nil
	}

	if err := validate.FormatOf("assignmentDate", "body", "date-time", m.AssignmentDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PublicIPAllocation) validatePublicIPAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.PublicIPAddress) { // not required
		return nil
	}

	if m.PublicIPAddress != nil {
		if err := m.PublicIPAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("publicIpAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("publicIpAddress")
			}
			return err
		}
	}

	return nil
}

func (m *PublicIPAllocation) validateServiceList(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceList) { // not required
		return nil
	}

	for i := 0; i < len(m.ServiceList); i++ {
		if swag.IsZero(m.ServiceList[i]) { // not required
			continue
		}

		if m.ServiceList[i] != nil {
			if err := m.ServiceList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serviceList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serviceList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this public Ip allocation based on the context it is used
func (m *PublicIPAllocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePublicIPAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServiceList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicIPAllocation) contextValidatePublicIPAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.PublicIPAddress != nil {
		if err := m.PublicIPAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("publicIpAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("publicIpAddress")
			}
			return err
		}
	}

	return nil
}

func (m *PublicIPAllocation) contextValidateServiceList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ServiceList); i++ {

		if m.ServiceList[i] != nil {
			if err := m.ServiceList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serviceList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serviceList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublicIPAllocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicIPAllocation) UnmarshalBinary(b []byte) error {
	var res PublicIPAllocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
