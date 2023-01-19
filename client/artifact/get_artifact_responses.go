// Code generated by go-swagger; DO NOT EDIT.

package artifact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/middlewaregruppen/harbor-operator/models"
)

// GetArtifactReader is a Reader for the GetArtifact structure.
type GetArtifactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetArtifactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetArtifactOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetArtifactBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetArtifactUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetArtifactForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetArtifactNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetArtifactInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetArtifactOK creates a GetArtifactOK with default headers values
func NewGetArtifactOK() *GetArtifactOK {
	return &GetArtifactOK{}
}

/*
GetArtifactOK describes a response with status code 200, with default header values.

Success
*/
type GetArtifactOK struct {
	Payload *models.Artifact
}

// IsSuccess returns true when this get artifact o k response has a 2xx status code
func (o *GetArtifactOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get artifact o k response has a 3xx status code
func (o *GetArtifactOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get artifact o k response has a 4xx status code
func (o *GetArtifactOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get artifact o k response has a 5xx status code
func (o *GetArtifactOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get artifact o k response a status code equal to that given
func (o *GetArtifactOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetArtifactOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}][%d] getArtifactOK  %+v", 200, o.Payload)
}

func (o *GetArtifactOK) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}][%d] getArtifactOK  %+v", 200, o.Payload)
}

func (o *GetArtifactOK) GetPayload() *models.Artifact {
	return o.Payload
}

func (o *GetArtifactOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Artifact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArtifactBadRequest creates a GetArtifactBadRequest with default headers values
func NewGetArtifactBadRequest() *GetArtifactBadRequest {
	return &GetArtifactBadRequest{}
}

/*
GetArtifactBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetArtifactBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get artifact bad request response has a 2xx status code
func (o *GetArtifactBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get artifact bad request response has a 3xx status code
func (o *GetArtifactBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get artifact bad request response has a 4xx status code
func (o *GetArtifactBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get artifact bad request response has a 5xx status code
func (o *GetArtifactBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get artifact bad request response a status code equal to that given
func (o *GetArtifactBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetArtifactBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}][%d] getArtifactBadRequest  %+v", 400, o.Payload)
}

func (o *GetArtifactBadRequest) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}][%d] getArtifactBadRequest  %+v", 400, o.Payload)
}

func (o *GetArtifactBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetArtifactBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArtifactUnauthorized creates a GetArtifactUnauthorized with default headers values
func NewGetArtifactUnauthorized() *GetArtifactUnauthorized {
	return &GetArtifactUnauthorized{}
}

/*
GetArtifactUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetArtifactUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get artifact unauthorized response has a 2xx status code
func (o *GetArtifactUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get artifact unauthorized response has a 3xx status code
func (o *GetArtifactUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get artifact unauthorized response has a 4xx status code
func (o *GetArtifactUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get artifact unauthorized response has a 5xx status code
func (o *GetArtifactUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get artifact unauthorized response a status code equal to that given
func (o *GetArtifactUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetArtifactUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}][%d] getArtifactUnauthorized  %+v", 401, o.Payload)
}

func (o *GetArtifactUnauthorized) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}][%d] getArtifactUnauthorized  %+v", 401, o.Payload)
}

func (o *GetArtifactUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetArtifactUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArtifactForbidden creates a GetArtifactForbidden with default headers values
func NewGetArtifactForbidden() *GetArtifactForbidden {
	return &GetArtifactForbidden{}
}

/*
GetArtifactForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetArtifactForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get artifact forbidden response has a 2xx status code
func (o *GetArtifactForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get artifact forbidden response has a 3xx status code
func (o *GetArtifactForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get artifact forbidden response has a 4xx status code
func (o *GetArtifactForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get artifact forbidden response has a 5xx status code
func (o *GetArtifactForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get artifact forbidden response a status code equal to that given
func (o *GetArtifactForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetArtifactForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}][%d] getArtifactForbidden  %+v", 403, o.Payload)
}

func (o *GetArtifactForbidden) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}][%d] getArtifactForbidden  %+v", 403, o.Payload)
}

func (o *GetArtifactForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetArtifactForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArtifactNotFound creates a GetArtifactNotFound with default headers values
func NewGetArtifactNotFound() *GetArtifactNotFound {
	return &GetArtifactNotFound{}
}

/*
GetArtifactNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetArtifactNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get artifact not found response has a 2xx status code
func (o *GetArtifactNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get artifact not found response has a 3xx status code
func (o *GetArtifactNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get artifact not found response has a 4xx status code
func (o *GetArtifactNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get artifact not found response has a 5xx status code
func (o *GetArtifactNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get artifact not found response a status code equal to that given
func (o *GetArtifactNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetArtifactNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}][%d] getArtifactNotFound  %+v", 404, o.Payload)
}

func (o *GetArtifactNotFound) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}][%d] getArtifactNotFound  %+v", 404, o.Payload)
}

func (o *GetArtifactNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetArtifactNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArtifactInternalServerError creates a GetArtifactInternalServerError with default headers values
func NewGetArtifactInternalServerError() *GetArtifactInternalServerError {
	return &GetArtifactInternalServerError{}
}

/*
GetArtifactInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetArtifactInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get artifact internal server error response has a 2xx status code
func (o *GetArtifactInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get artifact internal server error response has a 3xx status code
func (o *GetArtifactInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get artifact internal server error response has a 4xx status code
func (o *GetArtifactInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get artifact internal server error response has a 5xx status code
func (o *GetArtifactInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get artifact internal server error response a status code equal to that given
func (o *GetArtifactInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetArtifactInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}][%d] getArtifactInternalServerError  %+v", 500, o.Payload)
}

func (o *GetArtifactInternalServerError) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}][%d] getArtifactInternalServerError  %+v", 500, o.Payload)
}

func (o *GetArtifactInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetArtifactInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
