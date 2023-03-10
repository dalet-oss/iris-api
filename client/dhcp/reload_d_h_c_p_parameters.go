// Code generated by go-swagger; DO NOT EDIT.

package dhcp

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
)

// NewReloadDHCPParams creates a new ReloadDHCPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReloadDHCPParams() *ReloadDHCPParams {
	return &ReloadDHCPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReloadDHCPParamsWithTimeout creates a new ReloadDHCPParams object
// with the ability to set a timeout on a request.
func NewReloadDHCPParamsWithTimeout(timeout time.Duration) *ReloadDHCPParams {
	return &ReloadDHCPParams{
		timeout: timeout,
	}
}

// NewReloadDHCPParamsWithContext creates a new ReloadDHCPParams object
// with the ability to set a context for a request.
func NewReloadDHCPParamsWithContext(ctx context.Context) *ReloadDHCPParams {
	return &ReloadDHCPParams{
		Context: ctx,
	}
}

// NewReloadDHCPParamsWithHTTPClient creates a new ReloadDHCPParams object
// with the ability to set a custom HTTPClient for a request.
func NewReloadDHCPParamsWithHTTPClient(client *http.Client) *ReloadDHCPParams {
	return &ReloadDHCPParams{
		HTTPClient: client,
	}
}

/*
ReloadDHCPParams contains all the parameters to send to the API endpoint

	for the reload d h c p operation.

	Typically these are written to a http.Request.
*/
type ReloadDHCPParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the reload d h c p params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReloadDHCPParams) WithDefaults() *ReloadDHCPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the reload d h c p params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReloadDHCPParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the reload d h c p params
func (o *ReloadDHCPParams) WithTimeout(timeout time.Duration) *ReloadDHCPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reload d h c p params
func (o *ReloadDHCPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reload d h c p params
func (o *ReloadDHCPParams) WithContext(ctx context.Context) *ReloadDHCPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reload d h c p params
func (o *ReloadDHCPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reload d h c p params
func (o *ReloadDHCPParams) WithHTTPClient(client *http.Client) *ReloadDHCPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reload d h c p params
func (o *ReloadDHCPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ReloadDHCPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
