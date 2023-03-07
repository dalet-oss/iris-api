// Code generated by go-swagger; DO NOT EDIT.

package dns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteDNSZoneReader is a Reader for the DeleteDNSZone structure.
type DeleteDNSZoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDNSZoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDNSZoneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteDNSZoneNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteDNSZoneConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteDNSZoneInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteDNSZoneOK creates a DeleteDNSZoneOK with default headers values
func NewDeleteDNSZoneOK() *DeleteDNSZoneOK {
	return &DeleteDNSZoneOK{}
}

/*
DeleteDNSZoneOK describes a response with status code 200, with default header values.

The zone has been successfully removed.
*/
type DeleteDNSZoneOK struct {
}

// IsSuccess returns true when this delete Dns zone o k response has a 2xx status code
func (o *DeleteDNSZoneOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete Dns zone o k response has a 3xx status code
func (o *DeleteDNSZoneOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Dns zone o k response has a 4xx status code
func (o *DeleteDNSZoneOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Dns zone o k response has a 5xx status code
func (o *DeleteDNSZoneOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Dns zone o k response a status code equal to that given
func (o *DeleteDNSZoneOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete Dns zone o k response
func (o *DeleteDNSZoneOK) Code() int {
	return 200
}

func (o *DeleteDNSZoneOK) Error() string {
	return fmt.Sprintf("[DELETE /dns/zone/{zoneId}][%d] deleteDnsZoneOK ", 200)
}

func (o *DeleteDNSZoneOK) String() string {
	return fmt.Sprintf("[DELETE /dns/zone/{zoneId}][%d] deleteDnsZoneOK ", 200)
}

func (o *DeleteDNSZoneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDNSZoneNotFound creates a DeleteDNSZoneNotFound with default headers values
func NewDeleteDNSZoneNotFound() *DeleteDNSZoneNotFound {
	return &DeleteDNSZoneNotFound{}
}

/*
DeleteDNSZoneNotFound describes a response with status code 404, with default header values.

Invalid zone ID was provided.
*/
type DeleteDNSZoneNotFound struct {
}

// IsSuccess returns true when this delete Dns zone not found response has a 2xx status code
func (o *DeleteDNSZoneNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Dns zone not found response has a 3xx status code
func (o *DeleteDNSZoneNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Dns zone not found response has a 4xx status code
func (o *DeleteDNSZoneNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Dns zone not found response has a 5xx status code
func (o *DeleteDNSZoneNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Dns zone not found response a status code equal to that given
func (o *DeleteDNSZoneNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete Dns zone not found response
func (o *DeleteDNSZoneNotFound) Code() int {
	return 404
}

func (o *DeleteDNSZoneNotFound) Error() string {
	return fmt.Sprintf("[DELETE /dns/zone/{zoneId}][%d] deleteDnsZoneNotFound ", 404)
}

func (o *DeleteDNSZoneNotFound) String() string {
	return fmt.Sprintf("[DELETE /dns/zone/{zoneId}][%d] deleteDnsZoneNotFound ", 404)
}

func (o *DeleteDNSZoneNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDNSZoneConflict creates a DeleteDNSZoneConflict with default headers values
func NewDeleteDNSZoneConflict() *DeleteDNSZoneConflict {
	return &DeleteDNSZoneConflict{}
}

/*
DeleteDNSZoneConflict describes a response with status code 409, with default header values.

The zone is not empty or still has associated records.
*/
type DeleteDNSZoneConflict struct {
}

// IsSuccess returns true when this delete Dns zone conflict response has a 2xx status code
func (o *DeleteDNSZoneConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Dns zone conflict response has a 3xx status code
func (o *DeleteDNSZoneConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Dns zone conflict response has a 4xx status code
func (o *DeleteDNSZoneConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Dns zone conflict response has a 5xx status code
func (o *DeleteDNSZoneConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Dns zone conflict response a status code equal to that given
func (o *DeleteDNSZoneConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the delete Dns zone conflict response
func (o *DeleteDNSZoneConflict) Code() int {
	return 409
}

func (o *DeleteDNSZoneConflict) Error() string {
	return fmt.Sprintf("[DELETE /dns/zone/{zoneId}][%d] deleteDnsZoneConflict ", 409)
}

func (o *DeleteDNSZoneConflict) String() string {
	return fmt.Sprintf("[DELETE /dns/zone/{zoneId}][%d] deleteDnsZoneConflict ", 409)
}

func (o *DeleteDNSZoneConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDNSZoneInternalServerError creates a DeleteDNSZoneInternalServerError with default headers values
func NewDeleteDNSZoneInternalServerError() *DeleteDNSZoneInternalServerError {
	return &DeleteDNSZoneInternalServerError{}
}

/*
DeleteDNSZoneInternalServerError describes a response with status code 500, with default header values.

Unable to remove the requested zone.
*/
type DeleteDNSZoneInternalServerError struct {
}

// IsSuccess returns true when this delete Dns zone internal server error response has a 2xx status code
func (o *DeleteDNSZoneInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Dns zone internal server error response has a 3xx status code
func (o *DeleteDNSZoneInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Dns zone internal server error response has a 4xx status code
func (o *DeleteDNSZoneInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Dns zone internal server error response has a 5xx status code
func (o *DeleteDNSZoneInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete Dns zone internal server error response a status code equal to that given
func (o *DeleteDNSZoneInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete Dns zone internal server error response
func (o *DeleteDNSZoneInternalServerError) Code() int {
	return 500
}

func (o *DeleteDNSZoneInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /dns/zone/{zoneId}][%d] deleteDnsZoneInternalServerError ", 500)
}

func (o *DeleteDNSZoneInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /dns/zone/{zoneId}][%d] deleteDnsZoneInternalServerError ", 500)
}

func (o *DeleteDNSZoneInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
