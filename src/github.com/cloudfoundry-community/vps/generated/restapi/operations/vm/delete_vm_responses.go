package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cloudfoundry-community/vps/models"
)

/*DeleteVMNoContent vm removed successfully

swagger:response deleteVmNoContent
*/
type DeleteVMNoContent struct {

	// In: body
	Payload string `json:"body,omitempty"`
}

// NewDeleteVMNoContent creates DeleteVMNoContent with default headers values
func NewDeleteVMNoContent() *DeleteVMNoContent {
	return &DeleteVMNoContent{}
}

// WithPayload adds the payload to the delete Vm no content response
func (o *DeleteVMNoContent) WithPayload(payload string) *DeleteVMNoContent {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete Vm no content response
func (o *DeleteVMNoContent) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteVMNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*DeleteVMNotFound vm not found

swagger:response deleteVmNotFound
*/
type DeleteVMNotFound struct {
}

// NewDeleteVMNotFound creates DeleteVMNotFound with default headers values
func NewDeleteVMNotFound() *DeleteVMNotFound {
	return &DeleteVMNotFound{}
}

// WriteResponse to the client
func (o *DeleteVMNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

/*DeleteVMDefault unexpected error

swagger:response deleteVmDefault
*/
type DeleteVMDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteVMDefault creates DeleteVMDefault with default headers values
func NewDeleteVMDefault(code int) *DeleteVMDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteVMDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete Vm default response
func (o *DeleteVMDefault) WithStatusCode(code int) *DeleteVMDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete Vm default response
func (o *DeleteVMDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete Vm default response
func (o *DeleteVMDefault) WithPayload(payload *models.Error) *DeleteVMDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete Vm default response
func (o *DeleteVMDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteVMDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
