// Code generated by go-swagger; DO NOT EDIT.

package dns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAllDNSZonesHandlerFunc turns a function with the right signature into a get all DNS zones handler
type GetAllDNSZonesHandlerFunc func(GetAllDNSZonesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAllDNSZonesHandlerFunc) Handle(params GetAllDNSZonesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetAllDNSZonesHandler interface for that can handle valid get all DNS zones params
type GetAllDNSZonesHandler interface {
	Handle(GetAllDNSZonesParams, interface{}) middleware.Responder
}

// NewGetAllDNSZones creates a new http.Handler for the get all DNS zones operation
func NewGetAllDNSZones(ctx *middleware.Context, handler GetAllDNSZonesHandler) *GetAllDNSZones {
	return &GetAllDNSZones{Context: ctx, Handler: handler}
}

/*
	GetAllDNSZones swagger:route GET /dns/zone dns getAllDnsZones

Returns the IDs of DNS zones.
*/
type GetAllDNSZones struct {
	Context *middleware.Context
	Handler GetAllDNSZonesHandler
}

func (o *GetAllDNSZones) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAllDNSZonesParams()
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
