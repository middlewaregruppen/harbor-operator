// Code generated by go-swagger; DO NOT EDIT.

package artifact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new artifact API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for artifact API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CopyArtifact(params *CopyArtifactParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CopyArtifactCreated, error)

	AddLabel(params *AddLabelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddLabelOK, error)

	CreateTag(params *CreateTagParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateTagCreated, error)

	DeleteArtifact(params *DeleteArtifactParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteArtifactOK, error)

	DeleteTag(params *DeleteTagParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteTagOK, error)

	GetAddition(params *GetAdditionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAdditionOK, error)

	GetArtifact(params *GetArtifactParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetArtifactOK, error)

	GetVulnerabilitiesAddition(params *GetVulnerabilitiesAdditionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVulnerabilitiesAdditionOK, error)

	ListAccessories(params *ListAccessoriesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAccessoriesOK, error)

	ListArtifacts(params *ListArtifactsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListArtifactsOK, error)

	ListTags(params *ListTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListTagsOK, error)

	RemoveLabel(params *RemoveLabelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RemoveLabelOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CopyArtifact copies artifact

Copy the artifact specified in the "from" parameter to the repository.
*/
func (a *Client) CopyArtifact(params *CopyArtifactParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CopyArtifactCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCopyArtifactParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CopyArtifact",
		Method:             "POST",
		PathPattern:        "/projects/{project_name}/repositories/{repository_name}/artifacts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CopyArtifactReader{formats: a.formats},
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
	success, ok := result.(*CopyArtifactCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CopyArtifact: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AddLabel adds label to artifact

Add label to the specified artiact.
*/
func (a *Client) AddLabel(params *AddLabelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddLabelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddLabelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addLabel",
		Method:             "POST",
		PathPattern:        "/projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddLabelReader{formats: a.formats},
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
	success, ok := result.(*AddLabelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addLabel: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateTag creates tag

Create a tag for the specified artifact
*/
func (a *Client) CreateTag(params *CreateTagParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateTagCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateTagParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createTag",
		Method:             "POST",
		PathPattern:        "/projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateTagReader{formats: a.formats},
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
	success, ok := result.(*CreateTagCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createTag: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteArtifact deletes the specific artifact

Delete the artifact specified by the reference under the project and repository. The reference can be digest or tag
*/
func (a *Client) DeleteArtifact(params *DeleteArtifactParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteArtifactOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteArtifactParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteArtifact",
		Method:             "DELETE",
		PathPattern:        "/projects/{project_name}/repositories/{repository_name}/artifacts/{reference}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteArtifactReader{formats: a.formats},
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
	success, ok := result.(*DeleteArtifactOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteArtifact: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteTag deletes tag

Delete the tag of the specified artifact
*/
func (a *Client) DeleteTag(params *DeleteTagParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteTagOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTagParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteTag",
		Method:             "DELETE",
		PathPattern:        "/projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags/{tag_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteTagReader{formats: a.formats},
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
	success, ok := result.(*DeleteTagOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteTag: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAddition gets the addition of the specific artifact

Get the addition of the artifact specified by the reference under the project and repository.
*/
func (a *Client) GetAddition(params *GetAdditionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAdditionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAdditionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAddition",
		Method:             "GET",
		PathPattern:        "/projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/additions/{addition}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAdditionReader{formats: a.formats},
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
	success, ok := result.(*GetAdditionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAddition: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetArtifact gets the specific artifact

Get the artifact specified by the reference under the project and repository. The reference can be digest or tag.
*/
func (a *Client) GetArtifact(params *GetArtifactParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetArtifactOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetArtifactParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getArtifact",
		Method:             "GET",
		PathPattern:        "/projects/{project_name}/repositories/{repository_name}/artifacts/{reference}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetArtifactReader{formats: a.formats},
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
	success, ok := result.(*GetArtifactOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getArtifact: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetVulnerabilitiesAddition gets the vulnerabilities addition of the specific artifact

Get the vulnerabilities addition of the artifact specified by the reference under the project and repository.
*/
func (a *Client) GetVulnerabilitiesAddition(params *GetVulnerabilitiesAdditionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVulnerabilitiesAdditionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVulnerabilitiesAdditionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getVulnerabilitiesAddition",
		Method:             "GET",
		PathPattern:        "/projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/additions/vulnerabilities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetVulnerabilitiesAdditionReader{formats: a.formats},
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
	success, ok := result.(*GetVulnerabilitiesAdditionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVulnerabilitiesAddition: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListAccessories lists accessories

List accessories of the specific artifact
*/
func (a *Client) ListAccessories(params *ListAccessoriesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAccessoriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAccessoriesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAccessories",
		Method:             "GET",
		PathPattern:        "/projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/accessories",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListAccessoriesReader{formats: a.formats},
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
	success, ok := result.(*ListAccessoriesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listAccessories: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListArtifacts lists artifacts

List artifacts under the specific project and repository. Except the basic properties, the other supported queries in "q" includes "tags=*" to list only tagged artifacts, "tags=nil" to list only untagged artifacts, "tags=~v" to list artifacts whose tag fuzzy matches "v", "tags=v" to list artifact whose tag exactly matches "v", "labels=(id1, id2)" to list artifacts that both labels with id1 and id2 are added to
*/
func (a *Client) ListArtifacts(params *ListArtifactsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListArtifactsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListArtifactsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listArtifacts",
		Method:             "GET",
		PathPattern:        "/projects/{project_name}/repositories/{repository_name}/artifacts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListArtifactsReader{formats: a.formats},
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
	success, ok := result.(*ListArtifactsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listArtifacts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListTags lists tags

List tags of the specific artifact
*/
func (a *Client) ListTags(params *ListTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListTagsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListTagsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listTags",
		Method:             "GET",
		PathPattern:        "/projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListTagsReader{formats: a.formats},
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
	success, ok := result.(*ListTagsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listTags: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RemoveLabel removes label from artifact

Remove the label from the specified artiact.
*/
func (a *Client) RemoveLabel(params *RemoveLabelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RemoveLabelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveLabelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "removeLabel",
		Method:             "DELETE",
		PathPattern:        "/projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels/{label_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveLabelReader{formats: a.formats},
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
	success, ok := result.(*RemoveLabelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeLabel: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
