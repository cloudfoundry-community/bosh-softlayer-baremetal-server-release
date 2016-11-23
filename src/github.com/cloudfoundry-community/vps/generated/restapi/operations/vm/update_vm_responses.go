package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cloudfoundry-community/vps/models"
)

/*UpdateVMOK update vm in the pool

swagger:response updateVmOK
*/
type UpdateVMOK struct {

	// In: body
	Payload string `json:"body,omitempty"`
}

// NewUpdateVMOK creates UpdateVMOK with default headers values
func NewUpdateVMOK() *UpdateVMOK {
	return &UpdateVMOK{}
}

// WithPayload adds the payload to the update Vm o k response
func (o *UpdateVMOK) WithPayload(payload string) *UpdateVMOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update Vm o k response
func (o *UpdateVMOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateVMOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*UpdateVMNotFound vm not found

swagger:response updateVmNotFound
*/
type UpdateVMNotFound struct {
}

// NewUpdateVMNotFound creates UpdateVMNotFound with default headers values
func NewUpdateVMNotFound() *UpdateVMNotFound {
	return &UpdateVMNotFound{}
}

// WriteResponse to the client
func (o *UpdateVMNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

/*UpdateVMDefault unexpected error

swagger:response updateVmDefault
*/
type UpdateVMDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateVMDefault creates UpdateVMDefault with default headers values
func NewUpdateVMDefault(code int) *UpdateVMDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateVMDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update Vm default response
func (o *UpdateVMDefault) WithStatusCode(code int) *UpdateVMDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update Vm default response
func (o *UpdateVMDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update Vm default response
func (o *UpdateVMDefault) WithPayload(payload *models.Error) *UpdateVMDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update Vm default response
func (o *UpdateVMDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateVMDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
