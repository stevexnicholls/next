// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// HealthGetHandlerFunc turns a function with the right signature into a health get handler
type HealthGetHandlerFunc func(HealthGetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn HealthGetHandlerFunc) Handle(params HealthGetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// HealthGetHandler interface for that can handle valid health get params
type HealthGetHandler interface {
	Handle(HealthGetParams, interface{}) middleware.Responder
}

// NewHealthGet creates a new http.Handler for the health get operation
func NewHealthGet(ctx *middleware.Context, handler HealthGetHandler) *HealthGet {
	return &HealthGet{Context: ctx, Handler: handler}
}

/*HealthGet swagger:route GET /health health healthGet

get service health

*/
type HealthGet struct {
	Context *middleware.Context
	Handler HealthGetHandler
}

func (o *HealthGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewHealthGetParams()

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
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
