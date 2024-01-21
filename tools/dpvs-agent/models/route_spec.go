// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RouteSpec route spec
//
// swagger:model RouteSpec
type RouteSpec struct {

	// device
	Device string `json:"device,omitempty"`

	// dst
	Dst string `json:"dst,omitempty"`

	// gateway
	Gateway string `json:"gateway,omitempty"`

	// metric
	Metric uint32 `json:"metric,omitempty"`

	// mtu
	Mtu uint32 `json:"mtu,omitempty"`

	// prefix src
	PrefixSrc string `json:"prefixSrc,omitempty"`

	// scope
	Scope string `json:"scope,omitempty"`

	// src
	Src string `json:"src,omitempty"`
}

// Validate validates this route spec
func (m *RouteSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this route spec based on context it is used
func (m *RouteSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RouteSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RouteSpec) UnmarshalBinary(b []byte) error {
	var res RouteSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}