// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InetAddrSpec inet addr spec
//
// swagger:model InetAddrSpec
type InetAddrSpec struct {

	// addr
	Addr string `json:"addr,omitempty"`

	// broadcast
	Broadcast string `json:"broadcast,omitempty"`

	// scope
	Scope string `json:"scope,omitempty"`
}

// Validate validates this inet addr spec
func (m *InetAddrSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this inet addr spec based on context it is used
func (m *InetAddrSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InetAddrSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InetAddrSpec) UnmarshalBinary(b []byte) error {
	var res InetAddrSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}