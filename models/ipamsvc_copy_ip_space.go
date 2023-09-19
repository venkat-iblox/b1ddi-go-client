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

// IpamsvcCopyIPSpace ipamsvc copy IP space
//
// swagger:model ipamsvcCopyIPSpace
type IpamsvcCopyIPSpace struct {

	// The description for the copied IP space. May contain 0 to 1024 characters. Can include UTF-8.
	Comment string `json:"comment,omitempty"`

	// Indicates whether dhcp options should be copied or not when _ipam/ip_space_ object is copied.
	//
	// Defaults to _false_.
	CopyDhcpOptions bool `json:"copy_dhcp_options,omitempty"`

	// The resource identifier.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The name for the copied IP space. Must contain 1 to 256 characters. Can include UTF-8.
	// Required: true
	Name *string `json:"name"`

	// Indicates whether copying should skip an object in case of error and continue with next, or abort copying in case of error.
	//
	// Defaults to _false_.
	SkipOnError bool `json:"skip_on_error,omitempty"`
}

// Validate validates this ipamsvc copy IP space
func (m *IpamsvcCopyIPSpace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamsvcCopyIPSpace) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ipamsvc copy IP space based on the context it is used
func (m *IpamsvcCopyIPSpace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamsvcCopyIPSpace) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IpamsvcCopyIPSpace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpamsvcCopyIPSpace) UnmarshalBinary(b []byte) error {
	var res IpamsvcCopyIPSpace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
