// Code generated by go-swagger; DO NOT EDIT.

package email

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateEmailAddressHandlerFunc turns a function with the right signature into a update email address handler
type UpdateEmailAddressHandlerFunc func(UpdateEmailAddressParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateEmailAddressHandlerFunc) Handle(params UpdateEmailAddressParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateEmailAddressHandler interface for that can handle valid update email address params
type UpdateEmailAddressHandler interface {
	Handle(UpdateEmailAddressParams, interface{}) middleware.Responder
}

// NewUpdateEmailAddress creates a new http.Handler for the update email address operation
func NewUpdateEmailAddress(ctx *middleware.Context, handler UpdateEmailAddressHandler) *UpdateEmailAddress {
	return &UpdateEmailAddress{Context: ctx, Handler: handler}
}

/* UpdateEmailAddress swagger:route PUT /email/{emailAddress} email updateEmailAddress

UpdateEmailAddress update email address API

*/
type UpdateEmailAddress struct {
	Context *middleware.Context
	Handler UpdateEmailAddressHandler
}

func (o *UpdateEmailAddress) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateEmailAddressParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
