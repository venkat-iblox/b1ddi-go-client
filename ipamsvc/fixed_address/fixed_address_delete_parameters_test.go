package fixed_address

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewFixedAddressDeleteParams(t *testing.T) {
	p := NewFixedAddressDeleteParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewFixedAddressDeleteParamsWithTimeout(t *testing.T) {
	p := NewFixedAddressDeleteParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewFixedAddressDeleteParamsWithContext(t *testing.T) {
	p := NewFixedAddressDeleteParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewFixedAddressDeleteParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewFixedAddressDeleteParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestFixedAddressDeleteParams_WithDefaults(t *testing.T) {
	p := NewFixedAddressDeleteParams()
	p = p.WithDefaults()
}

func TestFixedAddressDeleteParams_WithTimeout(t *testing.T) {
	p := NewFixedAddressDeleteParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestFixedAddressDeleteParams_WithContext(t *testing.T) {
	p := NewFixedAddressDeleteParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestFixedAddressDeleteParams_WithHTTPClient(t *testing.T) {
	p := NewFixedAddressDeleteParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestFixedAddressDeleteParams_WithID(t *testing.T) {
	p := NewFixedAddressDeleteParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
