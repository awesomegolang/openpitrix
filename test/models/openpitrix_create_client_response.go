// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixCreateClientResponse openpitrix create client response
// swagger:model openpitrixCreateClientResponse
type OpenpitrixCreateClientResponse struct {

	// client id
	ClientID string `json:"client_id,omitempty"`

	// client secret
	ClientSecret string `json:"client_secret,omitempty"`

	// user id
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this openpitrix create client response
func (m *OpenpitrixCreateClientResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixCreateClientResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixCreateClientResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixCreateClientResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
