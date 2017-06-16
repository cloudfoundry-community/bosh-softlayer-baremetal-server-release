package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cloudfoundry-community/vps/generated/models"
)

/*GetVMByCidOK successful operation

swagger:response getVmByCidOK
*/
type GetVMByCidOK struct {

	// In: body
	Payload *models.VMResponse `json:"body,omitempty"`
}

// NewGetVMByCidOK creates GetVMByCidOK with default headers values
func NewGetVMByCidOK() *GetVMByCidOK {
	return &GetVMByCidOK{}
}

// WithPayload adds the payload to the get Vm by cid o k response
func (o *GetVMByCidOK) WithPayload(payload *models.VMResponse) *GetVMByCidOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Vm by cid o k response
func (o *GetVMByCidOK) SetPayload(payload *models.VMResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVMByCidOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetVMByCidNotFound vm not found

swagger:response getVmByCidNotFound
*/
type GetVMByCidNotFound struct {
}

// NewGetVMByCidNotFound creates GetVMByCidNotFound with default headers values
func NewGetVMByCidNotFound() *GetVMByCidNotFound {
	return &GetVMByCidNotFound{}
}

// WriteResponse to the client
func (o *GetVMByCidNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

/*GetVMByCidDefault unexpected error

swagger:response getVmByCidDefault
*/
type GetVMByCidDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetVMByCidDefault creates GetVMByCidDefault with default headers values
func NewGetVMByCidDefault(code int) *GetVMByCidDefault {
	if code <= 0 {
		code = 500
	}

	return &GetVMByCidDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get Vm by cid default response
func (o *GetVMByCidDefault) WithStatusCode(code int) *GetVMByCidDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get Vm by cid default response
func (o *GetVMByCidDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get Vm by cid default response
func (o *GetVMByCidDefault) WithPayload(payload *models.Error) *GetVMByCidDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Vm by cid default response
func (o *GetVMByCidDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVMByCidDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}