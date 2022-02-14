package auth_zone

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewAuthZoneDeleteParams(t *testing.T) {
	p := NewAuthZoneDeleteParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAuthZoneDeleteParamsWithTimeout(t *testing.T) {
	p := NewAuthZoneDeleteParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAuthZoneDeleteParamsWithContext(t *testing.T) {
	p := NewAuthZoneDeleteParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAuthZoneDeleteParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAuthZoneDeleteParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAuthZoneDeleteParams_WithDefaults(t *testing.T) {
	p := NewAuthZoneDeleteParams()
	p = p.WithDefaults()
}

func TestAuthZoneDeleteParams_WithTimeout(t *testing.T) {
	p := NewAuthZoneDeleteParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAuthZoneDeleteParams_WithContext(t *testing.T) {
	p := NewAuthZoneDeleteParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAuthZoneDeleteParams_WithHTTPClient(t *testing.T) {
	p := NewAuthZoneDeleteParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAuthZoneDeleteParams_WithID(t *testing.T) {
	p := NewAuthZoneDeleteParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
