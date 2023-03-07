// Code generated by go-swagger; DO NOT EDIT.

package dns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DisableDNSZoneRecordCreatedCode is the HTTP code returned for type DisableDNSZoneRecordCreated
const DisableDNSZoneRecordCreatedCode int = 201

/*
DisableDNSZoneRecordCreated The DNS zone's record has been disabled.

swagger:response disableDnsZoneRecordCreated
*/
type DisableDNSZoneRecordCreated struct {
}

// NewDisableDNSZoneRecordCreated creates DisableDNSZoneRecordCreated with default headers values
func NewDisableDNSZoneRecordCreated() *DisableDNSZoneRecordCreated {

	return &DisableDNSZoneRecordCreated{}
}

// WriteResponse to the client
func (o *DisableDNSZoneRecordCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// DisableDNSZoneRecordConflictCode is the HTTP code returned for type DisableDNSZoneRecordConflict
const DisableDNSZoneRecordConflictCode int = 409

/*
DisableDNSZoneRecordConflict The DNS zone's record was already disable.

swagger:response disableDnsZoneRecordConflict
*/
type DisableDNSZoneRecordConflict struct {
}

// NewDisableDNSZoneRecordConflict creates DisableDNSZoneRecordConflict with default headers values
func NewDisableDNSZoneRecordConflict() *DisableDNSZoneRecordConflict {

	return &DisableDNSZoneRecordConflict{}
}

// WriteResponse to the client
func (o *DisableDNSZoneRecordConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}

// DisableDNSZoneRecordInternalServerErrorCode is the HTTP code returned for type DisableDNSZoneRecordInternalServerError
const DisableDNSZoneRecordInternalServerErrorCode int = 500

/*
DisableDNSZoneRecordInternalServerError Unable to disable the DNS zone's record.

swagger:response disableDnsZoneRecordInternalServerError
*/
type DisableDNSZoneRecordInternalServerError struct {
}

// NewDisableDNSZoneRecordInternalServerError creates DisableDNSZoneRecordInternalServerError with default headers values
func NewDisableDNSZoneRecordInternalServerError() *DisableDNSZoneRecordInternalServerError {

	return &DisableDNSZoneRecordInternalServerError{}
}

// WriteResponse to the client
func (o *DisableDNSZoneRecordInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
