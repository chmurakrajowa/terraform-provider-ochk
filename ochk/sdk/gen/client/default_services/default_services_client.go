// Code generated by go-swagger; DO NOT EDIT.

package default_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new default services API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for default services API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ServiceGetUsingGET(params *ServiceGetUsingGETParams) (*ServiceGetUsingGETOK, error)

	ServiceListUsingGET(params *ServiceListUsingGETParams) (*ServiceListUsingGETOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ServiceGetUsingGET gets

  Get default service from NSX-T
*/
func (a *Client) ServiceGetUsingGET(params *ServiceGetUsingGETParams) (*ServiceGetUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceGetUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serviceGetUsingGET",
		Method:             "GET",
		PathPattern:        "/network/default-services/{serviceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceGetUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ServiceGetUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for serviceGetUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ServiceListUsingGET lists

  List default services from NSX-T
*/
func (a *Client) ServiceListUsingGET(params *ServiceListUsingGETParams) (*ServiceListUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceListUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serviceListUsingGET",
		Method:             "GET",
		PathPattern:        "/network/default-services",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceListUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ServiceListUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for serviceListUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
