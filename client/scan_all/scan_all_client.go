// Code generated by go-swagger; DO NOT EDIT.

package scan_all

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new scan all API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for scan all API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateScanAllSchedule(params *CreateScanAllScheduleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateScanAllScheduleCreated, error)

	GetLatestScanAllMetrics(params *GetLatestScanAllMetricsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetLatestScanAllMetricsOK, error)

	GetLatestScheduledScanAllMetrics(params *GetLatestScheduledScanAllMetricsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetLatestScheduledScanAllMetricsOK, error)

	GetScanAllSchedule(params *GetScanAllScheduleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScanAllScheduleOK, error)

	StopScanAll(params *StopScanAllParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StopScanAllAccepted, error)

	UpdateScanAllSchedule(params *UpdateScanAllScheduleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateScanAllScheduleOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateScanAllSchedule creates a schedule or a manual trigger for the scan all job

This endpoint is for creating a schedule or a manual trigger for the scan all job, which scans all of images in Harbor.
*/
func (a *Client) CreateScanAllSchedule(params *CreateScanAllScheduleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateScanAllScheduleCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateScanAllScheduleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createScanAllSchedule",
		Method:             "POST",
		PathPattern:        "/system/scanAll/schedule",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateScanAllScheduleReader{formats: a.formats},
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
	success, ok := result.(*CreateScanAllScheduleCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createScanAllSchedule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetLatestScanAllMetrics gets the metrics of the latest scan all process

Get the metrics of the latest scan all process
*/
func (a *Client) GetLatestScanAllMetrics(params *GetLatestScanAllMetricsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetLatestScanAllMetricsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLatestScanAllMetricsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getLatestScanAllMetrics",
		Method:             "GET",
		PathPattern:        "/scans/all/metrics",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLatestScanAllMetricsReader{formats: a.formats},
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
	success, ok := result.(*GetLatestScanAllMetricsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getLatestScanAllMetrics: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetLatestScheduledScanAllMetrics gets the metrics of the latest scheduled scan all process

Get the metrics of the latest scheduled scan all process
*/
func (a *Client) GetLatestScheduledScanAllMetrics(params *GetLatestScheduledScanAllMetricsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetLatestScheduledScanAllMetricsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLatestScheduledScanAllMetricsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getLatestScheduledScanAllMetrics",
		Method:             "GET",
		PathPattern:        "/scans/schedule/metrics",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLatestScheduledScanAllMetricsReader{formats: a.formats},
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
	success, ok := result.(*GetLatestScheduledScanAllMetricsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getLatestScheduledScanAllMetrics: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetScanAllSchedule gets scan all s schedule

This endpoint is for getting a schedule for the scan all job, which scans all of images in Harbor.
*/
func (a *Client) GetScanAllSchedule(params *GetScanAllScheduleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScanAllScheduleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetScanAllScheduleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getScanAllSchedule",
		Method:             "GET",
		PathPattern:        "/system/scanAll/schedule",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetScanAllScheduleReader{formats: a.formats},
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
	success, ok := result.(*GetScanAllScheduleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getScanAllSchedule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StopScanAll stops scan all job execution

Stop scanAll job execution
*/
func (a *Client) StopScanAll(params *StopScanAllParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StopScanAllAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopScanAllParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "stopScanAll",
		Method:             "POST",
		PathPattern:        "/system/scanAll/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopScanAllReader{formats: a.formats},
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
	success, ok := result.(*StopScanAllAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for stopScanAll: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateScanAllSchedule updates scan all s schedule

This endpoint is for updating the schedule of scan all job, which scans all of images in Harbor.
*/
func (a *Client) UpdateScanAllSchedule(params *UpdateScanAllScheduleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateScanAllScheduleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateScanAllScheduleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateScanAllSchedule",
		Method:             "PUT",
		PathPattern:        "/system/scanAll/schedule",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateScanAllScheduleReader{formats: a.formats},
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
	success, ok := result.(*UpdateScanAllScheduleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateScanAllSchedule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
