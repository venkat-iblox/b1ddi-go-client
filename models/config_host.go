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

// ConfigHost Host
//
// A DNS Host (_dns/host_) object associates DNS configuraton with hosts.
//
// Automatically created and destroyed based on the hosts known to the platform.
//
// swagger:model configHost
type ConfigHost struct {

	// Host FQDN.
	// Read Only: true
	AbsoluteName string `json:"absolute_name,omitempty"`

	// Host's primary IP Address.
	// Read Only: true
	Address string `json:"address,omitempty"`

	// Anycast address configured to the host. Order is not significant.
	// Read Only: true
	AnycastAddresses []string `json:"anycast_addresses,omitempty"`

	// Host associated server configuration.
	AssociatedServer *ConfigHostAssociatedServer `json:"associated_server,omitempty"`

	// Host description.
	// Read Only: true
	Comment string `json:"comment,omitempty"`

	// Host current version.
	// Read Only: true
	CurrentVersion string `json:"current_version,omitempty"`

	// Below _dfp_ field is deprecated and not supported anymore.
	// The indication whether or not BloxOne DDI DNS and BloxOne TD DFP are both active on the host will be migrated into the new _dfp_service_ field.
	// Read Only: true
	Dfp bool `json:"dfp,omitempty"`

	// DFP service indicates whether or not BloxOne DDI DNS and BloxOne TD DFP are both active on the host.
	// If so, BloxOne DDI DNS will augment recursive queries and forward them to BloxOne TD DFP.
	// Allowed values:
	//  * _unavailable_: BloxOne TD DFP application is not available,
	//  * _enabled_: BloxOne TD DFP application is available and enabled,
	//  * _disabled_: BloxOne TD DFP application is available but disabled.
	// Read Only: true
	DfpService string `json:"dfp_service,omitempty"`

	// The resource identifier.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Optional. Inheritance configuration.
	InheritanceSources *ConfigHostInheritance `json:"inheritance_sources,omitempty"`

	// Optional. _kerberos_keys_ contains a list of keys for GSS-TSIG signed dynamic updates.
	//
	// Defaults to empty.
	KerberosKeys []*ConfigKerberosKey `json:"kerberos_keys,omitempty"`

	// Host display name.
	// Read Only: true
	Name string `json:"name,omitempty"`

	// On-Prem Host ID.
	// Read Only: true
	Ophid string `json:"ophid,omitempty"`

	// Host FQDN in punycode.
	// Read Only: true
	ProtocolAbsoluteName string `json:"protocol_absolute_name,omitempty"`

	// External provider identifier.
	// Read Only: true
	ProviderID string `json:"provider_id,omitempty"`

	// The resource identifier.
	Server string `json:"server,omitempty"`

	// Host site ID.
	// Read Only: true
	SiteID string `json:"site_id,omitempty"`

	// Host tagging specifics.
	Tags interface{} `json:"tags,omitempty"`

	// Defines the type of host.
	// Allowed values:
	//  * _bloxone_ddi_: host type is BloxOne DDI,
	//  * _microsoft_azure_: host type is Microsoft Azure,
	//  * _amazon_web_service_: host type is Amazon Web Services,
	//  * _microsoft_active_directory_: host type is Microsoft Active Directory,
	//  * _google_cloud_platform_: host type is Google Cloud Platform.
	// Read Only: true
	Type string `json:"type,omitempty"`
}

// Validate validates this config host
func (m *ConfigHost) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssociatedServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInheritanceSources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKerberosKeys(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigHost) validateAssociatedServer(formats strfmt.Registry) error {
	if swag.IsZero(m.AssociatedServer) { // not required
		return nil
	}

	if m.AssociatedServer != nil {
		if err := m.AssociatedServer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("associated_server")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("associated_server")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigHost) validateInheritanceSources(formats strfmt.Registry) error {
	if swag.IsZero(m.InheritanceSources) { // not required
		return nil
	}

	if m.InheritanceSources != nil {
		if err := m.InheritanceSources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inheritance_sources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("inheritance_sources")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigHost) validateKerberosKeys(formats strfmt.Registry) error {
	if swag.IsZero(m.KerberosKeys) { // not required
		return nil
	}

	for i := 0; i < len(m.KerberosKeys); i++ {
		if swag.IsZero(m.KerberosKeys[i]) { // not required
			continue
		}

		if m.KerberosKeys[i] != nil {
			if err := m.KerberosKeys[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("kerberos_keys" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("kerberos_keys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this config host based on the context it is used
func (m *ConfigHost) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAbsoluteName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAnycastAddresses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAssociatedServer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateComment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCurrentVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDfp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDfpService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInheritanceSources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKerberosKeys(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOphid(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtocolAbsoluteName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProviderID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSiteID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigHost) contextValidateAbsoluteName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "absolute_name", "body", string(m.AbsoluteName)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigHost) contextValidateAddress(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "address", "body", string(m.Address)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigHost) contextValidateAnycastAddresses(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "anycast_addresses", "body", []string(m.AnycastAddresses)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigHost) contextValidateAssociatedServer(ctx context.Context, formats strfmt.Registry) error {

	if m.AssociatedServer != nil {

		if swag.IsZero(m.AssociatedServer) { // not required
			return nil
		}

		if err := m.AssociatedServer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("associated_server")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("associated_server")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigHost) contextValidateComment(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "comment", "body", string(m.Comment)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigHost) contextValidateCurrentVersion(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "current_version", "body", string(m.CurrentVersion)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigHost) contextValidateDfp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dfp", "body", bool(m.Dfp)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigHost) contextValidateDfpService(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dfp_service", "body", string(m.DfpService)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigHost) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigHost) contextValidateInheritanceSources(ctx context.Context, formats strfmt.Registry) error {

	if m.InheritanceSources != nil {

		if swag.IsZero(m.InheritanceSources) { // not required
			return nil
		}

		if err := m.InheritanceSources.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inheritance_sources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("inheritance_sources")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigHost) contextValidateKerberosKeys(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.KerberosKeys); i++ {

		if m.KerberosKeys[i] != nil {

			if swag.IsZero(m.KerberosKeys[i]) { // not required
				return nil
			}

			if err := m.KerberosKeys[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("kerberos_keys" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("kerberos_keys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigHost) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigHost) contextValidateOphid(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "ophid", "body", string(m.Ophid)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigHost) contextValidateProtocolAbsoluteName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "protocol_absolute_name", "body", string(m.ProtocolAbsoluteName)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigHost) contextValidateProviderID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "provider_id", "body", string(m.ProviderID)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigHost) contextValidateSiteID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "site_id", "body", string(m.SiteID)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigHost) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "type", "body", string(m.Type)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigHost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigHost) UnmarshalBinary(b []byte) error {
	var res ConfigHost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
