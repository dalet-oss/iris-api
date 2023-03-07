// Code generated by go-swagger; DO NOT EDIT.

package dns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewDisableDNSZoneRecordParams creates a new DisableDNSZoneRecordParams object
//
// There are no default values defined in the spec.
func NewDisableDNSZoneRecordParams() DisableDNSZoneRecordParams {

	return DisableDNSZoneRecordParams{}
}

// DisableDNSZoneRecordParams contains all the bound params for the disable DNS zone record operation
// typically these are obtained from a http.Request
//
// swagger:parameters DisableDNSZoneRecord
type DisableDNSZoneRecordParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The ID of the DNS record to disable.
	  Required: true
	  In: path
	*/
	RecordID string
	/*The ID of the zone's record to disable.
	  Required: true
	  In: path
	*/
	ZoneID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDisableDNSZoneRecordParams() beforehand.
func (o *DisableDNSZoneRecordParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rRecordID, rhkRecordID, _ := route.Params.GetOK("recordId")
	if err := o.bindRecordID(rRecordID, rhkRecordID, route.Formats); err != nil {
		res = append(res, err)
	}

	rZoneID, rhkZoneID, _ := route.Params.GetOK("zoneId")
	if err := o.bindZoneID(rZoneID, rhkZoneID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindRecordID binds and validates parameter RecordID from path.
func (o *DisableDNSZoneRecordParams) bindRecordID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.RecordID = raw

	return nil
}

// bindZoneID binds and validates parameter ZoneID from path.
func (o *DisableDNSZoneRecordParams) bindZoneID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ZoneID = raw

	return nil
}
