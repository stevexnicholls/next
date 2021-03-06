// Code generated by go-swagger; DO NOT EDIT.

package kv

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/stevexnicholls/next/models"
)

// ValueUpdateReader is a Reader for the ValueUpdate structure.
type ValueUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValueUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewValueUpdateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewValueUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewValueUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewValueUpdateCreated creates a ValueUpdateCreated with default headers values
func NewValueUpdateCreated() *ValueUpdateCreated {
	return &ValueUpdateCreated{}
}

/*ValueUpdateCreated handles this case with default header values.

Updated successfully
*/
type ValueUpdateCreated struct {
	Payload *models.KeyValue
}

func (o *ValueUpdateCreated) Error() string {
	return fmt.Sprintf("[PUT /v1alpha/kv][%d] valueUpdateCreated  %+v", 201, o.Payload)
}

func (o *ValueUpdateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KeyValue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValueUpdateNotFound creates a ValueUpdateNotFound with default headers values
func NewValueUpdateNotFound() *ValueUpdateNotFound {
	return &ValueUpdateNotFound{}
}

/*ValueUpdateNotFound handles this case with default header values.

The entry was not found
*/
type ValueUpdateNotFound struct {
	Payload *models.Error
}

func (o *ValueUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1alpha/kv][%d] valueUpdateNotFound  %+v", 404, o.Payload)
}

func (o *ValueUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValueUpdateDefault creates a ValueUpdateDefault with default headers values
func NewValueUpdateDefault(code int) *ValueUpdateDefault {
	return &ValueUpdateDefault{
		_statusCode: code,
	}
}

/*ValueUpdateDefault handles this case with default header values.

Error
*/
type ValueUpdateDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the value update default response
func (o *ValueUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ValueUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /v1alpha/kv][%d] ValueUpdate default  %+v", o._statusCode, o.Payload)
}

func (o *ValueUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
