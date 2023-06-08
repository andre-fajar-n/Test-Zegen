// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"zegen/gen/models"
	"zegen/gen/restapi/operations"
	"zegen/gen/restapi/operations/authentication"
	"zegen/gen/restapi/operations/author"
	"zegen/gen/restapi/operations/health"
)

//go:generate swagger generate server --target ../../gen --name Server --spec ../../api/zegen/result.yml --principal models.Principal --exclude-main

func configureFlags(api *operations.ServerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ServerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()
	api.MultipartformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Authorization" header is set
	if api.AuthorizationAuth == nil {
		api.AuthorizationAuth = func(token string) (*models.Principal, error) {
			return nil, errors.NotImplemented("api key auth (authorization) Authorization from header param [Authorization] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.AuthorCreateAuthorHandler == nil {
		api.AuthorCreateAuthorHandler = author.CreateAuthorHandlerFunc(func(params author.CreateAuthorParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation author.CreateAuthor has not yet been implemented")
		})
	}
	if api.HealthHealthHandler == nil {
		api.HealthHealthHandler = health.HealthHandlerFunc(func(params health.HealthParams) middleware.Responder {
			return middleware.NotImplemented("operation health.Health has not yet been implemented")
		})
	}
	if api.AuthenticationLoginHandler == nil {
		api.AuthenticationLoginHandler = authentication.LoginHandlerFunc(func(params authentication.LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation authentication.Login has not yet been implemented")
		})
	}
	if api.AuthenticationRegisterHandler == nil {
		api.AuthenticationRegisterHandler = authentication.RegisterHandlerFunc(func(params authentication.RegisterParams) middleware.Responder {
			return middleware.NotImplemented("operation authentication.Register has not yet been implemented")
		})
	}
	if api.AuthorSoftDeleteAuthorHandler == nil {
		api.AuthorSoftDeleteAuthorHandler = author.SoftDeleteAuthorHandlerFunc(func(params author.SoftDeleteAuthorParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation author.SoftDeleteAuthor has not yet been implemented")
		})
	}
	if api.AuthorUpdateAuthorHandler == nil {
		api.AuthorUpdateAuthorHandler = author.UpdateAuthorHandlerFunc(func(params author.UpdateAuthorParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation author.UpdateAuthor has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
