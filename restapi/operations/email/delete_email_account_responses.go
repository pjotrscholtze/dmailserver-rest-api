// Code generated by go-swagger; DO NOT EDIT.

package email

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteEmailAccountOKCode is the HTTP code returned for type DeleteEmailAccountOK
const DeleteEmailAccountOKCode int = 200

/*DeleteEmailAccountOK Successful operation

swagger:response deleteEmailAccountOK
*/
type DeleteEmailAccountOK struct {
}

// NewDeleteEmailAccountOK creates DeleteEmailAccountOK with default headers values
func NewDeleteEmailAccountOK() *DeleteEmailAccountOK {

	return &DeleteEmailAccountOK{}
}

// WriteResponse to the client
func (o *DeleteEmailAccountOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteEmailAccountNotFoundCode is the HTTP code returned for type DeleteEmailAccountNotFound
const DeleteEmailAccountNotFoundCode int = 404

/*DeleteEmailAccountNotFound Account not found

swagger:response deleteEmailAccountNotFound
*/
type DeleteEmailAccountNotFound struct {
}

// NewDeleteEmailAccountNotFound creates DeleteEmailAccountNotFound with default headers values
func NewDeleteEmailAccountNotFound() *DeleteEmailAccountNotFound {

	return &DeleteEmailAccountNotFound{}
}

// WriteResponse to the client
func (o *DeleteEmailAccountNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// DeleteEmailAccountMethodNotAllowedCode is the HTTP code returned for type DeleteEmailAccountMethodNotAllowed
const DeleteEmailAccountMethodNotAllowedCode int = 405

/*DeleteEmailAccountMethodNotAllowed Invalid input

swagger:response deleteEmailAccountMethodNotAllowed
*/
type DeleteEmailAccountMethodNotAllowed struct {
}

// NewDeleteEmailAccountMethodNotAllowed creates DeleteEmailAccountMethodNotAllowed with default headers values
func NewDeleteEmailAccountMethodNotAllowed() *DeleteEmailAccountMethodNotAllowed {

	return &DeleteEmailAccountMethodNotAllowed{}
}

// WriteResponse to the client
func (o *DeleteEmailAccountMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}

// DeleteEmailAccountInternalServerErrorCode is the HTTP code returned for type DeleteEmailAccountInternalServerError
const DeleteEmailAccountInternalServerErrorCode int = 500

/*DeleteEmailAccountInternalServerError Internal error

swagger:response deleteEmailAccountInternalServerError
*/
type DeleteEmailAccountInternalServerError struct {
}

// NewDeleteEmailAccountInternalServerError creates DeleteEmailAccountInternalServerError with default headers values
func NewDeleteEmailAccountInternalServerError() *DeleteEmailAccountInternalServerError {

	return &DeleteEmailAccountInternalServerError{}
}

// WriteResponse to the client
func (o *DeleteEmailAccountInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
