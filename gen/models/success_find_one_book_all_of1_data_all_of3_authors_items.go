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

// SuccessFindOneBookAllOf1DataAllOf3AuthorsItems success find one book all of1 data all of3 authors items
//
// swagger:model successFindOneBookAllOf1DataAllOf3AuthorsItems
type SuccessFindOneBookAllOf1DataAllOf3AuthorsItems struct {
	ModelIdentifier

	AuthorData
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SuccessFindOneBookAllOf1DataAllOf3AuthorsItems) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ModelIdentifier
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ModelIdentifier = aO0

	// AO1
	var aO1 AuthorData
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.AuthorData = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SuccessFindOneBookAllOf1DataAllOf3AuthorsItems) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ModelIdentifier)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.AuthorData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this success find one book all of1 data all of3 authors items
func (m *SuccessFindOneBookAllOf1DataAllOf3AuthorsItems) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ModelIdentifier
	if err := m.ModelIdentifier.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with AuthorData
	if err := m.AuthorData.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this success find one book all of1 data all of3 authors items based on the context it is used
func (m *SuccessFindOneBookAllOf1DataAllOf3AuthorsItems) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ModelIdentifier
	if err := m.ModelIdentifier.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with AuthorData
	if err := m.AuthorData.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SuccessFindOneBookAllOf1DataAllOf3AuthorsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SuccessFindOneBookAllOf1DataAllOf3AuthorsItems) UnmarshalBinary(b []byte) error {
	var res SuccessFindOneBookAllOf1DataAllOf3AuthorsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
