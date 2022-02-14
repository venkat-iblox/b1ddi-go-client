package address

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewAddressDeleteParams(t *testing.T) {
	p := NewAddressDeleteParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAddressDeleteParamsWithTimeout(t *testing.T) {
	p := NewAddressDeleteParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAddressDeleteParamsWithContext(t *testing.T) {
	p := NewAddressDeleteParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAddressDeleteParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAddressDeleteParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressDeleteParams_WithDefaults(t *testing.T) {
	p := NewAddressDeleteParams()
	p = p.WithDefaults()
}

func TestAddressDeleteParams_WithTimeout(t *testing.T) {
	p := NewAddressDeleteParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAddressDeleteParams_WithContext(t *testing.T) {
	p := NewAddressDeleteParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAddressDeleteParams_WithHTTPClient(t *testing.T) {
	p := NewAddressDeleteParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressDeleteParams_WithID(t *testing.T) {
	p := NewAddressDeleteParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
