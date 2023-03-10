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

	"github.com/dalet-oss/iris-api/models"
)

// NewUpdateDHCPSubnetParams creates a new UpdateDHCPSubnetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDHCPSubnetParams() *UpdateDHCPSubnetParams {
	return &UpdateDHCPSubnetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDHCPSubnetParamsWithTimeout creates a new UpdateDHCPSubnetParams object
// with the ability to set a timeout on a request.
func NewUpdateDHCPSubnetParamsWithTimeout(timeout time.Duration) *UpdateDHCPSubnetParams {
	return &UpdateDHCPSubnetParams{
		timeout: timeout,
	}
}

// NewUpdateDHCPSubnetParamsWithContext creates a new UpdateDHCPSubnetParams object
// with the ability to set a context for a request.
func NewUpdateDHCPSubnetParamsWithContext(ctx context.Context) *UpdateDHCPSubnetParams {
	return &UpdateDHCPSubnetParams{
		Context: ctx,
	}
}

// NewUpdateDHCPSubnetParamsWithHTTPClient creates a new UpdateDHCPSubnetParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDHCPSubnetParamsWithHTTPClient(client *http.Client) *UpdateDHCPSubnetParams {
	return &UpdateDHCPSubnetParams{
		HTTPClient: client,
	}
}

/*
UpdateDHCPSubnetParams contains all the parameters to send to the API endpoint

	for the update d h c p subnet operation.

	Typically these are written to a http.Request.
*/
type UpdateDHCPSubnetParams struct {

	// Body.
	Body *models.Subnet

	/* SubnetID.

	   The ID of the subnet to update.
	*/
	SubnetID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update d h c p subnet params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDHCPSubnetParams) WithDefaults() *UpdateDHCPSubnetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update d h c p subnet params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDHCPSubnetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update d h c p subnet params
func (o *UpdateDHCPSubnetParams) WithTimeout(timeout time.Duration) *UpdateDHCPSubnetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update d h c p subnet params
func (o *UpdateDHCPSubnetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update d h c p subnet params
func (o *UpdateDHCPSubnetParams) WithContext(ctx context.Context) *UpdateDHCPSubnetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update d h c p subnet params
func (o *UpdateDHCPSubnetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update d h c p subnet params
func (o *UpdateDHCPSubnetParams) WithHTTPClient(client *http.Client) *UpdateDHCPSubnetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update d h c p subnet params
func (o *UpdateDHCPSubnetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update d h c p subnet params
func (o *UpdateDHCPSubnetParams) WithBody(body *models.Subnet) *UpdateDHCPSubnetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update d h c p subnet params
func (o *UpdateDHCPSubnetParams) SetBody(body *models.Subnet) {
	o.Body = body
}

// WithSubnetID adds the subnetID to the update d h c p subnet params
func (o *UpdateDHCPSubnetParams) WithSubnetID(subnetID string) *UpdateDHCPSubnetParams {
	o.SetSubnetID(subnetID)
	return o
}

// SetSubnetID adds the subnetId to the update d h c p subnet params
func (o *UpdateDHCPSubnetParams) SetSubnetID(subnetID string) {
	o.SubnetID = subnetID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDHCPSubnetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param subnetId
	if err := r.SetPathParam("subnetId", o.SubnetID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
