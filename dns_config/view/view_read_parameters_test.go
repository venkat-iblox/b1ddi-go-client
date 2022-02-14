package view

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

func TestNewViewReadParams(t *testing.T) {
	p := NewViewReadParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewViewReadParamsWithTimeout(t *testing.T) {
	p := NewViewReadParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewViewReadParamsWithContext(t *testing.T) {
	p := NewViewReadParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewViewReadParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewViewReadParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestViewReadParams_WithDefaults(t *testing.T) {
	p := NewViewReadParams()
	p = p.WithDefaults()
}

func TestViewReadParams_WithTimeout(t *testing.T) {
	p := NewViewReadParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestViewReadParams_WithContext(t *testing.T) {
	p := NewViewReadParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestViewReadParams_WithHTTPClient(t *testing.T) {
	p := NewViewReadParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestViewReadParams_WithFields(t *testing.T) {
	p := NewViewReadParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestViewReadParams_WithID(t *testing.T) {
	p := NewViewReadParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
