package auth_zone

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewAuthZoneReadParams(t *testing.T) {
	p := NewAuthZoneReadParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAuthZoneReadParamsWithTimeout(t *testing.T) {
	p := NewAuthZoneReadParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAuthZoneReadParamsWithContext(t *testing.T) {
	p := NewAuthZoneReadParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAuthZoneReadParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAuthZoneReadParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAuthZoneReadParams_WithDefaults(t *testing.T) {
	p := NewAuthZoneReadParams()
	p = p.WithDefaults()
}

func TestAuthZoneReadParams_WithTimeout(t *testing.T) {
	p := NewAuthZoneReadParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAuthZoneReadParams_WithContext(t *testing.T) {
	p := NewAuthZoneReadParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAuthZoneReadParams_WithHTTPClient(t *testing.T) {
	p := NewAuthZoneReadParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAuthZoneReadParams_WithFields(t *testing.T) {
	p := NewAuthZoneReadParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestAuthZoneReadParams_WithID(t *testing.T) {
	p := NewAuthZoneReadParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
