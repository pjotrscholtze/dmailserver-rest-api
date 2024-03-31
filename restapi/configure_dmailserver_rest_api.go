// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/pjotrscholtze/dmailserver-rest-api/restapi/operations"
	"github.com/pjotrscholtze/dmailserver-rest-api/restapi/operations/email"
	"github.com/pjotrscholtze/dmailserver-rest-api/restapi/operations/fail2ban"
)

//go:generate swagger generate server --target ../../dmailserver-rest-api --name DmailserverRestAPI --spec ../openapi2.0.yaml --principal interface{}

func configureFlags(api *operations.DmailserverRestAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.DmailserverRestAPIAPI) http.Handler {
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

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "X-API-Key" header is set
	if api.APIKeyAuth == nil {
		api.APIKeyAuth = func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (api_key) X-API-Key from header param [X-API-Key] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.EmailAddEmailAccountHandler == nil {
		api.EmailAddEmailAccountHandler = email.AddEmailAccountHandlerFunc(func(params email.AddEmailAccountParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation email.AddEmailAccount has not yet been implemented")
		})
	}
	if api.EmailDeleteEmailAccountHandler == nil {
		api.EmailDeleteEmailAccountHandler = email.DeleteEmailAccountHandlerFunc(func(params email.DeleteEmailAccountParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation email.DeleteEmailAccount has not yet been implemented")
		})
	}
	if api.Fail2banDeleteFail2banIPHandler == nil {
		api.Fail2banDeleteFail2banIPHandler = fail2ban.DeleteFail2banIPHandlerFunc(func(params fail2ban.DeleteFail2banIPParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation fail2ban.DeleteFail2banIP has not yet been implemented")
		})
	}
	if api.Fail2banGetFail2banIpsHandler == nil {
		api.Fail2banGetFail2banIpsHandler = fail2ban.GetFail2banIpsHandlerFunc(func(params fail2ban.GetFail2banIpsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation fail2ban.GetFail2banIps has not yet been implemented")
		})
	}
	if api.EmailListEmailAccountsHandler == nil {
		api.EmailListEmailAccountsHandler = email.ListEmailAccountsHandlerFunc(func(params email.ListEmailAccountsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation email.ListEmailAccounts has not yet been implemented")
		})
	}
	if api.Fail2banPostFail2banIPHandler == nil {
		api.Fail2banPostFail2banIPHandler = fail2ban.PostFail2banIPHandlerFunc(func(params fail2ban.PostFail2banIPParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation fail2ban.PostFail2banIP has not yet been implemented")
		})
	}
	if api.EmailUpdateEmailAddressHandler == nil {
		api.EmailUpdateEmailAddressHandler = email.UpdateEmailAddressHandlerFunc(func(params email.UpdateEmailAddressParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation email.UpdateEmailAddress has not yet been implemented")
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
