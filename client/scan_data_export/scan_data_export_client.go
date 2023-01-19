// Code generated by go-swagger; DO NOT EDIT.

package scan_data_export

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new scan data export API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for scan data export API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DownloadScanData(params *DownloadScanDataParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer, opts ...ClientOption) (*DownloadScanDataOK, error)

	ExportScanData(params *ExportScanDataParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExportScanDataOK, error)

	GetScanDataExportExecution(params *GetScanDataExportExecutionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScanDataExportExecutionOK, error)

	GetScanDataExportExecutionList(params *GetScanDataExportExecutionListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScanDataExportExecutionListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DownloadScanData downloads the scan data export file

Download the scan data report. Default format is CSV
*/
func (a *Client) DownloadScanData(params *DownloadScanDataParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer, opts ...ClientOption) (*DownloadScanDataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDownloadScanDataParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "downloadScanData",
		Method:             "GET",
		PathPattern:        "/export/cve/download/{execution_id}",
		ProducesMediaTypes: []string{"text/csv"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DownloadScanDataReader{formats: a.formats, writer: writer},
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
	success, ok := result.(*DownloadScanDataOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for downloadScanData: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ExportScanData exports scan data for selected projects

Export scan data for selected projects
*/
func (a *Client) ExportScanData(params *ExportScanDataParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExportScanDataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExportScanDataParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "exportScanData",
		Method:             "POST",
		PathPattern:        "/export/cve",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ExportScanDataReader{formats: a.formats},
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
	success, ok := result.(*ExportScanDataOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for exportScanData: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetScanDataExportExecution gets the specific scan data export execution

Get the scan data export execution specified by ID
*/
func (a *Client) GetScanDataExportExecution(params *GetScanDataExportExecutionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScanDataExportExecutionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetScanDataExportExecutionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getScanDataExportExecution",
		Method:             "GET",
		PathPattern:        "/export/cve/execution/{execution_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetScanDataExportExecutionReader{formats: a.formats},
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
	success, ok := result.(*GetScanDataExportExecutionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getScanDataExportExecution: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetScanDataExportExecutionList gets a list of specific scan data export execution jobs for a specified user

Get a list of specific scan data export execution jobs for a specified user
*/
func (a *Client) GetScanDataExportExecutionList(params *GetScanDataExportExecutionListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScanDataExportExecutionListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetScanDataExportExecutionListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getScanDataExportExecutionList",
		Method:             "GET",
		PathPattern:        "/export/cve/executions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetScanDataExportExecutionListReader{formats: a.formats},
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
	success, ok := result.(*GetScanDataExportExecutionListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getScanDataExportExecutionList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
