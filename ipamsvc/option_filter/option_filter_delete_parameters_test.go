package option_filter

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewOptionFilterDeleteParams(t *testing.T) {
	p := NewOptionFilterDeleteParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewOptionFilterDeleteParamsWithTimeout(t *testing.T) {
	p := NewOptionFilterDeleteParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewOptionFilterDeleteParamsWithContext(t *testing.T) {
	p := NewOptionFilterDeleteParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewOptionFilterDeleteParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewOptionFilterDeleteParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionFilterDeleteParams_WithDefaults(t *testing.T) {
	p := NewOptionFilterDeleteParams()
	p = p.WithDefaults()
}

func TestOptionFilterDeleteParams_WithTimeout(t *testing.T) {
	p := NewOptionFilterDeleteParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestOptionFilterDeleteParams_WithContext(t *testing.T) {
	p := NewOptionFilterDeleteParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestOptionFilterDeleteParams_WithHTTPClient(t *testing.T) {
	p := NewOptionFilterDeleteParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionFilterDeleteParams_WithID(t *testing.T) {
	p := NewOptionFilterDeleteParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
