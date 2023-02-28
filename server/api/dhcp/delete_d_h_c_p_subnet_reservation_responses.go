// Code generated by go-swagger; DO NOT EDIT.

package dhcp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteDHCPSubnetReservationOKCode is the HTTP code returned for type DeleteDHCPSubnetReservationOK
const DeleteDHCPSubnetReservationOKCode int = 200

/*
DeleteDHCPSubnetReservationOK The reservation has been successfully removed.

swagger:response deleteDHCPSubnetReservationOK
*/
type DeleteDHCPSubnetReservationOK struct {
}

// NewDeleteDHCPSubnetReservationOK creates DeleteDHCPSubnetReservationOK with default headers values
func NewDeleteDHCPSubnetReservationOK() *DeleteDHCPSubnetReservationOK {

	return &DeleteDHCPSubnetReservationOK{}
}

// WriteResponse to the client
func (o *DeleteDHCPSubnetReservationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteDHCPSubnetReservationNotFoundCode is the HTTP code returned for type DeleteDHCPSubnetReservationNotFound
const DeleteDHCPSubnetReservationNotFoundCode int = 404

/*
DeleteDHCPSubnetReservationNotFound Invalid subnet ID or reservation MAC address was provided.

swagger:response deleteDHCPSubnetReservationNotFound
*/
type DeleteDHCPSubnetReservationNotFound struct {
}

// NewDeleteDHCPSubnetReservationNotFound creates DeleteDHCPSubnetReservationNotFound with default headers values
func NewDeleteDHCPSubnetReservationNotFound() *DeleteDHCPSubnetReservationNotFound {

	return &DeleteDHCPSubnetReservationNotFound{}
}

// WriteResponse to the client
func (o *DeleteDHCPSubnetReservationNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// DeleteDHCPSubnetReservationInternalServerErrorCode is the HTTP code returned for type DeleteDHCPSubnetReservationInternalServerError
const DeleteDHCPSubnetReservationInternalServerErrorCode int = 500

/*
DeleteDHCPSubnetReservationInternalServerError Unable to remove the requested reservation.

swagger:response deleteDHCPSubnetReservationInternalServerError
*/
type DeleteDHCPSubnetReservationInternalServerError struct {
}

// NewDeleteDHCPSubnetReservationInternalServerError creates DeleteDHCPSubnetReservationInternalServerError with default headers values
func NewDeleteDHCPSubnetReservationInternalServerError() *DeleteDHCPSubnetReservationInternalServerError {

	return &DeleteDHCPSubnetReservationInternalServerError{}
}

// WriteResponse to the client
func (o *DeleteDHCPSubnetReservationInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}