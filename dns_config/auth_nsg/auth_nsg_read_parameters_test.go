package auth_nsg

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

func TestNewAuthNsgReadParams(t *testing.T) {
	p := NewAuthNsgReadParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAuthNsgReadParamsWithTimeout(t *testing.T) {
	p := NewAuthNsgReadParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAuthNsgReadParamsWithContext(t *testing.T) {
	p := NewAuthNsgReadParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAuthNsgReadParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAuthNsgReadParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAuthNsgReadParams_WithDefaults(t *testing.T) {
	p := NewAuthNsgReadParams()
	p = p.WithDefaults()
}

func TestAuthNsgReadParams_WithTimeout(t *testing.T) {
	p := NewAuthNsgReadParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAuthNsgReadParams_WithContext(t *testing.T) {
	p := NewAuthNsgReadParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAuthNsgReadParams_WithHTTPClient(t *testing.T) {
	p := NewAuthNsgReadParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAuthNsgReadParams_WithFields(t *testing.T) {
	p := NewAuthNsgReadParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestAuthNsgReadParams_WithID(t *testing.T) {
	p := NewAuthNsgReadParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
