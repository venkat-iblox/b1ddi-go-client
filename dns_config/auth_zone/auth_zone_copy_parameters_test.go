package auth_zone

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

func TestNewAuthZoneCopyParams(t *testing.T) {
	p := NewAuthZoneCopyParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAuthZoneCopyParamsWithTimeout(t *testing.T) {
	p := NewAuthZoneCopyParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAuthZoneCopyParamsWithContext(t *testing.T) {
	p := NewAuthZoneCopyParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAuthZoneCopyParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAuthZoneCopyParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAuthZoneCopyParams_WithDefaults(t *testing.T) {
	p := NewAuthZoneCopyParams()
	p = p.WithDefaults()
}

func TestAuthZoneCopyParams_WithTimeout(t *testing.T) {
	p := NewAuthZoneCopyParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAuthZoneCopyParams_WithContext(t *testing.T) {
	p := NewAuthZoneCopyParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAuthZoneCopyParams_WithHTTPClient(t *testing.T) {
	p := NewAuthZoneCopyParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAuthZoneCopyParams_WithBody(t *testing.T) {
	p := NewAuthZoneCopyParams()
	b := &models.ConfigCopyAuthZone{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}
