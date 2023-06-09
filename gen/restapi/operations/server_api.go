// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"zegen/gen/models"
	"zegen/gen/restapi/operations/authentication"
	"zegen/gen/restapi/operations/author"
	"zegen/gen/restapi/operations/book"
	"zegen/gen/restapi/operations/health"
)

// NewServerAPI creates a new Server instance
func NewServerAPI(spec *loads.Document) *ServerAPI {
	return &ServerAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer:          runtime.JSONConsumer(),
		MultipartformConsumer: runtime.DiscardConsumer,

		JSONProducer: runtime.JSONProducer(),

		AuthorCreateAuthorHandler: author.CreateAuthorHandlerFunc(func(params author.CreateAuthorParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation author.CreateAuthor has not yet been implemented")
		}),
		BookCreateBookHandler: book.CreateBookHandlerFunc(func(params book.CreateBookParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation book.CreateBook has not yet been implemented")
		}),
		AuthorFindOneAuthorHandler: author.FindOneAuthorHandlerFunc(func(params author.FindOneAuthorParams) middleware.Responder {
			return middleware.NotImplemented("operation author.FindOneAuthor has not yet been implemented")
		}),
		BookFindOneBookHandler: book.FindOneBookHandlerFunc(func(params book.FindOneBookParams) middleware.Responder {
			return middleware.NotImplemented("operation book.FindOneBook has not yet been implemented")
		}),
		HealthHealthHandler: health.HealthHandlerFunc(func(params health.HealthParams) middleware.Responder {
			return middleware.NotImplemented("operation health.Health has not yet been implemented")
		}),
		AuthenticationLoginHandler: authentication.LoginHandlerFunc(func(params authentication.LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation authentication.Login has not yet been implemented")
		}),
		AuthenticationRegisterHandler: authentication.RegisterHandlerFunc(func(params authentication.RegisterParams) middleware.Responder {
			return middleware.NotImplemented("operation authentication.Register has not yet been implemented")
		}),
		AuthorSoftDeleteAuthorHandler: author.SoftDeleteAuthorHandlerFunc(func(params author.SoftDeleteAuthorParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation author.SoftDeleteAuthor has not yet been implemented")
		}),
		BookSoftDeleteBookHandler: book.SoftDeleteBookHandlerFunc(func(params book.SoftDeleteBookParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation book.SoftDeleteBook has not yet been implemented")
		}),
		AuthorUpdateAuthorHandler: author.UpdateAuthorHandlerFunc(func(params author.UpdateAuthorParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation author.UpdateAuthor has not yet been implemented")
		}),
		BookUpdateBookHandler: book.UpdateBookHandlerFunc(func(params book.UpdateBookParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation book.UpdateBook has not yet been implemented")
		}),

		// Applies when the "Authorization" header is set
		AuthorizationAuth: func(token string) (*models.Principal, error) {
			return nil, errors.NotImplemented("api key auth (authorization) Authorization from header param [Authorization] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*ServerAPI the server API */
type ServerAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for the following mime types:
	//   - multipart/form-data
	MultipartformConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// AuthorizationAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	AuthorizationAuth func(string) (*models.Principal, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// AuthorCreateAuthorHandler sets the operation handler for the create author operation
	AuthorCreateAuthorHandler author.CreateAuthorHandler
	// BookCreateBookHandler sets the operation handler for the create book operation
	BookCreateBookHandler book.CreateBookHandler
	// AuthorFindOneAuthorHandler sets the operation handler for the find one author operation
	AuthorFindOneAuthorHandler author.FindOneAuthorHandler
	// BookFindOneBookHandler sets the operation handler for the find one book operation
	BookFindOneBookHandler book.FindOneBookHandler
	// HealthHealthHandler sets the operation handler for the health operation
	HealthHealthHandler health.HealthHandler
	// AuthenticationLoginHandler sets the operation handler for the login operation
	AuthenticationLoginHandler authentication.LoginHandler
	// AuthenticationRegisterHandler sets the operation handler for the register operation
	AuthenticationRegisterHandler authentication.RegisterHandler
	// AuthorSoftDeleteAuthorHandler sets the operation handler for the soft delete author operation
	AuthorSoftDeleteAuthorHandler author.SoftDeleteAuthorHandler
	// BookSoftDeleteBookHandler sets the operation handler for the soft delete book operation
	BookSoftDeleteBookHandler book.SoftDeleteBookHandler
	// AuthorUpdateAuthorHandler sets the operation handler for the update author operation
	AuthorUpdateAuthorHandler author.UpdateAuthorHandler
	// BookUpdateBookHandler sets the operation handler for the update book operation
	BookUpdateBookHandler book.UpdateBookHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *ServerAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *ServerAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *ServerAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *ServerAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *ServerAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *ServerAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *ServerAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *ServerAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *ServerAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the ServerAPI
func (o *ServerAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}
	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.AuthorizationAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.AuthorCreateAuthorHandler == nil {
		unregistered = append(unregistered, "author.CreateAuthorHandler")
	}
	if o.BookCreateBookHandler == nil {
		unregistered = append(unregistered, "book.CreateBookHandler")
	}
	if o.AuthorFindOneAuthorHandler == nil {
		unregistered = append(unregistered, "author.FindOneAuthorHandler")
	}
	if o.BookFindOneBookHandler == nil {
		unregistered = append(unregistered, "book.FindOneBookHandler")
	}
	if o.HealthHealthHandler == nil {
		unregistered = append(unregistered, "health.HealthHandler")
	}
	if o.AuthenticationLoginHandler == nil {
		unregistered = append(unregistered, "authentication.LoginHandler")
	}
	if o.AuthenticationRegisterHandler == nil {
		unregistered = append(unregistered, "authentication.RegisterHandler")
	}
	if o.AuthorSoftDeleteAuthorHandler == nil {
		unregistered = append(unregistered, "author.SoftDeleteAuthorHandler")
	}
	if o.BookSoftDeleteBookHandler == nil {
		unregistered = append(unregistered, "book.SoftDeleteBookHandler")
	}
	if o.AuthorUpdateAuthorHandler == nil {
		unregistered = append(unregistered, "author.UpdateAuthorHandler")
	}
	if o.BookUpdateBookHandler == nil {
		unregistered = append(unregistered, "book.UpdateBookHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *ServerAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *ServerAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "authorization":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.AuthorizationAuth(token)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *ServerAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *ServerAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *ServerAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *ServerAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the server API
func (o *ServerAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *ServerAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v1/author"] = author.NewCreateAuthor(o.context, o.AuthorCreateAuthorHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v1/book"] = book.NewCreateBook(o.context, o.BookCreateBookHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/author/{author_id}"] = author.NewFindOneAuthor(o.context, o.AuthorFindOneAuthorHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/book/{book_id}"] = book.NewFindOneBook(o.context, o.BookFindOneBookHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/health"] = health.NewHealth(o.context, o.HealthHealthHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v1/login"] = authentication.NewLogin(o.context, o.AuthenticationLoginHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v1/register"] = authentication.NewRegister(o.context, o.AuthenticationRegisterHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/v1/author/{author_id}"] = author.NewSoftDeleteAuthor(o.context, o.AuthorSoftDeleteAuthorHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/v1/book/{book_id}"] = book.NewSoftDeleteBook(o.context, o.BookSoftDeleteBookHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/v1/author/{author_id}"] = author.NewUpdateAuthor(o.context, o.AuthorUpdateAuthorHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/v1/book/{book_id}"] = book.NewUpdateBook(o.context, o.BookUpdateBookHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *ServerAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *ServerAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *ServerAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *ServerAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *ServerAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
