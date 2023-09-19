// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IpamsvcInteger128 ipamsvc integer128
//
// swagger:model ipamsvcInteger128
type IpamsvcInteger128 struct {

	// raw value
	// Format: byte
	RawValue strfmt.Base64 `json:"raw_value,omitempty"`
}

// Validate validates this ipamsvc integer128
func (m *IpamsvcInteger128) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ipamsvc integer128 based on context it is used
func (m *IpamsvcInteger128) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IpamsvcInteger128) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpamsvcInteger128) UnmarshalBinary(b []byte) error {
	var res IpamsvcInteger128
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
