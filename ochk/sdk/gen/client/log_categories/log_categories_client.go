// Code generated by go-swagger; DO NOT EDIT.

package log_categories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new log categories API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for log categories API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetLogsUsingPOST(params *GetLogsUsingPOSTParams, opts ...ClientOption) (*GetLogsUsingPOSTOK, error)

	LogCategoriesListUsingGET(params *LogCategoriesListUsingGETParams, opts ...ClientOption) (*LogCategoriesListUsingGETOK, error)

	LogCategoryGetUsingGET(params *LogCategoryGetUsingGETParams, opts ...ClientOption) (*LogCategoryGetUsingGETOK, error)

	LogCategoryUpdateUsingPUT(params *LogCategoryUpdateUsingPUTParams, opts ...ClientOption) (*LogCategoryUpdateUsingPUTOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetLogsUsingPOST gets logs

  Get logs
*/
func (a *Client) GetLogsUsingPOST(params *GetLogsUsingPOSTParams, opts ...ClientOption) (*GetLogsUsingPOSTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLogsUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getLogsUsingPOST",
		Method:             "POST",
		PathPattern:        "/log/categories/{logCategoryId}/generate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLogsUsingPOSTReader{formats: a.formats},
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
	success, ok := result.(*GetLogsUsingPOSTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getLogsUsingPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LogCategoriesListUsingGET lists

  List log categories
*/
func (a *Client) LogCategoriesListUsingGET(params *LogCategoriesListUsingGETParams, opts ...ClientOption) (*LogCategoriesListUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLogCategoriesListUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "logCategoriesListUsingGET",
		Method:             "GET",
		PathPattern:        "/log/categories",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LogCategoriesListUsingGETReader{formats: a.formats},
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
	success, ok := result.(*LogCategoriesListUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for logCategoriesListUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LogCategoryGetUsingGET gets

  Get log category
*/
func (a *Client) LogCategoryGetUsingGET(params *LogCategoryGetUsingGETParams, opts ...ClientOption) (*LogCategoryGetUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLogCategoryGetUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "logCategoryGetUsingGET",
		Method:             "GET",
		PathPattern:        "/log/categories/{logCategoryId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LogCategoryGetUsingGETReader{formats: a.formats},
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
	success, ok := result.(*LogCategoryGetUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for logCategoryGetUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LogCategoryUpdateUsingPUT updates

  Update log category
*/
func (a *Client) LogCategoryUpdateUsingPUT(params *LogCategoryUpdateUsingPUTParams, opts ...ClientOption) (*LogCategoryUpdateUsingPUTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLogCategoryUpdateUsingPUTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "logCategoryUpdateUsingPUT",
		Method:             "PUT",
		PathPattern:        "/log/categories/{logCategoryId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LogCategoryUpdateUsingPUTReader{formats: a.formats},
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
	success, ok := result.(*LogCategoryUpdateUsingPUTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for logCategoryUpdateUsingPUT: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
