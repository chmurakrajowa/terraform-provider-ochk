// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ADUserInstance ADUserInstance
//
// swagger:model ADUserInstance
type ADUserInstance struct {

	// account expiration date
	// Format: date-time
	AccountExpirationDate *strfmt.DateTime `json:"accountExpirationDate,omitempty"`

	// account lockout time
	// Format: date-time
	AccountLockoutTime *strfmt.DateTime `json:"accountLockoutTime,omitempty"`

	// bad logon count
	BadLogonCount int32 `json:"badLogonCount,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// dn
	Dn string `json:"dn,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// first name
	FirstName string `json:"firstName,omitempty"`

	// last bad password attempt
	// Format: date-time
	LastBadPasswordAttempt *strfmt.DateTime `json:"lastBadPasswordAttempt,omitempty"`

	// last logon
	// Format: date-time
	LastLogon *strfmt.DateTime `json:"lastLogon,omitempty"`

	// last name
	LastName string `json:"lastName,omitempty"`

	// last password set
	// Format: date-time
	LastPasswordSet *strfmt.DateTime `json:"lastPasswordSet,omitempty"`

	// password never expires
	PasswordNeverExpires bool `json:"passwordNeverExpires,omitempty"`

	// sam account name
	SamAccountName string `json:"samAccountName,omitempty"`

	// user principal name
	UserPrincipalName string `json:"userPrincipalName,omitempty"`
}

// Validate validates this a d user instance
func (m *ADUserInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountExpirationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountLockoutTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastBadPasswordAttempt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastLogon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastPasswordSet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ADUserInstance) validateAccountExpirationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.AccountExpirationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("accountExpirationDate", "body", "date-time", m.AccountExpirationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ADUserInstance) validateAccountLockoutTime(formats strfmt.Registry) error {
	if swag.IsZero(m.AccountLockoutTime) { // not required
		return nil
	}

	if err := validate.FormatOf("accountLockoutTime", "body", "date-time", m.AccountLockoutTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ADUserInstance) validateLastBadPasswordAttempt(formats strfmt.Registry) error {
	if swag.IsZero(m.LastBadPasswordAttempt) { // not required
		return nil
	}

	if err := validate.FormatOf("lastBadPasswordAttempt", "body", "date-time", m.LastBadPasswordAttempt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ADUserInstance) validateLastLogon(formats strfmt.Registry) error {
	if swag.IsZero(m.LastLogon) { // not required
		return nil
	}

	if err := validate.FormatOf("lastLogon", "body", "date-time", m.LastLogon.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ADUserInstance) validateLastPasswordSet(formats strfmt.Registry) error {
	if swag.IsZero(m.LastPasswordSet) { // not required
		return nil
	}

	if err := validate.FormatOf("lastPasswordSet", "body", "date-time", m.LastPasswordSet.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this a d user instance based on context it is used
func (m *ADUserInstance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ADUserInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ADUserInstance) UnmarshalBinary(b []byte) error {
	var res ADUserInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
