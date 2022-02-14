package server

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewServerCreateParams(t *testing.T) {
	p := NewServerCreateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewServerCreateParamsWithTimeout(t *testing.T) {
	p := NewServerCreateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewServerCreateParamsWithContext(t *testing.T) {
	p := NewServerCreateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewServerCreateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewServerCreateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestServerCreateParams_WithDefaults(t *testing.T) {
	p := NewServerCreateParams()
	p = p.WithDefaults()
}

func TestServerCreateParams_WithTimeout(t *testing.T) {
	p := NewServerCreateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestServerCreateParams_WithContext(t *testing.T) {
	p := NewServerCreateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestServerCreateParams_WithHTTPClient(t *testing.T) {
	p := NewServerCreateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestServerCreateParams_WithBody(t *testing.T) {
	p := NewServerCreateParams()
	b := &models.ConfigServer{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}
