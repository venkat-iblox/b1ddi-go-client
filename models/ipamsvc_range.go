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

// IpamsvcRange Range
//
// A __Range__ object (_ipam/range_) represents a set of contiguous IP addresses in the same IP space with no gap, expressed as a (start, end) pair within a given subnet that are grouped together for administrative purpose and protocol management. The start and end values are not required to align with CIDR boundaries.
//
// swagger:model ipamsvcRange
type IpamsvcRange struct {

	// The description for the range. May contain 0 to 1024 characters. Can include UTF-8.
	Comment string `json:"comment,omitempty"`

	// Time when the object has been created.
	// Read Only: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// The resource identifier.
	DhcpHost string `json:"dhcp_host,omitempty"`

	// The list of DHCP options. May be either a specific option or a group of options.
	DhcpOptions []*IpamsvcOptionItem `json:"dhcp_options,omitempty"`

	// Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.
	//
	// Defaults to _false_.
	DisableDhcp bool `json:"disable_dhcp,omitempty"`

	// The end IP address of the range.
	// Required: true
	End *string `json:"end"`

	// The list of all exclusion ranges in the scope of the range.
	ExclusionRanges []*IpamsvcExclusionRange `json:"exclusion_ranges,omitempty"`

	// The list of all allow/deny filters of the range.
	Filters []*IpamsvcAccessFilter `json:"filters"`

	// The resource identifier.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The list of the inheritance assigned hosts of the object.
	// Read Only: true
	InheritanceAssignedHosts []*InheritanceAssignedHost `json:"inheritance_assigned_hosts,omitempty"`

	// The resource identifier.
	InheritanceParent string `json:"inheritance_parent,omitempty"`

	// The DHCP inheritance configuration for the range.
	InheritanceSources *IpamsvcDHCPOptionsInheritance `json:"inheritance_sources,omitempty"`

	// The name of the range. May contain 1 to 256 characters. Can include UTF-8.
	Name string `json:"name,omitempty"`

	// The resource identifier.
	Parent string `json:"parent,omitempty"`

	// The type of protocol (_ip4_ or _ip6_).
	// Read Only: true
	Protocol string `json:"protocol,omitempty"`

	// The resource identifier.
	// Required: true
	Space *string `json:"space"`

	// The start IP address of the range.
	// Required: true
	Start *string `json:"start"`

	// The tags for the range in JSON format.
	Tags interface{} `json:"tags,omitempty"`

	// The utilization threshold settings for the range.
	Threshold *IpamsvcUtilizationThreshold `json:"threshold,omitempty"`

	// Time when the object has been updated. Equals to _created_at_ if not updated after creation.
	// Read Only: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updated_at,omitempty"`

	// The utilization statistics of IPV4 addresses for the range.
	// Read Only: true
	Utilization *IpamsvcUtilization `json:"utilization,omitempty"`

	// The utilization of IPV6 addresses in the range.
	// Read Only: true
	UtilizationV6 *IpamsvcUtilizationV6 `json:"utilization_v6,omitempty"`
}

// Validate validates this ipamsvc range
func (m *IpamsvcRange) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDhcpOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExclusionRanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInheritanceAssignedHosts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInheritanceSources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThreshold(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUtilization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUtilizationV6(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamsvcRange) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcRange) validateDhcpOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.DhcpOptions) { // not required
		return nil
	}

	for i := 0; i < len(m.DhcpOptions); i++ {
		if swag.IsZero(m.DhcpOptions[i]) { // not required
			continue
		}

		if m.DhcpOptions[i] != nil {
			if err := m.DhcpOptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dhcp_options" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dhcp_options" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IpamsvcRange) validateEnd(formats strfmt.Registry) error {

	if err := validate.Required("end", "body", m.End); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcRange) validateExclusionRanges(formats strfmt.Registry) error {
	if swag.IsZero(m.ExclusionRanges) { // not required
		return nil
	}

	for i := 0; i < len(m.ExclusionRanges); i++ {
		if swag.IsZero(m.ExclusionRanges[i]) { // not required
			continue
		}

		if m.ExclusionRanges[i] != nil {
			if err := m.ExclusionRanges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("exclusion_ranges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("exclusion_ranges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IpamsvcRange) validateFilters(formats strfmt.Registry) error {
	if swag.IsZero(m.Filters) { // not required
		return nil
	}

	for i := 0; i < len(m.Filters); i++ {
		if swag.IsZero(m.Filters[i]) { // not required
			continue
		}

		if m.Filters[i] != nil {
			if err := m.Filters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IpamsvcRange) validateInheritanceAssignedHosts(formats strfmt.Registry) error {
	if swag.IsZero(m.InheritanceAssignedHosts) { // not required
		return nil
	}

	for i := 0; i < len(m.InheritanceAssignedHosts); i++ {
		if swag.IsZero(m.InheritanceAssignedHosts[i]) { // not required
			continue
		}

		if m.InheritanceAssignedHosts[i] != nil {
			if err := m.InheritanceAssignedHosts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inheritance_assigned_hosts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inheritance_assigned_hosts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IpamsvcRange) validateInheritanceSources(formats strfmt.Registry) error {
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

func (m *IpamsvcRange) validateSpace(formats strfmt.Registry) error {

	if err := validate.Required("space", "body", m.Space); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcRange) validateStart(formats strfmt.Registry) error {

	if err := validate.Required("start", "body", m.Start); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcRange) validateThreshold(formats strfmt.Registry) error {
	if swag.IsZero(m.Threshold) { // not required
		return nil
	}

	if m.Threshold != nil {
		if err := m.Threshold.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("threshold")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("threshold")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcRange) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcRange) validateUtilization(formats strfmt.Registry) error {
	if swag.IsZero(m.Utilization) { // not required
		return nil
	}

	if m.Utilization != nil {
		if err := m.Utilization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("utilization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("utilization")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcRange) validateUtilizationV6(formats strfmt.Registry) error {
	if swag.IsZero(m.UtilizationV6) { // not required
		return nil
	}

	if m.UtilizationV6 != nil {
		if err := m.UtilizationV6.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("utilization_v6")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("utilization_v6")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ipamsvc range based on the context it is used
func (m *IpamsvcRange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDhcpOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExclusionRanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInheritanceAssignedHosts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInheritanceSources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtocol(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThreshold(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUtilization(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUtilizationV6(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamsvcRange) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcRange) contextValidateDhcpOptions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DhcpOptions); i++ {

		if m.DhcpOptions[i] != nil {

			if swag.IsZero(m.DhcpOptions[i]) { // not required
				return nil
			}

			if err := m.DhcpOptions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dhcp_options" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dhcp_options" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IpamsvcRange) contextValidateExclusionRanges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExclusionRanges); i++ {

		if m.ExclusionRanges[i] != nil {

			if swag.IsZero(m.ExclusionRanges[i]) { // not required
				return nil
			}

			if err := m.ExclusionRanges[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("exclusion_ranges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("exclusion_ranges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IpamsvcRange) contextValidateFilters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Filters); i++ {

		if m.Filters[i] != nil {

			if swag.IsZero(m.Filters[i]) { // not required
				return nil
			}

			if err := m.Filters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IpamsvcRange) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcRange) contextValidateInheritanceAssignedHosts(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "inheritance_assigned_hosts", "body", []*InheritanceAssignedHost(m.InheritanceAssignedHosts)); err != nil {
		return err
	}

	for i := 0; i < len(m.InheritanceAssignedHosts); i++ {

		if m.InheritanceAssignedHosts[i] != nil {

			if swag.IsZero(m.InheritanceAssignedHosts[i]) { // not required
				return nil
			}

			if err := m.InheritanceAssignedHosts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inheritance_assigned_hosts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inheritance_assigned_hosts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IpamsvcRange) contextValidateInheritanceSources(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IpamsvcRange) contextValidateProtocol(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "protocol", "body", string(m.Protocol)); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcRange) contextValidateThreshold(ctx context.Context, formats strfmt.Registry) error {

	if m.Threshold != nil {

		if swag.IsZero(m.Threshold) { // not required
			return nil
		}

		if err := m.Threshold.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("threshold")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("threshold")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcRange) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcRange) contextValidateUtilization(ctx context.Context, formats strfmt.Registry) error {

	if m.Utilization != nil {

		if swag.IsZero(m.Utilization) { // not required
			return nil
		}

		if err := m.Utilization.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("utilization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("utilization")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcRange) contextValidateUtilizationV6(ctx context.Context, formats strfmt.Registry) error {

	if m.UtilizationV6 != nil {

		if swag.IsZero(m.UtilizationV6) { // not required
			return nil
		}

		if err := m.UtilizationV6.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("utilization_v6")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("utilization_v6")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IpamsvcRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpamsvcRange) UnmarshalBinary(b []byte) error {
	var res IpamsvcRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
