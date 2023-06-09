// Code generated by go-swagger; DO NOT EDIT.

package author

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewSoftDeleteAuthorParams creates a new SoftDeleteAuthorParams object
//
// There are no default values defined in the spec.
func NewSoftDeleteAuthorParams() SoftDeleteAuthorParams {

	return SoftDeleteAuthorParams{}
}

// SoftDeleteAuthorParams contains all the bound params for the soft delete author operation
// typically these are obtained from a http.Request
//
// swagger:parameters softDeleteAuthor
type SoftDeleteAuthorParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	AuthorID uint64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewSoftDeleteAuthorParams() beforehand.
func (o *SoftDeleteAuthorParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rAuthorID, rhkAuthorID, _ := route.Params.GetOK("author_id")
	if err := o.bindAuthorID(rAuthorID, rhkAuthorID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAuthorID binds and validates parameter AuthorID from path.
func (o *SoftDeleteAuthorParams) bindAuthorID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertUint64(raw)
	if err != nil {
		return errors.InvalidType("author_id", "path", "uint64", raw)
	}
	o.AuthorID = value

	return nil
}
