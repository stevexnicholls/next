// Code generated by go-swagger; DO NOT EDIT.

package kv

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// ValueGetURL generates an URL for the value get operation
type ValueGetURL struct {
	Key string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ValueGetURL) WithBasePath(bp string) *ValueGetURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ValueGetURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *ValueGetURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/v1alpha/kv/{key}"

	key := o.Key
	if key != "" {
		_path = strings.Replace(_path, "{key}", key, -1)
	} else {
		return nil, errors.New("Key is required on ValueGetURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *ValueGetURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *ValueGetURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *ValueGetURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on ValueGetURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on ValueGetURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *ValueGetURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
