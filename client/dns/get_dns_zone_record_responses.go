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

// GetDNSZoneRecordReader is a Reader for the GetDNSZoneRecord structure.
type GetDNSZoneRecordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDNSZoneRecordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDNSZoneRecordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetDNSZoneRecordNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDNSZoneRecordOK creates a GetDNSZoneRecordOK with default headers values
func NewGetDNSZoneRecordOK() *GetDNSZoneRecordOK {
	return &GetDNSZoneRecordOK{}
}

/*
GetDNSZoneRecordOK describes a response with status code 200, with default header values.

Returns the DNS Record object.
*/
type GetDNSZoneRecordOK struct {
	Payload *models.Record
}

// IsSuccess returns true when this get Dns zone record o k response has a 2xx status code
func (o *GetDNSZoneRecordOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Dns zone record o k response has a 3xx status code
func (o *GetDNSZoneRecordOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Dns zone record o k response has a 4xx status code
func (o *GetDNSZoneRecordOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Dns zone record o k response has a 5xx status code
func (o *GetDNSZoneRecordOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Dns zone record o k response a status code equal to that given
func (o *GetDNSZoneRecordOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Dns zone record o k response
func (o *GetDNSZoneRecordOK) Code() int {
	return 200
}

func (o *GetDNSZoneRecordOK) Error() string {
	return fmt.Sprintf("[GET /dns/zone/{zoneId}/record/{recordId}][%d] getDnsZoneRecordOK  %+v", 200, o.Payload)
}

func (o *GetDNSZoneRecordOK) String() string {
	return fmt.Sprintf("[GET /dns/zone/{zoneId}/record/{recordId}][%d] getDnsZoneRecordOK  %+v", 200, o.Payload)
}

func (o *GetDNSZoneRecordOK) GetPayload() *models.Record {
	return o.Payload
}

func (o *GetDNSZoneRecordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Record)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDNSZoneRecordNotFound creates a GetDNSZoneRecordNotFound with default headers values
func NewGetDNSZoneRecordNotFound() *GetDNSZoneRecordNotFound {
	return &GetDNSZoneRecordNotFound{}
}

/*
GetDNSZoneRecordNotFound describes a response with status code 404, with default header values.

Invalid zone ID or record ID was provided.
*/
type GetDNSZoneRecordNotFound struct {
}

// IsSuccess returns true when this get Dns zone record not found response has a 2xx status code
func (o *GetDNSZoneRecordNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Dns zone record not found response has a 3xx status code
func (o *GetDNSZoneRecordNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Dns zone record not found response has a 4xx status code
func (o *GetDNSZoneRecordNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Dns zone record not found response has a 5xx status code
func (o *GetDNSZoneRecordNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get Dns zone record not found response a status code equal to that given
func (o *GetDNSZoneRecordNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get Dns zone record not found response
func (o *GetDNSZoneRecordNotFound) Code() int {
	return 404
}

func (o *GetDNSZoneRecordNotFound) Error() string {
	return fmt.Sprintf("[GET /dns/zone/{zoneId}/record/{recordId}][%d] getDnsZoneRecordNotFound ", 404)
}

func (o *GetDNSZoneRecordNotFound) String() string {
	return fmt.Sprintf("[GET /dns/zone/{zoneId}/record/{recordId}][%d] getDnsZoneRecordNotFound ", 404)
}

func (o *GetDNSZoneRecordNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
