// Code generated by go-swagger; DO NOT EDIT.

package dhcp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAllDHCPSubnetsHandlerFunc turns a function with the right signature into a get all d h c p subnets handler
type GetAllDHCPSubnetsHandlerFunc func(GetAllDHCPSubnetsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAllDHCPSubnetsHandlerFunc) Handle(params GetAllDHCPSubnetsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetAllDHCPSubnetsHandler interface for that can handle valid get all d h c p subnets params
type GetAllDHCPSubnetsHandler interface {
	Handle(GetAllDHCPSubnetsParams, interface{}) middleware.Responder
}

// NewGetAllDHCPSubnets creates a new http.Handler for the get all d h c p subnets operation
func NewGetAllDHCPSubnets(ctx *middleware.Context, handler GetAllDHCPSubnetsHandler) *GetAllDHCPSubnets {
	return &GetAllDHCPSubnets{Context: ctx, Handler: handler}
}

/*
	GetAllDHCPSubnets swagger:route GET /dhcp/subnet dhcp getAllDHCPSubnets

Returns the IDs of DHCPv4 registered subnets.
*/
type GetAllDHCPSubnets struct {
	Context *middleware.Context
	Handler GetAllDHCPSubnetsHandler
}

func (o *GetAllDHCPSubnets) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAllDHCPSubnetsParams()
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
