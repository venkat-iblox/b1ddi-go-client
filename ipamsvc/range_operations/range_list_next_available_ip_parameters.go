// Code generated by go-swagger; DO NOT EDIT.

package range_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewRangeListNextAvailableIPParams creates a new RangeListNextAvailableIPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRangeListNextAvailableIPParams() *RangeListNextAvailableIPParams {
	return &RangeListNextAvailableIPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRangeListNextAvailableIPParamsWithTimeout creates a new RangeListNextAvailableIPParams object
// with the ability to set a timeout on a request.
func NewRangeListNextAvailableIPParamsWithTimeout(timeout time.Duration) *RangeListNextAvailableIPParams {
	return &RangeListNextAvailableIPParams{
		timeout: timeout,
	}
}

// NewRangeListNextAvailableIPParamsWithContext creates a new RangeListNextAvailableIPParams object
// with the ability to set a context for a request.
func NewRangeListNextAvailableIPParamsWithContext(ctx context.Context) *RangeListNextAvailableIPParams {
	return &RangeListNextAvailableIPParams{
		Context: ctx,
	}
}

// NewRangeListNextAvailableIPParamsWithHTTPClient creates a new RangeListNextAvailableIPParams object
// with the ability to set a custom HTTPClient for a request.
func NewRangeListNextAvailableIPParamsWithHTTPClient(client *http.Client) *RangeListNextAvailableIPParams {
	return &RangeListNextAvailableIPParams{
		HTTPClient: client,
	}
}

/*
RangeListNextAvailableIPParams contains all the parameters to send to the API endpoint

	for the range list next available IP operation.

	Typically these are written to a http.Request.
*/
type RangeListNextAvailableIPParams struct {

	/* Contiguous.

	     Indicates whether the IP addresses should belong to a contiguous block.

	Defaults to _false_.

	     Format: boolean
	*/
	Contiguous *bool

	/* Count.

	     The number of IP addresses requested.

	Defaults to 1.

	     Format: int32
	*/
	Count *int32

	/* ID.

	   An application specific resource identity of a resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the range list next available IP params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RangeListNextAvailableIPParams) WithDefaults() *RangeListNextAvailableIPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the range list next available IP params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RangeListNextAvailableIPParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the range list next available IP params
func (o *RangeListNextAvailableIPParams) WithTimeout(timeout time.Duration) *RangeListNextAvailableIPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the range list next available IP params
func (o *RangeListNextAvailableIPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the range list next available IP params
func (o *RangeListNextAvailableIPParams) WithContext(ctx context.Context) *RangeListNextAvailableIPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the range list next available IP params
func (o *RangeListNextAvailableIPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the range list next available IP params
func (o *RangeListNextAvailableIPParams) WithHTTPClient(client *http.Client) *RangeListNextAvailableIPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the range list next available IP params
func (o *RangeListNextAvailableIPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContiguous adds the contiguous to the range list next available IP params
func (o *RangeListNextAvailableIPParams) WithContiguous(contiguous *bool) *RangeListNextAvailableIPParams {
	o.SetContiguous(contiguous)
	return o
}

// SetContiguous adds the contiguous to the range list next available IP params
func (o *RangeListNextAvailableIPParams) SetContiguous(contiguous *bool) {
	o.Contiguous = contiguous
}

// WithCount adds the count to the range list next available IP params
func (o *RangeListNextAvailableIPParams) WithCount(count *int32) *RangeListNextAvailableIPParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the range list next available IP params
func (o *RangeListNextAvailableIPParams) SetCount(count *int32) {
	o.Count = count
}

// WithID adds the id to the range list next available IP params
func (o *RangeListNextAvailableIPParams) WithID(id string) *RangeListNextAvailableIPParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the range list next available IP params
func (o *RangeListNextAvailableIPParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RangeListNextAvailableIPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Contiguous != nil {

		// query param contiguous
		var qrContiguous bool

		if o.Contiguous != nil {
			qrContiguous = *o.Contiguous
		}
		qContiguous := swag.FormatBool(qrContiguous)
		if qContiguous != "" {

			if err := r.SetQueryParam("contiguous", qContiguous); err != nil {
				return err
			}
		}
	}

	if o.Count != nil {

		// query param count
		var qrCount int32

		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt32(qrCount)
		if qCount != "" {

			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
