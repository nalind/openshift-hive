// Code generated by go-swagger; DO NOT EDIT.

package internal_operations_shared_processor_pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new internal operations shared processor pools API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new internal operations shared processor pools API client with basic auth credentials.
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

// New creates a new internal operations shared processor pools API client with a bearer token for authentication.
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
Client for internal operations shared processor pools API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	InternalV1OperationsSharedprocessorpoolsDelete(params *InternalV1OperationsSharedprocessorpoolsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InternalV1OperationsSharedprocessorpoolsDeleteNoContent, error)

	InternalV1OperationsSharedprocessorpoolsPost(params *InternalV1OperationsSharedprocessorpoolsPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InternalV1OperationsSharedprocessorpoolsPostCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
InternalV1OperationsSharedprocessorpoolsDelete deletes a shared processor pool c r n
*/
func (a *Client) InternalV1OperationsSharedprocessorpoolsDelete(params *InternalV1OperationsSharedprocessorpoolsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InternalV1OperationsSharedprocessorpoolsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInternalV1OperationsSharedprocessorpoolsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "internal.v1.operations.sharedprocessorpools.delete",
		Method:             "DELETE",
		PathPattern:        "/internal/v1/operations/shared-processor-pools/{resource_crn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &InternalV1OperationsSharedprocessorpoolsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*InternalV1OperationsSharedprocessorpoolsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for internal.v1.operations.sharedprocessorpools.delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
InternalV1OperationsSharedprocessorpoolsPost creates a c r n for a shared processor pool
*/
func (a *Client) InternalV1OperationsSharedprocessorpoolsPost(params *InternalV1OperationsSharedprocessorpoolsPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InternalV1OperationsSharedprocessorpoolsPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInternalV1OperationsSharedprocessorpoolsPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "internal.v1.operations.sharedprocessorpools.post",
		Method:             "POST",
		PathPattern:        "/internal/v1/operations/shared-processor-pools",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &InternalV1OperationsSharedprocessorpoolsPostReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*InternalV1OperationsSharedprocessorpoolsPostCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for internal.v1.operations.sharedprocessorpools.post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
