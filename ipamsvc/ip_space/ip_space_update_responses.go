// Code generated by go-swagger; DO NOT EDIT.

package ip_space

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/infobloxopen/b1ddi-go-client/models"
	b1cliruntime "github.com/infobloxopen/b1ddi-go-client/runtime"
)

// IPSpaceUpdateReader is a Reader for the IPSpaceUpdate structure.
type IPSpaceUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPSpaceUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates client error", response.Body(), response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates server error", response.Body(), response.Code())
	}

	result := NewIPSpaceUpdateOK()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewIPSpaceUpdateOK creates a IPSpaceUpdateOK with default headers values
func NewIPSpaceUpdateOK() *IPSpaceUpdateOK {
	return &IPSpaceUpdateOK{}
}

/*
IPSpaceUpdateOK describes a response with status code 200, with default header values.

PATCH operation response
*/
type IPSpaceUpdateOK struct {
	Payload *models.IpamsvcUpdateIPSpaceResponse
}

func (o *IPSpaceUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/ip_space/{id}][%d] ipSpaceUpdateOK  %+v", 200, o.Payload)
}
func (o *IPSpaceUpdateOK) GetPayload() *models.IpamsvcUpdateIPSpaceResponse {
	return o.Payload
}

func (o *IPSpaceUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpamsvcUpdateIPSpaceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
