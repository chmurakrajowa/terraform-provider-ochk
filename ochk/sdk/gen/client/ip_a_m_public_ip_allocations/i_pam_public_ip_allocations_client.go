// Code generated by go-swagger; DO NOT EDIT.

package ip_a_m_public_ip_allocations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new ip a m public ip allocations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ip a m public ip allocations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AllocationCreateUsingPUT(params *AllocationCreateUsingPUTParams, opts ...ClientOption) (*AllocationCreateUsingPUTOK, *AllocationCreateUsingPUTCreated, error)

	AllocationDeleteUsingDELETE(params *AllocationDeleteUsingDELETEParams, opts ...ClientOption) (*AllocationDeleteUsingDELETEOK, error)

	AllocationGetUsingGET(params *AllocationGetUsingGETParams, opts ...ClientOption) (*AllocationGetUsingGETOK, error)

	AllocationListUsingGET(params *AllocationListUsingGETParams, opts ...ClientOption) (*AllocationListUsingGETOK, error)

	AllocationUpdateUsingPUT(params *AllocationUpdateUsingPUTParams, opts ...ClientOption) (*AllocationUpdateUsingPUTOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AllocationCreateUsingPUT creates

Create public IP allocation object
*/
func (a *Client) AllocationCreateUsingPUT(params *AllocationCreateUsingPUTParams, opts ...ClientOption) (*AllocationCreateUsingPUTOK, *AllocationCreateUsingPUTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllocationCreateUsingPUTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "allocationCreateUsingPUT",
		Method:             "PUT",
		PathPattern:        "/ipam/ipaddress/public/allocation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AllocationCreateUsingPUTReader{formats: a.formats},
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
	case *AllocationCreateUsingPUTOK:
		return value, nil, nil
	case *AllocationCreateUsingPUTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ip_a_m_public_ip_allocations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AllocationDeleteUsingDELETE deletes

Delete public IP allocation object
*/
func (a *Client) AllocationDeleteUsingDELETE(params *AllocationDeleteUsingDELETEParams, opts ...ClientOption) (*AllocationDeleteUsingDELETEOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllocationDeleteUsingDELETEParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "allocationDeleteUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/ipam/ipaddress/public/allocation/{allocationId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AllocationDeleteUsingDELETEReader{formats: a.formats},
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
	success, ok := result.(*AllocationDeleteUsingDELETEOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for allocationDeleteUsingDELETE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AllocationGetUsingGET gets

Get allocation of public IP
*/
func (a *Client) AllocationGetUsingGET(params *AllocationGetUsingGETParams, opts ...ClientOption) (*AllocationGetUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllocationGetUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "allocationGetUsingGET",
		Method:             "GET",
		PathPattern:        "/ipam/ipaddress/public/allocation/{allocationId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AllocationGetUsingGETReader{formats: a.formats},
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
	success, ok := result.(*AllocationGetUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for allocationGetUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AllocationListUsingGET lists

List allocations of public IPs
*/
func (a *Client) AllocationListUsingGET(params *AllocationListUsingGETParams, opts ...ClientOption) (*AllocationListUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllocationListUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "allocationListUsingGET",
		Method:             "GET",
		PathPattern:        "/ipam/ipaddress/public/allocation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AllocationListUsingGETReader{formats: a.formats},
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
	success, ok := result.(*AllocationListUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for allocationListUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AllocationUpdateUsingPUT updates

Update public IP allocation object
*/
func (a *Client) AllocationUpdateUsingPUT(params *AllocationUpdateUsingPUTParams, opts ...ClientOption) (*AllocationUpdateUsingPUTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllocationUpdateUsingPUTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "allocationUpdateUsingPUT",
		Method:             "PUT",
		PathPattern:        "/ipam/ipaddress/public/allocation/{allocationId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AllocationUpdateUsingPUTReader{formats: a.formats},
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
	success, ok := result.(*AllocationUpdateUsingPUTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for allocationUpdateUsingPUT: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
