// Code generated by go-swagger; DO NOT EDIT.

package dns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/dalet-oss/iris-api/models"
)

// CreateDNSZoneReader is a Reader for the CreateDNSZone structure.
type CreateDNSZoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDNSZoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateDNSZoneCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDNSZoneBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateDNSZoneConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateDNSZoneInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateDNSZoneCreated creates a CreateDNSZoneCreated with default headers values
func NewCreateDNSZoneCreated() *CreateDNSZoneCreated {
	return &CreateDNSZoneCreated{}
}

/*
CreateDNSZoneCreated describes a response with status code 201, with default header values.

Returns the newly created zone object.
*/
type CreateDNSZoneCreated struct {
	Payload *models.Zone
}

// IsSuccess returns true when this create Dns zone created response has a 2xx status code
func (o *CreateDNSZoneCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create Dns zone created response has a 3xx status code
func (o *CreateDNSZoneCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create Dns zone created response has a 4xx status code
func (o *CreateDNSZoneCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create Dns zone created response has a 5xx status code
func (o *CreateDNSZoneCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create Dns zone created response a status code equal to that given
func (o *CreateDNSZoneCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create Dns zone created response
func (o *CreateDNSZoneCreated) Code() int {
	return 201
}

func (o *CreateDNSZoneCreated) Error() string {
	return fmt.Sprintf("[POST /dns/zone][%d] createDnsZoneCreated  %+v", 201, o.Payload)
}

func (o *CreateDNSZoneCreated) String() string {
	return fmt.Sprintf("[POST /dns/zone][%d] createDnsZoneCreated  %+v", 201, o.Payload)
}

func (o *CreateDNSZoneCreated) GetPayload() *models.Zone {
	return o.Payload
}

func (o *CreateDNSZoneCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Zone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDNSZoneBadRequest creates a CreateDNSZoneBadRequest with default headers values
func NewCreateDNSZoneBadRequest() *CreateDNSZoneBadRequest {
	return &CreateDNSZoneBadRequest{}
}

/*
CreateDNSZoneBadRequest describes a response with status code 400, with default header values.

Bad parameters were provided.
*/
type CreateDNSZoneBadRequest struct {
}

// IsSuccess returns true when this create Dns zone bad request response has a 2xx status code
func (o *CreateDNSZoneBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create Dns zone bad request response has a 3xx status code
func (o *CreateDNSZoneBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create Dns zone bad request response has a 4xx status code
func (o *CreateDNSZoneBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create Dns zone bad request response has a 5xx status code
func (o *CreateDNSZoneBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create Dns zone bad request response a status code equal to that given
func (o *CreateDNSZoneBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create Dns zone bad request response
func (o *CreateDNSZoneBadRequest) Code() int {
	return 400
}

func (o *CreateDNSZoneBadRequest) Error() string {
	return fmt.Sprintf("[POST /dns/zone][%d] createDnsZoneBadRequest ", 400)
}

func (o *CreateDNSZoneBadRequest) String() string {
	return fmt.Sprintf("[POST /dns/zone][%d] createDnsZoneBadRequest ", 400)
}

func (o *CreateDNSZoneBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateDNSZoneConflict creates a CreateDNSZoneConflict with default headers values
func NewCreateDNSZoneConflict() *CreateDNSZoneConflict {
	return &CreateDNSZoneConflict{}
}

/*
CreateDNSZoneConflict describes a response with status code 409, with default header values.

Zone already exists.
*/
type CreateDNSZoneConflict struct {
}

// IsSuccess returns true when this create Dns zone conflict response has a 2xx status code
func (o *CreateDNSZoneConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create Dns zone conflict response has a 3xx status code
func (o *CreateDNSZoneConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create Dns zone conflict response has a 4xx status code
func (o *CreateDNSZoneConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create Dns zone conflict response has a 5xx status code
func (o *CreateDNSZoneConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create Dns zone conflict response a status code equal to that given
func (o *CreateDNSZoneConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create Dns zone conflict response
func (o *CreateDNSZoneConflict) Code() int {
	return 409
}

func (o *CreateDNSZoneConflict) Error() string {
	return fmt.Sprintf("[POST /dns/zone][%d] createDnsZoneConflict ", 409)
}

func (o *CreateDNSZoneConflict) String() string {
	return fmt.Sprintf("[POST /dns/zone][%d] createDnsZoneConflict ", 409)
}

func (o *CreateDNSZoneConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateDNSZoneInternalServerError creates a CreateDNSZoneInternalServerError with default headers values
func NewCreateDNSZoneInternalServerError() *CreateDNSZoneInternalServerError {
	return &CreateDNSZoneInternalServerError{}
}

/*
CreateDNSZoneInternalServerError describes a response with status code 500, with default header values.

Unable to create zone.
*/
type CreateDNSZoneInternalServerError struct {
}

// IsSuccess returns true when this create Dns zone internal server error response has a 2xx status code
func (o *CreateDNSZoneInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create Dns zone internal server error response has a 3xx status code
func (o *CreateDNSZoneInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create Dns zone internal server error response has a 4xx status code
func (o *CreateDNSZoneInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create Dns zone internal server error response has a 5xx status code
func (o *CreateDNSZoneInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create Dns zone internal server error response a status code equal to that given
func (o *CreateDNSZoneInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create Dns zone internal server error response
func (o *CreateDNSZoneInternalServerError) Code() int {
	return 500
}

func (o *CreateDNSZoneInternalServerError) Error() string {
	return fmt.Sprintf("[POST /dns/zone][%d] createDnsZoneInternalServerError ", 500)
}

func (o *CreateDNSZoneInternalServerError) String() string {
	return fmt.Sprintf("[POST /dns/zone][%d] createDnsZoneInternalServerError ", 500)
}

func (o *CreateDNSZoneInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}