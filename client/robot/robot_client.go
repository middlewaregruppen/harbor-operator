// Code generated by go-swagger; DO NOT EDIT.

package robot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new robot API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for robot API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateRobot(params *CreateRobotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateRobotCreated, error)

	DeleteRobot(params *DeleteRobotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteRobotOK, error)

	GetRobotByID(params *GetRobotByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRobotByIDOK, error)

	ListRobot(params *ListRobotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListRobotOK, error)

	RefreshSec(params *RefreshSecParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RefreshSecOK, error)

	UpdateRobot(params *UpdateRobotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateRobotOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateRobot creates a robot account

Create a robot account
*/
func (a *Client) CreateRobot(params *CreateRobotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateRobotCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRobotParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateRobot",
		Method:             "POST",
		PathPattern:        "/robots",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateRobotReader{formats: a.formats},
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
	success, ok := result.(*CreateRobotCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateRobot: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteRobot deletes a robot account

This endpoint deletes specific robot account information by robot ID.
*/
func (a *Client) DeleteRobot(params *DeleteRobotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteRobotOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRobotParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteRobot",
		Method:             "DELETE",
		PathPattern:        "/robots/{robot_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteRobotReader{formats: a.formats},
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
	success, ok := result.(*DeleteRobotOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteRobot: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetRobotByID gets a robot account

This endpoint returns specific robot account information by robot ID.
*/
func (a *Client) GetRobotByID(params *GetRobotByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRobotByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRobotByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetRobotByID",
		Method:             "GET",
		PathPattern:        "/robots/{robot_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRobotByIDReader{formats: a.formats},
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
	success, ok := result.(*GetRobotByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetRobotByID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListRobot gets robot account

List the robot accounts with the specified level and project.
*/
func (a *Client) ListRobot(params *ListRobotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListRobotOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRobotParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListRobot",
		Method:             "GET",
		PathPattern:        "/robots",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListRobotReader{formats: a.formats},
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
	success, ok := result.(*ListRobotOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListRobot: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RefreshSec refreshes the robot secret

Refresh the robot secret
*/
func (a *Client) RefreshSec(params *RefreshSecParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RefreshSecOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRefreshSecParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RefreshSec",
		Method:             "PATCH",
		PathPattern:        "/robots/{robot_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RefreshSecReader{formats: a.formats},
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
	success, ok := result.(*RefreshSecOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RefreshSec: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateRobot updates a robot account

This endpoint updates specific robot account information by robot ID.
*/
func (a *Client) UpdateRobot(params *UpdateRobotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateRobotOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRobotParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateRobot",
		Method:             "PUT",
		PathPattern:        "/robots/{robot_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateRobotReader{formats: a.formats},
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
	success, ok := result.(*UpdateRobotOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateRobot: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
