// Code generated by go-swagger; DO NOT EDIT.

package email

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AddEmailAliasReader is a Reader for the AddEmailAlias structure.
type AddEmailAliasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddEmailAliasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddEmailAliasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 405:
		result := NewAddEmailAliasMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewAddEmailAliasNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddEmailAliasInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddEmailAliasOK creates a AddEmailAliasOK with default headers values
func NewAddEmailAliasOK() *AddEmailAliasOK {
	return &AddEmailAliasOK{}
}

/* AddEmailAliasOK describes a response with status code 200, with default header values.

Successful operation
*/
type AddEmailAliasOK struct {
}

func (o *AddEmailAliasOK) Error() string {
	return fmt.Sprintf("[POST /email/{emailAddress}/aliasses][%d] addEmailAliasOK ", 200)
}

func (o *AddEmailAliasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddEmailAliasMethodNotAllowed creates a AddEmailAliasMethodNotAllowed with default headers values
func NewAddEmailAliasMethodNotAllowed() *AddEmailAliasMethodNotAllowed {
	return &AddEmailAliasMethodNotAllowed{}
}

/* AddEmailAliasMethodNotAllowed describes a response with status code 405, with default header values.

Invalid input
*/
type AddEmailAliasMethodNotAllowed struct {
}

func (o *AddEmailAliasMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /email/{emailAddress}/aliasses][%d] addEmailAliasMethodNotAllowed ", 405)
}

func (o *AddEmailAliasMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddEmailAliasNotAcceptable creates a AddEmailAliasNotAcceptable with default headers values
func NewAddEmailAliasNotAcceptable() *AddEmailAliasNotAcceptable {
	return &AddEmailAliasNotAcceptable{}
}

/* AddEmailAliasNotAcceptable describes a response with status code 406, with default header values.

Email alias already exists!
*/
type AddEmailAliasNotAcceptable struct {
}

func (o *AddEmailAliasNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /email/{emailAddress}/aliasses][%d] addEmailAliasNotAcceptable ", 406)
}

func (o *AddEmailAliasNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddEmailAliasInternalServerError creates a AddEmailAliasInternalServerError with default headers values
func NewAddEmailAliasInternalServerError() *AddEmailAliasInternalServerError {
	return &AddEmailAliasInternalServerError{}
}

/* AddEmailAliasInternalServerError describes a response with status code 500, with default header values.

Internal error
*/
type AddEmailAliasInternalServerError struct {
}

func (o *AddEmailAliasInternalServerError) Error() string {
	return fmt.Sprintf("[POST /email/{emailAddress}/aliasses][%d] addEmailAliasInternalServerError ", 500)
}

func (o *AddEmailAliasInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
