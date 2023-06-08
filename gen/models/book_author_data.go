// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BookAuthorData book author data
//
// swagger:model bookAuthorData
type BookAuthorData struct {

	// author id
	AuthorID string `json:"author_id,omitempty"`

	// book id
	BookID string `json:"book_id,omitempty"`
}

// Validate validates this book author data
func (m *BookAuthorData) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this book author data based on context it is used
func (m *BookAuthorData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BookAuthorData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BookAuthorData) UnmarshalBinary(b []byte) error {
	var res BookAuthorData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}