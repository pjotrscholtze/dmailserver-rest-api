// Code generated by go-swagger; DO NOT EDIT.

package email

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pjotrscholtze/dmailserver-rest-api/models"
)

// GetQuotaOfEmailAccountOKCode is the HTTP code returned for type GetQuotaOfEmailAccountOK
const GetQuotaOfEmailAccountOKCode int = 200

/*GetQuotaOfEmailAccountOK Successful retrieval of email account quota

swagger:response getQuotaOfEmailAccountOK
*/
type GetQuotaOfEmailAccountOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Quota `json:"body,omitempty"`
}

// NewGetQuotaOfEmailAccountOK creates GetQuotaOfEmailAccountOK with default headers values
func NewGetQuotaOfEmailAccountOK() *GetQuotaOfEmailAccountOK {

	return &GetQuotaOfEmailAccountOK{}
}

// WithPayload adds the payload to the get quota of email account o k response
func (o *GetQuotaOfEmailAccountOK) WithPayload(payload []*models.Quota) *GetQuotaOfEmailAccountOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get quota of email account o k response
func (o *GetQuotaOfEmailAccountOK) SetPayload(payload []*models.Quota) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetQuotaOfEmailAccountOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Quota, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetQuotaOfEmailAccountNotFoundCode is the HTTP code returned for type GetQuotaOfEmailAccountNotFound
const GetQuotaOfEmailAccountNotFoundCode int = 404

/*GetQuotaOfEmailAccountNotFound Email account not found

swagger:response getQuotaOfEmailAccountNotFound
*/
type GetQuotaOfEmailAccountNotFound struct {
}

// NewGetQuotaOfEmailAccountNotFound creates GetQuotaOfEmailAccountNotFound with default headers values
func NewGetQuotaOfEmailAccountNotFound() *GetQuotaOfEmailAccountNotFound {

	return &GetQuotaOfEmailAccountNotFound{}
}

// WriteResponse to the client
func (o *GetQuotaOfEmailAccountNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetQuotaOfEmailAccountInternalServerErrorCode is the HTTP code returned for type GetQuotaOfEmailAccountInternalServerError
const GetQuotaOfEmailAccountInternalServerErrorCode int = 500

/*GetQuotaOfEmailAccountInternalServerError Internal error

swagger:response getQuotaOfEmailAccountInternalServerError
*/
type GetQuotaOfEmailAccountInternalServerError struct {
}

// NewGetQuotaOfEmailAccountInternalServerError creates GetQuotaOfEmailAccountInternalServerError with default headers values
func NewGetQuotaOfEmailAccountInternalServerError() *GetQuotaOfEmailAccountInternalServerError {

	return &GetQuotaOfEmailAccountInternalServerError{}
}

// WriteResponse to the client
func (o *GetQuotaOfEmailAccountInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
