// Code generated by go-swagger; DO NOT EDIT.

package dhcp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetAllDHCPSubnetsOKCode is the HTTP code returned for type GetAllDHCPSubnetsOK
const GetAllDHCPSubnetsOKCode int = 200

/*
GetAllDHCPSubnetsOK Returns the an array of DHCPv4 registered subnets.

swagger:response getAllDHCPSubnetsOK
*/
type GetAllDHCPSubnetsOK struct {

	/*
	  In: Body
	*/
	Payload []string `json:"body,omitempty"`
}

// NewGetAllDHCPSubnetsOK creates GetAllDHCPSubnetsOK with default headers values
func NewGetAllDHCPSubnetsOK() *GetAllDHCPSubnetsOK {

	return &GetAllDHCPSubnetsOK{}
}

// WithPayload adds the payload to the get all d h c p subnets o k response
func (o *GetAllDHCPSubnetsOK) WithPayload(payload []string) *GetAllDHCPSubnetsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all d h c p subnets o k response
func (o *GetAllDHCPSubnetsOK) SetPayload(payload []string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllDHCPSubnetsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
