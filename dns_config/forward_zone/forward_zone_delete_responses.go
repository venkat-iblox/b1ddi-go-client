// Code generated by go-swagger; DO NOT EDIT.

package forward_zone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	b1cliruntime "github.com/infobloxopen/b1ddi-go-client/runtime"
)

// ForwardZoneDeleteReader is a Reader for the ForwardZoneDelete structure.
type ForwardZoneDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ForwardZoneDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates client error", response.Body(), response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates server error", response.Body(), response.Code())
	}

	result := NewForwardZoneDeleteOK()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewForwardZoneDeleteOK creates a ForwardZoneDeleteOK with default headers values
func NewForwardZoneDeleteOK() *ForwardZoneDeleteOK {
	return &ForwardZoneDeleteOK{}
}

/*
ForwardZoneDeleteOK describes a response with status code 200, with default header values.

OK
*/
type ForwardZoneDeleteOK struct {
}

func (o *ForwardZoneDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /dns/forward_zone/{id}][%d] forwardZoneDeleteOK ", 200)
}

func (o *ForwardZoneDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
