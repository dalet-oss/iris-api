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

// UpdateDHCPSubnetReader is a Reader for the UpdateDHCPSubnet structure.
type UpdateDHCPSubnetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDHCPSubnetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDHCPSubnetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateDHCPSubnetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateDHCPSubnetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateDHCPSubnetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateDHCPSubnetOK creates a UpdateDHCPSubnetOK with default headers values
func NewUpdateDHCPSubnetOK() *UpdateDHCPSubnetOK {
	return &UpdateDHCPSubnetOK{}
}

/*
UpdateDHCPSubnetOK describes a response with status code 200, with default header values.

Returns the updated subnet object.
*/
type UpdateDHCPSubnetOK struct {
	Payload *models.Subnet
}

// IsSuccess returns true when this update d h c p subnet o k response has a 2xx status code
func (o *UpdateDHCPSubnetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update d h c p subnet o k response has a 3xx status code
func (o *UpdateDHCPSubnetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update d h c p subnet o k response has a 4xx status code
func (o *UpdateDHCPSubnetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update d h c p subnet o k response has a 5xx status code
func (o *UpdateDHCPSubnetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update d h c p subnet o k response a status code equal to that given
func (o *UpdateDHCPSubnetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update d h c p subnet o k response
func (o *UpdateDHCPSubnetOK) Code() int {
	return 200
}

func (o *UpdateDHCPSubnetOK) Error() string {
	return fmt.Sprintf("[PUT /dhcp/subnet/{subnetId}][%d] updateDHCPSubnetOK  %+v", 200, o.Payload)
}

func (o *UpdateDHCPSubnetOK) String() string {
	return fmt.Sprintf("[PUT /dhcp/subnet/{subnetId}][%d] updateDHCPSubnetOK  %+v", 200, o.Payload)
}

func (o *UpdateDHCPSubnetOK) GetPayload() *models.Subnet {
	return o.Payload
}

func (o *UpdateDHCPSubnetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Subnet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDHCPSubnetBadRequest creates a UpdateDHCPSubnetBadRequest with default headers values
func NewUpdateDHCPSubnetBadRequest() *UpdateDHCPSubnetBadRequest {
	return &UpdateDHCPSubnetBadRequest{}
}

/*
UpdateDHCPSubnetBadRequest describes a response with status code 400, with default header values.

Bad parameters were provided.
*/
type UpdateDHCPSubnetBadRequest struct {
}

// IsSuccess returns true when this update d h c p subnet bad request response has a 2xx status code
func (o *UpdateDHCPSubnetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update d h c p subnet bad request response has a 3xx status code
func (o *UpdateDHCPSubnetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update d h c p subnet bad request response has a 4xx status code
func (o *UpdateDHCPSubnetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update d h c p subnet bad request response has a 5xx status code
func (o *UpdateDHCPSubnetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update d h c p subnet bad request response a status code equal to that given
func (o *UpdateDHCPSubnetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update d h c p subnet bad request response
func (o *UpdateDHCPSubnetBadRequest) Code() int {
	return 400
}

func (o *UpdateDHCPSubnetBadRequest) Error() string {
	return fmt.Sprintf("[PUT /dhcp/subnet/{subnetId}][%d] updateDHCPSubnetBadRequest ", 400)
}

func (o *UpdateDHCPSubnetBadRequest) String() string {
	return fmt.Sprintf("[PUT /dhcp/subnet/{subnetId}][%d] updateDHCPSubnetBadRequest ", 400)
}

func (o *UpdateDHCPSubnetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateDHCPSubnetNotFound creates a UpdateDHCPSubnetNotFound with default headers values
func NewUpdateDHCPSubnetNotFound() *UpdateDHCPSubnetNotFound {
	return &UpdateDHCPSubnetNotFound{}
}

/*
UpdateDHCPSubnetNotFound describes a response with status code 404, with default header values.

Invalid subnet ID was provided.
*/
type UpdateDHCPSubnetNotFound struct {
}

// IsSuccess returns true when this update d h c p subnet not found response has a 2xx status code
func (o *UpdateDHCPSubnetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update d h c p subnet not found response has a 3xx status code
func (o *UpdateDHCPSubnetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update d h c p subnet not found response has a 4xx status code
func (o *UpdateDHCPSubnetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update d h c p subnet not found response has a 5xx status code
func (o *UpdateDHCPSubnetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update d h c p subnet not found response a status code equal to that given
func (o *UpdateDHCPSubnetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update d h c p subnet not found response
func (o *UpdateDHCPSubnetNotFound) Code() int {
	return 404
}

func (o *UpdateDHCPSubnetNotFound) Error() string {
	return fmt.Sprintf("[PUT /dhcp/subnet/{subnetId}][%d] updateDHCPSubnetNotFound ", 404)
}

func (o *UpdateDHCPSubnetNotFound) String() string {
	return fmt.Sprintf("[PUT /dhcp/subnet/{subnetId}][%d] updateDHCPSubnetNotFound ", 404)
}

func (o *UpdateDHCPSubnetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateDHCPSubnetInternalServerError creates a UpdateDHCPSubnetInternalServerError with default headers values
func NewUpdateDHCPSubnetInternalServerError() *UpdateDHCPSubnetInternalServerError {
	return &UpdateDHCPSubnetInternalServerError{}
}

/*
UpdateDHCPSubnetInternalServerError describes a response with status code 500, with default header values.

Unable to update subnet.
*/
type UpdateDHCPSubnetInternalServerError struct {
}

// IsSuccess returns true when this update d h c p subnet internal server error response has a 2xx status code
func (o *UpdateDHCPSubnetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update d h c p subnet internal server error response has a 3xx status code
func (o *UpdateDHCPSubnetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update d h c p subnet internal server error response has a 4xx status code
func (o *UpdateDHCPSubnetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update d h c p subnet internal server error response has a 5xx status code
func (o *UpdateDHCPSubnetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update d h c p subnet internal server error response a status code equal to that given
func (o *UpdateDHCPSubnetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update d h c p subnet internal server error response
func (o *UpdateDHCPSubnetInternalServerError) Code() int {
	return 500
}

func (o *UpdateDHCPSubnetInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /dhcp/subnet/{subnetId}][%d] updateDHCPSubnetInternalServerError ", 500)
}

func (o *UpdateDHCPSubnetInternalServerError) String() string {
	return fmt.Sprintf("[PUT /dhcp/subnet/{subnetId}][%d] updateDHCPSubnetInternalServerError ", 500)
}

func (o *UpdateDHCPSubnetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
