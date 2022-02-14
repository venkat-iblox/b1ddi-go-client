package server

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewServerDeleteParams(t *testing.T) {
	p := NewServerDeleteParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewServerDeleteParamsWithTimeout(t *testing.T) {
	p := NewServerDeleteParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewServerDeleteParamsWithContext(t *testing.T) {
	p := NewServerDeleteParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewServerDeleteParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewServerDeleteParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestServerDeleteParams_WithDefaults(t *testing.T) {
	p := NewServerDeleteParams()
	p = p.WithDefaults()
}

func TestServerDeleteParams_WithTimeout(t *testing.T) {
	p := NewServerDeleteParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestServerDeleteParams_WithContext(t *testing.T) {
	p := NewServerDeleteParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestServerDeleteParams_WithHTTPClient(t *testing.T) {
	p := NewServerDeleteParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestServerDeleteParams_WithID(t *testing.T) {
	p := NewServerDeleteParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
