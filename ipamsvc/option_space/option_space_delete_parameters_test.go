package option_space

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewOptionSpaceDeleteParams(t *testing.T) {
	p := NewOptionSpaceDeleteParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewOptionSpaceDeleteParamsWithTimeout(t *testing.T) {
	p := NewOptionSpaceDeleteParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewOptionSpaceDeleteParamsWithContext(t *testing.T) {
	p := NewOptionSpaceDeleteParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewOptionSpaceDeleteParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewOptionSpaceDeleteParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionSpaceDeleteParams_WithDefaults(t *testing.T) {
	p := NewOptionSpaceDeleteParams()
	p = p.WithDefaults()
}

func TestOptionSpaceDeleteParams_WithTimeout(t *testing.T) {
	p := NewOptionSpaceDeleteParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestOptionSpaceDeleteParams_WithContext(t *testing.T) {
	p := NewOptionSpaceDeleteParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestOptionSpaceDeleteParams_WithHTTPClient(t *testing.T) {
	p := NewOptionSpaceDeleteParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionSpaceDeleteParams_WithID(t *testing.T) {
	p := NewOptionSpaceDeleteParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
