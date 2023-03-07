// Code generated by go-swagger; DO NOT EDIT.

package dns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// EnableDNSZoneRecordHandlerFunc turns a function with the right signature into a enable DNS zone record handler
type EnableDNSZoneRecordHandlerFunc func(EnableDNSZoneRecordParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn EnableDNSZoneRecordHandlerFunc) Handle(params EnableDNSZoneRecordParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// EnableDNSZoneRecordHandler interface for that can handle valid enable DNS zone record params
type EnableDNSZoneRecordHandler interface {
	Handle(EnableDNSZoneRecordParams, interface{}) middleware.Responder
}

// NewEnableDNSZoneRecord creates a new http.Handler for the enable DNS zone record operation
func NewEnableDNSZoneRecord(ctx *middleware.Context, handler EnableDNSZoneRecordHandler) *EnableDNSZoneRecord {
	return &EnableDNSZoneRecord{Context: ctx, Handler: handler}
}

/*
	EnableDNSZoneRecord swagger:route POST /dns/zone/{zoneId}/record/{recordId}/enable dns enableDnsZoneRecord

Enable a given DNS Zone record.
*/
type EnableDNSZoneRecord struct {
	Context *middleware.Context
	Handler EnableDNSZoneRecordHandler
}

func (o *EnableDNSZoneRecord) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewEnableDNSZoneRecordParams()
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