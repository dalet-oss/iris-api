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

// NewGetDNSZoneParams creates a new GetDNSZoneParams object
//
// There are no default values defined in the spec.
func NewGetDNSZoneParams() GetDNSZoneParams {

	return GetDNSZoneParams{}
}

// GetDNSZoneParams contains all the bound params for the get DNS zone operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetDNSZone
type GetDNSZoneParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The ID of the DNS zone to query.
	  Required: true
	  In: path
	*/
	ZoneID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetDNSZoneParams() beforehand.
func (o *GetDNSZoneParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rZoneID, rhkZoneID, _ := route.Params.GetOK("zoneId")
	if err := o.bindZoneID(rZoneID, rhkZoneID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindZoneID binds and validates parameter ZoneID from path.
func (o *GetDNSZoneParams) bindZoneID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ZoneID = raw

	return nil
}
