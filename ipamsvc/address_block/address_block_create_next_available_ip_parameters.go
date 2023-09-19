// Code generated by go-swagger; DO NOT EDIT.

package address_block

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
)

// NewAddressBlockCreateNextAvailableIPParams creates a new AddressBlockCreateNextAvailableIPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddressBlockCreateNextAvailableIPParams() *AddressBlockCreateNextAvailableIPParams {
	return &AddressBlockCreateNextAvailableIPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddressBlockCreateNextAvailableIPParamsWithTimeout creates a new AddressBlockCreateNextAvailableIPParams object
// with the ability to set a timeout on a request.
func NewAddressBlockCreateNextAvailableIPParamsWithTimeout(timeout time.Duration) *AddressBlockCreateNextAvailableIPParams {
	return &AddressBlockCreateNextAvailableIPParams{
		timeout: timeout,
	}
}

// NewAddressBlockCreateNextAvailableIPParamsWithContext creates a new AddressBlockCreateNextAvailableIPParams object
// with the ability to set a context for a request.
func NewAddressBlockCreateNextAvailableIPParamsWithContext(ctx context.Context) *AddressBlockCreateNextAvailableIPParams {
	return &AddressBlockCreateNextAvailableIPParams{
		Context: ctx,
	}
}

// NewAddressBlockCreateNextAvailableIPParamsWithHTTPClient creates a new AddressBlockCreateNextAvailableIPParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddressBlockCreateNextAvailableIPParamsWithHTTPClient(client *http.Client) *AddressBlockCreateNextAvailableIPParams {
	return &AddressBlockCreateNextAvailableIPParams{
		HTTPClient: client,
	}
}

/*
AddressBlockCreateNextAvailableIPParams contains all the parameters to send to the API endpoint

	for the address block create next available IP operation.

	Typically these are written to a http.Request.
*/
type AddressBlockCreateNextAvailableIPParams struct {

	/* ID.

	   An application specific resource identity of a resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the address block create next available IP params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddressBlockCreateNextAvailableIPParams) WithDefaults() *AddressBlockCreateNextAvailableIPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the address block create next available IP params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddressBlockCreateNextAvailableIPParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the address block create next available IP params
func (o *AddressBlockCreateNextAvailableIPParams) WithTimeout(timeout time.Duration) *AddressBlockCreateNextAvailableIPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the address block create next available IP params
func (o *AddressBlockCreateNextAvailableIPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the address block create next available IP params
func (o *AddressBlockCreateNextAvailableIPParams) WithContext(ctx context.Context) *AddressBlockCreateNextAvailableIPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the address block create next available IP params
func (o *AddressBlockCreateNextAvailableIPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the address block create next available IP params
func (o *AddressBlockCreateNextAvailableIPParams) WithHTTPClient(client *http.Client) *AddressBlockCreateNextAvailableIPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the address block create next available IP params
func (o *AddressBlockCreateNextAvailableIPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the address block create next available IP params
func (o *AddressBlockCreateNextAvailableIPParams) WithID(id string) *AddressBlockCreateNextAvailableIPParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the address block create next available IP params
func (o *AddressBlockCreateNextAvailableIPParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AddressBlockCreateNextAvailableIPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
