// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Reservation reservation
// Example: {"hostname":"my-awesome-server","ip":"192.168.0.10","mac":"1a:1b:1c:1d:1e:1f"}
//
// swagger:model Reservation
type Reservation struct {

	// The host name
	Hostname string `json:"hostname,omitempty"`

	// The host IP address.
	IP string `json:"ip,omitempty"`

	// The host MAC address.
	Mac string `json:"mac,omitempty"`
}

// Validate validates this reservation
func (m *Reservation) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this reservation based on context it is used
func (m *Reservation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Reservation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Reservation) UnmarshalBinary(b []byte) error {
	var res Reservation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
