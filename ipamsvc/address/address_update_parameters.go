// Code generated by go-swagger; DO NOT EDIT.

package address

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

	"github.com/infobloxopen/b1ddi-go-client/models"
)

// NewAddressUpdateParams creates a new AddressUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddressUpdateParams() *AddressUpdateParams {
	return &AddressUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddressUpdateParamsWithTimeout creates a new AddressUpdateParams object
// with the ability to set a timeout on a request.
func NewAddressUpdateParamsWithTimeout(timeout time.Duration) *AddressUpdateParams {
	return &AddressUpdateParams{
		timeout: timeout,
	}
}

// NewAddressUpdateParamsWithContext creates a new AddressUpdateParams object
// with the ability to set a context for a request.
func NewAddressUpdateParamsWithContext(ctx context.Context) *AddressUpdateParams {
	return &AddressUpdateParams{
		Context: ctx,
	}
}

// NewAddressUpdateParamsWithHTTPClient creates a new AddressUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddressUpdateParamsWithHTTPClient(client *http.Client) *AddressUpdateParams {
	return &AddressUpdateParams{
		HTTPClient: client,
	}
}

/*
AddressUpdateParams contains all the parameters to send to the API endpoint

	for the address update operation.

	Typically these are written to a http.Request.
*/
type AddressUpdateParams struct {

	// Body.
	Body *models.IpamsvcAddress

	/* ID.

	   An application specific resource identity of a resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the address update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddressUpdateParams) WithDefaults() *AddressUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the address update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddressUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the address update params
func (o *AddressUpdateParams) WithTimeout(timeout time.Duration) *AddressUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the address update params
func (o *AddressUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the address update params
func (o *AddressUpdateParams) WithContext(ctx context.Context) *AddressUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the address update params
func (o *AddressUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the address update params
func (o *AddressUpdateParams) WithHTTPClient(client *http.Client) *AddressUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the address update params
func (o *AddressUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the address update params
func (o *AddressUpdateParams) WithBody(body *models.IpamsvcAddress) *AddressUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the address update params
func (o *AddressUpdateParams) SetBody(body *models.IpamsvcAddress) {
	o.Body = body
}

// WithID adds the id to the address update params
func (o *AddressUpdateParams) WithID(id string) *AddressUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the address update params
func (o *AddressUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AddressUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
