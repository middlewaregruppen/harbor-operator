// Code generated by go-swagger; DO NOT EDIT.

package replication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/middlewaregruppen/harbor-operator/models"
)

// UpdateReplicationPolicyReader is a Reader for the UpdateReplicationPolicy structure.
type UpdateReplicationPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateReplicationPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateReplicationPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateReplicationPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateReplicationPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateReplicationPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateReplicationPolicyConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateReplicationPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateReplicationPolicyOK creates a UpdateReplicationPolicyOK with default headers values
func NewUpdateReplicationPolicyOK() *UpdateReplicationPolicyOK {
	return &UpdateReplicationPolicyOK{}
}

/*
UpdateReplicationPolicyOK describes a response with status code 200, with default header values.

Success
*/
type UpdateReplicationPolicyOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this update replication policy o k response has a 2xx status code
func (o *UpdateReplicationPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update replication policy o k response has a 3xx status code
func (o *UpdateReplicationPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update replication policy o k response has a 4xx status code
func (o *UpdateReplicationPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update replication policy o k response has a 5xx status code
func (o *UpdateReplicationPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update replication policy o k response a status code equal to that given
func (o *UpdateReplicationPolicyOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateReplicationPolicyOK) Error() string {
	return fmt.Sprintf("[PUT /replication/policies/{id}][%d] updateReplicationPolicyOK ", 200)
}

func (o *UpdateReplicationPolicyOK) String() string {
	return fmt.Sprintf("[PUT /replication/policies/{id}][%d] updateReplicationPolicyOK ", 200)
}

func (o *UpdateReplicationPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewUpdateReplicationPolicyUnauthorized creates a UpdateReplicationPolicyUnauthorized with default headers values
func NewUpdateReplicationPolicyUnauthorized() *UpdateReplicationPolicyUnauthorized {
	return &UpdateReplicationPolicyUnauthorized{}
}

/*
UpdateReplicationPolicyUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateReplicationPolicyUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update replication policy unauthorized response has a 2xx status code
func (o *UpdateReplicationPolicyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update replication policy unauthorized response has a 3xx status code
func (o *UpdateReplicationPolicyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update replication policy unauthorized response has a 4xx status code
func (o *UpdateReplicationPolicyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update replication policy unauthorized response has a 5xx status code
func (o *UpdateReplicationPolicyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update replication policy unauthorized response a status code equal to that given
func (o *UpdateReplicationPolicyUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateReplicationPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /replication/policies/{id}][%d] updateReplicationPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateReplicationPolicyUnauthorized) String() string {
	return fmt.Sprintf("[PUT /replication/policies/{id}][%d] updateReplicationPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateReplicationPolicyUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateReplicationPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateReplicationPolicyForbidden creates a UpdateReplicationPolicyForbidden with default headers values
func NewUpdateReplicationPolicyForbidden() *UpdateReplicationPolicyForbidden {
	return &UpdateReplicationPolicyForbidden{}
}

/*
UpdateReplicationPolicyForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateReplicationPolicyForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update replication policy forbidden response has a 2xx status code
func (o *UpdateReplicationPolicyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update replication policy forbidden response has a 3xx status code
func (o *UpdateReplicationPolicyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update replication policy forbidden response has a 4xx status code
func (o *UpdateReplicationPolicyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update replication policy forbidden response has a 5xx status code
func (o *UpdateReplicationPolicyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update replication policy forbidden response a status code equal to that given
func (o *UpdateReplicationPolicyForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateReplicationPolicyForbidden) Error() string {
	return fmt.Sprintf("[PUT /replication/policies/{id}][%d] updateReplicationPolicyForbidden  %+v", 403, o.Payload)
}

func (o *UpdateReplicationPolicyForbidden) String() string {
	return fmt.Sprintf("[PUT /replication/policies/{id}][%d] updateReplicationPolicyForbidden  %+v", 403, o.Payload)
}

func (o *UpdateReplicationPolicyForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateReplicationPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateReplicationPolicyNotFound creates a UpdateReplicationPolicyNotFound with default headers values
func NewUpdateReplicationPolicyNotFound() *UpdateReplicationPolicyNotFound {
	return &UpdateReplicationPolicyNotFound{}
}

/*
UpdateReplicationPolicyNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateReplicationPolicyNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update replication policy not found response has a 2xx status code
func (o *UpdateReplicationPolicyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update replication policy not found response has a 3xx status code
func (o *UpdateReplicationPolicyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update replication policy not found response has a 4xx status code
func (o *UpdateReplicationPolicyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update replication policy not found response has a 5xx status code
func (o *UpdateReplicationPolicyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update replication policy not found response a status code equal to that given
func (o *UpdateReplicationPolicyNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateReplicationPolicyNotFound) Error() string {
	return fmt.Sprintf("[PUT /replication/policies/{id}][%d] updateReplicationPolicyNotFound  %+v", 404, o.Payload)
}

func (o *UpdateReplicationPolicyNotFound) String() string {
	return fmt.Sprintf("[PUT /replication/policies/{id}][%d] updateReplicationPolicyNotFound  %+v", 404, o.Payload)
}

func (o *UpdateReplicationPolicyNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateReplicationPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateReplicationPolicyConflict creates a UpdateReplicationPolicyConflict with default headers values
func NewUpdateReplicationPolicyConflict() *UpdateReplicationPolicyConflict {
	return &UpdateReplicationPolicyConflict{}
}

/*
UpdateReplicationPolicyConflict describes a response with status code 409, with default header values.

Conflict
*/
type UpdateReplicationPolicyConflict struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update replication policy conflict response has a 2xx status code
func (o *UpdateReplicationPolicyConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update replication policy conflict response has a 3xx status code
func (o *UpdateReplicationPolicyConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update replication policy conflict response has a 4xx status code
func (o *UpdateReplicationPolicyConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update replication policy conflict response has a 5xx status code
func (o *UpdateReplicationPolicyConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update replication policy conflict response a status code equal to that given
func (o *UpdateReplicationPolicyConflict) IsCode(code int) bool {
	return code == 409
}

func (o *UpdateReplicationPolicyConflict) Error() string {
	return fmt.Sprintf("[PUT /replication/policies/{id}][%d] updateReplicationPolicyConflict  %+v", 409, o.Payload)
}

func (o *UpdateReplicationPolicyConflict) String() string {
	return fmt.Sprintf("[PUT /replication/policies/{id}][%d] updateReplicationPolicyConflict  %+v", 409, o.Payload)
}

func (o *UpdateReplicationPolicyConflict) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateReplicationPolicyConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateReplicationPolicyInternalServerError creates a UpdateReplicationPolicyInternalServerError with default headers values
func NewUpdateReplicationPolicyInternalServerError() *UpdateReplicationPolicyInternalServerError {
	return &UpdateReplicationPolicyInternalServerError{}
}

/*
UpdateReplicationPolicyInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type UpdateReplicationPolicyInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update replication policy internal server error response has a 2xx status code
func (o *UpdateReplicationPolicyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update replication policy internal server error response has a 3xx status code
func (o *UpdateReplicationPolicyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update replication policy internal server error response has a 4xx status code
func (o *UpdateReplicationPolicyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update replication policy internal server error response has a 5xx status code
func (o *UpdateReplicationPolicyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update replication policy internal server error response a status code equal to that given
func (o *UpdateReplicationPolicyInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UpdateReplicationPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /replication/policies/{id}][%d] updateReplicationPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateReplicationPolicyInternalServerError) String() string {
	return fmt.Sprintf("[PUT /replication/policies/{id}][%d] updateReplicationPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateReplicationPolicyInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateReplicationPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
