// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// OptionsContainersGetDeleteHandlerFunc turns a function with the right signature into a options containers get delete handler
type OptionsContainersGetDeleteHandlerFunc func(OptionsContainersGetDeleteParams) middleware.Responder

// Handle executing the request and returning a response
func (fn OptionsContainersGetDeleteHandlerFunc) Handle(params OptionsContainersGetDeleteParams) middleware.Responder {
	return fn(params)
}

// OptionsContainersGetDeleteHandler interface for that can handle valid options containers get delete params
type OptionsContainersGetDeleteHandler interface {
	Handle(OptionsContainersGetDeleteParams) middleware.Responder
}

// NewOptionsContainersGetDelete creates a new http.Handler for the options containers get delete operation
func NewOptionsContainersGetDelete(ctx *middleware.Context, handler OptionsContainersGetDeleteHandler) *OptionsContainersGetDelete {
	return &OptionsContainersGetDelete{Context: ctx, Handler: handler}
}

/* OptionsContainersGetDelete swagger:route OPTIONS /containers/{containerId} optionsContainersGetDelete

OptionsContainersGetDelete options containers get delete API

*/
type OptionsContainersGetDelete struct {
	Context *middleware.Context
	Handler OptionsContainersGetDeleteHandler
}

func (o *OptionsContainersGetDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewOptionsContainersGetDeleteParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}