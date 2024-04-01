// Code generated by go-swagger; DO NOT EDIT.

package email

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteEmailAliasReader is a Reader for the DeleteEmailAlias structure.
type DeleteEmailAliasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEmailAliasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteEmailAliasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteEmailAliasNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewDeleteEmailAliasMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteEmailAliasInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteEmailAliasOK creates a DeleteEmailAliasOK with default headers values
func NewDeleteEmailAliasOK() *DeleteEmailAliasOK {
	return &DeleteEmailAliasOK{}
}

/* DeleteEmailAliasOK describes a response with status code 200, with default header values.

Successful operation
*/
type DeleteEmailAliasOK struct {
}

func (o *DeleteEmailAliasOK) Error() string {
	return fmt.Sprintf("[DELETE /email/{emailAddress}/aliasses][%d] deleteEmailAliasOK ", 200)
}

func (o *DeleteEmailAliasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEmailAliasNotFound creates a DeleteEmailAliasNotFound with default headers values
func NewDeleteEmailAliasNotFound() *DeleteEmailAliasNotFound {
	return &DeleteEmailAliasNotFound{}
}

/* DeleteEmailAliasNotFound describes a response with status code 404, with default header values.

Account or alias not found
*/
type DeleteEmailAliasNotFound struct {
}

func (o *DeleteEmailAliasNotFound) Error() string {
	return fmt.Sprintf("[DELETE /email/{emailAddress}/aliasses][%d] deleteEmailAliasNotFound ", 404)
}

func (o *DeleteEmailAliasNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEmailAliasMethodNotAllowed creates a DeleteEmailAliasMethodNotAllowed with default headers values
func NewDeleteEmailAliasMethodNotAllowed() *DeleteEmailAliasMethodNotAllowed {
	return &DeleteEmailAliasMethodNotAllowed{}
}

/* DeleteEmailAliasMethodNotAllowed describes a response with status code 405, with default header values.

Invalid input
*/
type DeleteEmailAliasMethodNotAllowed struct {
}

func (o *DeleteEmailAliasMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /email/{emailAddress}/aliasses][%d] deleteEmailAliasMethodNotAllowed ", 405)
}

func (o *DeleteEmailAliasMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEmailAliasInternalServerError creates a DeleteEmailAliasInternalServerError with default headers values
func NewDeleteEmailAliasInternalServerError() *DeleteEmailAliasInternalServerError {
	return &DeleteEmailAliasInternalServerError{}
}

/* DeleteEmailAliasInternalServerError describes a response with status code 500, with default header values.

Internal error
*/
type DeleteEmailAliasInternalServerError struct {
}

func (o *DeleteEmailAliasInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /email/{emailAddress}/aliasses][%d] deleteEmailAliasInternalServerError ", 500)
}

func (o *DeleteEmailAliasInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
