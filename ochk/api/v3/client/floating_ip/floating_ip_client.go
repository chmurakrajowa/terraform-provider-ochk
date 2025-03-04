// Code generated by go-swagger; DO NOT EDIT.

package floating_ip

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new floating ip API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new floating ip API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new floating ip API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for floating ip API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// This client is generated with a few options you might find useful for your swagger spec.
//
// Feel free to add you own set of options.

// WithContentType allows the client to force the Content-Type header
// to negotiate a specific Consumer from the server.
//
// You may use this option to set arbitrary extensions to your MIME media type.
func WithContentType(mime string) ClientOption {
	return func(r *runtime.ClientOperation) {
		r.ConsumesMediaTypes = []string{mime}
	}
}

// WithContentTypeApplicationStarJSON sets the Content-Type header to "application/*+json".
func WithContentTypeApplicationStarJSON(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"application/*+json"}
}

// WithContentTypeApplicationJSON sets the Content-Type header to "application/json".
func WithContentTypeApplicationJSON(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"application/json"}
}

// WithContentTypeTextJSON sets the Content-Type header to "text/json".
func WithContentTypeTextJSON(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"text/json"}
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteNetworkFloatingIpsFloatingIPID(params *DeleteNetworkFloatingIpsFloatingIPIDParams, opts ...ClientOption) (*DeleteNetworkFloatingIpsFloatingIPIDOK, error)

	GetNetworkFloatingIps(params *GetNetworkFloatingIpsParams, opts ...ClientOption) (*GetNetworkFloatingIpsOK, error)

	GetNetworkFloatingIpsFloatingIPID(params *GetNetworkFloatingIpsFloatingIPIDParams, opts ...ClientOption) (*GetNetworkFloatingIpsFloatingIPIDOK, error)

	PutNetworkFloatingIps(params *PutNetworkFloatingIpsParams, opts ...ClientOption) (*PutNetworkFloatingIpsOK, error)

	PutNetworkFloatingIpsFloatingIPID(params *PutNetworkFloatingIpsFloatingIPIDParams, opts ...ClientOption) (*PutNetworkFloatingIpsFloatingIPIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteNetworkFloatingIpsFloatingIPID deletes firewall rule o s c

Delete firewall rule (OSC)
*/
func (a *Client) DeleteNetworkFloatingIpsFloatingIPID(params *DeleteNetworkFloatingIpsFloatingIPIDParams, opts ...ClientOption) (*DeleteNetworkFloatingIpsFloatingIPIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNetworkFloatingIpsFloatingIPIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteNetworkFloatingIpsFloatingIPID",
		Method:             "DELETE",
		PathPattern:        "/network/floating-ips/{floatingIpId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteNetworkFloatingIpsFloatingIPIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteNetworkFloatingIpsFloatingIPIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteNetworkFloatingIpsFloatingIPID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetNetworkFloatingIps lists floating ips o s c

List floating ips (OSC)
*/
func (a *Client) GetNetworkFloatingIps(params *GetNetworkFloatingIpsParams, opts ...ClientOption) (*GetNetworkFloatingIpsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkFloatingIpsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetNetworkFloatingIps",
		Method:             "GET",
		PathPattern:        "/network/floating-ips",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetNetworkFloatingIpsReader{formats: a.formats},
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
	success, ok := result.(*GetNetworkFloatingIpsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetNetworkFloatingIps: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetNetworkFloatingIpsFloatingIPID gets floating ip o s c

Get floating ip (OSC)
*/
func (a *Client) GetNetworkFloatingIpsFloatingIPID(params *GetNetworkFloatingIpsFloatingIPIDParams, opts ...ClientOption) (*GetNetworkFloatingIpsFloatingIPIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkFloatingIpsFloatingIPIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetNetworkFloatingIpsFloatingIPID",
		Method:             "GET",
		PathPattern:        "/network/floating-ips/{floatingIpId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetNetworkFloatingIpsFloatingIPIDReader{formats: a.formats},
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
	success, ok := result.(*GetNetworkFloatingIpsFloatingIPIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetNetworkFloatingIpsFloatingIPID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutNetworkFloatingIps creates floating ip o s c

Create floating ip (OSC)
*/
func (a *Client) PutNetworkFloatingIps(params *PutNetworkFloatingIpsParams, opts ...ClientOption) (*PutNetworkFloatingIpsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutNetworkFloatingIpsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutNetworkFloatingIps",
		Method:             "PUT",
		PathPattern:        "/network/floating-ips",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/json", "application/*+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutNetworkFloatingIpsReader{formats: a.formats},
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
	success, ok := result.(*PutNetworkFloatingIpsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutNetworkFloatingIps: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutNetworkFloatingIpsFloatingIPID updates firewall rule o s c

Update firewall rule (OSC)
*/
func (a *Client) PutNetworkFloatingIpsFloatingIPID(params *PutNetworkFloatingIpsFloatingIPIDParams, opts ...ClientOption) (*PutNetworkFloatingIpsFloatingIPIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutNetworkFloatingIpsFloatingIPIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutNetworkFloatingIpsFloatingIPID",
		Method:             "PUT",
		PathPattern:        "/network/floating-ips/{floatingIpId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/json", "application/*+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutNetworkFloatingIpsFloatingIPIDReader{formats: a.formats},
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
	success, ok := result.(*PutNetworkFloatingIpsFloatingIPIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutNetworkFloatingIpsFloatingIPID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
