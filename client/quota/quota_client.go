// Code generated by go-swagger; DO NOT EDIT.

package quota

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new quota API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for quota API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetQuota(params *GetQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetQuotaOK, error)

	ListQuotas(params *ListQuotasParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListQuotasOK, error)

	UpdateQuota(params *UpdateQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateQuotaOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetQuota gets the specified quota

Get the specified quota
*/
func (a *Client) GetQuota(params *GetQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetQuotaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetQuotaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getQuota",
		Method:             "GET",
		PathPattern:        "/quotas/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetQuotaReader{formats: a.formats},
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
	success, ok := result.(*GetQuotaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getQuota: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListQuotas lists quotas

List quotas
*/
func (a *Client) ListQuotas(params *ListQuotasParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListQuotasOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListQuotasParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listQuotas",
		Method:             "GET",
		PathPattern:        "/quotas",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListQuotasReader{formats: a.formats},
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
	success, ok := result.(*ListQuotasOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listQuotas: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateQuota updates the specified quota

Update hard limits of the specified quota
*/
func (a *Client) UpdateQuota(params *UpdateQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateQuotaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateQuotaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateQuota",
		Method:             "PUT",
		PathPattern:        "/quotas/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateQuotaReader{formats: a.formats},
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
	success, ok := result.(*UpdateQuotaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateQuota: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
