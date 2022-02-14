package ip_space

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewIPSpaceDeleteParams(t *testing.T) {
	p := NewIPSpaceDeleteParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewIPSpaceDeleteParamsWithTimeout(t *testing.T) {
	p := NewIPSpaceDeleteParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewIPSpaceDeleteParamsWithContext(t *testing.T) {
	p := NewIPSpaceDeleteParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewIPSpaceDeleteParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewIPSpaceDeleteParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestIPSpaceDeleteParams_WithDefaults(t *testing.T) {
	p := NewIPSpaceDeleteParams()
	p = p.WithDefaults()
}

func TestIPSpaceDeleteParams_WithTimeout(t *testing.T) {
	p := NewIPSpaceDeleteParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestIPSpaceDeleteParams_WithContext(t *testing.T) {
	p := NewIPSpaceDeleteParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestIPSpaceDeleteParams_WithHTTPClient(t *testing.T) {
	p := NewIPSpaceDeleteParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestIPSpaceDeleteParams_WithID(t *testing.T) {
	p := NewIPSpaceDeleteParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
