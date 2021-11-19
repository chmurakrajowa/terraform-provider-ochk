// Code generated by go-swagger; DO NOT EDIT.

package system_tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new system tags API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for system tags API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	SystemTagCreateUsingPUT(params *SystemTagCreateUsingPUTParams, opts ...ClientOption) (*SystemTagCreateUsingPUTOK, *SystemTagCreateUsingPUTCreated, error)

	SystemTagDeleteUsingDELETE(params *SystemTagDeleteUsingDELETEParams, opts ...ClientOption) (*SystemTagDeleteUsingDELETEOK, error)

	SystemTagGetUsingGET(params *SystemTagGetUsingGETParams, opts ...ClientOption) (*SystemTagGetUsingGETOK, error)

	SystemTagListUsingGET(params *SystemTagListUsingGETParams, opts ...ClientOption) (*SystemTagListUsingGETOK, error)

	SystemTagUpdateUsingPUT(params *SystemTagUpdateUsingPUTParams, opts ...ClientOption) (*SystemTagUpdateUsingPUTOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  SystemTagCreateUsingPUT creates

  Create system tag
*/
func (a *Client) SystemTagCreateUsingPUT(params *SystemTagCreateUsingPUTParams, opts ...ClientOption) (*SystemTagCreateUsingPUTOK, *SystemTagCreateUsingPUTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSystemTagCreateUsingPUTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "systemTagCreateUsingPUT",
		Method:             "PUT",
		PathPattern:        "/tags/systemTags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SystemTagCreateUsingPUTReader{formats: a.formats},
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
	case *SystemTagCreateUsingPUTOK:
		return value, nil, nil
	case *SystemTagCreateUsingPUTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for system_tags: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SystemTagDeleteUsingDELETE deletes

  Delete system tag
*/
func (a *Client) SystemTagDeleteUsingDELETE(params *SystemTagDeleteUsingDELETEParams, opts ...ClientOption) (*SystemTagDeleteUsingDELETEOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSystemTagDeleteUsingDELETEParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "systemTagDeleteUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/tags/systemTags/{systemTagId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SystemTagDeleteUsingDELETEReader{formats: a.formats},
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
	success, ok := result.(*SystemTagDeleteUsingDELETEOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for systemTagDeleteUsingDELETE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SystemTagGetUsingGET gets

  Get System tag collection
*/
func (a *Client) SystemTagGetUsingGET(params *SystemTagGetUsingGETParams, opts ...ClientOption) (*SystemTagGetUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSystemTagGetUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "systemTagGetUsingGET",
		Method:             "GET",
		PathPattern:        "/tags/systemTags/{systemTagId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SystemTagGetUsingGETReader{formats: a.formats},
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
	success, ok := result.(*SystemTagGetUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for systemTagGetUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SystemTagListUsingGET lists

  List system tags
*/
func (a *Client) SystemTagListUsingGET(params *SystemTagListUsingGETParams, opts ...ClientOption) (*SystemTagListUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSystemTagListUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "systemTagListUsingGET",
		Method:             "GET",
		PathPattern:        "/tags/systemTags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SystemTagListUsingGETReader{formats: a.formats},
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
	success, ok := result.(*SystemTagListUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for systemTagListUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SystemTagUpdateUsingPUT updates

  Update system tag
*/
func (a *Client) SystemTagUpdateUsingPUT(params *SystemTagUpdateUsingPUTParams, opts ...ClientOption) (*SystemTagUpdateUsingPUTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSystemTagUpdateUsingPUTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "systemTagUpdateUsingPUT",
		Method:             "PUT",
		PathPattern:        "/tags/systemTags/{systemTagId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SystemTagUpdateUsingPUTReader{formats: a.formats},
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
	success, ok := result.(*SystemTagUpdateUsingPUTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for systemTagUpdateUsingPUT: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
