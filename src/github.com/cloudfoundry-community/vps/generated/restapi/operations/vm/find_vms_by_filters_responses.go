package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cloudfoundry-community/vps/generated/models"
)

/*FindVmsByFiltersOK successful operation

swagger:response findVmsByFiltersOK
*/
type FindVmsByFiltersOK struct {

	// In: body
	Payload *models.VmsResponse `json:"body,omitempty"`
}

// NewFindVmsByFiltersOK creates FindVmsByFiltersOK with default headers values
func NewFindVmsByFiltersOK() *FindVmsByFiltersOK {
	return &FindVmsByFiltersOK{}
}

// WithPayload adds the payload to the find vms by filters o k response
func (o *FindVmsByFiltersOK) WithPayload(payload *models.VmsResponse) *FindVmsByFiltersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find vms by filters o k response
func (o *FindVmsByFiltersOK) SetPayload(payload *models.VmsResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindVmsByFiltersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*FindVmsByFiltersNotFound vm not found

swagger:response findVmsByFiltersNotFound
*/
type FindVmsByFiltersNotFound struct {
}

// NewFindVmsByFiltersNotFound creates FindVmsByFiltersNotFound with default headers values
func NewFindVmsByFiltersNotFound() *FindVmsByFiltersNotFound {
	return &FindVmsByFiltersNotFound{}
}

// WriteResponse to the client
func (o *FindVmsByFiltersNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

/*FindVmsByFiltersDefault unexpected error

swagger:response findVmsByFiltersDefault
*/
type FindVmsByFiltersDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewFindVmsByFiltersDefault creates FindVmsByFiltersDefault with default headers values
func NewFindVmsByFiltersDefault(code int) *FindVmsByFiltersDefault {
	if code <= 0 {
		code = 500
	}

	return &FindVmsByFiltersDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the find vms by filters default response
func (o *FindVmsByFiltersDefault) WithStatusCode(code int) *FindVmsByFiltersDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the find vms by filters default response
func (o *FindVmsByFiltersDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the find vms by filters default response
func (o *FindVmsByFiltersDefault) WithPayload(payload *models.Error) *FindVmsByFiltersDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find vms by filters default response
func (o *FindVmsByFiltersDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindVmsByFiltersDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
