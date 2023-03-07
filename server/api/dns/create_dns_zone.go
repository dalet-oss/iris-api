// Code generated by go-swagger; DO NOT EDIT.

package dns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateDNSZoneHandlerFunc turns a function with the right signature into a create DNS zone handler
type CreateDNSZoneHandlerFunc func(CreateDNSZoneParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateDNSZoneHandlerFunc) Handle(params CreateDNSZoneParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateDNSZoneHandler interface for that can handle valid create DNS zone params
type CreateDNSZoneHandler interface {
	Handle(CreateDNSZoneParams, interface{}) middleware.Responder
}

// NewCreateDNSZone creates a new http.Handler for the create DNS zone operation
func NewCreateDNSZone(ctx *middleware.Context, handler CreateDNSZoneHandler) *CreateDNSZone {
	return &CreateDNSZone{Context: ctx, Handler: handler}
}

/*
	CreateDNSZone swagger:route POST /dns/zone dns createDnsZone

Creates a new DNS (sub-)zone.
*/
type CreateDNSZone struct {
	Context *middleware.Context
	Handler CreateDNSZoneHandler
}

func (o *CreateDNSZone) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateDNSZoneParams()
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
