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

// RangeCreateReader is a Reader for the RangeCreate structure.
type RangeCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RangeCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates client error", response.Body(), response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates server error", response.Body(), response.Code())
	}

	result := NewRangeCreateCreated()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewRangeCreateCreated creates a RangeCreateCreated with default headers values
func NewRangeCreateCreated() *RangeCreateCreated {
	return &RangeCreateCreated{}
}

/*
RangeCreateCreated describes a response with status code 201, with default header values.

POST operation response
*/
type RangeCreateCreated struct {
	Payload *models.IpamsvcCreateRangeResponse
}

func (o *RangeCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/range][%d] rangeCreateCreated  %+v", 201, o.Payload)
}
func (o *RangeCreateCreated) GetPayload() *models.IpamsvcCreateRangeResponse {
	return o.Payload
}

func (o *RangeCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpamsvcCreateRangeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
