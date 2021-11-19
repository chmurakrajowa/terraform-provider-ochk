// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// QueryFilter QueryFilter
//
// swagger:model QueryFilter
type QueryFilter struct {

	// end date time
	// Format: date-time
	EndDateTime *strfmt.DateTime `json:"endDateTime,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// log severity
	// Enum: [ALERT CRITICAL DEBUG DEFAULT EMERGENCY ERROR FATAL INFO NOTICE WARN WARNING]
	LogSeverity string `json:"logSeverity,omitempty"`

	// query string
	QueryString string `json:"queryString,omitempty"`

	// resource name
	ResourceName string `json:"resourceName,omitempty"`

	// start date time
	// Format: date-time
	StartDateTime *strfmt.DateTime `json:"startDateTime,omitempty"`

	// user name
	UserName string `json:"userName,omitempty"`
}

// Validate validates this query filter
func (m *QueryFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogSeverity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryFilter) validateEndDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("endDateTime", "body", "date-time", m.EndDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var queryFilterTypeLogSeverityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ALERT","CRITICAL","DEBUG","DEFAULT","EMERGENCY","ERROR","FATAL","INFO","NOTICE","WARN","WARNING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		queryFilterTypeLogSeverityPropEnum = append(queryFilterTypeLogSeverityPropEnum, v)
	}
}

const (

	// QueryFilterLogSeverityALERT captures enum value "ALERT"
	QueryFilterLogSeverityALERT string = "ALERT"

	// QueryFilterLogSeverityCRITICAL captures enum value "CRITICAL"
	QueryFilterLogSeverityCRITICAL string = "CRITICAL"

	// QueryFilterLogSeverityDEBUG captures enum value "DEBUG"
	QueryFilterLogSeverityDEBUG string = "DEBUG"

	// QueryFilterLogSeverityDEFAULT captures enum value "DEFAULT"
	QueryFilterLogSeverityDEFAULT string = "DEFAULT"

	// QueryFilterLogSeverityEMERGENCY captures enum value "EMERGENCY"
	QueryFilterLogSeverityEMERGENCY string = "EMERGENCY"

	// QueryFilterLogSeverityERROR captures enum value "ERROR"
	QueryFilterLogSeverityERROR string = "ERROR"

	// QueryFilterLogSeverityFATAL captures enum value "FATAL"
	QueryFilterLogSeverityFATAL string = "FATAL"

	// QueryFilterLogSeverityINFO captures enum value "INFO"
	QueryFilterLogSeverityINFO string = "INFO"

	// QueryFilterLogSeverityNOTICE captures enum value "NOTICE"
	QueryFilterLogSeverityNOTICE string = "NOTICE"

	// QueryFilterLogSeverityWARN captures enum value "WARN"
	QueryFilterLogSeverityWARN string = "WARN"

	// QueryFilterLogSeverityWARNING captures enum value "WARNING"
	QueryFilterLogSeverityWARNING string = "WARNING"
)

// prop value enum
func (m *QueryFilter) validateLogSeverityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, queryFilterTypeLogSeverityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *QueryFilter) validateLogSeverity(formats strfmt.Registry) error {
	if swag.IsZero(m.LogSeverity) { // not required
		return nil
	}

	// value enum
	if err := m.validateLogSeverityEnum("logSeverity", "body", m.LogSeverity); err != nil {
		return err
	}

	return nil
}

func (m *QueryFilter) validateStartDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startDateTime", "body", "date-time", m.StartDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this query filter based on context it is used
func (m *QueryFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QueryFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryFilter) UnmarshalBinary(b []byte) error {
	var res QueryFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
