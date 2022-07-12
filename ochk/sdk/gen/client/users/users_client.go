// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for users API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	UserGetUsingGET(params *UserGetUsingGETParams, opts ...ClientOption) (*UserGetUsingGETOK, error)

	UserListUsingGET(params *UserListUsingGETParams, opts ...ClientOption) (*UserListUsingGETOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  UserGetUsingGET gets

  Get user
*/
func (a *Client) UserGetUsingGET(params *UserGetUsingGETParams, opts ...ClientOption) (*UserGetUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserGetUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "userGetUsingGET",
		Method:             "GET",
		PathPattern:        "/users/{userId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserGetUsingGETReader{formats: a.formats},
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
	success, ok := result.(*UserGetUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for userGetUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UserListUsingGET lists

  List users from vra/core
*/
func (a *Client) UserListUsingGET(params *UserListUsingGETParams, opts ...ClientOption) (*UserListUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserListUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "userListUsingGET",
		Method:             "GET",
		PathPattern:        "/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserListUsingGETReader{formats: a.formats},
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
	success, ok := result.(*UserListUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for userListUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
