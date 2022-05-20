// Code generated by go-swagger; DO NOT EDIT.

package k_m_s_key_rotation_scheduler

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new k m s key rotation scheduler API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for k m s key rotation scheduler API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	KeyRotationScheduleCreateUsingPUT(params *KeyRotationScheduleCreateUsingPUTParams, opts ...ClientOption) (*KeyRotationScheduleCreateUsingPUTOK, *KeyRotationScheduleCreateUsingPUTCreated, error)

	KeyRotationScheduleDeleteUsingDELETE(params *KeyRotationScheduleDeleteUsingDELETEParams, opts ...ClientOption) (*KeyRotationScheduleDeleteUsingDELETEOK, error)

	KeyRotationScheduleGetUsingGET(params *KeyRotationScheduleGetUsingGETParams, opts ...ClientOption) (*KeyRotationScheduleGetUsingGETOK, error)

	KeyRotationScheduleUpdateUsingPUT(params *KeyRotationScheduleUpdateUsingPUTParams, opts ...ClientOption) (*KeyRotationScheduleUpdateUsingPUTOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  KeyRotationScheduleCreateUsingPUT creates

  Create key rotation schedule (KMS)
*/
func (a *Client) KeyRotationScheduleCreateUsingPUT(params *KeyRotationScheduleCreateUsingPUTParams, opts ...ClientOption) (*KeyRotationScheduleCreateUsingPUTOK, *KeyRotationScheduleCreateUsingPUTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKeyRotationScheduleCreateUsingPUTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "keyRotationScheduleCreateUsingPUT",
		Method:             "PUT",
		PathPattern:        "/kms/schedule",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &KeyRotationScheduleCreateUsingPUTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *KeyRotationScheduleCreateUsingPUTOK:
		return value, nil, nil
	case *KeyRotationScheduleCreateUsingPUTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for k_m_s_key_rotation_scheduler: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  KeyRotationScheduleDeleteUsingDELETE deletes

  Delete key rotation schedule (KMS)
*/
func (a *Client) KeyRotationScheduleDeleteUsingDELETE(params *KeyRotationScheduleDeleteUsingDELETEParams, opts ...ClientOption) (*KeyRotationScheduleDeleteUsingDELETEOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKeyRotationScheduleDeleteUsingDELETEParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "keyRotationScheduleDeleteUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/kms/schedule/{keyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &KeyRotationScheduleDeleteUsingDELETEReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*KeyRotationScheduleDeleteUsingDELETEOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for keyRotationScheduleDeleteUsingDELETE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  KeyRotationScheduleGetUsingGET gets

  Get key rotation schedule (KMS)
*/
func (a *Client) KeyRotationScheduleGetUsingGET(params *KeyRotationScheduleGetUsingGETParams, opts ...ClientOption) (*KeyRotationScheduleGetUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKeyRotationScheduleGetUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "keyRotationScheduleGetUsingGET",
		Method:             "GET",
		PathPattern:        "/kms/schedule/{keyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &KeyRotationScheduleGetUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*KeyRotationScheduleGetUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for keyRotationScheduleGetUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  KeyRotationScheduleUpdateUsingPUT updates

  Update key rotation schedule (KMS)
*/
func (a *Client) KeyRotationScheduleUpdateUsingPUT(params *KeyRotationScheduleUpdateUsingPUTParams, opts ...ClientOption) (*KeyRotationScheduleUpdateUsingPUTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKeyRotationScheduleUpdateUsingPUTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "keyRotationScheduleUpdateUsingPUT",
		Method:             "PUT",
		PathPattern:        "/kms/schedule/{keyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &KeyRotationScheduleUpdateUsingPUTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*KeyRotationScheduleUpdateUsingPUTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for keyRotationScheduleUpdateUsingPUT: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
