package address_block

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewAddressBlockCreateNextAvailableSubnetParams(t *testing.T) {
	p := NewAddressBlockCreateNextAvailableSubnetParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAddressBlockCreateNextAvailableSubnetParamsWithTimeout(t *testing.T) {
	p := NewAddressBlockCreateNextAvailableSubnetParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAddressBlockCreateNextAvailableSubnetParamsWithContext(t *testing.T) {
	p := NewAddressBlockCreateNextAvailableSubnetParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAddressBlockCreateNextAvailableSubnetParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAddressBlockCreateNextAvailableSubnetParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressBlockCreateNextAvailableSubnetParams_WithDefaults(t *testing.T) {
	p := NewAddressBlockCreateNextAvailableSubnetParams()
	p = p.WithDefaults()
}

func TestAddressBlockCreateNextAvailableSubnetParams_WithTimeout(t *testing.T) {
	p := NewAddressBlockCreateNextAvailableSubnetParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAddressBlockCreateNextAvailableSubnetParams_WithContext(t *testing.T) {
	p := NewAddressBlockCreateNextAvailableSubnetParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAddressBlockCreateNextAvailableSubnetParams_WithHTTPClient(t *testing.T) {
	p := NewAddressBlockCreateNextAvailableSubnetParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressBlockCreateNextAvailableSubnetParams_WithID(t *testing.T) {
	p := NewAddressBlockCreateNextAvailableSubnetParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
