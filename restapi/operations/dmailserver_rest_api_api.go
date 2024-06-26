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

	"github.com/pjotrscholtze/dmailserver-rest-api/restapi/operations/email"
	"github.com/pjotrscholtze/dmailserver-rest-api/restapi/operations/fail2ban"
)

// NewDmailserverRestAPIAPI creates a new DmailserverRestAPI instance
func NewDmailserverRestAPIAPI(spec *loads.Document) *DmailserverRestAPIAPI {
	return &DmailserverRestAPIAPI{
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

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		EmailAddEmailAccountHandler: email.AddEmailAccountHandlerFunc(func(params email.AddEmailAccountParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation email.AddEmailAccount has not yet been implemented")
		}),
		EmailAddEmailAliasHandler: email.AddEmailAliasHandlerFunc(func(params email.AddEmailAliasParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation email.AddEmailAlias has not yet been implemented")
		}),
		EmailDeleteEmailAccountHandler: email.DeleteEmailAccountHandlerFunc(func(params email.DeleteEmailAccountParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation email.DeleteEmailAccount has not yet been implemented")
		}),
		EmailDeleteEmailAliasHandler: email.DeleteEmailAliasHandlerFunc(func(params email.DeleteEmailAliasParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation email.DeleteEmailAlias has not yet been implemented")
		}),
		Fail2banDeleteFail2banIPHandler: fail2ban.DeleteFail2banIPHandlerFunc(func(params fail2ban.DeleteFail2banIPParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation fail2ban.DeleteFail2banIP has not yet been implemented")
		}),
		Fail2banGetFail2banIpsHandler: fail2ban.GetFail2banIpsHandlerFunc(func(params fail2ban.GetFail2banIpsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation fail2ban.GetFail2banIps has not yet been implemented")
		}),
		EmailListEmailAccountsHandler: email.ListEmailAccountsHandlerFunc(func(params email.ListEmailAccountsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation email.ListEmailAccounts has not yet been implemented")
		}),
		Fail2banPostFail2banIPHandler: fail2ban.PostFail2banIPHandlerFunc(func(params fail2ban.PostFail2banIPParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation fail2ban.PostFail2banIP has not yet been implemented")
		}),
		EmailUpdateEmailAddressHandler: email.UpdateEmailAddressHandlerFunc(func(params email.UpdateEmailAddressParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation email.UpdateEmailAddress has not yet been implemented")
		}),

		// Applies when the "X-API-Key" header is set
		APIKeyAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (api_key) X-API-Key from header param [X-API-Key] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*DmailserverRestAPIAPI the dmailserver rest API API */
type DmailserverRestAPIAPI struct {
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

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// APIKeyAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key X-API-Key provided in the header
	APIKeyAuth func(string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// EmailAddEmailAccountHandler sets the operation handler for the add email account operation
	EmailAddEmailAccountHandler email.AddEmailAccountHandler
	// EmailAddEmailAliasHandler sets the operation handler for the add email alias operation
	EmailAddEmailAliasHandler email.AddEmailAliasHandler
	// EmailDeleteEmailAccountHandler sets the operation handler for the delete email account operation
	EmailDeleteEmailAccountHandler email.DeleteEmailAccountHandler
	// EmailDeleteEmailAliasHandler sets the operation handler for the delete email alias operation
	EmailDeleteEmailAliasHandler email.DeleteEmailAliasHandler
	// Fail2banDeleteFail2banIPHandler sets the operation handler for the delete fail2ban Ip operation
	Fail2banDeleteFail2banIPHandler fail2ban.DeleteFail2banIPHandler
	// Fail2banGetFail2banIpsHandler sets the operation handler for the get fail2ban ips operation
	Fail2banGetFail2banIpsHandler fail2ban.GetFail2banIpsHandler
	// EmailListEmailAccountsHandler sets the operation handler for the list email accounts operation
	EmailListEmailAccountsHandler email.ListEmailAccountsHandler
	// Fail2banPostFail2banIPHandler sets the operation handler for the post fail2ban Ip operation
	Fail2banPostFail2banIPHandler fail2ban.PostFail2banIPHandler
	// EmailUpdateEmailAddressHandler sets the operation handler for the update email address operation
	EmailUpdateEmailAddressHandler email.UpdateEmailAddressHandler

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
func (o *DmailserverRestAPIAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *DmailserverRestAPIAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *DmailserverRestAPIAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *DmailserverRestAPIAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *DmailserverRestAPIAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *DmailserverRestAPIAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *DmailserverRestAPIAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *DmailserverRestAPIAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *DmailserverRestAPIAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the DmailserverRestAPIAPI
func (o *DmailserverRestAPIAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.APIKeyAuth == nil {
		unregistered = append(unregistered, "XAPIKeyAuth")
	}

	if o.EmailAddEmailAccountHandler == nil {
		unregistered = append(unregistered, "email.AddEmailAccountHandler")
	}
	if o.EmailAddEmailAliasHandler == nil {
		unregistered = append(unregistered, "email.AddEmailAliasHandler")
	}
	if o.EmailDeleteEmailAccountHandler == nil {
		unregistered = append(unregistered, "email.DeleteEmailAccountHandler")
	}
	if o.EmailDeleteEmailAliasHandler == nil {
		unregistered = append(unregistered, "email.DeleteEmailAliasHandler")
	}
	if o.Fail2banDeleteFail2banIPHandler == nil {
		unregistered = append(unregistered, "fail2ban.DeleteFail2banIPHandler")
	}
	if o.Fail2banGetFail2banIpsHandler == nil {
		unregistered = append(unregistered, "fail2ban.GetFail2banIpsHandler")
	}
	if o.EmailListEmailAccountsHandler == nil {
		unregistered = append(unregistered, "email.ListEmailAccountsHandler")
	}
	if o.Fail2banPostFail2banIPHandler == nil {
		unregistered = append(unregistered, "fail2ban.PostFail2banIPHandler")
	}
	if o.EmailUpdateEmailAddressHandler == nil {
		unregistered = append(unregistered, "email.UpdateEmailAddressHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *DmailserverRestAPIAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *DmailserverRestAPIAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "api_key":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.APIKeyAuth)

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *DmailserverRestAPIAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *DmailserverRestAPIAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *DmailserverRestAPIAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
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
func (o *DmailserverRestAPIAPI) HandlerFor(method, path string) (http.Handler, bool) {
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

// Context returns the middleware context for the dmailserver rest API API
func (o *DmailserverRestAPIAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *DmailserverRestAPIAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/email"] = email.NewAddEmailAccount(o.context, o.EmailAddEmailAccountHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/email/{emailAddress}/aliasses"] = email.NewAddEmailAlias(o.context, o.EmailAddEmailAliasHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/email/{emailAddress}"] = email.NewDeleteEmailAccount(o.context, o.EmailDeleteEmailAccountHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/email/{emailAddress}/aliasses"] = email.NewDeleteEmailAlias(o.context, o.EmailDeleteEmailAliasHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/fail2ban/{ip}"] = fail2ban.NewDeleteFail2banIP(o.context, o.Fail2banDeleteFail2banIPHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/fail2ban"] = fail2ban.NewGetFail2banIps(o.context, o.Fail2banGetFail2banIpsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/email"] = email.NewListEmailAccounts(o.context, o.EmailListEmailAccountsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/fail2ban"] = fail2ban.NewPostFail2banIP(o.context, o.Fail2banPostFail2banIPHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/email/{emailAddress}"] = email.NewUpdateEmailAddress(o.context, o.EmailUpdateEmailAddressHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *DmailserverRestAPIAPI) Serve(builder middleware.Builder) http.Handler {
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
func (o *DmailserverRestAPIAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *DmailserverRestAPIAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *DmailserverRestAPIAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *DmailserverRestAPIAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
