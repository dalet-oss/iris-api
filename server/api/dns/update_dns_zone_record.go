// Code generated by go-swagger; DO NOT EDIT.

package dns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateDNSZoneRecordHandlerFunc turns a function with the right signature into a update DNS zone record handler
type UpdateDNSZoneRecordHandlerFunc func(UpdateDNSZoneRecordParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateDNSZoneRecordHandlerFunc) Handle(params UpdateDNSZoneRecordParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateDNSZoneRecordHandler interface for that can handle valid update DNS zone record params
type UpdateDNSZoneRecordHandler interface {
	Handle(UpdateDNSZoneRecordParams, interface{}) middleware.Responder
}

// NewUpdateDNSZoneRecord creates a new http.Handler for the update DNS zone record operation
func NewUpdateDNSZoneRecord(ctx *middleware.Context, handler UpdateDNSZoneRecordHandler) *UpdateDNSZoneRecord {
	return &UpdateDNSZoneRecord{Context: ctx, Handler: handler}
}

/*
	UpdateDNSZoneRecord swagger:route PUT /dns/zone/{zoneId}/record/{recordId} dns updateDnsZoneRecord

Updates a DNS zone record.
*/
type UpdateDNSZoneRecord struct {
	Context *middleware.Context
	Handler UpdateDNSZoneRecordHandler
}

func (o *UpdateDNSZoneRecord) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateDNSZoneRecordParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
