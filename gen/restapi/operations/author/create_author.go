// Code generated by go-swagger; DO NOT EDIT.

package author

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"zegen/gen/models"
)

// CreateAuthorHandlerFunc turns a function with the right signature into a create author handler
type CreateAuthorHandlerFunc func(CreateAuthorParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateAuthorHandlerFunc) Handle(params CreateAuthorParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// CreateAuthorHandler interface for that can handle valid create author params
type CreateAuthorHandler interface {
	Handle(CreateAuthorParams, *models.Principal) middleware.Responder
}

// NewCreateAuthor creates a new http.Handler for the create author operation
func NewCreateAuthor(ctx *middleware.Context, handler CreateAuthorHandler) *CreateAuthor {
	return &CreateAuthor{Context: ctx, Handler: handler}
}

/*
	CreateAuthor swagger:route POST /v1/author author createAuthor

# Create

Create author
*/
type CreateAuthor struct {
	Context *middleware.Context
	Handler CreateAuthorHandler
}

func (o *CreateAuthor) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateAuthorParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}