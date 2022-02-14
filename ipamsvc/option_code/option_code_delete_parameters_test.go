package option_code

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewOptionCodeDeleteParams(t *testing.T) {
	p := NewOptionCodeDeleteParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewOptionCodeDeleteParamsWithTimeout(t *testing.T) {
	p := NewOptionCodeDeleteParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewOptionCodeDeleteParamsWithContext(t *testing.T) {
	p := NewOptionCodeDeleteParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewOptionCodeDeleteParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewOptionCodeDeleteParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionCodeDeleteParams_WithDefaults(t *testing.T) {
	p := NewOptionCodeDeleteParams()
	p = p.WithDefaults()
}

func TestOptionCodeDeleteParams_WithTimeout(t *testing.T) {
	p := NewOptionCodeDeleteParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestOptionCodeDeleteParams_WithContext(t *testing.T) {
	p := NewOptionCodeDeleteParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestOptionCodeDeleteParams_WithHTTPClient(t *testing.T) {
	p := NewOptionCodeDeleteParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionCodeDeleteParams_WithID(t *testing.T) {
	p := NewOptionCodeDeleteParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
