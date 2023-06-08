// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateBookParamsBody update book params body
//
// swagger:model updateBookParamsBody
type UpdateBookParamsBody struct {

	// author ids
	// Required: true
	AuthorIds []float64 `json:"author_ids"`

	// isbn
	// Required: true
	// Min Length: 1
	Isbn *string `json:"isbn"`

	// published year
	// Required: true
	PublishedYear *int64 `json:"published_year"`

	// title
	// Required: true
	// Min Length: 1
	Title *string `json:"title"`
}

// Validate validates this update book params body
func (m *UpdateBookParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsbn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublishedYear(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateBookParamsBody) validateAuthorIds(formats strfmt.Registry) error {

	if err := validate.Required("author_ids", "body", m.AuthorIds); err != nil {
		return err
	}

	return nil
}

func (m *UpdateBookParamsBody) validateIsbn(formats strfmt.Registry) error {

	if err := validate.Required("isbn", "body", m.Isbn); err != nil {
		return err
	}

	if err := validate.MinLength("isbn", "body", *m.Isbn, 1); err != nil {
		return err
	}

	return nil
}

func (m *UpdateBookParamsBody) validatePublishedYear(formats strfmt.Registry) error {

	if err := validate.Required("published_year", "body", m.PublishedYear); err != nil {
		return err
	}

	return nil
}

func (m *UpdateBookParamsBody) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	if err := validate.MinLength("title", "body", *m.Title, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update book params body based on context it is used
func (m *UpdateBookParamsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateBookParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateBookParamsBody) UnmarshalBinary(b []byte) error {
	var res UpdateBookParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}