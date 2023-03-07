// Code generated by go-swagger; DO NOT EDIT.

package dns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"github.com/dalet-oss/iris-api/models"
)

// NewUpdateDNSZoneRecordParams creates a new UpdateDNSZoneRecordParams object
//
// There are no default values defined in the spec.
func NewUpdateDNSZoneRecordParams() UpdateDNSZoneRecordParams {

	return UpdateDNSZoneRecordParams{}
}

// UpdateDNSZoneRecordParams contains all the bound params for the update DNS zone record operation
// typically these are obtained from a http.Request
//
// swagger:parameters UpdateDNSZoneRecord
type UpdateDNSZoneRecordParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	Body *models.Record
	/*The ID of the DNS record to query.
	  Required: true
	  In: path
	*/
	RecordID string
	/*The ID of the zone's record to update.
	  Required: true
	  In: path
	*/
	ZoneID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateDNSZoneRecordParams() beforehand.
func (o *UpdateDNSZoneRecordParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Record
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	}

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
func (o *UpdateDNSZoneRecordParams) bindRecordID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *UpdateDNSZoneRecordParams) bindZoneID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ZoneID = raw

	return nil
}