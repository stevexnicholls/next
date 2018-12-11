// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// HealthGetOKCode is the HTTP code returned for type HealthGetOK
const HealthGetOKCode int = 200

/*HealthGetOK confirm that the service is healthy

swagger:response healthGetOK
*/
type HealthGetOK struct {
}

// NewHealthGetOK creates HealthGetOK with default headers values
func NewHealthGetOK() *HealthGetOK {

	return &HealthGetOK{}
}

// WriteResponse to the client
func (o *HealthGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*HealthGetDefault generic error response

swagger:response healthGetDefault
*/
type HealthGetDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewHealthGetDefault creates HealthGetDefault with default headers values
func NewHealthGetDefault(code int) *HealthGetDefault {
	if code <= 0 {
		code = 500
	}

	return &HealthGetDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the health get default response
func (o *HealthGetDefault) WithStatusCode(code int) *HealthGetDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the health get default response
func (o *HealthGetDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the health get default response
func (o *HealthGetDefault) WithPayload(payload interface{}) *HealthGetDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the health get default response
func (o *HealthGetDefault) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *HealthGetDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}