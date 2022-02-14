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

func TestNewAddressBlockReadParams(t *testing.T) {
	p := NewAddressBlockReadParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAddressBlockReadParamsWithTimeout(t *testing.T) {
	p := NewAddressBlockReadParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAddressBlockReadParamsWithContext(t *testing.T) {
	p := NewAddressBlockReadParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAddressBlockReadParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAddressBlockReadParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressBlockReadParams_WithDefaults(t *testing.T) {
	p := NewAddressBlockReadParams()
	p = p.WithDefaults()
}

func TestAddressBlockReadParams_WithTimeout(t *testing.T) {
	p := NewAddressBlockReadParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAddressBlockReadParams_WithContext(t *testing.T) {
	p := NewAddressBlockReadParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAddressBlockReadParams_WithHTTPClient(t *testing.T) {
	p := NewAddressBlockReadParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressBlockReadParams_WithID(t *testing.T) {
	p := NewAddressBlockReadParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
