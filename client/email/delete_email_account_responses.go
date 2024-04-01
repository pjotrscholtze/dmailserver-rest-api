// Code generated by go-swagger; DO NOT EDIT.

package email

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteEmailAccountReader is a Reader for the DeleteEmailAccount structure.
type DeleteEmailAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEmailAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteEmailAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteEmailAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewDeleteEmailAccountMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteEmailAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteEmailAccountOK creates a DeleteEmailAccountOK with default headers values
func NewDeleteEmailAccountOK() *DeleteEmailAccountOK {
	return &DeleteEmailAccountOK{}
}

/* DeleteEmailAccountOK describes a response with status code 200, with default header values.

Successful operation
*/
type DeleteEmailAccountOK struct {
}

func (o *DeleteEmailAccountOK) Error() string {
	return fmt.Sprintf("[DELETE /email/{emailAddress}][%d] deleteEmailAccountOK ", 200)
}

func (o *DeleteEmailAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEmailAccountNotFound creates a DeleteEmailAccountNotFound with default headers values
func NewDeleteEmailAccountNotFound() *DeleteEmailAccountNotFound {
	return &DeleteEmailAccountNotFound{}
}

/* DeleteEmailAccountNotFound describes a response with status code 404, with default header values.

Account not found
*/
type DeleteEmailAccountNotFound struct {
}

func (o *DeleteEmailAccountNotFound) Error() string {
	return fmt.Sprintf("[DELETE /email/{emailAddress}][%d] deleteEmailAccountNotFound ", 404)
}

func (o *DeleteEmailAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEmailAccountMethodNotAllowed creates a DeleteEmailAccountMethodNotAllowed with default headers values
func NewDeleteEmailAccountMethodNotAllowed() *DeleteEmailAccountMethodNotAllowed {
	return &DeleteEmailAccountMethodNotAllowed{}
}

/* DeleteEmailAccountMethodNotAllowed describes a response with status code 405, with default header values.

Invalid input
*/
type DeleteEmailAccountMethodNotAllowed struct {
}

func (o *DeleteEmailAccountMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /email/{emailAddress}][%d] deleteEmailAccountMethodNotAllowed ", 405)
}

func (o *DeleteEmailAccountMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEmailAccountInternalServerError creates a DeleteEmailAccountInternalServerError with default headers values
func NewDeleteEmailAccountInternalServerError() *DeleteEmailAccountInternalServerError {
	return &DeleteEmailAccountInternalServerError{}
}

/* DeleteEmailAccountInternalServerError describes a response with status code 500, with default header values.

Internal error
*/
type DeleteEmailAccountInternalServerError struct {
}

func (o *DeleteEmailAccountInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /email/{emailAddress}][%d] deleteEmailAccountInternalServerError ", 500)
}

func (o *DeleteEmailAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}