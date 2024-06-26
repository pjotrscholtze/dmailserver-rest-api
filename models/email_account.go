// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EmailAccount email account
//
// swagger:model EmailAccount
type EmailAccount struct {

	// address
	// Example: example@example.com
	Address string `json:"address,omitempty"`

	// password
	Password string `json:"password,omitempty"`
}

// Validate validates this email account
func (m *EmailAccount) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this email account based on context it is used
func (m *EmailAccount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EmailAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmailAccount) UnmarshalBinary(b []byte) error {
	var res EmailAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
