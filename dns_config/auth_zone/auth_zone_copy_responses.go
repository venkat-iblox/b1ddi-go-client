// Code generated by go-swagger; DO NOT EDIT.

package auth_zone

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

// AuthZoneCopyReader is a Reader for the AuthZoneCopy structure.
type AuthZoneCopyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthZoneCopyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates client error", response.Body(), response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates server error", response.Body(), response.Code())
	}

	result := NewAuthZoneCopyOK()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewAuthZoneCopyOK creates a AuthZoneCopyOK with default headers values
func NewAuthZoneCopyOK() *AuthZoneCopyOK {
	return &AuthZoneCopyOK{}
}

/*
AuthZoneCopyOK describes a response with status code 200, with default header values.

POST operation response
*/
type AuthZoneCopyOK struct {
	Payload *models.ConfigCopyAuthZoneResponse
}

func (o *AuthZoneCopyOK) Error() string {
	return fmt.Sprintf("[POST /dns/auth_zone/copy][%d] authZoneCopyOK  %+v", 200, o.Payload)
}
func (o *AuthZoneCopyOK) GetPayload() *models.ConfigCopyAuthZoneResponse {
	return o.Payload
}

func (o *AuthZoneCopyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigCopyAuthZoneResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
