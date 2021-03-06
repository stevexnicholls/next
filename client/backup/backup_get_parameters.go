// Code generated by go-swagger; DO NOT EDIT.

package backup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewBackupGetParams creates a new BackupGetParams object
// with the default values initialized.
func NewBackupGetParams() *BackupGetParams {

	return &BackupGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBackupGetParamsWithTimeout creates a new BackupGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBackupGetParamsWithTimeout(timeout time.Duration) *BackupGetParams {

	return &BackupGetParams{

		timeout: timeout,
	}
}

// NewBackupGetParamsWithContext creates a new BackupGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewBackupGetParamsWithContext(ctx context.Context) *BackupGetParams {

	return &BackupGetParams{

		Context: ctx,
	}
}

// NewBackupGetParamsWithHTTPClient creates a new BackupGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBackupGetParamsWithHTTPClient(client *http.Client) *BackupGetParams {

	return &BackupGetParams{
		HTTPClient: client,
	}
}

/*BackupGetParams contains all the parameters to send to the API endpoint
for the backup get operation typically these are written to a http.Request
*/
type BackupGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the backup get params
func (o *BackupGetParams) WithTimeout(timeout time.Duration) *BackupGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the backup get params
func (o *BackupGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the backup get params
func (o *BackupGetParams) WithContext(ctx context.Context) *BackupGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the backup get params
func (o *BackupGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the backup get params
func (o *BackupGetParams) WithHTTPClient(client *http.Client) *BackupGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the backup get params
func (o *BackupGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *BackupGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
