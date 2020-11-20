// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VcsVirtualMachineDeleteResponse VcsVirtualMachineDeleteResponse
//
// swagger:model VcsVirtualMachineDeleteResponse
type VcsVirtualMachineDeleteResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// request instance
	RequestInstance *RequestInstance `json:"requestInstance,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`

	// vcs virtual machine instance
	VcsVirtualMachineInstance *VcsVirtualMachineInstance `json:"vcsVirtualMachineInstance,omitempty"`
}

// Validate validates this vcs virtual machine delete response
func (m *VcsVirtualMachineDeleteResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequestInstance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcsVirtualMachineInstance(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VcsVirtualMachineDeleteResponse) validateRequestInstance(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestInstance) { // not required
		return nil
	}

	if m.RequestInstance != nil {
		if err := m.RequestInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requestInstance")
			}
			return err
		}
	}

	return nil
}

func (m *VcsVirtualMachineDeleteResponse) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VcsVirtualMachineDeleteResponse) validateVcsVirtualMachineInstance(formats strfmt.Registry) error {

	if swag.IsZero(m.VcsVirtualMachineInstance) { // not required
		return nil
	}

	if m.VcsVirtualMachineInstance != nil {
		if err := m.VcsVirtualMachineInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vcsVirtualMachineInstance")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VcsVirtualMachineDeleteResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VcsVirtualMachineDeleteResponse) UnmarshalBinary(b []byte) error {
	var res VcsVirtualMachineDeleteResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}