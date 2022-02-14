package server

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

func TestNewServerReadParams(t *testing.T) {
	p := NewServerReadParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewServerReadParamsWithTimeout(t *testing.T) {
	p := NewServerReadParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewServerReadParamsWithContext(t *testing.T) {
	p := NewServerReadParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewServerReadParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewServerReadParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestServerReadParams_WithDefaults(t *testing.T) {
	p := NewServerReadParams()
	p = p.WithDefaults()
}

func TestServerReadParams_WithTimeout(t *testing.T) {
	p := NewServerReadParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestServerReadParams_WithContext(t *testing.T) {
	p := NewServerReadParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestServerReadParams_WithHTTPClient(t *testing.T) {
	p := NewServerReadParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestServerReadParams_WithFields(t *testing.T) {
	p := NewServerReadParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestServerReadParams_WithID(t *testing.T) {
	p := NewServerReadParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
