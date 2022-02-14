package global

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

func TestNewGlobalReadParams(t *testing.T) {
	p := NewGlobalReadParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewGlobalReadParamsWithTimeout(t *testing.T) {
	p := NewGlobalReadParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewGlobalReadParamsWithContext(t *testing.T) {
	p := NewGlobalReadParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewGlobalReadParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewGlobalReadParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestGlobalReadParams_WithDefaults(t *testing.T) {
	p := NewGlobalReadParams()
	p = p.WithDefaults()
}

func TestGlobalReadParams_WithTimeout(t *testing.T) {
	p := NewGlobalReadParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestGlobalReadParams_WithContext(t *testing.T) {
	p := NewGlobalReadParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestGlobalReadParams_WithHTTPClient(t *testing.T) {
	p := NewGlobalReadParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestGlobalReadParams_WithFields(t *testing.T) {
	p := NewGlobalReadParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}
