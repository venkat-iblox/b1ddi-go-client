// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConfigTTLInheritance TTLInheritance
//
// The inheritance configuration specifies how the object inherits the _ttl_ field.
//
// swagger:model configTTLInheritance
type ConfigTTLInheritance struct {

	// The field config for the _ttl_ field
	TTL *Inheritance2InheritedUInt32 `json:"ttl,omitempty"`
}

// Validate validates this config TTL inheritance
func (m *ConfigTTLInheritance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTTL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigTTLInheritance) validateTTL(formats strfmt.Registry) error {
	if swag.IsZero(m.TTL) { // not required
		return nil
	}

	if m.TTL != nil {
		if err := m.TTL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ttl")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ttl")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config TTL inheritance based on the context it is used
func (m *ConfigTTLInheritance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTTL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigTTLInheritance) contextValidateTTL(ctx context.Context, formats strfmt.Registry) error {

	if m.TTL != nil {

		if swag.IsZero(m.TTL) { // not required
			return nil
		}

		if err := m.TTL.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ttl")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ttl")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigTTLInheritance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigTTLInheritance) UnmarshalBinary(b []byte) error {
	var res ConfigTTLInheritance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
