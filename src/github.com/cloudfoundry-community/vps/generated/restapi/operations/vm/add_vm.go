package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AddVMHandlerFunc turns a function with the right signature into a add Vm handler
type AddVMHandlerFunc func(AddVMParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddVMHandlerFunc) Handle(params AddVMParams) middleware.Responder {
	return fn(params)
}

// AddVMHandler interface for that can handle valid add Vm params
type AddVMHandler interface {
	Handle(AddVMParams) middleware.Responder
}

// NewAddVM creates a new http.Handler for the add Vm operation
func NewAddVM(ctx *middleware.Context, handler AddVMHandler) *AddVM {
	return &AddVM{Context: ctx, Handler: handler}
}

/*AddVM swagger:route POST /vms vm addVm

Add a new vm to the pool

*/
type AddVM struct {
	Context *middleware.Context
	Handler AddVMHandler
}

func (o *AddVM) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewAddVMParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
