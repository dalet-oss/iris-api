// Code generated by go-swagger; DO NOT EDIT.

package dns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetDNSServerHandlerFunc turns a function with the right signature into a get DNS server handler
type GetDNSServerHandlerFunc func(GetDNSServerParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDNSServerHandlerFunc) Handle(params GetDNSServerParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetDNSServerHandler interface for that can handle valid get DNS server params
type GetDNSServerHandler interface {
	Handle(GetDNSServerParams, interface{}) middleware.Responder
}

// NewGetDNSServer creates a new http.Handler for the get DNS server operation
func NewGetDNSServer(ctx *middleware.Context, handler GetDNSServerHandler) *GetDNSServer {
	return &GetDNSServer{Context: ctx, Handler: handler}
}

/*
	GetDNSServer swagger:route GET /dns/server/{serverId} dns getDnsServer

Returns the requested DNSServer object.
*/
type GetDNSServer struct {
	Context *middleware.Context
	Handler GetDNSServerHandler
}

func (o *GetDNSServer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetDNSServerParams()
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
