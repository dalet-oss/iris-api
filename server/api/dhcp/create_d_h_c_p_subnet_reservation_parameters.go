// Code generated by go-swagger; DO NOT EDIT.

package dhcp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"github.com/dalet-oss/iris-api/models"
)

// NewCreateDHCPSubnetReservationParams creates a new CreateDHCPSubnetReservationParams object
//
// There are no default values defined in the spec.
func NewCreateDHCPSubnetReservationParams() CreateDHCPSubnetReservationParams {

	return CreateDHCPSubnetReservationParams{}
}

// CreateDHCPSubnetReservationParams contains all the bound params for the create d h c p subnet reservation operation
// typically these are obtained from a http.Request
//
// swagger:parameters CreateDHCPSubnetReservation
type CreateDHCPSubnetReservationParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	Body *models.Reservation
	/*The ID of the DHCP subnet to create reservations on.
	  Required: true
	  In: path
	*/
	SubnetID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateDHCPSubnetReservationParams() beforehand.
func (o *CreateDHCPSubnetReservationParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Reservation
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body", ""))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	} else {
		res = append(res, errors.Required("body", "body", ""))
	}

	rSubnetID, rhkSubnetID, _ := route.Params.GetOK("subnetId")
	if err := o.bindSubnetID(rSubnetID, rhkSubnetID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindSubnetID binds and validates parameter SubnetID from path.
func (o *CreateDHCPSubnetReservationParams) bindSubnetID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.SubnetID = raw

	return nil
}
