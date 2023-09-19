// Code generated by go-swagger; DO NOT EDIT.

package fixed_address

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	b1cliruntime "github.com/infobloxopen/b1ddi-go-client/runtime"
)

// FixedAddressDeleteReader is a Reader for the FixedAddressDelete structure.
type FixedAddressDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FixedAddressDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates client error", response.Body(), response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates server error", response.Body(), response.Code())
	}

	result := NewFixedAddressDeleteOK()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewFixedAddressDeleteOK creates a FixedAddressDeleteOK with default headers values
func NewFixedAddressDeleteOK() *FixedAddressDeleteOK {
	return &FixedAddressDeleteOK{}
}

/*
FixedAddressDeleteOK describes a response with status code 200, with default header values.

OK
*/
type FixedAddressDeleteOK struct {
}

func (o *FixedAddressDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /dhcp/fixed_address/{id}][%d] fixedAddressDeleteOK ", 200)
}

func (o *FixedAddressDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
