package address_block

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewAddressBlockListNextAvailableSubnetParams(t *testing.T) {
	p := NewAddressBlockListNextAvailableSubnetParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAddressBlockListNextAvailableSubnetParamsWithTimeout(t *testing.T) {
	p := NewAddressBlockListNextAvailableSubnetParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAddressBlockListNextAvailableSubnetParamsWithContext(t *testing.T) {
	p := NewAddressBlockListNextAvailableSubnetParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAddressBlockListNextAvailableSubnetParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAddressBlockListNextAvailableSubnetParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressBlockListNextAvailableSubnetParams_WithDefaults(t *testing.T) {
	p := NewAddressBlockListNextAvailableSubnetParams()
	p = p.WithDefaults()
}

func TestAddressBlockListNextAvailableSubnetParams_WithTimeout(t *testing.T) {
	p := NewAddressBlockListNextAvailableSubnetParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAddressBlockListNextAvailableSubnetParams_WithContext(t *testing.T) {
	p := NewAddressBlockListNextAvailableSubnetParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAddressBlockListNextAvailableSubnetParams_WithHTTPClient(t *testing.T) {
	p := NewAddressBlockListNextAvailableSubnetParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressBlockListNextAvailableSubnetParams_WithID(t *testing.T) {
	p := NewAddressBlockListNextAvailableSubnetParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}

func TestAddressBlockListNextAvailableSubnetParams_WithName(t *testing.T) {
	p := NewAddressBlockListNextAvailableSubnetParams()
	p = p.WithName(swag.String("name"))
	require.NotNil(t, p.Name)
	assert.Equal(t, "name", *p.Name)
}
