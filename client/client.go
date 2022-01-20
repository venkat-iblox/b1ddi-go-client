package client

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/infobloxopen/b1ddi-go-client/dns_config"
	"github.com/infobloxopen/b1ddi-go-client/dns_data"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc"
)

type Client struct {
	IPAddressManagementAPI *ipamsvc.IPAddressManagementAPI
	DNSConfigurationAPI    *dns_config.DNSConfigurationAPI
	DNSDataAPI             *dns_data.DNSDataAPI
}

func NewClient(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{
		IPAddressManagementAPI: ipamsvc.New(transport, formats),
		DNSConfigurationAPI:    dns_config.New(transport, formats),
		DNSDataAPI:             dns_data.New(transport, formats),
	}
}

// B1DDIToken provides a header for the BloxOne DDI token authentication
func B1DDIToken(token string) runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		return r.SetHeaderParam("Authorization", "Token "+token)
	})
}
