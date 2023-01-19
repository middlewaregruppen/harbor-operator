// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/middlewaregruppen/harbor-operator/models"
)

// DeleteProjectMemberReader is a Reader for the DeleteProjectMember structure.
type DeleteProjectMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteProjectMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteProjectMemberBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteProjectMemberUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteProjectMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteProjectMemberInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteProjectMemberOK creates a DeleteProjectMemberOK with default headers values
func NewDeleteProjectMemberOK() *DeleteProjectMemberOK {
	return &DeleteProjectMemberOK{}
}

/*
DeleteProjectMemberOK describes a response with status code 200, with default header values.

Success
*/
type DeleteProjectMemberOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this delete project member o k response has a 2xx status code
func (o *DeleteProjectMemberOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete project member o k response has a 3xx status code
func (o *DeleteProjectMemberOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project member o k response has a 4xx status code
func (o *DeleteProjectMemberOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete project member o k response has a 5xx status code
func (o *DeleteProjectMemberOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project member o k response a status code equal to that given
func (o *DeleteProjectMemberOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteProjectMemberOK) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/members/{mid}][%d] deleteProjectMemberOK ", 200)
}

func (o *DeleteProjectMemberOK) String() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/members/{mid}][%d] deleteProjectMemberOK ", 200)
}

func (o *DeleteProjectMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewDeleteProjectMemberBadRequest creates a DeleteProjectMemberBadRequest with default headers values
func NewDeleteProjectMemberBadRequest() *DeleteProjectMemberBadRequest {
	return &DeleteProjectMemberBadRequest{}
}

/*
DeleteProjectMemberBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteProjectMemberBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete project member bad request response has a 2xx status code
func (o *DeleteProjectMemberBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete project member bad request response has a 3xx status code
func (o *DeleteProjectMemberBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project member bad request response has a 4xx status code
func (o *DeleteProjectMemberBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete project member bad request response has a 5xx status code
func (o *DeleteProjectMemberBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project member bad request response a status code equal to that given
func (o *DeleteProjectMemberBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteProjectMemberBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/members/{mid}][%d] deleteProjectMemberBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteProjectMemberBadRequest) String() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/members/{mid}][%d] deleteProjectMemberBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteProjectMemberBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteProjectMemberBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteProjectMemberUnauthorized creates a DeleteProjectMemberUnauthorized with default headers values
func NewDeleteProjectMemberUnauthorized() *DeleteProjectMemberUnauthorized {
	return &DeleteProjectMemberUnauthorized{}
}

/*
DeleteProjectMemberUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteProjectMemberUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete project member unauthorized response has a 2xx status code
func (o *DeleteProjectMemberUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete project member unauthorized response has a 3xx status code
func (o *DeleteProjectMemberUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project member unauthorized response has a 4xx status code
func (o *DeleteProjectMemberUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete project member unauthorized response has a 5xx status code
func (o *DeleteProjectMemberUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project member unauthorized response a status code equal to that given
func (o *DeleteProjectMemberUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteProjectMemberUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/members/{mid}][%d] deleteProjectMemberUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteProjectMemberUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/members/{mid}][%d] deleteProjectMemberUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteProjectMemberUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteProjectMemberUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteProjectMemberForbidden creates a DeleteProjectMemberForbidden with default headers values
func NewDeleteProjectMemberForbidden() *DeleteProjectMemberForbidden {
	return &DeleteProjectMemberForbidden{}
}

/*
DeleteProjectMemberForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteProjectMemberForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete project member forbidden response has a 2xx status code
func (o *DeleteProjectMemberForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete project member forbidden response has a 3xx status code
func (o *DeleteProjectMemberForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project member forbidden response has a 4xx status code
func (o *DeleteProjectMemberForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete project member forbidden response has a 5xx status code
func (o *DeleteProjectMemberForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project member forbidden response a status code equal to that given
func (o *DeleteProjectMemberForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteProjectMemberForbidden) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/members/{mid}][%d] deleteProjectMemberForbidden  %+v", 403, o.Payload)
}

func (o *DeleteProjectMemberForbidden) String() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/members/{mid}][%d] deleteProjectMemberForbidden  %+v", 403, o.Payload)
}

func (o *DeleteProjectMemberForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteProjectMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteProjectMemberInternalServerError creates a DeleteProjectMemberInternalServerError with default headers values
func NewDeleteProjectMemberInternalServerError() *DeleteProjectMemberInternalServerError {
	return &DeleteProjectMemberInternalServerError{}
}

/*
DeleteProjectMemberInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteProjectMemberInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete project member internal server error response has a 2xx status code
func (o *DeleteProjectMemberInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete project member internal server error response has a 3xx status code
func (o *DeleteProjectMemberInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project member internal server error response has a 4xx status code
func (o *DeleteProjectMemberInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete project member internal server error response has a 5xx status code
func (o *DeleteProjectMemberInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete project member internal server error response a status code equal to that given
func (o *DeleteProjectMemberInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteProjectMemberInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/members/{mid}][%d] deleteProjectMemberInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteProjectMemberInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/members/{mid}][%d] deleteProjectMemberInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteProjectMemberInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteProjectMemberInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
