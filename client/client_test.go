package client

import (
	"fmt"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClient(t *testing.T) {
	// create the transport
	transport := httptransport.New(
		"", "api/ddi/v1", nil,
	)

	cli := NewClient(transport, strfmt.Default)

	assert.NotNil(t, cli)
	assert.NotNil(t, cli.IPAddressManagementAPI)
	assert.NotNil(t, cli.DNSConfigurationAPI)
	assert.NotNil(t, cli.DNSDataAPI)
}

func TestBLOXONEAPIKey(t *testing.T) {
	key := "KxPlKBne1NtqkcgaUgiU29cME9Y0I13JtZ8QPyDLTalv0yXLriJZgXF4lDXYV31Ky"

	authFunc := BLOXONEAPIKey(key)

	req := &runtime.TestClientRequest{}

	err := authFunc.AuthenticateRequest(req, strfmt.Default)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(
		t,
		fmt.Sprintf("Token %s", key),
		req.Headers["Authorization"][0],
	)
}
