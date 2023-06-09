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

// BookAuthor book author
//
// swagger:model bookAuthor
type BookAuthor struct {
	ModelIdentifier

	ModelTrackTime

	BookAuthorData

	BookAuthorForeignKey
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *BookAuthor) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ModelIdentifier
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ModelIdentifier = aO0

	// AO1
	var aO1 ModelTrackTime
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.ModelTrackTime = aO1

	// AO2
	var aO2 BookAuthorData
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.BookAuthorData = aO2

	// AO3
	var aO3 BookAuthorForeignKey
	if err := swag.ReadJSON(raw, &aO3); err != nil {
		return err
	}
	m.BookAuthorForeignKey = aO3

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m BookAuthor) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 4)

	aO0, err := swag.WriteJSON(m.ModelIdentifier)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.ModelTrackTime)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.BookAuthorData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	aO3, err := swag.WriteJSON(m.BookAuthorForeignKey)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO3)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this book author
func (m *BookAuthor) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ModelIdentifier
	if err := m.ModelIdentifier.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with ModelTrackTime
	if err := m.ModelTrackTime.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with BookAuthorData
	if err := m.BookAuthorData.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with BookAuthorForeignKey
	if err := m.BookAuthorForeignKey.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this book author based on the context it is used
func (m *BookAuthor) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ModelIdentifier
	if err := m.ModelIdentifier.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with ModelTrackTime
	if err := m.ModelTrackTime.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with BookAuthorData
	if err := m.BookAuthorData.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with BookAuthorForeignKey
	if err := m.BookAuthorForeignKey.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *BookAuthor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BookAuthor) UnmarshalBinary(b []byte) error {
	var res BookAuthor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
