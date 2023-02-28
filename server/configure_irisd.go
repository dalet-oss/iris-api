// This file is safe to edit. Once it exists it will not be overwritten

package server

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	swaggerapi "github.com/dalet-oss/iris-api/server/api"
	"github.com/dalet-oss/iris-api/server/api/dhcp"
)

//go:generate swagger generate server --target ../../iris-api --name Irisd --spec ../swagger.generated.yml --api-package api --server-package server --principal interface{} --exclude-main

func configureFlags(api *swaggerapi.IrisdAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *swaggerapi.IrisdAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()
	api.TxtProducer = runtime.TextProducer()

	// Applies when the "x-token" header is set
	if api.KeyAuth == nil {
		api.KeyAuth = func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (key) x-token from header param [x-token] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.DhcpCreateDHCPSubnetHandler == nil {
		api.DhcpCreateDHCPSubnetHandler = dhcp.CreateDHCPSubnetHandlerFunc(func(params dhcp.CreateDHCPSubnetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation dhcp.CreateDHCPSubnet has not yet been implemented")
		})
	}
	if api.DhcpCreateDHCPSubnetReservationHandler == nil {
		api.DhcpCreateDHCPSubnetReservationHandler = dhcp.CreateDHCPSubnetReservationHandlerFunc(func(params dhcp.CreateDHCPSubnetReservationParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation dhcp.CreateDHCPSubnetReservation has not yet been implemented")
		})
	}
	if api.DhcpDeleteDHCPSubnetHandler == nil {
		api.DhcpDeleteDHCPSubnetHandler = dhcp.DeleteDHCPSubnetHandlerFunc(func(params dhcp.DeleteDHCPSubnetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation dhcp.DeleteDHCPSubnet has not yet been implemented")
		})
	}
	if api.DhcpDeleteDHCPSubnetReservationHandler == nil {
		api.DhcpDeleteDHCPSubnetReservationHandler = dhcp.DeleteDHCPSubnetReservationHandlerFunc(func(params dhcp.DeleteDHCPSubnetReservationParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation dhcp.DeleteDHCPSubnetReservation has not yet been implemented")
		})
	}
	if api.DhcpDisableDHCPHandler == nil {
		api.DhcpDisableDHCPHandler = dhcp.DisableDHCPHandlerFunc(func(params dhcp.DisableDHCPParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation dhcp.DisableDHCP has not yet been implemented")
		})
	}
	if api.DhcpEnableDHCPHandler == nil {
		api.DhcpEnableDHCPHandler = dhcp.EnableDHCPHandlerFunc(func(params dhcp.EnableDHCPParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation dhcp.EnableDHCP has not yet been implemented")
		})
	}
	if api.DhcpGetAllDHCPSubnetReservationsHandler == nil {
		api.DhcpGetAllDHCPSubnetReservationsHandler = dhcp.GetAllDHCPSubnetReservationsHandlerFunc(func(params dhcp.GetAllDHCPSubnetReservationsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation dhcp.GetAllDHCPSubnetReservations has not yet been implemented")
		})
	}
	if api.DhcpGetAllDHCPSubnetsHandler == nil {
		api.DhcpGetAllDHCPSubnetsHandler = dhcp.GetAllDHCPSubnetsHandlerFunc(func(params dhcp.GetAllDHCPSubnetsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation dhcp.GetAllDHCPSubnets has not yet been implemented")
		})
	}
	if api.DhcpGetDHCPStatusHandler == nil {
		api.DhcpGetDHCPStatusHandler = dhcp.GetDHCPStatusHandlerFunc(func(params dhcp.GetDHCPStatusParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation dhcp.GetDHCPStatus has not yet been implemented")
		})
	}
	if api.DhcpGetDHCPSubnetHandler == nil {
		api.DhcpGetDHCPSubnetHandler = dhcp.GetDHCPSubnetHandlerFunc(func(params dhcp.GetDHCPSubnetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation dhcp.GetDHCPSubnet has not yet been implemented")
		})
	}
	if api.DhcpGetDHCPSubnetReservationHandler == nil {
		api.DhcpGetDHCPSubnetReservationHandler = dhcp.GetDHCPSubnetReservationHandlerFunc(func(params dhcp.GetDHCPSubnetReservationParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation dhcp.GetDHCPSubnetReservation has not yet been implemented")
		})
	}
	if api.HealthzHandler == nil {
		api.HealthzHandler = swaggerapi.HealthzHandlerFunc(func(params swaggerapi.HealthzParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation api.Healthz has not yet been implemented")
		})
	}
	if api.DhcpReloadDHCPHandler == nil {
		api.DhcpReloadDHCPHandler = dhcp.ReloadDHCPHandlerFunc(func(params dhcp.ReloadDHCPParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation dhcp.ReloadDHCP has not yet been implemented")
		})
	}
	if api.DhcpUpdateDHCPSubnetHandler == nil {
		api.DhcpUpdateDHCPSubnetHandler = dhcp.UpdateDHCPSubnetHandlerFunc(func(params dhcp.UpdateDHCPSubnetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation dhcp.UpdateDHCPSubnet has not yet been implemented")
		})
	}
	if api.DhcpUpdateDHCPSubnetReservationHandler == nil {
		api.DhcpUpdateDHCPSubnetReservationHandler = dhcp.UpdateDHCPSubnetReservationHandlerFunc(func(params dhcp.UpdateDHCPSubnetReservationParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation dhcp.UpdateDHCPSubnetReservation has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
