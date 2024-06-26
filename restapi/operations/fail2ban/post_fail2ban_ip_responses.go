// Code generated by go-swagger; DO NOT EDIT.

package fail2ban

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostFail2banIPOKCode is the HTTP code returned for type PostFail2banIPOK
const PostFail2banIPOKCode int = 200

/*PostFail2banIPOK Successful operation

swagger:response postFail2banIpOK
*/
type PostFail2banIPOK struct {
}

// NewPostFail2banIPOK creates PostFail2banIPOK with default headers values
func NewPostFail2banIPOK() *PostFail2banIPOK {

	return &PostFail2banIPOK{}
}

// WriteResponse to the client
func (o *PostFail2banIPOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PostFail2banIPMethodNotAllowedCode is the HTTP code returned for type PostFail2banIPMethodNotAllowed
const PostFail2banIPMethodNotAllowedCode int = 405

/*PostFail2banIPMethodNotAllowed Invalid input

swagger:response postFail2banIpMethodNotAllowed
*/
type PostFail2banIPMethodNotAllowed struct {
}

// NewPostFail2banIPMethodNotAllowed creates PostFail2banIPMethodNotAllowed with default headers values
func NewPostFail2banIPMethodNotAllowed() *PostFail2banIPMethodNotAllowed {

	return &PostFail2banIPMethodNotAllowed{}
}

// WriteResponse to the client
func (o *PostFail2banIPMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}

// PostFail2banIPInternalServerErrorCode is the HTTP code returned for type PostFail2banIPInternalServerError
const PostFail2banIPInternalServerErrorCode int = 500

/*PostFail2banIPInternalServerError Internal error

swagger:response postFail2banIpInternalServerError
*/
type PostFail2banIPInternalServerError struct {
}

// NewPostFail2banIPInternalServerError creates PostFail2banIPInternalServerError with default headers values
func NewPostFail2banIPInternalServerError() *PostFail2banIPInternalServerError {

	return &PostFail2banIPInternalServerError{}
}

// WriteResponse to the client
func (o *PostFail2banIPInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
