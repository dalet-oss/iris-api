// Code generated by go-swagger; DO NOT EDIT.

package dhcp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/dalet-oss/iris-api/models"
)

// CreateDHCPSubnetReader is a Reader for the CreateDHCPSubnet structure.
type CreateDHCPSubnetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDHCPSubnetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateDHCPSubnetCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDHCPSubnetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateDHCPSubnetConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateDHCPSubnetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateDHCPSubnetCreated creates a CreateDHCPSubnetCreated with default headers values
func NewCreateDHCPSubnetCreated() *CreateDHCPSubnetCreated {
	return &CreateDHCPSubnetCreated{}
}

/*
CreateDHCPSubnetCreated describes a response with status code 201, with default header values.

Returns the newly created subnet object.
*/
type CreateDHCPSubnetCreated struct {
	Payload *models.Subnet
}

// IsSuccess returns true when this create d h c p subnet created response has a 2xx status code
func (o *CreateDHCPSubnetCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create d h c p subnet created response has a 3xx status code
func (o *CreateDHCPSubnetCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create d h c p subnet created response has a 4xx status code
func (o *CreateDHCPSubnetCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create d h c p subnet created response has a 5xx status code
func (o *CreateDHCPSubnetCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create d h c p subnet created response a status code equal to that given
func (o *CreateDHCPSubnetCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create d h c p subnet created response
func (o *CreateDHCPSubnetCreated) Code() int {
	return 201
}

func (o *CreateDHCPSubnetCreated) Error() string {
	return fmt.Sprintf("[POST /dhcp/subnet][%d] createDHCPSubnetCreated  %+v", 201, o.Payload)
}

func (o *CreateDHCPSubnetCreated) String() string {
	return fmt.Sprintf("[POST /dhcp/subnet][%d] createDHCPSubnetCreated  %+v", 201, o.Payload)
}

func (o *CreateDHCPSubnetCreated) GetPayload() *models.Subnet {
	return o.Payload
}

func (o *CreateDHCPSubnetCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Subnet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDHCPSubnetBadRequest creates a CreateDHCPSubnetBadRequest with default headers values
func NewCreateDHCPSubnetBadRequest() *CreateDHCPSubnetBadRequest {
	return &CreateDHCPSubnetBadRequest{}
}

/*
CreateDHCPSubnetBadRequest describes a response with status code 400, with default header values.

Bad parameters were provided.
*/
type CreateDHCPSubnetBadRequest struct {
}

// IsSuccess returns true when this create d h c p subnet bad request response has a 2xx status code
func (o *CreateDHCPSubnetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create d h c p subnet bad request response has a 3xx status code
func (o *CreateDHCPSubnetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create d h c p subnet bad request response has a 4xx status code
func (o *CreateDHCPSubnetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create d h c p subnet bad request response has a 5xx status code
func (o *CreateDHCPSubnetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create d h c p subnet bad request response a status code equal to that given
func (o *CreateDHCPSubnetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create d h c p subnet bad request response
func (o *CreateDHCPSubnetBadRequest) Code() int {
	return 400
}

func (o *CreateDHCPSubnetBadRequest) Error() string {
	return fmt.Sprintf("[POST /dhcp/subnet][%d] createDHCPSubnetBadRequest ", 400)
}

func (o *CreateDHCPSubnetBadRequest) String() string {
	return fmt.Sprintf("[POST /dhcp/subnet][%d] createDHCPSubnetBadRequest ", 400)
}

func (o *CreateDHCPSubnetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateDHCPSubnetConflict creates a CreateDHCPSubnetConflict with default headers values
func NewCreateDHCPSubnetConflict() *CreateDHCPSubnetConflict {
	return &CreateDHCPSubnetConflict{}
}

/*
CreateDHCPSubnetConflict describes a response with status code 409, with default header values.

Subnet already exists.
*/
type CreateDHCPSubnetConflict struct {
}

// IsSuccess returns true when this create d h c p subnet conflict response has a 2xx status code
func (o *CreateDHCPSubnetConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create d h c p subnet conflict response has a 3xx status code
func (o *CreateDHCPSubnetConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create d h c p subnet conflict response has a 4xx status code
func (o *CreateDHCPSubnetConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create d h c p subnet conflict response has a 5xx status code
func (o *CreateDHCPSubnetConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create d h c p subnet conflict response a status code equal to that given
func (o *CreateDHCPSubnetConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create d h c p subnet conflict response
func (o *CreateDHCPSubnetConflict) Code() int {
	return 409
}

func (o *CreateDHCPSubnetConflict) Error() string {
	return fmt.Sprintf("[POST /dhcp/subnet][%d] createDHCPSubnetConflict ", 409)
}

func (o *CreateDHCPSubnetConflict) String() string {
	return fmt.Sprintf("[POST /dhcp/subnet][%d] createDHCPSubnetConflict ", 409)
}

func (o *CreateDHCPSubnetConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateDHCPSubnetInternalServerError creates a CreateDHCPSubnetInternalServerError with default headers values
func NewCreateDHCPSubnetInternalServerError() *CreateDHCPSubnetInternalServerError {
	return &CreateDHCPSubnetInternalServerError{}
}

/*
CreateDHCPSubnetInternalServerError describes a response with status code 500, with default header values.

Unable to create Subnet's resevration.
*/
type CreateDHCPSubnetInternalServerError struct {
}

// IsSuccess returns true when this create d h c p subnet internal server error response has a 2xx status code
func (o *CreateDHCPSubnetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create d h c p subnet internal server error response has a 3xx status code
func (o *CreateDHCPSubnetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create d h c p subnet internal server error response has a 4xx status code
func (o *CreateDHCPSubnetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create d h c p subnet internal server error response has a 5xx status code
func (o *CreateDHCPSubnetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create d h c p subnet internal server error response a status code equal to that given
func (o *CreateDHCPSubnetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create d h c p subnet internal server error response
func (o *CreateDHCPSubnetInternalServerError) Code() int {
	return 500
}

func (o *CreateDHCPSubnetInternalServerError) Error() string {
	return fmt.Sprintf("[POST /dhcp/subnet][%d] createDHCPSubnetInternalServerError ", 500)
}

func (o *CreateDHCPSubnetInternalServerError) String() string {
	return fmt.Sprintf("[POST /dhcp/subnet][%d] createDHCPSubnetInternalServerError ", 500)
}

func (o *CreateDHCPSubnetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
