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

// ReservationInstance ReservationInstance
//
// swagger:model ReservationInstance
type ReservationInstance struct {

	// created date
	// Format: date-time
	CreatedDate *strfmt.DateTime `json:"createdDate,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// external Id
	ExternalID string `json:"externalId,omitempty"`

	// last updated
	// Format: date-time
	LastUpdated *strfmt.DateTime `json:"lastUpdated,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// network instances
	NetworkInstances []*NetworkInstance `json:"networkInstances"`

	// priority
	Priority int32 `json:"priority,omitempty"`

	// reservation Id
	ReservationID string `json:"reservationId,omitempty"`

	// reservation type
	ReservationType *ReservationType `json:"reservationType,omitempty"`

	// version
	Version int64 `json:"version,omitempty"`
}

// Validate validates this reservation instance
func (m *ReservationInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkInstances(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReservationType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReservationInstance) validateCreatedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createdDate", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReservationInstance) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReservationInstance) validateNetworkInstances(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkInstances) { // not required
		return nil
	}

	for i := 0; i < len(m.NetworkInstances); i++ {
		if swag.IsZero(m.NetworkInstances[i]) { // not required
			continue
		}

		if m.NetworkInstances[i] != nil {
			if err := m.NetworkInstances[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networkInstances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReservationInstance) validateReservationType(formats strfmt.Registry) error {

	if swag.IsZero(m.ReservationType) { // not required
		return nil
	}

	if m.ReservationType != nil {
		if err := m.ReservationType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reservationType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReservationInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReservationInstance) UnmarshalBinary(b []byte) error {
	var res ReservationInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
