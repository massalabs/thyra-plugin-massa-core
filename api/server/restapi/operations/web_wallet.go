// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WebWalletHandlerFunc turns a function with the right signature into a web wallet handler
type WebWalletHandlerFunc func(WebWalletParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WebWalletHandlerFunc) Handle(params WebWalletParams) middleware.Responder {
	return fn(params)
}

// WebWalletHandler interface for that can handle valid web wallet params
type WebWalletHandler interface {
	Handle(WebWalletParams) middleware.Responder
}

// NewWebWallet creates a new http.Handler for the web wallet operation
func NewWebWallet(ctx *middleware.Context, handler WebWalletHandler) *WebWallet {
	return &WebWallet{Context: ctx, Handler: handler}
}

/* WebWallet swagger:route GET /web/wallet/{resource} webWallet

WebWallet web wallet API

*/
type WebWallet struct {
	Context *middleware.Context
	Handler WebWalletHandler
}

func (o *WebWallet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewWebWalletParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}