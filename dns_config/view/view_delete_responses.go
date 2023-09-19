// Code generated by go-swagger; DO NOT EDIT.

package view

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	b1cliruntime "github.com/infobloxopen/b1ddi-go-client/runtime"
)

// ViewDeleteReader is a Reader for the ViewDelete structure.
type ViewDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ViewDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates client error", response.Body(), response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates server error", response.Body(), response.Code())
	}

	result := NewViewDeleteOK()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewViewDeleteOK creates a ViewDeleteOK with default headers values
func NewViewDeleteOK() *ViewDeleteOK {
	return &ViewDeleteOK{}
}

/*
ViewDeleteOK describes a response with status code 200, with default header values.

OK
*/
type ViewDeleteOK struct {
}

func (o *ViewDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /dns/view/{id}][%d] viewDeleteOK ", 200)
}

func (o *ViewDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
