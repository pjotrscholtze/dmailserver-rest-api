// Code generated by go-swagger; DO NOT EDIT.

package email

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

// NewAddEmailAliasParams creates a new AddEmailAliasParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddEmailAliasParams() *AddEmailAliasParams {
	return &AddEmailAliasParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddEmailAliasParamsWithTimeout creates a new AddEmailAliasParams object
// with the ability to set a timeout on a request.
func NewAddEmailAliasParamsWithTimeout(timeout time.Duration) *AddEmailAliasParams {
	return &AddEmailAliasParams{
		timeout: timeout,
	}
}

// NewAddEmailAliasParamsWithContext creates a new AddEmailAliasParams object
// with the ability to set a context for a request.
func NewAddEmailAliasParamsWithContext(ctx context.Context) *AddEmailAliasParams {
	return &AddEmailAliasParams{
		Context: ctx,
	}
}

// NewAddEmailAliasParamsWithHTTPClient creates a new AddEmailAliasParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddEmailAliasParamsWithHTTPClient(client *http.Client) *AddEmailAliasParams {
	return &AddEmailAliasParams{
		HTTPClient: client,
	}
}

/* AddEmailAliasParams contains all the parameters to send to the API endpoint
   for the add email alias operation.

   Typically these are written to a http.Request.
*/
type AddEmailAliasParams struct {

	/* Alias.

	   Alias to add to the email account
	*/
	Alias string

	/* EmailAddress.

	   Address of the email account to point the alias to.
	*/
	EmailAddress string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add email alias params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddEmailAliasParams) WithDefaults() *AddEmailAliasParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add email alias params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddEmailAliasParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add email alias params
func (o *AddEmailAliasParams) WithTimeout(timeout time.Duration) *AddEmailAliasParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add email alias params
func (o *AddEmailAliasParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add email alias params
func (o *AddEmailAliasParams) WithContext(ctx context.Context) *AddEmailAliasParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add email alias params
func (o *AddEmailAliasParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add email alias params
func (o *AddEmailAliasParams) WithHTTPClient(client *http.Client) *AddEmailAliasParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add email alias params
func (o *AddEmailAliasParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAlias adds the alias to the add email alias params
func (o *AddEmailAliasParams) WithAlias(alias string) *AddEmailAliasParams {
	o.SetAlias(alias)
	return o
}

// SetAlias adds the alias to the add email alias params
func (o *AddEmailAliasParams) SetAlias(alias string) {
	o.Alias = alias
}

// WithEmailAddress adds the emailAddress to the add email alias params
func (o *AddEmailAliasParams) WithEmailAddress(emailAddress string) *AddEmailAliasParams {
	o.SetEmailAddress(emailAddress)
	return o
}

// SetEmailAddress adds the emailAddress to the add email alias params
func (o *AddEmailAliasParams) SetEmailAddress(emailAddress string) {
	o.EmailAddress = emailAddress
}

// WriteToRequest writes these params to a swagger request
func (o *AddEmailAliasParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Alias); err != nil {
		return err
	}

	// path param emailAddress
	if err := r.SetPathParam("emailAddress", o.EmailAddress); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}