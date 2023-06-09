// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SuccessRegisterAllOf1 success register all of1
//
// swagger:model successRegisterAllOf1
type SuccessRegisterAllOf1 struct {

	// user id
	UserID uint64 `json:"user_id,omitempty"`
}

// Validate validates this success register all of1
func (m *SuccessRegisterAllOf1) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this success register all of1 based on context it is used
func (m *SuccessRegisterAllOf1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SuccessRegisterAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SuccessRegisterAllOf1) UnmarshalBinary(b []byte) error {
	var res SuccessRegisterAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
