// Code generated by go-swagger; DO NOT EDIT.

package billing_periods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new billing periods API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for billing periods API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetBillingPeriodUsingGET(params *GetBillingPeriodUsingGETParams, opts ...ClientOption) (*GetBillingPeriodUsingGETOK, error)

	ListBillingPeriodsUsingGET(params *ListBillingPeriodsUsingGETParams, opts ...ClientOption) (*ListBillingPeriodsUsingGETOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetBillingPeriodUsingGET gets billing period
*/
func (a *Client) GetBillingPeriodUsingGET(params *GetBillingPeriodUsingGETParams, opts ...ClientOption) (*GetBillingPeriodUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBillingPeriodUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBillingPeriodUsingGET",
		Method:             "GET",
		PathPattern:        "/billing/accounts/billing-periods/{billingPeriodId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBillingPeriodUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetBillingPeriodUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBillingPeriodUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListBillingPeriodsUsingGET lists billing periods
*/
func (a *Client) ListBillingPeriodsUsingGET(params *ListBillingPeriodsUsingGETParams, opts ...ClientOption) (*ListBillingPeriodsUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListBillingPeriodsUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listBillingPeriodsUsingGET",
		Method:             "GET",
		PathPattern:        "/billing/accounts/billing-periods",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListBillingPeriodsUsingGETReader{formats: a.formats},
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
	success, ok := result.(*ListBillingPeriodsUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listBillingPeriodsUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}