// Code generated by go-swagger; DO NOT EDIT.

package view

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

// ViewListReader is a Reader for the ViewList structure.
type ViewListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ViewListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates client error", response.Body(), response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates server error", response.Body(), response.Code())
	}

	result := NewViewListOK()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewViewListOK creates a ViewListOK with default headers values
func NewViewListOK() *ViewListOK {
	return &ViewListOK{}
}

/* ViewListOK describes a response with status code 200, with default header values.

GET operation response
*/
type ViewListOK struct {
	Payload *models.ConfigListViewResponse
}

func (o *ViewListOK) Error() string {
	return fmt.Sprintf("[GET /dns/view][%d] viewListOK  %+v", 200, o.Payload)
}
func (o *ViewListOK) GetPayload() *models.ConfigListViewResponse {
	return o.Payload
}

func (o *ViewListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigListViewResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
