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

// VirtualMachinePerformanceReportGetResponse VirtualMachinePerformanceReportGetResponse
//
// swagger:model VirtualMachinePerformanceReportGetResponse
type VirtualMachinePerformanceReportGetResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`

	// virtual machine performance report item array
	VirtualMachinePerformanceReportItemArray []*VirtualMachinePerformanceReport `json:"virtualMachinePerformanceReportItemArray"`
}

// Validate validates this virtual machine performance report get response
func (m *VirtualMachinePerformanceReportGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualMachinePerformanceReportItemArray(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualMachinePerformanceReportGetResponse) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VirtualMachinePerformanceReportGetResponse) validateVirtualMachinePerformanceReportItemArray(formats strfmt.Registry) error {
	if swag.IsZero(m.VirtualMachinePerformanceReportItemArray) { // not required
		return nil
	}

	for i := 0; i < len(m.VirtualMachinePerformanceReportItemArray); i++ {
		if swag.IsZero(m.VirtualMachinePerformanceReportItemArray[i]) { // not required
			continue
		}

		if m.VirtualMachinePerformanceReportItemArray[i] != nil {
			if err := m.VirtualMachinePerformanceReportItemArray[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("virtualMachinePerformanceReportItemArray" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this virtual machine performance report get response based on the context it is used
func (m *VirtualMachinePerformanceReportGetResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVirtualMachinePerformanceReportItemArray(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualMachinePerformanceReportGetResponse) contextValidateVirtualMachinePerformanceReportItemArray(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VirtualMachinePerformanceReportItemArray); i++ {

		if m.VirtualMachinePerformanceReportItemArray[i] != nil {
			if err := m.VirtualMachinePerformanceReportItemArray[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("virtualMachinePerformanceReportItemArray" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VirtualMachinePerformanceReportGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualMachinePerformanceReportGetResponse) UnmarshalBinary(b []byte) error {
	var res VirtualMachinePerformanceReportGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
