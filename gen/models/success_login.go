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

// SuccessLogin success login
//
// swagger:model successLogin
type SuccessLogin struct {
	Success

	SuccessLoginAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SuccessLogin) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Success
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Success = aO0

	// AO1
	var aO1 SuccessLoginAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.SuccessLoginAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SuccessLogin) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Success)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.SuccessLoginAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this success login
func (m *SuccessLogin) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Success
	if err := m.Success.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with SuccessLoginAllOf1
	if err := m.SuccessLoginAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this success login based on the context it is used
func (m *SuccessLogin) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Success
	if err := m.Success.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with SuccessLoginAllOf1
	if err := m.SuccessLoginAllOf1.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SuccessLogin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SuccessLogin) UnmarshalBinary(b []byte) error {
	var res SuccessLogin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
