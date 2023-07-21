// Package client provides useful primitives for working with BloxOne DDI APIs
package client

import (
	"crypto/tls"
	"fmt"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/infobloxopen/b1ddi-go-client/dns_config"
	"github.com/infobloxopen/b1ddi-go-client/dns_data"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc"
	"net/http"
)

// Client is an aggregation of different BloxOne DDI API clients.
type Client struct {
	IPAddressManagementAPI *ipamsvc.IPAddressManagementAPI
	DNSConfigurationAPI    *dns_config.DNSConfigurationAPI
	DNSDataAPI             *dns_data.DNSDataAPI
}

// NewClient creates a new BloxOne DDI API Client.
func NewClient(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{
		IPAddressManagementAPI: ipamsvc.New(transport, formats),
		DNSConfigurationAPI:    dns_config.New(transport, formats),
		DNSDataAPI:             dns_data.New(transport, formats),
	}
}

// BloxOneAPIKey provides a header for the BloxOne DDI API authentication.
//
// See https://docs.infoblox.com/display/BloxOneDDI/BloxOne+DDI+API+Guide learn how to get the API key.
func BloxOneAPIKey(apiKey string) runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		return r.SetHeaderParam("Authorization", "Token "+apiKey)
	})
}

func NewTransport(headers map[string]string) *customTransport {
	transport := http.DefaultTransport
	transport.(*http.Transport).TLSClientConfig = &tls.Config{
		MaxVersion: tls.VersionTLS12,
		MinVersion: tls.VersionTLS12,
	}
	return &customTransport{
		originalTransport: transport,
		headers:           headers,
	}
}

type customTransport struct {
	originalTransport http.RoundTripper
	headers           map[string]string
}

func (c *customTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	for key, val := range c.headers {
		r.Header.Add(key, val)
	}
	r.Header.Add("x-infoblox-sdk", fmt.Sprintf("golang-sdk/v%s", "1.0.0"))

	resp, err := c.originalTransport.RoundTrip(r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
