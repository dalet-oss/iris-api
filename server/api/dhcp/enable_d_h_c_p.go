// Code generated by go-swagger; DO NOT EDIT.

package dhcp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// EnableDHCPHandlerFunc turns a function with the right signature into a enable d h c p handler
type EnableDHCPHandlerFunc func(EnableDHCPParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn EnableDHCPHandlerFunc) Handle(params EnableDHCPParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// EnableDHCPHandler interface for that can handle valid enable d h c p params
type EnableDHCPHandler interface {
	Handle(EnableDHCPParams, interface{}) middleware.Responder
}

// NewEnableDHCP creates a new http.Handler for the enable d h c p operation
func NewEnableDHCP(ctx *middleware.Context, handler EnableDHCPHandler) *EnableDHCP {
	return &EnableDHCP{Context: ctx, Handler: handler}
}

/*
	EnableDHCP swagger:route POST /dhcp/enable dhcp enableDHCP

Enable the DHCPv4 service.
*/
type EnableDHCP struct {
	Context *middleware.Context
	Handler EnableDHCPHandler
}

func (o *EnableDHCP) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewEnableDHCPParams()
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