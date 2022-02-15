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

func TestNewAddressBlockListNextAvailableIPParams(t *testing.T) {
	p := NewAddressBlockListNextAvailableIPParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAddressBlockListNextAvailableIPParamsWithTimeout(t *testing.T) {
	p := NewAddressBlockListNextAvailableIPParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAddressBlockListNextAvailableIPParamsWithContext(t *testing.T) {
	p := NewAddressBlockListNextAvailableIPParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAddressBlockListNextAvailableIPParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAddressBlockListNextAvailableIPParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressBlockListNextAvailableIPParams_WithDefaults(t *testing.T) {
	p := NewAddressBlockListNextAvailableIPParams()
	p = p.WithDefaults()
}

func TestAddressBlockListNextAvailableIPParams_WithTimeout(t *testing.T) {
	p := NewAddressBlockListNextAvailableIPParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAddressBlockListNextAvailableIPParams_WithContext(t *testing.T) {
	p := NewAddressBlockListNextAvailableIPParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAddressBlockListNextAvailableIPParams_WithHTTPClient(t *testing.T) {
	p := NewAddressBlockListNextAvailableIPParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressBlockListNextAvailableIPParams_WithID(t *testing.T) {
	p := NewAddressBlockListNextAvailableIPParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
