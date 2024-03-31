// Code generated by go-swagger; DO NOT EDIT.

package email

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pjotrscholtze/dmailserver-rest-api/models"
)

// ListEmailAccountsOKCode is the HTTP code returned for type ListEmailAccountsOK
const ListEmailAccountsOKCode int = 200

/*ListEmailAccountsOK Successful listing of email addresses

swagger:response listEmailAccountsOK
*/
type ListEmailAccountsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.EmailAccountListItem `json:"body,omitempty"`
}

// NewListEmailAccountsOK creates ListEmailAccountsOK with default headers values
func NewListEmailAccountsOK() *ListEmailAccountsOK {

	return &ListEmailAccountsOK{}
}

// WithPayload adds the payload to the list email accounts o k response
func (o *ListEmailAccountsOK) WithPayload(payload []*models.EmailAccountListItem) *ListEmailAccountsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list email accounts o k response
func (o *ListEmailAccountsOK) SetPayload(payload []*models.EmailAccountListItem) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEmailAccountsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.EmailAccountListItem, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListEmailAccountsInternalServerErrorCode is the HTTP code returned for type ListEmailAccountsInternalServerError
const ListEmailAccountsInternalServerErrorCode int = 500

/*ListEmailAccountsInternalServerError Internal error

swagger:response listEmailAccountsInternalServerError
*/
type ListEmailAccountsInternalServerError struct {
}

// NewListEmailAccountsInternalServerError creates ListEmailAccountsInternalServerError with default headers values
func NewListEmailAccountsInternalServerError() *ListEmailAccountsInternalServerError {

	return &ListEmailAccountsInternalServerError{}
}

// WriteResponse to the client
func (o *ListEmailAccountsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}