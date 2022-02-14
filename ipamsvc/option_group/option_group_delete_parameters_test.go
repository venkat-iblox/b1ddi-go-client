package option_group

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewOptionGroupDeleteParams(t *testing.T) {
	p := NewOptionGroupDeleteParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewOptionGroupDeleteParamsWithTimeout(t *testing.T) {
	p := NewOptionGroupDeleteParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewOptionGroupDeleteParamsWithContext(t *testing.T) {
	p := NewOptionGroupDeleteParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewOptionGroupDeleteParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewOptionGroupDeleteParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionGroupDeleteParams_WithDefaults(t *testing.T) {
	p := NewOptionGroupDeleteParams()
	p = p.WithDefaults()
}

func TestOptionGroupDeleteParams_WithTimeout(t *testing.T) {
	p := NewOptionGroupDeleteParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestOptionGroupDeleteParams_WithContext(t *testing.T) {
	p := NewOptionGroupDeleteParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestOptionGroupDeleteParams_WithHTTPClient(t *testing.T) {
	p := NewOptionGroupDeleteParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionGroupDeleteParams_WithID(t *testing.T) {
	p := NewOptionGroupDeleteParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
