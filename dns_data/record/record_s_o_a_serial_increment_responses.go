// Code generated by go-swagger; DO NOT EDIT.

package record

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

// RecordSOASerialIncrementReader is a Reader for the RecordSOASerialIncrement structure.
type RecordSOASerialIncrementReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RecordSOASerialIncrementReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	if response.Code() >= 400 && response.Code() < 500 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates client error", response.Body(), response.Code())
	}

	if response.Code() >= 500 && response.Code() < 600 {
		return nil, b1cliruntime.NewAPIHTTPError("response status code indicates server error", response.Body(), response.Code())
	}

	result := NewRecordSOASerialIncrementCreated()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil
}

// NewRecordSOASerialIncrementCreated creates a RecordSOASerialIncrementCreated with default headers values
func NewRecordSOASerialIncrementCreated() *RecordSOASerialIncrementCreated {
	return &RecordSOASerialIncrementCreated{}
}

/* RecordSOASerialIncrementCreated describes a response with status code 201, with default header values.

POST operation response
*/
type RecordSOASerialIncrementCreated struct {
	Payload *models.DataSOASerialIncrementResponse
}

func (o *RecordSOASerialIncrementCreated) Error() string {
	return fmt.Sprintf("[POST /dns/record/{id}/serial_increment][%d] recordSOASerialIncrementCreated  %+v", 201, o.Payload)
}
func (o *RecordSOASerialIncrementCreated) GetPayload() *models.DataSOASerialIncrementResponse {
	return o.Payload
}

func (o *RecordSOASerialIncrementCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataSOASerialIncrementResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}