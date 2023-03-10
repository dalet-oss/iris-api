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

	"github.com/dalet-oss/iris-api/models"
)

// NewUpdateDNSZoneRecordParams creates a new UpdateDNSZoneRecordParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDNSZoneRecordParams() *UpdateDNSZoneRecordParams {
	return &UpdateDNSZoneRecordParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDNSZoneRecordParamsWithTimeout creates a new UpdateDNSZoneRecordParams object
// with the ability to set a timeout on a request.
func NewUpdateDNSZoneRecordParamsWithTimeout(timeout time.Duration) *UpdateDNSZoneRecordParams {
	return &UpdateDNSZoneRecordParams{
		timeout: timeout,
	}
}

// NewUpdateDNSZoneRecordParamsWithContext creates a new UpdateDNSZoneRecordParams object
// with the ability to set a context for a request.
func NewUpdateDNSZoneRecordParamsWithContext(ctx context.Context) *UpdateDNSZoneRecordParams {
	return &UpdateDNSZoneRecordParams{
		Context: ctx,
	}
}

// NewUpdateDNSZoneRecordParamsWithHTTPClient creates a new UpdateDNSZoneRecordParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDNSZoneRecordParamsWithHTTPClient(client *http.Client) *UpdateDNSZoneRecordParams {
	return &UpdateDNSZoneRecordParams{
		HTTPClient: client,
	}
}

/*
UpdateDNSZoneRecordParams contains all the parameters to send to the API endpoint

	for the update DNS zone record operation.

	Typically these are written to a http.Request.
*/
type UpdateDNSZoneRecordParams struct {

	// Body.
	Body *models.Record

	/* RecordID.

	   The ID of the DNS record to query.
	*/
	RecordID string

	/* ZoneID.

	   The ID of the zone's record to update.
	*/
	ZoneID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update DNS zone record params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDNSZoneRecordParams) WithDefaults() *UpdateDNSZoneRecordParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update DNS zone record params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDNSZoneRecordParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update DNS zone record params
func (o *UpdateDNSZoneRecordParams) WithTimeout(timeout time.Duration) *UpdateDNSZoneRecordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update DNS zone record params
func (o *UpdateDNSZoneRecordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update DNS zone record params
func (o *UpdateDNSZoneRecordParams) WithContext(ctx context.Context) *UpdateDNSZoneRecordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update DNS zone record params
func (o *UpdateDNSZoneRecordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update DNS zone record params
func (o *UpdateDNSZoneRecordParams) WithHTTPClient(client *http.Client) *UpdateDNSZoneRecordParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update DNS zone record params
func (o *UpdateDNSZoneRecordParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update DNS zone record params
func (o *UpdateDNSZoneRecordParams) WithBody(body *models.Record) *UpdateDNSZoneRecordParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update DNS zone record params
func (o *UpdateDNSZoneRecordParams) SetBody(body *models.Record) {
	o.Body = body
}

// WithRecordID adds the recordID to the update DNS zone record params
func (o *UpdateDNSZoneRecordParams) WithRecordID(recordID string) *UpdateDNSZoneRecordParams {
	o.SetRecordID(recordID)
	return o
}

// SetRecordID adds the recordId to the update DNS zone record params
func (o *UpdateDNSZoneRecordParams) SetRecordID(recordID string) {
	o.RecordID = recordID
}

// WithZoneID adds the zoneID to the update DNS zone record params
func (o *UpdateDNSZoneRecordParams) WithZoneID(zoneID string) *UpdateDNSZoneRecordParams {
	o.SetZoneID(zoneID)
	return o
}

// SetZoneID adds the zoneId to the update DNS zone record params
func (o *UpdateDNSZoneRecordParams) SetZoneID(zoneID string) {
	o.ZoneID = zoneID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDNSZoneRecordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
