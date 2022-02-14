package auth_nsg

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

func TestNewAuthNsgCreateParams(t *testing.T) {
	p := NewAuthNsgCreateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAuthNsgCreateParamsWithTimeout(t *testing.T) {
	p := NewAuthNsgCreateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAuthNsgCreateParamsWithContext(t *testing.T) {
	p := NewAuthNsgCreateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAuthNsgCreateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAuthNsgCreateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAuthNsgCreateParams_WithDefaults(t *testing.T) {
	p := NewAuthNsgCreateParams()
	p = p.WithDefaults()
}

func TestAuthNsgCreateParams_WithTimeout(t *testing.T) {
	p := NewAuthNsgCreateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAuthNsgCreateParams_WithContext(t *testing.T) {
	p := NewAuthNsgCreateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAuthNsgCreateParams_WithHTTPClient(t *testing.T) {
	p := NewAuthNsgCreateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAuthNsgCreateParams_WithBody(t *testing.T) {
	p := NewAuthNsgCreateParams()
	b := &models.ConfigAuthNSG{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}
