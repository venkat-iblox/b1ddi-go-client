// Code generated by go-swagger; DO NOT EDIT.

package range_operations

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

// RangeReadReader is a Reader for the RangeRead structure.
type RangeReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RangeReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates client error", response.Body(), response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates server error", response.Body(), response.Code())
	}

	result := NewRangeReadOK()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewRangeReadOK creates a RangeReadOK with default headers values
func NewRangeReadOK() *RangeReadOK {
	return &RangeReadOK{}
}

/* RangeReadOK describes a response with status code 200, with default header values.

GET operation response
*/
type RangeReadOK struct {
	Payload *models.IpamsvcReadRangeResponse
}

func (o *RangeReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/range/{id}][%d] rangeReadOK  %+v", 200, o.Payload)
}
func (o *RangeReadOK) GetPayload() *models.IpamsvcReadRangeResponse {
	return o.Payload
}

func (o *RangeReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpamsvcReadRangeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
