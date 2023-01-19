// Code generated by go-swagger; DO NOT EDIT.

package retention

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/middlewaregruppen/harbor-operator/models"
)

// TriggerRetentionExecutionReader is a Reader for the TriggerRetentionExecution structure.
type TriggerRetentionExecutionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TriggerRetentionExecutionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTriggerRetentionExecutionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewTriggerRetentionExecutionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewTriggerRetentionExecutionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTriggerRetentionExecutionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTriggerRetentionExecutionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTriggerRetentionExecutionOK creates a TriggerRetentionExecutionOK with default headers values
func NewTriggerRetentionExecutionOK() *TriggerRetentionExecutionOK {
	return &TriggerRetentionExecutionOK{}
}

/*
TriggerRetentionExecutionOK describes a response with status code 200, with default header values.

Trigger a Retention job successfully.
*/
type TriggerRetentionExecutionOK struct {
}

// IsSuccess returns true when this trigger retention execution o k response has a 2xx status code
func (o *TriggerRetentionExecutionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this trigger retention execution o k response has a 3xx status code
func (o *TriggerRetentionExecutionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this trigger retention execution o k response has a 4xx status code
func (o *TriggerRetentionExecutionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this trigger retention execution o k response has a 5xx status code
func (o *TriggerRetentionExecutionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this trigger retention execution o k response a status code equal to that given
func (o *TriggerRetentionExecutionOK) IsCode(code int) bool {
	return code == 200
}

func (o *TriggerRetentionExecutionOK) Error() string {
	return fmt.Sprintf("[POST /retentions/{id}/executions][%d] triggerRetentionExecutionOK ", 200)
}

func (o *TriggerRetentionExecutionOK) String() string {
	return fmt.Sprintf("[POST /retentions/{id}/executions][%d] triggerRetentionExecutionOK ", 200)
}

func (o *TriggerRetentionExecutionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTriggerRetentionExecutionCreated creates a TriggerRetentionExecutionCreated with default headers values
func NewTriggerRetentionExecutionCreated() *TriggerRetentionExecutionCreated {
	return &TriggerRetentionExecutionCreated{}
}

/*
TriggerRetentionExecutionCreated describes a response with status code 201, with default header values.

Created
*/
type TriggerRetentionExecutionCreated struct {

	/* The location of the resource
	 */
	Location string

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this trigger retention execution created response has a 2xx status code
func (o *TriggerRetentionExecutionCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this trigger retention execution created response has a 3xx status code
func (o *TriggerRetentionExecutionCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this trigger retention execution created response has a 4xx status code
func (o *TriggerRetentionExecutionCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this trigger retention execution created response has a 5xx status code
func (o *TriggerRetentionExecutionCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this trigger retention execution created response a status code equal to that given
func (o *TriggerRetentionExecutionCreated) IsCode(code int) bool {
	return code == 201
}

func (o *TriggerRetentionExecutionCreated) Error() string {
	return fmt.Sprintf("[POST /retentions/{id}/executions][%d] triggerRetentionExecutionCreated ", 201)
}

func (o *TriggerRetentionExecutionCreated) String() string {
	return fmt.Sprintf("[POST /retentions/{id}/executions][%d] triggerRetentionExecutionCreated ", 201)
}

func (o *TriggerRetentionExecutionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewTriggerRetentionExecutionUnauthorized creates a TriggerRetentionExecutionUnauthorized with default headers values
func NewTriggerRetentionExecutionUnauthorized() *TriggerRetentionExecutionUnauthorized {
	return &TriggerRetentionExecutionUnauthorized{}
}

/*
TriggerRetentionExecutionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type TriggerRetentionExecutionUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this trigger retention execution unauthorized response has a 2xx status code
func (o *TriggerRetentionExecutionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this trigger retention execution unauthorized response has a 3xx status code
func (o *TriggerRetentionExecutionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this trigger retention execution unauthorized response has a 4xx status code
func (o *TriggerRetentionExecutionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this trigger retention execution unauthorized response has a 5xx status code
func (o *TriggerRetentionExecutionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this trigger retention execution unauthorized response a status code equal to that given
func (o *TriggerRetentionExecutionUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *TriggerRetentionExecutionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /retentions/{id}/executions][%d] triggerRetentionExecutionUnauthorized  %+v", 401, o.Payload)
}

func (o *TriggerRetentionExecutionUnauthorized) String() string {
	return fmt.Sprintf("[POST /retentions/{id}/executions][%d] triggerRetentionExecutionUnauthorized  %+v", 401, o.Payload)
}

func (o *TriggerRetentionExecutionUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *TriggerRetentionExecutionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewTriggerRetentionExecutionForbidden creates a TriggerRetentionExecutionForbidden with default headers values
func NewTriggerRetentionExecutionForbidden() *TriggerRetentionExecutionForbidden {
	return &TriggerRetentionExecutionForbidden{}
}

/*
TriggerRetentionExecutionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type TriggerRetentionExecutionForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this trigger retention execution forbidden response has a 2xx status code
func (o *TriggerRetentionExecutionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this trigger retention execution forbidden response has a 3xx status code
func (o *TriggerRetentionExecutionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this trigger retention execution forbidden response has a 4xx status code
func (o *TriggerRetentionExecutionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this trigger retention execution forbidden response has a 5xx status code
func (o *TriggerRetentionExecutionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this trigger retention execution forbidden response a status code equal to that given
func (o *TriggerRetentionExecutionForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *TriggerRetentionExecutionForbidden) Error() string {
	return fmt.Sprintf("[POST /retentions/{id}/executions][%d] triggerRetentionExecutionForbidden  %+v", 403, o.Payload)
}

func (o *TriggerRetentionExecutionForbidden) String() string {
	return fmt.Sprintf("[POST /retentions/{id}/executions][%d] triggerRetentionExecutionForbidden  %+v", 403, o.Payload)
}

func (o *TriggerRetentionExecutionForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *TriggerRetentionExecutionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewTriggerRetentionExecutionInternalServerError creates a TriggerRetentionExecutionInternalServerError with default headers values
func NewTriggerRetentionExecutionInternalServerError() *TriggerRetentionExecutionInternalServerError {
	return &TriggerRetentionExecutionInternalServerError{}
}

/*
TriggerRetentionExecutionInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type TriggerRetentionExecutionInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this trigger retention execution internal server error response has a 2xx status code
func (o *TriggerRetentionExecutionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this trigger retention execution internal server error response has a 3xx status code
func (o *TriggerRetentionExecutionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this trigger retention execution internal server error response has a 4xx status code
func (o *TriggerRetentionExecutionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this trigger retention execution internal server error response has a 5xx status code
func (o *TriggerRetentionExecutionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this trigger retention execution internal server error response a status code equal to that given
func (o *TriggerRetentionExecutionInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *TriggerRetentionExecutionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /retentions/{id}/executions][%d] triggerRetentionExecutionInternalServerError  %+v", 500, o.Payload)
}

func (o *TriggerRetentionExecutionInternalServerError) String() string {
	return fmt.Sprintf("[POST /retentions/{id}/executions][%d] triggerRetentionExecutionInternalServerError  %+v", 500, o.Payload)
}

func (o *TriggerRetentionExecutionInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *TriggerRetentionExecutionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*
TriggerRetentionExecutionBody trigger retention execution body
swagger:model TriggerRetentionExecutionBody
*/
type TriggerRetentionExecutionBody struct {

	// dry run
	DryRun bool `json:"dry_run,omitempty"`
}

// Validate validates this trigger retention execution body
func (o *TriggerRetentionExecutionBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this trigger retention execution body based on context it is used
func (o *TriggerRetentionExecutionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TriggerRetentionExecutionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TriggerRetentionExecutionBody) UnmarshalBinary(b []byte) error {
	var res TriggerRetentionExecutionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
