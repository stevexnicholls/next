// Code generated by go-swagger; DO NOT EDIT.

package kv

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/stevexnicholls/next/models"
)

// ValueUpdateCreatedCode is the HTTP code returned for type ValueUpdateCreated
const ValueUpdateCreatedCode int = 201

/*ValueUpdateCreated Updated successfully

swagger:response valueUpdateCreated
*/
type ValueUpdateCreated struct {

	/*
	  In: Body
	*/
	Payload *models.KeyValue `json:"body,omitempty"`
}

// NewValueUpdateCreated creates ValueUpdateCreated with default headers values
func NewValueUpdateCreated() *ValueUpdateCreated {

	return &ValueUpdateCreated{}
}

// WithPayload adds the payload to the value update created response
func (o *ValueUpdateCreated) WithPayload(payload *models.KeyValue) *ValueUpdateCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the value update created response
func (o *ValueUpdateCreated) SetPayload(payload *models.KeyValue) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ValueUpdateCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ValueUpdateNotFoundCode is the HTTP code returned for type ValueUpdateNotFound
const ValueUpdateNotFoundCode int = 404

/*ValueUpdateNotFound The entry was not found

swagger:response valueUpdateNotFound
*/
type ValueUpdateNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewValueUpdateNotFound creates ValueUpdateNotFound with default headers values
func NewValueUpdateNotFound() *ValueUpdateNotFound {

	return &ValueUpdateNotFound{}
}

// WithPayload adds the payload to the value update not found response
func (o *ValueUpdateNotFound) WithPayload(payload *models.Error) *ValueUpdateNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the value update not found response
func (o *ValueUpdateNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ValueUpdateNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ValueUpdateDefault Error

swagger:response valueUpdateDefault
*/
type ValueUpdateDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewValueUpdateDefault creates ValueUpdateDefault with default headers values
func NewValueUpdateDefault(code int) *ValueUpdateDefault {
	if code <= 0 {
		code = 500
	}

	return &ValueUpdateDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the value update default response
func (o *ValueUpdateDefault) WithStatusCode(code int) *ValueUpdateDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the value update default response
func (o *ValueUpdateDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the value update default response
func (o *ValueUpdateDefault) WithPayload(payload *models.Error) *ValueUpdateDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the value update default response
func (o *ValueUpdateDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ValueUpdateDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
