// Code generated by go-swagger; DO NOT EDIT.

package kv

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new kv API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for kv API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
KeyDelete deletes an existing key
*/
func (a *Client) KeyDelete(params *KeyDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*KeyDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKeyDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "KeyDelete",
		Method:             "DELETE",
		PathPattern:        "/v1alpha/kv/{key}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &KeyDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*KeyDeleteNoContent), nil

}

/*
KeyList Lists all the keys
*/
func (a *Client) KeyList(params *KeyListParams, authInfo runtime.ClientAuthInfoWriter) (*KeyListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKeyListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "KeyList",
		Method:             "GET",
		PathPattern:        "/v1alpha/kv",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &KeyListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*KeyListOK), nil

}

/*
ValueGet gets a value
*/
func (a *Client) ValueGet(params *ValueGetParams, authInfo runtime.ClientAuthInfoWriter) (*ValueGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValueGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ValueGet",
		Method:             "GET",
		PathPattern:        "/v1alpha/kv/{key}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ValueGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ValueGetOK), nil

}

/*
ValueUpdate updates value
*/
func (a *Client) ValueUpdate(params *ValueUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*ValueUpdateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValueUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ValueUpdate",
		Method:             "PUT",
		PathPattern:        "/v1alpha/kv",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ValueUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ValueUpdateCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
