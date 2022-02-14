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

func TestNewAddressBlockDeleteParams(t *testing.T) {
	p := NewAddressBlockDeleteParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAddressBlockDeleteParamsWithTimeout(t *testing.T) {
	p := NewAddressBlockDeleteParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAddressBlockDeleteParamsWithContext(t *testing.T) {
	p := NewAddressBlockDeleteParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAddressBlockDeleteParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAddressBlockDeleteParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressBlockDeleteParams_WithDefaults(t *testing.T) {
	p := NewAddressBlockDeleteParams()
	p = p.WithDefaults()
}

func TestAddressBlockDeleteParams_WithTimeout(t *testing.T) {
	p := NewAddressBlockDeleteParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAddressBlockDeleteParams_WithContext(t *testing.T) {
	p := NewAddressBlockDeleteParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAddressBlockDeleteParams_WithHTTPClient(t *testing.T) {
	p := NewAddressBlockDeleteParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressBlockDeleteParams_WithID(t *testing.T) {
	p := NewAddressBlockDeleteParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
