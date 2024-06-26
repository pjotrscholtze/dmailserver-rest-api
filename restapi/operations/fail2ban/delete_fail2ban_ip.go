// Code generated by go-swagger; DO NOT EDIT.

package fail2ban

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteFail2banIPHandlerFunc turns a function with the right signature into a delete fail2ban Ip handler
type DeleteFail2banIPHandlerFunc func(DeleteFail2banIPParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteFail2banIPHandlerFunc) Handle(params DeleteFail2banIPParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteFail2banIPHandler interface for that can handle valid delete fail2ban Ip params
type DeleteFail2banIPHandler interface {
	Handle(DeleteFail2banIPParams, interface{}) middleware.Responder
}

// NewDeleteFail2banIP creates a new http.Handler for the delete fail2ban Ip operation
func NewDeleteFail2banIP(ctx *middleware.Context, handler DeleteFail2banIPHandler) *DeleteFail2banIP {
	return &DeleteFail2banIP{Context: ctx, Handler: handler}
}

/* DeleteFail2banIP swagger:route DELETE /fail2ban/{ip} fail2ban deleteFail2banIp

DeleteFail2banIP delete fail2ban Ip API

*/
type DeleteFail2banIP struct {
	Context *middleware.Context
	Handler DeleteFail2banIPHandler
}

func (o *DeleteFail2banIP) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteFail2banIPParams()
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
