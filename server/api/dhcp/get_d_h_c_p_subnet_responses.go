// Code generated by go-swagger; DO NOT EDIT.

package dhcp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/dalet-oss/iris-api/models"
)

// GetDHCPSubnetOKCode is the HTTP code returned for type GetDHCPSubnetOK
const GetDHCPSubnetOKCode int = 200

/*
GetDHCPSubnetOK Returns the DHCP Subnet object.

swagger:response getDHCPSubnetOK
*/
type GetDHCPSubnetOK struct {

	/*
	  In: Body
	*/
	Payload *models.Subnet `json:"body,omitempty"`
}

// NewGetDHCPSubnetOK creates GetDHCPSubnetOK with default headers values
func NewGetDHCPSubnetOK() *GetDHCPSubnetOK {

	return &GetDHCPSubnetOK{}
}

// WithPayload adds the payload to the get d h c p subnet o k response
func (o *GetDHCPSubnetOK) WithPayload(payload *models.Subnet) *GetDHCPSubnetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get d h c p subnet o k response
func (o *GetDHCPSubnetOK) SetPayload(payload *models.Subnet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDHCPSubnetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDHCPSubnetNotFoundCode is the HTTP code returned for type GetDHCPSubnetNotFound
const GetDHCPSubnetNotFoundCode int = 404

/*
GetDHCPSubnetNotFound Invalid subnet ID was provided.

swagger:response getDHCPSubnetNotFound
*/
type GetDHCPSubnetNotFound struct {
}

// NewGetDHCPSubnetNotFound creates GetDHCPSubnetNotFound with default headers values
func NewGetDHCPSubnetNotFound() *GetDHCPSubnetNotFound {

	return &GetDHCPSubnetNotFound{}
}

// WriteResponse to the client
func (o *GetDHCPSubnetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
