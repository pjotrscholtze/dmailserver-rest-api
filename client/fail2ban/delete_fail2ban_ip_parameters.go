// Code generated by go-swagger; DO NOT EDIT.

package fail2ban

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteFail2banIPParams creates a new DeleteFail2banIPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteFail2banIPParams() *DeleteFail2banIPParams {
	return &DeleteFail2banIPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFail2banIPParamsWithTimeout creates a new DeleteFail2banIPParams object
// with the ability to set a timeout on a request.
func NewDeleteFail2banIPParamsWithTimeout(timeout time.Duration) *DeleteFail2banIPParams {
	return &DeleteFail2banIPParams{
		timeout: timeout,
	}
}

// NewDeleteFail2banIPParamsWithContext creates a new DeleteFail2banIPParams object
// with the ability to set a context for a request.
func NewDeleteFail2banIPParamsWithContext(ctx context.Context) *DeleteFail2banIPParams {
	return &DeleteFail2banIPParams{
		Context: ctx,
	}
}

// NewDeleteFail2banIPParamsWithHTTPClient creates a new DeleteFail2banIPParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteFail2banIPParamsWithHTTPClient(client *http.Client) *DeleteFail2banIPParams {
	return &DeleteFail2banIPParams{
		HTTPClient: client,
	}
}

/* DeleteFail2banIPParams contains all the parameters to send to the API endpoint
   for the delete fail2ban Ip operation.

   Typically these are written to a http.Request.
*/
type DeleteFail2banIPParams struct {

	/* IP.

	   IP address to delete
	*/
	IP string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete fail2ban Ip params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteFail2banIPParams) WithDefaults() *DeleteFail2banIPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete fail2ban Ip params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteFail2banIPParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete fail2ban Ip params
func (o *DeleteFail2banIPParams) WithTimeout(timeout time.Duration) *DeleteFail2banIPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete fail2ban Ip params
func (o *DeleteFail2banIPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete fail2ban Ip params
func (o *DeleteFail2banIPParams) WithContext(ctx context.Context) *DeleteFail2banIPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete fail2ban Ip params
func (o *DeleteFail2banIPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete fail2ban Ip params
func (o *DeleteFail2banIPParams) WithHTTPClient(client *http.Client) *DeleteFail2banIPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete fail2ban Ip params
func (o *DeleteFail2banIPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIP adds the ip to the delete fail2ban Ip params
func (o *DeleteFail2banIPParams) WithIP(ip string) *DeleteFail2banIPParams {
	o.SetIP(ip)
	return o
}

// SetIP adds the ip to the delete fail2ban Ip params
func (o *DeleteFail2banIPParams) SetIP(ip string) {
	o.IP = ip
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFail2banIPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ip
	if err := r.SetPathParam("ip", o.IP); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}