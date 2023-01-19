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

// CreateTagReader is a Reader for the CreateTag structure.
type CreateTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateTagCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateTagBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateTagUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateTagForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateTagNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCreateTagMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateTagConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateTagInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateTagCreated creates a CreateTagCreated with default headers values
func NewCreateTagCreated() *CreateTagCreated {
	return &CreateTagCreated{}
}

/*
CreateTagCreated describes a response with status code 201, with default header values.

Created
*/
type CreateTagCreated struct {

	/* The location of the resource
	 */
	Location string

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this create tag created response has a 2xx status code
func (o *CreateTagCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create tag created response has a 3xx status code
func (o *CreateTagCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create tag created response has a 4xx status code
func (o *CreateTagCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create tag created response has a 5xx status code
func (o *CreateTagCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create tag created response a status code equal to that given
func (o *CreateTagCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateTagCreated) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags][%d] createTagCreated ", 201)
}

func (o *CreateTagCreated) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags][%d] createTagCreated ", 201)
}

func (o *CreateTagCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewCreateTagBadRequest creates a CreateTagBadRequest with default headers values
func NewCreateTagBadRequest() *CreateTagBadRequest {
	return &CreateTagBadRequest{}
}

/*
CreateTagBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateTagBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create tag bad request response has a 2xx status code
func (o *CreateTagBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create tag bad request response has a 3xx status code
func (o *CreateTagBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create tag bad request response has a 4xx status code
func (o *CreateTagBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create tag bad request response has a 5xx status code
func (o *CreateTagBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create tag bad request response a status code equal to that given
func (o *CreateTagBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateTagBadRequest) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags][%d] createTagBadRequest  %+v", 400, o.Payload)
}

func (o *CreateTagBadRequest) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags][%d] createTagBadRequest  %+v", 400, o.Payload)
}

func (o *CreateTagBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateTagBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateTagUnauthorized creates a CreateTagUnauthorized with default headers values
func NewCreateTagUnauthorized() *CreateTagUnauthorized {
	return &CreateTagUnauthorized{}
}

/*
CreateTagUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateTagUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create tag unauthorized response has a 2xx status code
func (o *CreateTagUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create tag unauthorized response has a 3xx status code
func (o *CreateTagUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create tag unauthorized response has a 4xx status code
func (o *CreateTagUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create tag unauthorized response has a 5xx status code
func (o *CreateTagUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create tag unauthorized response a status code equal to that given
func (o *CreateTagUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateTagUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags][%d] createTagUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateTagUnauthorized) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags][%d] createTagUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateTagUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateTagUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateTagForbidden creates a CreateTagForbidden with default headers values
func NewCreateTagForbidden() *CreateTagForbidden {
	return &CreateTagForbidden{}
}

/*
CreateTagForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateTagForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create tag forbidden response has a 2xx status code
func (o *CreateTagForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create tag forbidden response has a 3xx status code
func (o *CreateTagForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create tag forbidden response has a 4xx status code
func (o *CreateTagForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create tag forbidden response has a 5xx status code
func (o *CreateTagForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create tag forbidden response a status code equal to that given
func (o *CreateTagForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateTagForbidden) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags][%d] createTagForbidden  %+v", 403, o.Payload)
}

func (o *CreateTagForbidden) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags][%d] createTagForbidden  %+v", 403, o.Payload)
}

func (o *CreateTagForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateTagForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateTagNotFound creates a CreateTagNotFound with default headers values
func NewCreateTagNotFound() *CreateTagNotFound {
	return &CreateTagNotFound{}
}

/*
CreateTagNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateTagNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create tag not found response has a 2xx status code
func (o *CreateTagNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create tag not found response has a 3xx status code
func (o *CreateTagNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create tag not found response has a 4xx status code
func (o *CreateTagNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create tag not found response has a 5xx status code
func (o *CreateTagNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create tag not found response a status code equal to that given
func (o *CreateTagNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateTagNotFound) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags][%d] createTagNotFound  %+v", 404, o.Payload)
}

func (o *CreateTagNotFound) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags][%d] createTagNotFound  %+v", 404, o.Payload)
}

func (o *CreateTagNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateTagNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateTagMethodNotAllowed creates a CreateTagMethodNotAllowed with default headers values
func NewCreateTagMethodNotAllowed() *CreateTagMethodNotAllowed {
	return &CreateTagMethodNotAllowed{}
}

/*
CreateTagMethodNotAllowed describes a response with status code 405, with default header values.

Method not allowed
*/
type CreateTagMethodNotAllowed struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create tag method not allowed response has a 2xx status code
func (o *CreateTagMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create tag method not allowed response has a 3xx status code
func (o *CreateTagMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create tag method not allowed response has a 4xx status code
func (o *CreateTagMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this create tag method not allowed response has a 5xx status code
func (o *CreateTagMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this create tag method not allowed response a status code equal to that given
func (o *CreateTagMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *CreateTagMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags][%d] createTagMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *CreateTagMethodNotAllowed) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags][%d] createTagMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *CreateTagMethodNotAllowed) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateTagMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateTagConflict creates a CreateTagConflict with default headers values
func NewCreateTagConflict() *CreateTagConflict {
	return &CreateTagConflict{}
}

/*
CreateTagConflict describes a response with status code 409, with default header values.

Conflict
*/
type CreateTagConflict struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create tag conflict response has a 2xx status code
func (o *CreateTagConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create tag conflict response has a 3xx status code
func (o *CreateTagConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create tag conflict response has a 4xx status code
func (o *CreateTagConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create tag conflict response has a 5xx status code
func (o *CreateTagConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create tag conflict response a status code equal to that given
func (o *CreateTagConflict) IsCode(code int) bool {
	return code == 409
}

func (o *CreateTagConflict) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags][%d] createTagConflict  %+v", 409, o.Payload)
}

func (o *CreateTagConflict) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags][%d] createTagConflict  %+v", 409, o.Payload)
}

func (o *CreateTagConflict) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateTagConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateTagInternalServerError creates a CreateTagInternalServerError with default headers values
func NewCreateTagInternalServerError() *CreateTagInternalServerError {
	return &CreateTagInternalServerError{}
}

/*
CreateTagInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type CreateTagInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create tag internal server error response has a 2xx status code
func (o *CreateTagInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create tag internal server error response has a 3xx status code
func (o *CreateTagInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create tag internal server error response has a 4xx status code
func (o *CreateTagInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create tag internal server error response has a 5xx status code
func (o *CreateTagInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create tag internal server error response a status code equal to that given
func (o *CreateTagInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateTagInternalServerError) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags][%d] createTagInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateTagInternalServerError) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags][%d] createTagInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateTagInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateTagInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
