// Code generated by go-swagger; DO NOT EDIT.

package dhcp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAllDHCPSubnetReservationsHandlerFunc turns a function with the right signature into a get all d h c p subnet reservations handler
type GetAllDHCPSubnetReservationsHandlerFunc func(GetAllDHCPSubnetReservationsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAllDHCPSubnetReservationsHandlerFunc) Handle(params GetAllDHCPSubnetReservationsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetAllDHCPSubnetReservationsHandler interface for that can handle valid get all d h c p subnet reservations params
type GetAllDHCPSubnetReservationsHandler interface {
	Handle(GetAllDHCPSubnetReservationsParams, interface{}) middleware.Responder
}

// NewGetAllDHCPSubnetReservations creates a new http.Handler for the get all d h c p subnet reservations operation
func NewGetAllDHCPSubnetReservations(ctx *middleware.Context, handler GetAllDHCPSubnetReservationsHandler) *GetAllDHCPSubnetReservations {
	return &GetAllDHCPSubnetReservations{Context: ctx, Handler: handler}
}

/*
	GetAllDHCPSubnetReservations swagger:route GET /dhcp/subnet/{subnetId}/reservation dhcp getAllDHCPSubnetReservations

Returns the list of MAC hardware addresses of subnet's reservations.
*/
type GetAllDHCPSubnetReservations struct {
	Context *middleware.Context
	Handler GetAllDHCPSubnetReservationsHandler
}

func (o *GetAllDHCPSubnetReservations) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAllDHCPSubnetReservationsParams()
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
