// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RestWalletSignOperationHandlerFunc turns a function with the right signature into a rest wallet sign operation handler
type RestWalletSignOperationHandlerFunc func(RestWalletSignOperationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RestWalletSignOperationHandlerFunc) Handle(params RestWalletSignOperationParams) middleware.Responder {
	return fn(params)
}

// RestWalletSignOperationHandler interface for that can handle valid rest wallet sign operation params
type RestWalletSignOperationHandler interface {
	Handle(RestWalletSignOperationParams) middleware.Responder
}

// NewRestWalletSignOperation creates a new http.Handler for the rest wallet sign operation operation
func NewRestWalletSignOperation(ctx *middleware.Context, handler RestWalletSignOperationHandler) *RestWalletSignOperation {
	return &RestWalletSignOperation{Context: ctx, Handler: handler}
}

/* RestWalletSignOperation swagger:route POST /rest/wallet/{nickname}/signOperation restWalletSignOperation

RestWalletSignOperation rest wallet sign operation API

*/
type RestWalletSignOperation struct {
	Context *middleware.Context
	Handler RestWalletSignOperationHandler
}

func (o *RestWalletSignOperation) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewRestWalletSignOperationParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// RestWalletSignOperationBody rest wallet sign operation body
//
// swagger:model RestWalletSignOperationBody
type RestWalletSignOperationBody struct {

	// Serialized attributes of the operation to be signed with the key pair corresponding to the given nickname.
	// Required: true
	// Format: byte
	Operation *strfmt.Base64 `json:"operation"`
}

// Validate validates this rest wallet sign operation body
func (o *RestWalletSignOperationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RestWalletSignOperationBody) validateOperation(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"operation", "body", o.Operation); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rest wallet sign operation body based on context it is used
func (o *RestWalletSignOperationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RestWalletSignOperationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RestWalletSignOperationBody) UnmarshalBinary(b []byte) error {
	var res RestWalletSignOperationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
