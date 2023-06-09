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

// BookAuthorForeignKeyBook book author foreign key book
//
// swagger:model bookAuthorForeignKeyBook
type BookAuthorForeignKeyBook struct {
	ModelIdentifier

	ModelTrackTime

	BookData

	BookForeignKey
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *BookAuthorForeignKeyBook) UnmarshalJSON(raw []byte) error {
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
	var aO2 BookData
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.BookData = aO2

	// AO3
	var aO3 BookForeignKey
	if err := swag.ReadJSON(raw, &aO3); err != nil {
		return err
	}
	m.BookForeignKey = aO3

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m BookAuthorForeignKeyBook) MarshalJSON() ([]byte, error) {
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

	aO2, err := swag.WriteJSON(m.BookData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	aO3, err := swag.WriteJSON(m.BookForeignKey)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO3)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this book author foreign key book
func (m *BookAuthorForeignKeyBook) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ModelIdentifier
	if err := m.ModelIdentifier.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with ModelTrackTime
	if err := m.ModelTrackTime.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with BookData
	if err := m.BookData.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with BookForeignKey
	if err := m.BookForeignKey.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this book author foreign key book based on the context it is used
func (m *BookAuthorForeignKeyBook) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ModelIdentifier
	if err := m.ModelIdentifier.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with ModelTrackTime
	if err := m.ModelTrackTime.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with BookData
	if err := m.BookData.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with BookForeignKey
	if err := m.BookForeignKey.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *BookAuthorForeignKeyBook) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BookAuthorForeignKeyBook) UnmarshalBinary(b []byte) error {
	var res BookAuthorForeignKeyBook
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}