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

// IpamsvcUtilizationV6 UtilizationV6
//
// The __UtilizationV6__ object represents IPV6 address usage statistics for an object.
//
// swagger:model ipamsvcUtilizationV6
type IpamsvcUtilizationV6 struct {

	// The number of IPV6 addresses in the scope of the object which are in the abandoned state (issued by a DHCP server and then declined by the client).
	// Read Only: true
	Abandoned *IpamsvcInteger128 `json:"abandoned,omitempty"`

	// The number of IPV6 addresses handed out by DHCP in the scope of the object. This includes all leased addresses, fixed addresses that are defined but not currently leased and abandoned leases.
	// Read Only: true
	Dynamic *IpamsvcInteger128 `json:"dynamic,omitempty"`

	// The number of defined IPV6 addresses such as reservations or DNS records. It can be computed as _static_ = _used_ - _dynamic_.
	// Read Only: true
	Static *IpamsvcInteger128 `json:"static,omitempty"`

	// The total number of IPV6 addresses available in the scope of the object.
	// Read Only: true
	Total *IpamsvcInteger128 `json:"total,omitempty"`

	// The number of IPV6 addresses used in the scope of the object.
	// Read Only: true
	Used *IpamsvcInteger128 `json:"used,omitempty"`
}

// Validate validates this ipamsvc utilization v6
func (m *IpamsvcUtilizationV6) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbandoned(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDynamic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamsvcUtilizationV6) validateAbandoned(formats strfmt.Registry) error {
	if swag.IsZero(m.Abandoned) { // not required
		return nil
	}

	if m.Abandoned != nil {
		if err := m.Abandoned.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("abandoned")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("abandoned")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcUtilizationV6) validateDynamic(formats strfmt.Registry) error {
	if swag.IsZero(m.Dynamic) { // not required
		return nil
	}

	if m.Dynamic != nil {
		if err := m.Dynamic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dynamic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dynamic")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcUtilizationV6) validateStatic(formats strfmt.Registry) error {
	if swag.IsZero(m.Static) { // not required
		return nil
	}

	if m.Static != nil {
		if err := m.Static.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("static")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("static")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcUtilizationV6) validateTotal(formats strfmt.Registry) error {
	if swag.IsZero(m.Total) { // not required
		return nil
	}

	if m.Total != nil {
		if err := m.Total.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("total")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("total")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcUtilizationV6) validateUsed(formats strfmt.Registry) error {
	if swag.IsZero(m.Used) { // not required
		return nil
	}

	if m.Used != nil {
		if err := m.Used.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("used")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("used")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ipamsvc utilization v6 based on the context it is used
func (m *IpamsvcUtilizationV6) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAbandoned(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDynamic(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatic(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamsvcUtilizationV6) contextValidateAbandoned(ctx context.Context, formats strfmt.Registry) error {

	if m.Abandoned != nil {

		if swag.IsZero(m.Abandoned) { // not required
			return nil
		}

		if err := m.Abandoned.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("abandoned")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("abandoned")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcUtilizationV6) contextValidateDynamic(ctx context.Context, formats strfmt.Registry) error {

	if m.Dynamic != nil {

		if swag.IsZero(m.Dynamic) { // not required
			return nil
		}

		if err := m.Dynamic.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dynamic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dynamic")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcUtilizationV6) contextValidateStatic(ctx context.Context, formats strfmt.Registry) error {

	if m.Static != nil {

		if swag.IsZero(m.Static) { // not required
			return nil
		}

		if err := m.Static.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("static")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("static")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcUtilizationV6) contextValidateTotal(ctx context.Context, formats strfmt.Registry) error {

	if m.Total != nil {

		if swag.IsZero(m.Total) { // not required
			return nil
		}

		if err := m.Total.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("total")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("total")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcUtilizationV6) contextValidateUsed(ctx context.Context, formats strfmt.Registry) error {

	if m.Used != nil {

		if swag.IsZero(m.Used) { // not required
			return nil
		}

		if err := m.Used.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("used")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("used")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IpamsvcUtilizationV6) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpamsvcUtilizationV6) UnmarshalBinary(b []byte) error {
	var res IpamsvcUtilizationV6
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
