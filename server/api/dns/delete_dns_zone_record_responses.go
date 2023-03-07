// Code generated by go-swagger; DO NOT EDIT.

package dns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteDNSZoneRecordOKCode is the HTTP code returned for type DeleteDNSZoneRecordOK
const DeleteDNSZoneRecordOKCode int = 200

/*
DeleteDNSZoneRecordOK The record has been successfully removed.

swagger:response deleteDnsZoneRecordOK
*/
type DeleteDNSZoneRecordOK struct {
}

// NewDeleteDNSZoneRecordOK creates DeleteDNSZoneRecordOK with default headers values
func NewDeleteDNSZoneRecordOK() *DeleteDNSZoneRecordOK {

	return &DeleteDNSZoneRecordOK{}
}

// WriteResponse to the client
func (o *DeleteDNSZoneRecordOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteDNSZoneRecordNotFoundCode is the HTTP code returned for type DeleteDNSZoneRecordNotFound
const DeleteDNSZoneRecordNotFoundCode int = 404

/*
DeleteDNSZoneRecordNotFound Invalid zone ID or record ID was provided.

swagger:response deleteDnsZoneRecordNotFound
*/
type DeleteDNSZoneRecordNotFound struct {
}

// NewDeleteDNSZoneRecordNotFound creates DeleteDNSZoneRecordNotFound with default headers values
func NewDeleteDNSZoneRecordNotFound() *DeleteDNSZoneRecordNotFound {

	return &DeleteDNSZoneRecordNotFound{}
}

// WriteResponse to the client
func (o *DeleteDNSZoneRecordNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// DeleteDNSZoneRecordInternalServerErrorCode is the HTTP code returned for type DeleteDNSZoneRecordInternalServerError
const DeleteDNSZoneRecordInternalServerErrorCode int = 500

/*
DeleteDNSZoneRecordInternalServerError Unable to remove the requested record.

swagger:response deleteDnsZoneRecordInternalServerError
*/
type DeleteDNSZoneRecordInternalServerError struct {
}

// NewDeleteDNSZoneRecordInternalServerError creates DeleteDNSZoneRecordInternalServerError with default headers values
func NewDeleteDNSZoneRecordInternalServerError() *DeleteDNSZoneRecordInternalServerError {

	return &DeleteDNSZoneRecordInternalServerError{}
}

// WriteResponse to the client
func (o *DeleteDNSZoneRecordInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
