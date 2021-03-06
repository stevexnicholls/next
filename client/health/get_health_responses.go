// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetHealthReader is a Reader for the GetHealth structure.
type GetHealthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHealthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetHealthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetHealthDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHealthOK creates a GetHealthOK with default headers values
func NewGetHealthOK() *GetHealthOK {
	return &GetHealthOK{}
}

/*GetHealthOK handles this case with default header values.

confirm that the service is healthy
*/
type GetHealthOK struct {
}

func (o *GetHealthOK) Error() string {
	return fmt.Sprintf("[GET /health][%d] getHealthOK ", 200)
}

func (o *GetHealthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetHealthDefault creates a GetHealthDefault with default headers values
func NewGetHealthDefault(code int) *GetHealthDefault {
	return &GetHealthDefault{
		_statusCode: code,
	}
}

/*GetHealthDefault handles this case with default header values.

generic error response
*/
type GetHealthDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the get health default response
func (o *GetHealthDefault) Code() int {
	return o._statusCode
}

func (o *GetHealthDefault) Error() string {
	return fmt.Sprintf("[GET /health][%d] GetHealth default  %+v", o._statusCode, o.Payload)
}

func (o *GetHealthDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
