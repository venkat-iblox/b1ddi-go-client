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

// ConfigInheritedZoneAuthority InheritedZoneAuthority
//
// Inheritance configuration for a field of type _ZoneAuthority_.
//
// swagger:model configInheritedZoneAuthority
type ConfigInheritedZoneAuthority struct {

	// Optional. Field config for _default_ttl_ field from _ZoneAuthority_ object.
	DefaultTTL *Inheritance2InheritedUInt32 `json:"default_ttl,omitempty"`

	// Optional. Field config for _expire_ field from _ZoneAuthority_ object.
	Expire *Inheritance2InheritedUInt32 `json:"expire,omitempty"`

	// Optional. Field config for _mname_ block from _ZoneAuthority_ object.
	MnameBlock *ConfigInheritedZoneAuthorityMNameBlock `json:"mname_block,omitempty"`

	// Optional. Field config for _negative_ttl_ field from _ZoneAuthority_ object.
	NegativeTTL *Inheritance2InheritedUInt32 `json:"negative_ttl,omitempty"`

	// Optional. Field config for _protocol_rname_ field from _ZoneAuthority_ object.
	ProtocolRname *Inheritance2InheritedString `json:"protocol_rname,omitempty"`

	// Optional. Field config for _refresh_ field from _ZoneAuthority_ object.
	Refresh *Inheritance2InheritedUInt32 `json:"refresh,omitempty"`

	// Optional. Field config for _retry_ field from _ZoneAuthority_ object.
	Retry *Inheritance2InheritedUInt32 `json:"retry,omitempty"`

	// Optional. Field config for _rname_ field from _ZoneAuthority_ object.
	Rname *Inheritance2InheritedString `json:"rname,omitempty"`
}

// Validate validates this config inherited zone authority
func (m *ConfigInheritedZoneAuthority) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultTTL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpire(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMnameBlock(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNegativeTTL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocolRname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefresh(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRetry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRname(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigInheritedZoneAuthority) validateDefaultTTL(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultTTL) { // not required
		return nil
	}

	if m.DefaultTTL != nil {
		if err := m.DefaultTTL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_ttl")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("default_ttl")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigInheritedZoneAuthority) validateExpire(formats strfmt.Registry) error {
	if swag.IsZero(m.Expire) { // not required
		return nil
	}

	if m.Expire != nil {
		if err := m.Expire.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("expire")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("expire")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigInheritedZoneAuthority) validateMnameBlock(formats strfmt.Registry) error {
	if swag.IsZero(m.MnameBlock) { // not required
		return nil
	}

	if m.MnameBlock != nil {
		if err := m.MnameBlock.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mname_block")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mname_block")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigInheritedZoneAuthority) validateNegativeTTL(formats strfmt.Registry) error {
	if swag.IsZero(m.NegativeTTL) { // not required
		return nil
	}

	if m.NegativeTTL != nil {
		if err := m.NegativeTTL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("negative_ttl")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("negative_ttl")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigInheritedZoneAuthority) validateProtocolRname(formats strfmt.Registry) error {
	if swag.IsZero(m.ProtocolRname) { // not required
		return nil
	}

	if m.ProtocolRname != nil {
		if err := m.ProtocolRname.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protocol_rname")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protocol_rname")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigInheritedZoneAuthority) validateRefresh(formats strfmt.Registry) error {
	if swag.IsZero(m.Refresh) { // not required
		return nil
	}

	if m.Refresh != nil {
		if err := m.Refresh.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("refresh")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("refresh")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigInheritedZoneAuthority) validateRetry(formats strfmt.Registry) error {
	if swag.IsZero(m.Retry) { // not required
		return nil
	}

	if m.Retry != nil {
		if err := m.Retry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("retry")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("retry")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigInheritedZoneAuthority) validateRname(formats strfmt.Registry) error {
	if swag.IsZero(m.Rname) { // not required
		return nil
	}

	if m.Rname != nil {
		if err := m.Rname.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rname")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rname")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config inherited zone authority based on the context it is used
func (m *ConfigInheritedZoneAuthority) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDefaultTTL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExpire(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMnameBlock(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNegativeTTL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtocolRname(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRefresh(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRetry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRname(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigInheritedZoneAuthority) contextValidateDefaultTTL(ctx context.Context, formats strfmt.Registry) error {

	if m.DefaultTTL != nil {

		if swag.IsZero(m.DefaultTTL) { // not required
			return nil
		}

		if err := m.DefaultTTL.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_ttl")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("default_ttl")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigInheritedZoneAuthority) contextValidateExpire(ctx context.Context, formats strfmt.Registry) error {

	if m.Expire != nil {

		if swag.IsZero(m.Expire) { // not required
			return nil
		}

		if err := m.Expire.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("expire")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("expire")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigInheritedZoneAuthority) contextValidateMnameBlock(ctx context.Context, formats strfmt.Registry) error {

	if m.MnameBlock != nil {

		if swag.IsZero(m.MnameBlock) { // not required
			return nil
		}

		if err := m.MnameBlock.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mname_block")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mname_block")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigInheritedZoneAuthority) contextValidateNegativeTTL(ctx context.Context, formats strfmt.Registry) error {

	if m.NegativeTTL != nil {

		if swag.IsZero(m.NegativeTTL) { // not required
			return nil
		}

		if err := m.NegativeTTL.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("negative_ttl")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("negative_ttl")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigInheritedZoneAuthority) contextValidateProtocolRname(ctx context.Context, formats strfmt.Registry) error {

	if m.ProtocolRname != nil {

		if swag.IsZero(m.ProtocolRname) { // not required
			return nil
		}

		if err := m.ProtocolRname.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protocol_rname")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protocol_rname")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigInheritedZoneAuthority) contextValidateRefresh(ctx context.Context, formats strfmt.Registry) error {

	if m.Refresh != nil {

		if swag.IsZero(m.Refresh) { // not required
			return nil
		}

		if err := m.Refresh.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("refresh")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("refresh")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigInheritedZoneAuthority) contextValidateRetry(ctx context.Context, formats strfmt.Registry) error {

	if m.Retry != nil {

		if swag.IsZero(m.Retry) { // not required
			return nil
		}

		if err := m.Retry.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("retry")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("retry")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigInheritedZoneAuthority) contextValidateRname(ctx context.Context, formats strfmt.Registry) error {

	if m.Rname != nil {

		if swag.IsZero(m.Rname) { // not required
			return nil
		}

		if err := m.Rname.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rname")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rname")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigInheritedZoneAuthority) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigInheritedZoneAuthority) UnmarshalBinary(b []byte) error {
	var res ConfigInheritedZoneAuthority
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
