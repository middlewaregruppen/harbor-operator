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

// NewSetUserSysAdminParams creates a new SetUserSysAdminParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetUserSysAdminParams() *SetUserSysAdminParams {
	return &SetUserSysAdminParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetUserSysAdminParamsWithTimeout creates a new SetUserSysAdminParams object
// with the ability to set a timeout on a request.
func NewSetUserSysAdminParamsWithTimeout(timeout time.Duration) *SetUserSysAdminParams {
	return &SetUserSysAdminParams{
		timeout: timeout,
	}
}

// NewSetUserSysAdminParamsWithContext creates a new SetUserSysAdminParams object
// with the ability to set a context for a request.
func NewSetUserSysAdminParamsWithContext(ctx context.Context) *SetUserSysAdminParams {
	return &SetUserSysAdminParams{
		Context: ctx,
	}
}

// NewSetUserSysAdminParamsWithHTTPClient creates a new SetUserSysAdminParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetUserSysAdminParamsWithHTTPClient(client *http.Client) *SetUserSysAdminParams {
	return &SetUserSysAdminParams{
		HTTPClient: client,
	}
}

/*
SetUserSysAdminParams contains all the parameters to send to the API endpoint

	for the set user sys admin operation.

	Typically these are written to a http.Request.
*/
type SetUserSysAdminParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* SysadminFlag.

	   Toggle a user to admin or not.
	*/
	SysadminFlag *models.UserSysAdminFlag

	// UserID.
	//
	// Format: int
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set user sys admin params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetUserSysAdminParams) WithDefaults() *SetUserSysAdminParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set user sys admin params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetUserSysAdminParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set user sys admin params
func (o *SetUserSysAdminParams) WithTimeout(timeout time.Duration) *SetUserSysAdminParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set user sys admin params
func (o *SetUserSysAdminParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set user sys admin params
func (o *SetUserSysAdminParams) WithContext(ctx context.Context) *SetUserSysAdminParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set user sys admin params
func (o *SetUserSysAdminParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set user sys admin params
func (o *SetUserSysAdminParams) WithHTTPClient(client *http.Client) *SetUserSysAdminParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set user sys admin params
func (o *SetUserSysAdminParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the set user sys admin params
func (o *SetUserSysAdminParams) WithXRequestID(xRequestID *string) *SetUserSysAdminParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the set user sys admin params
func (o *SetUserSysAdminParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithSysadminFlag adds the sysadminFlag to the set user sys admin params
func (o *SetUserSysAdminParams) WithSysadminFlag(sysadminFlag *models.UserSysAdminFlag) *SetUserSysAdminParams {
	o.SetSysadminFlag(sysadminFlag)
	return o
}

// SetSysadminFlag adds the sysadminFlag to the set user sys admin params
func (o *SetUserSysAdminParams) SetSysadminFlag(sysadminFlag *models.UserSysAdminFlag) {
	o.SysadminFlag = sysadminFlag
}

// WithUserID adds the userID to the set user sys admin params
func (o *SetUserSysAdminParams) WithUserID(userID int64) *SetUserSysAdminParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the set user sys admin params
func (o *SetUserSysAdminParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *SetUserSysAdminParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.SysadminFlag != nil {
		if err := r.SetBodyParam(o.SysadminFlag); err != nil {
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
