package auth_nsg

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewAuthNsgDeleteParams(t *testing.T) {
	p := NewAuthNsgDeleteParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAuthNsgDeleteParamsWithTimeout(t *testing.T) {
	p := NewAuthNsgDeleteParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAuthNsgDeleteParamsWithContext(t *testing.T) {
	p := NewAuthNsgDeleteParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAuthNsgDeleteParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAuthNsgDeleteParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAuthNsgDeleteParams_WithDefaults(t *testing.T) {
	p := NewAuthNsgDeleteParams()
	p = p.WithDefaults()
}

func TestAuthNsgDeleteParams_WithTimeout(t *testing.T) {
	p := NewAuthNsgDeleteParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAuthNsgDeleteParams_WithContext(t *testing.T) {
	p := NewAuthNsgDeleteParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAuthNsgDeleteParams_WithHTTPClient(t *testing.T) {
	p := NewAuthNsgDeleteParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAuthNsgDeleteParams_WithID(t *testing.T) {
	p := NewAuthNsgDeleteParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
