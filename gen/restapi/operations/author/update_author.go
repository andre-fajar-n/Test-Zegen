// Code generated by go-swagger; DO NOT EDIT.

package author

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"zegen/gen/models"
)

// UpdateAuthorHandlerFunc turns a function with the right signature into a update author handler
type UpdateAuthorHandlerFunc func(UpdateAuthorParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateAuthorHandlerFunc) Handle(params UpdateAuthorParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// UpdateAuthorHandler interface for that can handle valid update author params
type UpdateAuthorHandler interface {
	Handle(UpdateAuthorParams, *models.Principal) middleware.Responder
}

// NewUpdateAuthor creates a new http.Handler for the update author operation
func NewUpdateAuthor(ctx *middleware.Context, handler UpdateAuthorHandler) *UpdateAuthor {
	return &UpdateAuthor{Context: ctx, Handler: handler}
}

/*
	UpdateAuthor swagger:route PUT /v1/author/{author_id} author updateAuthor

# Update

Update author
*/
type UpdateAuthor struct {
	Context *middleware.Context
	Handler UpdateAuthorHandler
}

func (o *UpdateAuthor) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateAuthorParams()
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
