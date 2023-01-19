// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/middlewaregruppen/harbor-operator/models"
)

// UpdateUserProfileReader is a Reader for the UpdateUserProfile structure.
type UpdateUserProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUserProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateUserProfileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateUserProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateUserProfileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateUserProfileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateUserProfileOK creates a UpdateUserProfileOK with default headers values
func NewUpdateUserProfileOK() *UpdateUserProfileOK {
	return &UpdateUserProfileOK{}
}

/*
UpdateUserProfileOK describes a response with status code 200, with default header values.

Success
*/
type UpdateUserProfileOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this update user profile o k response has a 2xx status code
func (o *UpdateUserProfileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update user profile o k response has a 3xx status code
func (o *UpdateUserProfileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user profile o k response has a 4xx status code
func (o *UpdateUserProfileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update user profile o k response has a 5xx status code
func (o *UpdateUserProfileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update user profile o k response a status code equal to that given
func (o *UpdateUserProfileOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateUserProfileOK) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] updateUserProfileOK ", 200)
}

func (o *UpdateUserProfileOK) String() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] updateUserProfileOK ", 200)
}

func (o *UpdateUserProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewUpdateUserProfileUnauthorized creates a UpdateUserProfileUnauthorized with default headers values
func NewUpdateUserProfileUnauthorized() *UpdateUserProfileUnauthorized {
	return &UpdateUserProfileUnauthorized{}
}

/*
UpdateUserProfileUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateUserProfileUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update user profile unauthorized response has a 2xx status code
func (o *UpdateUserProfileUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user profile unauthorized response has a 3xx status code
func (o *UpdateUserProfileUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user profile unauthorized response has a 4xx status code
func (o *UpdateUserProfileUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update user profile unauthorized response has a 5xx status code
func (o *UpdateUserProfileUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update user profile unauthorized response a status code equal to that given
func (o *UpdateUserProfileUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateUserProfileUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] updateUserProfileUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateUserProfileUnauthorized) String() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] updateUserProfileUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateUserProfileUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateUserProfileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateUserProfileForbidden creates a UpdateUserProfileForbidden with default headers values
func NewUpdateUserProfileForbidden() *UpdateUserProfileForbidden {
	return &UpdateUserProfileForbidden{}
}

/*
UpdateUserProfileForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateUserProfileForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update user profile forbidden response has a 2xx status code
func (o *UpdateUserProfileForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user profile forbidden response has a 3xx status code
func (o *UpdateUserProfileForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user profile forbidden response has a 4xx status code
func (o *UpdateUserProfileForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update user profile forbidden response has a 5xx status code
func (o *UpdateUserProfileForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update user profile forbidden response a status code equal to that given
func (o *UpdateUserProfileForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateUserProfileForbidden) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] updateUserProfileForbidden  %+v", 403, o.Payload)
}

func (o *UpdateUserProfileForbidden) String() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] updateUserProfileForbidden  %+v", 403, o.Payload)
}

func (o *UpdateUserProfileForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateUserProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateUserProfileNotFound creates a UpdateUserProfileNotFound with default headers values
func NewUpdateUserProfileNotFound() *UpdateUserProfileNotFound {
	return &UpdateUserProfileNotFound{}
}

/*
UpdateUserProfileNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateUserProfileNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update user profile not found response has a 2xx status code
func (o *UpdateUserProfileNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user profile not found response has a 3xx status code
func (o *UpdateUserProfileNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user profile not found response has a 4xx status code
func (o *UpdateUserProfileNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update user profile not found response has a 5xx status code
func (o *UpdateUserProfileNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update user profile not found response a status code equal to that given
func (o *UpdateUserProfileNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateUserProfileNotFound) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] updateUserProfileNotFound  %+v", 404, o.Payload)
}

func (o *UpdateUserProfileNotFound) String() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] updateUserProfileNotFound  %+v", 404, o.Payload)
}

func (o *UpdateUserProfileNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateUserProfileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateUserProfileInternalServerError creates a UpdateUserProfileInternalServerError with default headers values
func NewUpdateUserProfileInternalServerError() *UpdateUserProfileInternalServerError {
	return &UpdateUserProfileInternalServerError{}
}

/*
UpdateUserProfileInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type UpdateUserProfileInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update user profile internal server error response has a 2xx status code
func (o *UpdateUserProfileInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user profile internal server error response has a 3xx status code
func (o *UpdateUserProfileInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user profile internal server error response has a 4xx status code
func (o *UpdateUserProfileInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update user profile internal server error response has a 5xx status code
func (o *UpdateUserProfileInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update user profile internal server error response a status code equal to that given
func (o *UpdateUserProfileInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UpdateUserProfileInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] updateUserProfileInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateUserProfileInternalServerError) String() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] updateUserProfileInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateUserProfileInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateUserProfileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
