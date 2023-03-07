// Code generated by go-swagger; DO NOT EDIT.

package dns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetDNSZoneRecordHandlerFunc turns a function with the right signature into a get DNS zone record handler
type GetDNSZoneRecordHandlerFunc func(GetDNSZoneRecordParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDNSZoneRecordHandlerFunc) Handle(params GetDNSZoneRecordParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetDNSZoneRecordHandler interface for that can handle valid get DNS zone record params
type GetDNSZoneRecordHandler interface {
	Handle(GetDNSZoneRecordParams, interface{}) middleware.Responder
}

// NewGetDNSZoneRecord creates a new http.Handler for the get DNS zone record operation
func NewGetDNSZoneRecord(ctx *middleware.Context, handler GetDNSZoneRecordHandler) *GetDNSZoneRecord {
	return &GetDNSZoneRecord{Context: ctx, Handler: handler}
}

/*
	GetDNSZoneRecord swagger:route GET /dns/zone/{zoneId}/record/{recordId} dns getDnsZoneRecord

Returns the zone record object.
*/
type GetDNSZoneRecord struct {
	Context *middleware.Context
	Handler GetDNSZoneRecordHandler
}

func (o *GetDNSZoneRecord) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetDNSZoneRecordParams()
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