// Code generated by go-swagger; DO NOT EDIT.

package dhcp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ReloadDHCPCreatedCode is the HTTP code returned for type ReloadDHCPCreated
const ReloadDHCPCreatedCode int = 201

/*
ReloadDHCPCreated The DHCPv4 service configuration has been reloaded.

swagger:response reloadDHCPCreated
*/
type ReloadDHCPCreated struct {
}

// NewReloadDHCPCreated creates ReloadDHCPCreated with default headers values
func NewReloadDHCPCreated() *ReloadDHCPCreated {

	return &ReloadDHCPCreated{}
}

// WriteResponse to the client
func (o *ReloadDHCPCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// ReloadDHCPInternalServerErrorCode is the HTTP code returned for type ReloadDHCPInternalServerError
const ReloadDHCPInternalServerErrorCode int = 500

/*
ReloadDHCPInternalServerError Unable to reload the DHCPv4 service configuration.

swagger:response reloadDHCPInternalServerError
*/
type ReloadDHCPInternalServerError struct {
}

// NewReloadDHCPInternalServerError creates ReloadDHCPInternalServerError with default headers values
func NewReloadDHCPInternalServerError() *ReloadDHCPInternalServerError {

	return &ReloadDHCPInternalServerError{}
}

// WriteResponse to the client
func (o *ReloadDHCPInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}