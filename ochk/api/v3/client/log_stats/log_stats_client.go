// Code generated by go-swagger; DO NOT EDIT.

package log_stats

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new log stats API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new log stats API client with basic auth credentials.
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

// New creates a new log stats API client with a bearer token for authentication.
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
Client for log stats API
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
	PostLogStats(params *PostLogStatsParams, opts ...ClientOption) (*PostLogStatsOK, error)

	PostLogStatsLogCategoryID(params *PostLogStatsLogCategoryIDParams, opts ...ClientOption) (*PostLogStatsLogCategoryIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
PostLogStats gets logs stats

Get logs stats
*/
func (a *Client) PostLogStats(params *PostLogStatsParams, opts ...ClientOption) (*PostLogStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLogStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostLogStats",
		Method:             "POST",
		PathPattern:        "/log/stats",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/json", "application/*+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostLogStatsReader{formats: a.formats},
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
	success, ok := result.(*PostLogStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostLogStats: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostLogStatsLogCategoryID gets log stats per category

Get log stats per category
*/
func (a *Client) PostLogStatsLogCategoryID(params *PostLogStatsLogCategoryIDParams, opts ...ClientOption) (*PostLogStatsLogCategoryIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLogStatsLogCategoryIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostLogStatsLogCategoryID",
		Method:             "POST",
		PathPattern:        "/log/stats/{logCategoryId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/json", "application/*+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostLogStatsLogCategoryIDReader{formats: a.formats},
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
	success, ok := result.(*PostLogStatsLogCategoryIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostLogStatsLogCategoryID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
