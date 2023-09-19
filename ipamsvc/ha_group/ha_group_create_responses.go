// Code generated by go-swagger; DO NOT EDIT.

package ha_group

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

// HaGroupCreateReader is a Reader for the HaGroupCreate structure.
type HaGroupCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HaGroupCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates client error", response.Body(), response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates server error", response.Body(), response.Code())
	}

	result := NewHaGroupCreateOK()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewHaGroupCreateOK creates a HaGroupCreateOK with default headers values
func NewHaGroupCreateOK() *HaGroupCreateOK {
	return &HaGroupCreateOK{}
}

/*
HaGroupCreateOK describes a response with status code 200, with default header values.

POST operation response
*/
type HaGroupCreateOK struct {
	Payload *models.IpamsvcCreateHAGroupResponse
}

func (o *HaGroupCreateOK) Error() string {
	return fmt.Sprintf("[POST /dhcp/ha_group][%d] haGroupCreateOK  %+v", 200, o.Payload)
}
func (o *HaGroupCreateOK) GetPayload() *models.IpamsvcCreateHAGroupResponse {
	return o.Payload
}

func (o *HaGroupCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpamsvcCreateHAGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
