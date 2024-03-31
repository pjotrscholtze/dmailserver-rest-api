// Code generated by go-swagger; DO NOT EDIT.

package email

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewDeleteQuotaOfEmailAccountParams creates a new DeleteQuotaOfEmailAccountParams object
//
// There are no default values defined in the spec.
func NewDeleteQuotaOfEmailAccountParams() DeleteQuotaOfEmailAccountParams {

	return DeleteQuotaOfEmailAccountParams{}
}

// DeleteQuotaOfEmailAccountParams contains all the bound params for the delete quota of email account operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteQuotaOfEmailAccount
type DeleteQuotaOfEmailAccountParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Address of the email account to delete
	  Required: true
	  In: path
	*/
	EmailAddress string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteQuotaOfEmailAccountParams() beforehand.
func (o *DeleteQuotaOfEmailAccountParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rEmailAddress, rhkEmailAddress, _ := route.Params.GetOK("emailAddress")
	if err := o.bindEmailAddress(rEmailAddress, rhkEmailAddress, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindEmailAddress binds and validates parameter EmailAddress from path.
func (o *DeleteQuotaOfEmailAccountParams) bindEmailAddress(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.EmailAddress = raw

	return nil
}
