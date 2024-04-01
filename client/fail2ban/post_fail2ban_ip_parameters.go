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

// NewPostFail2banIPParams creates a new PostFail2banIPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostFail2banIPParams() *PostFail2banIPParams {
	return &PostFail2banIPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostFail2banIPParamsWithTimeout creates a new PostFail2banIPParams object
// with the ability to set a timeout on a request.
func NewPostFail2banIPParamsWithTimeout(timeout time.Duration) *PostFail2banIPParams {
	return &PostFail2banIPParams{
		timeout: timeout,
	}
}

// NewPostFail2banIPParamsWithContext creates a new PostFail2banIPParams object
// with the ability to set a context for a request.
func NewPostFail2banIPParamsWithContext(ctx context.Context) *PostFail2banIPParams {
	return &PostFail2banIPParams{
		Context: ctx,
	}
}

// NewPostFail2banIPParamsWithHTTPClient creates a new PostFail2banIPParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostFail2banIPParamsWithHTTPClient(client *http.Client) *PostFail2banIPParams {
	return &PostFail2banIPParams{
		HTTPClient: client,
	}
}

/* PostFail2banIPParams contains all the parameters to send to the API endpoint
   for the post fail2ban Ip operation.

   Typically these are written to a http.Request.
*/
type PostFail2banIPParams struct {

	/* Ipaddress.

	   IP address to add to fail2ban
	*/
	Ipaddress string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post fail2ban Ip params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostFail2banIPParams) WithDefaults() *PostFail2banIPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post fail2ban Ip params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostFail2banIPParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post fail2ban Ip params
func (o *PostFail2banIPParams) WithTimeout(timeout time.Duration) *PostFail2banIPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post fail2ban Ip params
func (o *PostFail2banIPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post fail2ban Ip params
func (o *PostFail2banIPParams) WithContext(ctx context.Context) *PostFail2banIPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post fail2ban Ip params
func (o *PostFail2banIPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post fail2ban Ip params
func (o *PostFail2banIPParams) WithHTTPClient(client *http.Client) *PostFail2banIPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post fail2ban Ip params
func (o *PostFail2banIPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIpaddress adds the ipaddress to the post fail2ban Ip params
func (o *PostFail2banIPParams) WithIpaddress(ipaddress string) *PostFail2banIPParams {
	o.SetIpaddress(ipaddress)
	return o
}

// SetIpaddress adds the ipaddress to the post fail2ban Ip params
func (o *PostFail2banIPParams) SetIpaddress(ipaddress string) {
	o.Ipaddress = ipaddress
}

// WriteToRequest writes these params to a swagger request
func (o *PostFail2banIPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Ipaddress); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
