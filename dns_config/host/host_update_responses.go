// Code generated by go-swagger; DO NOT EDIT.

package host

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

// HostUpdateReader is a Reader for the HostUpdate structure.
type HostUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HostUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates client error", response.Body(), response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates server error", response.Body(), response.Code())
	}

	result := NewHostUpdateOK()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewHostUpdateOK creates a HostUpdateOK with default headers values
func NewHostUpdateOK() *HostUpdateOK {
	return &HostUpdateOK{}
}

/*
HostUpdateOK describes a response with status code 200, with default header values.

PATCH operation response
*/
type HostUpdateOK struct {
	Payload *models.ConfigUpdateHostResponse
}

func (o *HostUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dns/host/{id}][%d] hostUpdateOK  %+v", 200, o.Payload)
}
func (o *HostUpdateOK) GetPayload() *models.ConfigUpdateHostResponse {
	return o.Payload
}

func (o *HostUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigUpdateHostResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
