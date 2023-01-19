// Code generated by go-swagger; DO NOT EDIT.

package purge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/middlewaregruppen/harbor-operator/models"
)

// StopPurgeReader is a Reader for the StopPurge structure.
type StopPurgeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopPurgeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopPurgeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewStopPurgeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStopPurgeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStopPurgeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStopPurgeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStopPurgeOK creates a StopPurgeOK with default headers values
func NewStopPurgeOK() *StopPurgeOK {
	return &StopPurgeOK{}
}

/*
StopPurgeOK describes a response with status code 200, with default header values.

Success
*/
type StopPurgeOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this stop purge o k response has a 2xx status code
func (o *StopPurgeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stop purge o k response has a 3xx status code
func (o *StopPurgeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop purge o k response has a 4xx status code
func (o *StopPurgeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop purge o k response has a 5xx status code
func (o *StopPurgeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stop purge o k response a status code equal to that given
func (o *StopPurgeOK) IsCode(code int) bool {
	return code == 200
}

func (o *StopPurgeOK) Error() string {
	return fmt.Sprintf("[PUT /system/purgeaudit/{purge_id}][%d] stopPurgeOK ", 200)
}

func (o *StopPurgeOK) String() string {
	return fmt.Sprintf("[PUT /system/purgeaudit/{purge_id}][%d] stopPurgeOK ", 200)
}

func (o *StopPurgeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewStopPurgeUnauthorized creates a StopPurgeUnauthorized with default headers values
func NewStopPurgeUnauthorized() *StopPurgeUnauthorized {
	return &StopPurgeUnauthorized{}
}

/*
StopPurgeUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StopPurgeUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this stop purge unauthorized response has a 2xx status code
func (o *StopPurgeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop purge unauthorized response has a 3xx status code
func (o *StopPurgeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop purge unauthorized response has a 4xx status code
func (o *StopPurgeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop purge unauthorized response has a 5xx status code
func (o *StopPurgeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stop purge unauthorized response a status code equal to that given
func (o *StopPurgeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *StopPurgeUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /system/purgeaudit/{purge_id}][%d] stopPurgeUnauthorized  %+v", 401, o.Payload)
}

func (o *StopPurgeUnauthorized) String() string {
	return fmt.Sprintf("[PUT /system/purgeaudit/{purge_id}][%d] stopPurgeUnauthorized  %+v", 401, o.Payload)
}

func (o *StopPurgeUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *StopPurgeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStopPurgeForbidden creates a StopPurgeForbidden with default headers values
func NewStopPurgeForbidden() *StopPurgeForbidden {
	return &StopPurgeForbidden{}
}

/*
StopPurgeForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StopPurgeForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this stop purge forbidden response has a 2xx status code
func (o *StopPurgeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop purge forbidden response has a 3xx status code
func (o *StopPurgeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop purge forbidden response has a 4xx status code
func (o *StopPurgeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop purge forbidden response has a 5xx status code
func (o *StopPurgeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stop purge forbidden response a status code equal to that given
func (o *StopPurgeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *StopPurgeForbidden) Error() string {
	return fmt.Sprintf("[PUT /system/purgeaudit/{purge_id}][%d] stopPurgeForbidden  %+v", 403, o.Payload)
}

func (o *StopPurgeForbidden) String() string {
	return fmt.Sprintf("[PUT /system/purgeaudit/{purge_id}][%d] stopPurgeForbidden  %+v", 403, o.Payload)
}

func (o *StopPurgeForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *StopPurgeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStopPurgeNotFound creates a StopPurgeNotFound with default headers values
func NewStopPurgeNotFound() *StopPurgeNotFound {
	return &StopPurgeNotFound{}
}

/*
StopPurgeNotFound describes a response with status code 404, with default header values.

Not found
*/
type StopPurgeNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this stop purge not found response has a 2xx status code
func (o *StopPurgeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop purge not found response has a 3xx status code
func (o *StopPurgeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop purge not found response has a 4xx status code
func (o *StopPurgeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop purge not found response has a 5xx status code
func (o *StopPurgeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stop purge not found response a status code equal to that given
func (o *StopPurgeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *StopPurgeNotFound) Error() string {
	return fmt.Sprintf("[PUT /system/purgeaudit/{purge_id}][%d] stopPurgeNotFound  %+v", 404, o.Payload)
}

func (o *StopPurgeNotFound) String() string {
	return fmt.Sprintf("[PUT /system/purgeaudit/{purge_id}][%d] stopPurgeNotFound  %+v", 404, o.Payload)
}

func (o *StopPurgeNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *StopPurgeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStopPurgeInternalServerError creates a StopPurgeInternalServerError with default headers values
func NewStopPurgeInternalServerError() *StopPurgeInternalServerError {
	return &StopPurgeInternalServerError{}
}

/*
StopPurgeInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type StopPurgeInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this stop purge internal server error response has a 2xx status code
func (o *StopPurgeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop purge internal server error response has a 3xx status code
func (o *StopPurgeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop purge internal server error response has a 4xx status code
func (o *StopPurgeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop purge internal server error response has a 5xx status code
func (o *StopPurgeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stop purge internal server error response a status code equal to that given
func (o *StopPurgeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *StopPurgeInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /system/purgeaudit/{purge_id}][%d] stopPurgeInternalServerError  %+v", 500, o.Payload)
}

func (o *StopPurgeInternalServerError) String() string {
	return fmt.Sprintf("[PUT /system/purgeaudit/{purge_id}][%d] stopPurgeInternalServerError  %+v", 500, o.Payload)
}

func (o *StopPurgeInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *StopPurgeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
