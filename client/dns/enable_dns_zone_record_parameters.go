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

// NewEnableDNSZoneRecordParams creates a new EnableDNSZoneRecordParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEnableDNSZoneRecordParams() *EnableDNSZoneRecordParams {
	return &EnableDNSZoneRecordParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEnableDNSZoneRecordParamsWithTimeout creates a new EnableDNSZoneRecordParams object
// with the ability to set a timeout on a request.
func NewEnableDNSZoneRecordParamsWithTimeout(timeout time.Duration) *EnableDNSZoneRecordParams {
	return &EnableDNSZoneRecordParams{
		timeout: timeout,
	}
}

// NewEnableDNSZoneRecordParamsWithContext creates a new EnableDNSZoneRecordParams object
// with the ability to set a context for a request.
func NewEnableDNSZoneRecordParamsWithContext(ctx context.Context) *EnableDNSZoneRecordParams {
	return &EnableDNSZoneRecordParams{
		Context: ctx,
	}
}

// NewEnableDNSZoneRecordParamsWithHTTPClient creates a new EnableDNSZoneRecordParams object
// with the ability to set a custom HTTPClient for a request.
func NewEnableDNSZoneRecordParamsWithHTTPClient(client *http.Client) *EnableDNSZoneRecordParams {
	return &EnableDNSZoneRecordParams{
		HTTPClient: client,
	}
}

/*
EnableDNSZoneRecordParams contains all the parameters to send to the API endpoint

	for the enable DNS zone record operation.

	Typically these are written to a http.Request.
*/
type EnableDNSZoneRecordParams struct {

	/* RecordID.

	   The ID of the DNS record to enable.
	*/
	RecordID string

	/* ZoneID.

	   The ID of the zone's record to enable.
	*/
	ZoneID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the enable DNS zone record params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnableDNSZoneRecordParams) WithDefaults() *EnableDNSZoneRecordParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the enable DNS zone record params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnableDNSZoneRecordParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the enable DNS zone record params
func (o *EnableDNSZoneRecordParams) WithTimeout(timeout time.Duration) *EnableDNSZoneRecordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enable DNS zone record params
func (o *EnableDNSZoneRecordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enable DNS zone record params
func (o *EnableDNSZoneRecordParams) WithContext(ctx context.Context) *EnableDNSZoneRecordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enable DNS zone record params
func (o *EnableDNSZoneRecordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enable DNS zone record params
func (o *EnableDNSZoneRecordParams) WithHTTPClient(client *http.Client) *EnableDNSZoneRecordParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enable DNS zone record params
func (o *EnableDNSZoneRecordParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRecordID adds the recordID to the enable DNS zone record params
func (o *EnableDNSZoneRecordParams) WithRecordID(recordID string) *EnableDNSZoneRecordParams {
	o.SetRecordID(recordID)
	return o
}

// SetRecordID adds the recordId to the enable DNS zone record params
func (o *EnableDNSZoneRecordParams) SetRecordID(recordID string) {
	o.RecordID = recordID
}

// WithZoneID adds the zoneID to the enable DNS zone record params
func (o *EnableDNSZoneRecordParams) WithZoneID(zoneID string) *EnableDNSZoneRecordParams {
	o.SetZoneID(zoneID)
	return o
}

// SetZoneID adds the zoneId to the enable DNS zone record params
func (o *EnableDNSZoneRecordParams) SetZoneID(zoneID string) {
	o.ZoneID = zoneID
}

// WriteToRequest writes these params to a swagger request
func (o *EnableDNSZoneRecordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param recordId
	if err := r.SetPathParam("recordId", o.RecordID); err != nil {
		return err
	}

	// path param zoneId
	if err := r.SetPathParam("zoneId", o.ZoneID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
