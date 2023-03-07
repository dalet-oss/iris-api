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

// NewGetDNSZoneRecordParams creates a new GetDNSZoneRecordParams object
//
// There are no default values defined in the spec.
func NewGetDNSZoneRecordParams() GetDNSZoneRecordParams {

	return GetDNSZoneRecordParams{}
}

// GetDNSZoneRecordParams contains all the bound params for the get DNS zone record operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetDNSZoneRecord
type GetDNSZoneRecordParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The ID of the DNS record to query.
	  Required: true
	  In: path
	*/
	RecordID string
	/*The ID of the DNS zone to query.
	  Required: true
	  In: path
	*/
	ZoneID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetDNSZoneRecordParams() beforehand.
func (o *GetDNSZoneRecordParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
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
func (o *GetDNSZoneRecordParams) bindRecordID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetDNSZoneRecordParams) bindZoneID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ZoneID = raw

	return nil
}
