// Code generated by go-swagger; DO NOT EDIT.

package address_block

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

// AddressBlockCreateNextAvailableIPReader is a Reader for the AddressBlockCreateNextAvailableIP structure.
type AddressBlockCreateNextAvailableIPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddressBlockCreateNextAvailableIPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates client error", response.Body(), response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates server error", response.Body(), response.Code())
	}

	result := NewAddressBlockCreateNextAvailableIPOK()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewAddressBlockCreateNextAvailableIPOK creates a AddressBlockCreateNextAvailableIPOK with default headers values
func NewAddressBlockCreateNextAvailableIPOK() *AddressBlockCreateNextAvailableIPOK {
	return &AddressBlockCreateNextAvailableIPOK{}
}

/*
AddressBlockCreateNextAvailableIPOK describes a response with status code 200, with default header values.

POST operation response
*/
type AddressBlockCreateNextAvailableIPOK struct {
	Payload *models.IpamsvcCreateNextAvailableIPResponse
}

func (o *AddressBlockCreateNextAvailableIPOK) Error() string {
	return fmt.Sprintf("[POST /ipam/address_block/{id}/nextavailableip][%d] addressBlockCreateNextAvailableIpOK  %+v", 200, o.Payload)
}
func (o *AddressBlockCreateNextAvailableIPOK) GetPayload() *models.IpamsvcCreateNextAvailableIPResponse {
	return o.Payload
}

func (o *AddressBlockCreateNextAvailableIPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpamsvcCreateNextAvailableIPResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
