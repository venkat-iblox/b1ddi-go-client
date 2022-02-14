package forward_nsg

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewForwardNsgDeleteParams(t *testing.T) {
	p := NewForwardNsgDeleteParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewForwardNsgDeleteParamsWithTimeout(t *testing.T) {
	p := NewForwardNsgDeleteParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewForwardNsgDeleteParamsWithContext(t *testing.T) {
	p := NewForwardNsgDeleteParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewForwardNsgDeleteParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewForwardNsgDeleteParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestForwardNsgDeleteParams_WithDefaults(t *testing.T) {
	p := NewForwardNsgDeleteParams()
	p = p.WithDefaults()
}

func TestForwardNsgDeleteParams_WithTimeout(t *testing.T) {
	p := NewForwardNsgDeleteParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestForwardNsgDeleteParams_WithContext(t *testing.T) {
	p := NewForwardNsgDeleteParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestForwardNsgDeleteParams_WithHTTPClient(t *testing.T) {
	p := NewForwardNsgDeleteParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestForwardNsgDeleteParams_WithID(t *testing.T) {
	p := NewForwardNsgDeleteParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
