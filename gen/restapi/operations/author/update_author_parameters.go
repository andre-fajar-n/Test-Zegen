// Code generated by go-swagger; DO NOT EDIT.

package author

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"zegen/gen/models"
)

// NewUpdateAuthorParams creates a new UpdateAuthorParams object
//
// There are no default values defined in the spec.
func NewUpdateAuthorParams() UpdateAuthorParams {

	return UpdateAuthorParams{}
}

// UpdateAuthorParams contains all the bound params for the update author operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateAuthor
type UpdateAuthorParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	AuthorID uint64
	/*
	  Required: true
	  In: body
	*/
	Data *models.UpdateAuthorParamsBody
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateAuthorParams() beforehand.
func (o *UpdateAuthorParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rAuthorID, rhkAuthorID, _ := route.Params.GetOK("author_id")
	if err := o.bindAuthorID(rAuthorID, rhkAuthorID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.UpdateAuthorParamsBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("data", "body", ""))
			} else {
				res = append(res, errors.NewParseError("data", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Data = &body
			}
		}
	} else {
		res = append(res, errors.Required("data", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAuthorID binds and validates parameter AuthorID from path.
func (o *UpdateAuthorParams) bindAuthorID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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