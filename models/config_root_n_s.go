// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConfigRootNS RootNS
//
// # Root nameserver
//
// swagger:model configRootNS
type ConfigRootNS struct {

	// IPv4 address.
	// Required: true
	Address *string `json:"address"`

	// FQDN.
	// Required: true
	Fqdn *string `json:"fqdn"`

	// FQDN in punycode.
	// Read Only: true
	ProtocolFqdn string `json:"protocol_fqdn,omitempty"`
}

// Validate validates this config root n s
func (m *ConfigRootNS) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFqdn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigRootNS) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *ConfigRootNS) validateFqdn(formats strfmt.Registry) error {

	if err := validate.Required("fqdn", "body", m.Fqdn); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this config root n s based on the context it is used
func (m *ConfigRootNS) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProtocolFqdn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigRootNS) contextValidateProtocolFqdn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "protocol_fqdn", "body", string(m.ProtocolFqdn)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigRootNS) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigRootNS) UnmarshalBinary(b []byte) error {
	var res ConfigRootNS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
