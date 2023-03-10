// Code generated by go-swagger; DO NOT EDIT.

package dhcp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetAllDHCPSubnetsReader is a Reader for the GetAllDHCPSubnets structure.
type GetAllDHCPSubnetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllDHCPSubnetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllDHCPSubnetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllDHCPSubnetsOK creates a GetAllDHCPSubnetsOK with default headers values
func NewGetAllDHCPSubnetsOK() *GetAllDHCPSubnetsOK {
	return &GetAllDHCPSubnetsOK{}
}

/*
GetAllDHCPSubnetsOK describes a response with status code 200, with default header values.

Returns the an array of DHCPv4 registered subnets.
*/
type GetAllDHCPSubnetsOK struct {
	Payload []string
}

// IsSuccess returns true when this get all d h c p subnets o k response has a 2xx status code
func (o *GetAllDHCPSubnetsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all d h c p subnets o k response has a 3xx status code
func (o *GetAllDHCPSubnetsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all d h c p subnets o k response has a 4xx status code
func (o *GetAllDHCPSubnetsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all d h c p subnets o k response has a 5xx status code
func (o *GetAllDHCPSubnetsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all d h c p subnets o k response a status code equal to that given
func (o *GetAllDHCPSubnetsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get all d h c p subnets o k response
func (o *GetAllDHCPSubnetsOK) Code() int {
	return 200
}

func (o *GetAllDHCPSubnetsOK) Error() string {
	return fmt.Sprintf("[GET /dhcp/subnet][%d] getAllDHCPSubnetsOK  %+v", 200, o.Payload)
}

func (o *GetAllDHCPSubnetsOK) String() string {
	return fmt.Sprintf("[GET /dhcp/subnet][%d] getAllDHCPSubnetsOK  %+v", 200, o.Payload)
}

func (o *GetAllDHCPSubnetsOK) GetPayload() []string {
	return o.Payload
}

func (o *GetAllDHCPSubnetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
