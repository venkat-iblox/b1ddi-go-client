// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	b1cliruntime "github.com/infobloxopen/b1ddi-go-client/runtime"
)

// New creates a new server API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for server API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ServerCreate(params *ServerCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServerCreateOK, error)

	ServerDelete(params *ServerDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServerDeleteOK, error)

	ServerList(params *ServerListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServerListOK, error)

	ServerRead(params *ServerReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServerReadOK, error)

	ServerUpdate(params *ServerUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServerUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	ServerCreate creates the server object

	Use this method to create a Server object.

A DNS Config Profile is a named configuration profile that can be shared for specified list of hosts.
*/
func (a *Client) ServerCreate(params *ServerCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServerCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServerCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "serverCreate",
		Method:             "POST",
		PathPattern:        "/dns/server",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ServerCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ServerCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for serverCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	ServerDelete moves the server object to recyclebin

	Use this method to move a Server object to Recyclebin.

A DNS Config Profile is a named configuration profile that can be shared for specified list of hosts.
*/
func (a *Client) ServerDelete(params *ServerDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServerDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServerDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "serverDelete",
		Method:             "DELETE",
		PathPattern:        "/dns/server/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ServerDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	params.ID = b1cliruntime.TrimIDPrefix(op.PathPattern, params.ID)
	op.Params = params

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ServerDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for serverDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	ServerList lists server objects

	Use this method to list Server objects.

A DNS Config Profile is a named configuration profile that can be shared for specified list of hosts.
*/
func (a *Client) ServerList(params *ServerListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServerListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServerListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "serverList",
		Method:             "GET",
		PathPattern:        "/dns/server",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ServerListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ServerListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for serverList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	ServerRead reads the server object

	Use this method to read a Server object.

A DNS Config Profile is a named configuration profile that can be shared for specified list of hosts.
*/
func (a *Client) ServerRead(params *ServerReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServerReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServerReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "serverRead",
		Method:             "GET",
		PathPattern:        "/dns/server/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ServerReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	params.ID = b1cliruntime.TrimIDPrefix(op.PathPattern, params.ID)
	op.Params = params

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ServerReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for serverRead: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	ServerUpdate updates the server object

	Use this method to update a Server object.

A DNS Config Profile is a named configuration profile that can be shared for specified list of hosts.
*/
func (a *Client) ServerUpdate(params *ServerUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServerUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServerUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "serverUpdate",
		Method:             "PATCH",
		PathPattern:        "/dns/server/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ServerUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	params.ID = b1cliruntime.TrimIDPrefix(op.PathPattern, params.ID)
	op.Params = params

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ServerUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for serverUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
