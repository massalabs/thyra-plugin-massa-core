// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
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
)

// NewMassaCoreAPI creates a new MassaCore instance
func NewMassaCoreAPI(spec *loads.Document) *MassaCoreAPI {
	return &MassaCoreAPI{
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

		BinProducer: runtime.ByteStreamProducer(),
		CSSProducer: runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
			return errors.NotImplemented("css producer has not yet been implemented")
		}),
		HTMLProducer: runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
			return errors.NotImplemented("html producer has not yet been implemented")
		}),
		JsProducer: runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
			return errors.NotImplemented("js producer has not yet been implemented")
		}),
		JSONProducer: runtime.JSONProducer(),
		TextWebpProducer: runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
			return errors.NotImplemented("textWebp producer has not yet been implemented")
		}),

		RestWalletCreateHandler: RestWalletCreateHandlerFunc(func(params RestWalletCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation RestWalletCreate has not yet been implemented")
		}),
		RestWalletDeleteHandler: RestWalletDeleteHandlerFunc(func(params RestWalletDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation RestWalletDelete has not yet been implemented")
		}),
		RestWalletGetHandler: RestWalletGetHandlerFunc(func(params RestWalletGetParams) middleware.Responder {
			return middleware.NotImplemented("operation RestWalletGet has not yet been implemented")
		}),
		RestWalletImportHandler: RestWalletImportHandlerFunc(func(params RestWalletImportParams) middleware.Responder {
			return middleware.NotImplemented("operation RestWalletImport has not yet been implemented")
		}),
		RestWalletListHandler: RestWalletListHandlerFunc(func(params RestWalletListParams) middleware.Responder {
			return middleware.NotImplemented("operation RestWalletList has not yet been implemented")
		}),
		WebWalletHandler: WebWalletHandlerFunc(func(params WebWalletParams) middleware.Responder {
			return middleware.NotImplemented("operation WebWallet has not yet been implemented")
		}),
	}
}

/*MassaCoreAPI Massa core Thyra plugin. */
type MassaCoreAPI struct {
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

	// BinProducer registers a producer for the following mime types:
	//   - image/png
	BinProducer runtime.Producer
	// CSSProducer registers a producer for the following mime types:
	//   - text/css
	CSSProducer runtime.Producer
	// HTMLProducer registers a producer for the following mime types:
	//   - text/html
	HTMLProducer runtime.Producer
	// JsProducer registers a producer for the following mime types:
	//   - text/javascript
	JsProducer runtime.Producer
	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer
	// TextWebpProducer registers a producer for the following mime types:
	//   - text/webp
	TextWebpProducer runtime.Producer

	// RestWalletCreateHandler sets the operation handler for the rest wallet create operation
	RestWalletCreateHandler RestWalletCreateHandler
	// RestWalletDeleteHandler sets the operation handler for the rest wallet delete operation
	RestWalletDeleteHandler RestWalletDeleteHandler
	// RestWalletGetHandler sets the operation handler for the rest wallet get operation
	RestWalletGetHandler RestWalletGetHandler
	// RestWalletImportHandler sets the operation handler for the rest wallet import operation
	RestWalletImportHandler RestWalletImportHandler
	// RestWalletListHandler sets the operation handler for the rest wallet list operation
	RestWalletListHandler RestWalletListHandler
	// WebWalletHandler sets the operation handler for the web wallet operation
	WebWalletHandler WebWalletHandler

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
func (o *MassaCoreAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *MassaCoreAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *MassaCoreAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *MassaCoreAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *MassaCoreAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *MassaCoreAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *MassaCoreAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *MassaCoreAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *MassaCoreAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the MassaCoreAPI
func (o *MassaCoreAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.BinProducer == nil {
		unregistered = append(unregistered, "BinProducer")
	}
	if o.CSSProducer == nil {
		unregistered = append(unregistered, "CSSProducer")
	}
	if o.HTMLProducer == nil {
		unregistered = append(unregistered, "HTMLProducer")
	}
	if o.JsProducer == nil {
		unregistered = append(unregistered, "JsProducer")
	}
	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}
	if o.TextWebpProducer == nil {
		unregistered = append(unregistered, "TextWebpProducer")
	}

	if o.RestWalletCreateHandler == nil {
		unregistered = append(unregistered, "RestWalletCreateHandler")
	}
	if o.RestWalletDeleteHandler == nil {
		unregistered = append(unregistered, "RestWalletDeleteHandler")
	}
	if o.RestWalletGetHandler == nil {
		unregistered = append(unregistered, "RestWalletGetHandler")
	}
	if o.RestWalletImportHandler == nil {
		unregistered = append(unregistered, "RestWalletImportHandler")
	}
	if o.RestWalletListHandler == nil {
		unregistered = append(unregistered, "RestWalletListHandler")
	}
	if o.WebWalletHandler == nil {
		unregistered = append(unregistered, "WebWalletHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *MassaCoreAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *MassaCoreAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *MassaCoreAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *MassaCoreAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
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
func (o *MassaCoreAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "image/png":
			result["image/png"] = o.BinProducer
		case "text/css":
			result["text/css"] = o.CSSProducer
		case "text/html":
			result["text/html"] = o.HTMLProducer
		case "text/javascript":
			result["text/javascript"] = o.JsProducer
		case "application/json":
			result["application/json"] = o.JSONProducer
		case "text/webp":
			result["text/webp"] = o.TextWebpProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *MassaCoreAPI) HandlerFor(method, path string) (http.Handler, bool) {
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

// Context returns the middleware context for the massa core API
func (o *MassaCoreAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *MassaCoreAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/rest/wallet"] = NewRestWalletCreate(o.context, o.RestWalletCreateHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/rest/wallet/{nickname}"] = NewRestWalletDelete(o.context, o.RestWalletDeleteHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/rest/wallet/{nickname}"] = NewRestWalletGet(o.context, o.RestWalletGetHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/rest/wallet"] = NewRestWalletImport(o.context, o.RestWalletImportHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/rest/wallet"] = NewRestWalletList(o.context, o.RestWalletListHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/web/wallet/{resource}"] = NewWebWallet(o.context, o.WebWalletHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *MassaCoreAPI) Serve(builder middleware.Builder) http.Handler {
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
func (o *MassaCoreAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *MassaCoreAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *MassaCoreAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *MassaCoreAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
