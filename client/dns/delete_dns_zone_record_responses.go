// Code generated by go-swagger; DO NOT EDIT.

package dns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteDNSZoneRecordReader is a Reader for the DeleteDNSZoneRecord structure.
type DeleteDNSZoneRecordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDNSZoneRecordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDNSZoneRecordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteDNSZoneRecordNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteDNSZoneRecordInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteDNSZoneRecordOK creates a DeleteDNSZoneRecordOK with default headers values
func NewDeleteDNSZoneRecordOK() *DeleteDNSZoneRecordOK {
	return &DeleteDNSZoneRecordOK{}
}

/*
DeleteDNSZoneRecordOK describes a response with status code 200, with default header values.

The record has been successfully removed.
*/
type DeleteDNSZoneRecordOK struct {
}

// IsSuccess returns true when this delete Dns zone record o k response has a 2xx status code
func (o *DeleteDNSZoneRecordOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete Dns zone record o k response has a 3xx status code
func (o *DeleteDNSZoneRecordOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Dns zone record o k response has a 4xx status code
func (o *DeleteDNSZoneRecordOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Dns zone record o k response has a 5xx status code
func (o *DeleteDNSZoneRecordOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Dns zone record o k response a status code equal to that given
func (o *DeleteDNSZoneRecordOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete Dns zone record o k response
func (o *DeleteDNSZoneRecordOK) Code() int {
	return 200
}

func (o *DeleteDNSZoneRecordOK) Error() string {
	return fmt.Sprintf("[DELETE /dns/zone/{zoneId}/record/{recordId}][%d] deleteDnsZoneRecordOK ", 200)
}

func (o *DeleteDNSZoneRecordOK) String() string {
	return fmt.Sprintf("[DELETE /dns/zone/{zoneId}/record/{recordId}][%d] deleteDnsZoneRecordOK ", 200)
}

func (o *DeleteDNSZoneRecordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDNSZoneRecordNotFound creates a DeleteDNSZoneRecordNotFound with default headers values
func NewDeleteDNSZoneRecordNotFound() *DeleteDNSZoneRecordNotFound {
	return &DeleteDNSZoneRecordNotFound{}
}

/*
DeleteDNSZoneRecordNotFound describes a response with status code 404, with default header values.

Invalid zone ID or record ID was provided.
*/
type DeleteDNSZoneRecordNotFound struct {
}

// IsSuccess returns true when this delete Dns zone record not found response has a 2xx status code
func (o *DeleteDNSZoneRecordNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Dns zone record not found response has a 3xx status code
func (o *DeleteDNSZoneRecordNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Dns zone record not found response has a 4xx status code
func (o *DeleteDNSZoneRecordNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Dns zone record not found response has a 5xx status code
func (o *DeleteDNSZoneRecordNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Dns zone record not found response a status code equal to that given
func (o *DeleteDNSZoneRecordNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete Dns zone record not found response
func (o *DeleteDNSZoneRecordNotFound) Code() int {
	return 404
}

func (o *DeleteDNSZoneRecordNotFound) Error() string {
	return fmt.Sprintf("[DELETE /dns/zone/{zoneId}/record/{recordId}][%d] deleteDnsZoneRecordNotFound ", 404)
}

func (o *DeleteDNSZoneRecordNotFound) String() string {
	return fmt.Sprintf("[DELETE /dns/zone/{zoneId}/record/{recordId}][%d] deleteDnsZoneRecordNotFound ", 404)
}

func (o *DeleteDNSZoneRecordNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDNSZoneRecordInternalServerError creates a DeleteDNSZoneRecordInternalServerError with default headers values
func NewDeleteDNSZoneRecordInternalServerError() *DeleteDNSZoneRecordInternalServerError {
	return &DeleteDNSZoneRecordInternalServerError{}
}

/*
DeleteDNSZoneRecordInternalServerError describes a response with status code 500, with default header values.

Unable to remove the requested record.
*/
type DeleteDNSZoneRecordInternalServerError struct {
}

// IsSuccess returns true when this delete Dns zone record internal server error response has a 2xx status code
func (o *DeleteDNSZoneRecordInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Dns zone record internal server error response has a 3xx status code
func (o *DeleteDNSZoneRecordInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Dns zone record internal server error response has a 4xx status code
func (o *DeleteDNSZoneRecordInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Dns zone record internal server error response has a 5xx status code
func (o *DeleteDNSZoneRecordInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete Dns zone record internal server error response a status code equal to that given
func (o *DeleteDNSZoneRecordInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete Dns zone record internal server error response
func (o *DeleteDNSZoneRecordInternalServerError) Code() int {
	return 500
}

func (o *DeleteDNSZoneRecordInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /dns/zone/{zoneId}/record/{recordId}][%d] deleteDnsZoneRecordInternalServerError ", 500)
}

func (o *DeleteDNSZoneRecordInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /dns/zone/{zoneId}/record/{recordId}][%d] deleteDnsZoneRecordInternalServerError ", 500)
}

func (o *DeleteDNSZoneRecordInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
