// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProblemDetails problem details
//
// swagger:model ProblemDetails
type ProblemDetails struct {

	// detail
	Detail string `json:"detail,omitempty"`

	// instance
	Instance string `json:"instance,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// problem details
	ProblemDetails map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *ProblemDetails) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// detail
		Detail string `json:"detail,omitempty"`

		// instance
		Instance string `json:"instance,omitempty"`

		// status
		Status int32 `json:"status,omitempty"`

		// title
		Title string `json:"title,omitempty"`

		// type
		Type string `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv ProblemDetails

	rcv.Detail = stage1.Detail
	rcv.Instance = stage1.Instance
	rcv.Status = stage1.Status
	rcv.Title = stage1.Title
	rcv.Type = stage1.Type
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "detail")
	delete(stage2, "instance")
	delete(stage2, "status")
	delete(stage2, "title")
	delete(stage2, "type")
	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]interface{})
		for k, v := range stage2 {
			var toadd interface{}
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		m.ProblemDetails = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m ProblemDetails) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// detail
		Detail string `json:"detail,omitempty"`

		// instance
		Instance string `json:"instance,omitempty"`

		// status
		Status int32 `json:"status,omitempty"`

		// title
		Title string `json:"title,omitempty"`

		// type
		Type string `json:"type,omitempty"`
	}

	stage1.Detail = m.Detail
	stage1.Instance = m.Instance
	stage1.Status = m.Status
	stage1.Title = m.Title
	stage1.Type = m.Type

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.ProblemDetails) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.ProblemDetails)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this problem details
func (m *ProblemDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this problem details based on context it is used
func (m *ProblemDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProblemDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProblemDetails) UnmarshalBinary(b []byte) error {
	var res ProblemDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
