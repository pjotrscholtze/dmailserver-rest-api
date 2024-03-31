// Code generated by go-swagger; DO NOT EDIT.

package fail2ban

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// NewPostFail2banIPParams creates a new PostFail2banIPParams object
//
// There are no default values defined in the spec.
func NewPostFail2banIPParams() PostFail2banIPParams {

	return PostFail2banIPParams{}
}

// PostFail2banIPParams contains all the bound params for the post fail2ban Ip operation
// typically these are obtained from a http.Request
//
// swagger:parameters postFail2banIp
type PostFail2banIPParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*IP address to add to fail2ban
	  Required: true
	  In: body
	*/
	Ipaddress string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostFail2banIPParams() beforehand.
func (o *PostFail2banIPParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body string
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("ipaddress", "body", ""))
			} else {
				res = append(res, errors.NewParseError("ipaddress", "body", "", err))
			}
		} else {
			// no validation required on inline body
			o.Ipaddress = body
		}
	} else {
		res = append(res, errors.Required("ipaddress", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}