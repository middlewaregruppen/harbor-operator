// Code generated by go-swagger; DO NOT EDIT.

package schedule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new schedule API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for schedule API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetSchedulePaused(params *GetSchedulePausedParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSchedulePausedOK, error)

	ListSchedules(params *ListSchedulesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSchedulesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetSchedulePaused Get scheduler paused status
*/
func (a *Client) GetSchedulePaused(params *GetSchedulePausedParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSchedulePausedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSchedulePausedParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSchedulePaused",
		Method:             "GET",
		PathPattern:        "/schedules/{job_type}/paused",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSchedulePausedReader{formats: a.formats},
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
	success, ok := result.(*GetSchedulePausedOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSchedulePaused: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListSchedules List schedules
*/
func (a *Client) ListSchedules(params *ListSchedulesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSchedulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSchedulesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listSchedules",
		Method:             "GET",
		PathPattern:        "/schedules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListSchedulesReader{formats: a.formats},
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
	success, ok := result.(*ListSchedulesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listSchedules: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
