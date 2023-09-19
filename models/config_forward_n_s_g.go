// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConfigForwardNSG ForwardNSG
//
// Forward DNS Server Group for forward zones.
//
// swagger:model configForwardNSG
type ConfigForwardNSG struct {

	// Optional. Comment for the object.
	Comment string `json:"comment,omitempty"`

	// Optional. External DNS servers to forward to. Order is not significant.
	ExternalForwarders []*ConfigForwarder `json:"external_forwarders,omitempty"`

	// Optional. _true_ to only forward.
	ForwardersOnly bool `json:"forwarders_only,omitempty"`

	// The resource identifier.
	Hosts []string `json:"hosts,omitempty"`

	// The resource identifier.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The resource identifier.
	InternalForwarders []string `json:"internal_forwarders,omitempty"`

	// Name of the object.
	// Required: true
	Name *string `json:"name"`

	// The resource identifier.
	Nsgs []string `json:"nsgs,omitempty"`

	// Tagging specifics.
	Tags interface{} `json:"tags,omitempty"`
}

// Validate validates this config forward n s g
func (m *ConfigForwardNSG) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExternalForwarders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigForwardNSG) validateExternalForwarders(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalForwarders) { // not required
		return nil
	}

	for i := 0; i < len(m.ExternalForwarders); i++ {
		if swag.IsZero(m.ExternalForwarders[i]) { // not required
			continue
		}

		if m.ExternalForwarders[i] != nil {
			if err := m.ExternalForwarders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("external_forwarders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("external_forwarders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigForwardNSG) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this config forward n s g based on the context it is used
func (m *ConfigForwardNSG) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExternalForwarders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigForwardNSG) contextValidateExternalForwarders(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExternalForwarders); i++ {

		if m.ExternalForwarders[i] != nil {

			if swag.IsZero(m.ExternalForwarders[i]) { // not required
				return nil
			}

			if err := m.ExternalForwarders[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("external_forwarders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("external_forwarders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigForwardNSG) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigForwardNSG) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigForwardNSG) UnmarshalBinary(b []byte) error {
	var res ConfigForwardNSG
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
