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

func TestNewAddressBlockCreateNextAvailableIPParams(t *testing.T) {
	p := NewAddressBlockCreateNextAvailableIPParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAddressBlockCreateNextAvailableIPParamsWithTimeout(t *testing.T) {
	p := NewAddressBlockCreateNextAvailableIPParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAddressBlockCreateNextAvailableIPParamsWithContext(t *testing.T) {
	p := NewAddressBlockCreateNextAvailableIPParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAddressBlockCreateNextAvailableIPParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAddressBlockCreateNextAvailableIPParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressBlockCreateNextAvailableIPParams_WithDefaults(t *testing.T) {
	p := NewAddressBlockCreateNextAvailableIPParams()
	p = p.WithDefaults()
}

func TestAddressBlockCreateNextAvailableIPParams_WithTimeout(t *testing.T) {
	p := NewAddressBlockCreateNextAvailableIPParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAddressBlockCreateNextAvailableIPParams_WithContext(t *testing.T) {
	p := NewAddressBlockCreateNextAvailableIPParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAddressBlockCreateNextAvailableIPParams_WithHTTPClient(t *testing.T) {
	p := NewAddressBlockCreateNextAvailableIPParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressBlockCreateNextAvailableIPParams_WithID(t *testing.T) {
	p := NewAddressBlockCreateNextAvailableIPParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
