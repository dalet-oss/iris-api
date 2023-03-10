// Code generated by go-swagger; DO NOT EDIT.

package dns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetAllDNSZonesOKCode is the HTTP code returned for type GetAllDNSZonesOK
const GetAllDNSZonesOKCode int = 200

/*
GetAllDNSZonesOK Returns the an array of DNS zones.

swagger:response getAllDnsZonesOK
*/
type GetAllDNSZonesOK struct {

	/*
	  In: Body
	*/
	Payload []string `json:"body,omitempty"`
}

// NewGetAllDNSZonesOK creates GetAllDNSZonesOK with default headers values
func NewGetAllDNSZonesOK() *GetAllDNSZonesOK {

	return &GetAllDNSZonesOK{}
}

// WithPayload adds the payload to the get all Dns zones o k response
func (o *GetAllDNSZonesOK) WithPayload(payload []string) *GetAllDNSZonesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all Dns zones o k response
func (o *GetAllDNSZonesOK) SetPayload(payload []string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllDNSZonesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]string, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
