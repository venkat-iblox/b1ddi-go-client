package view

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewViewDeleteParams(t *testing.T) {
	p := NewViewDeleteParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewViewDeleteParamsWithTimeout(t *testing.T) {
	p := NewViewDeleteParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewViewDeleteParamsWithContext(t *testing.T) {
	p := NewViewDeleteParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewViewDeleteParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewViewDeleteParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestViewDeleteParams_WithDefaults(t *testing.T) {
	p := NewViewDeleteParams()
	p = p.WithDefaults()
}

func TestViewDeleteParams_WithTimeout(t *testing.T) {
	p := NewViewDeleteParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestViewDeleteParams_WithContext(t *testing.T) {
	p := NewViewDeleteParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestViewDeleteParams_WithHTTPClient(t *testing.T) {
	p := NewViewDeleteParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestViewDeleteParams_WithID(t *testing.T) {
	p := NewViewDeleteParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
