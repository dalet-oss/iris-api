// Code generated by go-swagger; DO NOT EDIT.

package dns

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

// NewDeleteDNSZoneParams creates a new DeleteDNSZoneParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDNSZoneParams() *DeleteDNSZoneParams {
	return &DeleteDNSZoneParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDNSZoneParamsWithTimeout creates a new DeleteDNSZoneParams object
// with the ability to set a timeout on a request.
func NewDeleteDNSZoneParamsWithTimeout(timeout time.Duration) *DeleteDNSZoneParams {
	return &DeleteDNSZoneParams{
		timeout: timeout,
	}
}

// NewDeleteDNSZoneParamsWithContext creates a new DeleteDNSZoneParams object
// with the ability to set a context for a request.
func NewDeleteDNSZoneParamsWithContext(ctx context.Context) *DeleteDNSZoneParams {
	return &DeleteDNSZoneParams{
		Context: ctx,
	}
}

// NewDeleteDNSZoneParamsWithHTTPClient creates a new DeleteDNSZoneParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDNSZoneParamsWithHTTPClient(client *http.Client) *DeleteDNSZoneParams {
	return &DeleteDNSZoneParams{
		HTTPClient: client,
	}
}

/*
DeleteDNSZoneParams contains all the parameters to send to the API endpoint

	for the delete DNS zone operation.

	Typically these are written to a http.Request.
*/
type DeleteDNSZoneParams struct {

	/* ZoneID.

	   The ID of the zone to delete.
	*/
	ZoneID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete DNS zone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDNSZoneParams) WithDefaults() *DeleteDNSZoneParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete DNS zone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDNSZoneParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete DNS zone params
func (o *DeleteDNSZoneParams) WithTimeout(timeout time.Duration) *DeleteDNSZoneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete DNS zone params
func (o *DeleteDNSZoneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete DNS zone params
func (o *DeleteDNSZoneParams) WithContext(ctx context.Context) *DeleteDNSZoneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete DNS zone params
func (o *DeleteDNSZoneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete DNS zone params
func (o *DeleteDNSZoneParams) WithHTTPClient(client *http.Client) *DeleteDNSZoneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete DNS zone params
func (o *DeleteDNSZoneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithZoneID adds the zoneID to the delete DNS zone params
func (o *DeleteDNSZoneParams) WithZoneID(zoneID string) *DeleteDNSZoneParams {
	o.SetZoneID(zoneID)
	return o
}

// SetZoneID adds the zoneId to the delete DNS zone params
func (o *DeleteDNSZoneParams) SetZoneID(zoneID string) {
	o.ZoneID = zoneID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDNSZoneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param zoneId
	if err := r.SetPathParam("zoneId", o.ZoneID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
