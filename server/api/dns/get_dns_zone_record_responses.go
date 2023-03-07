// Code generated by go-swagger; DO NOT EDIT.

package dns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/dalet-oss/iris-api/models"
)

// GetDNSZoneRecordOKCode is the HTTP code returned for type GetDNSZoneRecordOK
const GetDNSZoneRecordOKCode int = 200

/*
GetDNSZoneRecordOK Returns the DNS Record object.

swagger:response getDnsZoneRecordOK
*/
type GetDNSZoneRecordOK struct {

	/*
	  In: Body
	*/
	Payload *models.Record `json:"body,omitempty"`
}

// NewGetDNSZoneRecordOK creates GetDNSZoneRecordOK with default headers values
func NewGetDNSZoneRecordOK() *GetDNSZoneRecordOK {

	return &GetDNSZoneRecordOK{}
}

// WithPayload adds the payload to the get Dns zone record o k response
func (o *GetDNSZoneRecordOK) WithPayload(payload *models.Record) *GetDNSZoneRecordOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Dns zone record o k response
func (o *GetDNSZoneRecordOK) SetPayload(payload *models.Record) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDNSZoneRecordOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDNSZoneRecordNotFoundCode is the HTTP code returned for type GetDNSZoneRecordNotFound
const GetDNSZoneRecordNotFoundCode int = 404

/*
GetDNSZoneRecordNotFound Invalid zone ID or record ID was provided.

swagger:response getDnsZoneRecordNotFound
*/
type GetDNSZoneRecordNotFound struct {
}

// NewGetDNSZoneRecordNotFound creates GetDNSZoneRecordNotFound with default headers values
func NewGetDNSZoneRecordNotFound() *GetDNSZoneRecordNotFound {

	return &GetDNSZoneRecordNotFound{}
}

// WriteResponse to the client
func (o *GetDNSZoneRecordNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
