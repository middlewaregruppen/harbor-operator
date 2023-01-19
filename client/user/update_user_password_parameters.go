// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/middlewaregruppen/harbor-operator/models"
)

// NewUpdateUserPasswordParams creates a new UpdateUserPasswordParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateUserPasswordParams() *UpdateUserPasswordParams {
	return &UpdateUserPasswordParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUserPasswordParamsWithTimeout creates a new UpdateUserPasswordParams object
// with the ability to set a timeout on a request.
func NewUpdateUserPasswordParamsWithTimeout(timeout time.Duration) *UpdateUserPasswordParams {
	return &UpdateUserPasswordParams{
		timeout: timeout,
	}
}

// NewUpdateUserPasswordParamsWithContext creates a new UpdateUserPasswordParams object
// with the ability to set a context for a request.
func NewUpdateUserPasswordParamsWithContext(ctx context.Context) *UpdateUserPasswordParams {
	return &UpdateUserPasswordParams{
		Context: ctx,
	}
}

// NewUpdateUserPasswordParamsWithHTTPClient creates a new UpdateUserPasswordParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateUserPasswordParamsWithHTTPClient(client *http.Client) *UpdateUserPasswordParams {
	return &UpdateUserPasswordParams{
		HTTPClient: client,
	}
}

/*
UpdateUserPasswordParams contains all the parameters to send to the API endpoint

	for the update user password operation.

	Typically these are written to a http.Request.
*/
type UpdateUserPasswordParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* Password.

	   Password to be updated, the attribute 'old_password' is optional when the API is called by the system administrator.
	*/
	Password *models.PasswordReq

	// UserID.
	//
	// Format: int
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update user password params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUserPasswordParams) WithDefaults() *UpdateUserPasswordParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update user password params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUserPasswordParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update user password params
func (o *UpdateUserPasswordParams) WithTimeout(timeout time.Duration) *UpdateUserPasswordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update user password params
func (o *UpdateUserPasswordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update user password params
func (o *UpdateUserPasswordParams) WithContext(ctx context.Context) *UpdateUserPasswordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update user password params
func (o *UpdateUserPasswordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update user password params
func (o *UpdateUserPasswordParams) WithHTTPClient(client *http.Client) *UpdateUserPasswordParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update user password params
func (o *UpdateUserPasswordParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the update user password params
func (o *UpdateUserPasswordParams) WithXRequestID(xRequestID *string) *UpdateUserPasswordParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the update user password params
func (o *UpdateUserPasswordParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithPassword adds the password to the update user password params
func (o *UpdateUserPasswordParams) WithPassword(password *models.PasswordReq) *UpdateUserPasswordParams {
	o.SetPassword(password)
	return o
}

// SetPassword adds the password to the update user password params
func (o *UpdateUserPasswordParams) SetPassword(password *models.PasswordReq) {
	o.Password = password
}

// WithUserID adds the userID to the update user password params
func (o *UpdateUserPasswordParams) WithUserID(userID int64) *UpdateUserPasswordParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the update user password params
func (o *UpdateUserPasswordParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUserPasswordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}
	if o.Password != nil {
		if err := r.SetBodyParam(o.Password); err != nil {
			return err
		}
	}

	// path param user_id
	if err := r.SetPathParam("user_id", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
