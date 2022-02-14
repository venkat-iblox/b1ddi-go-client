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

func TestNewAuthZoneCreateParams(t *testing.T) {
	p := NewAuthZoneCreateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAuthZoneCreateParamsWithTimeout(t *testing.T) {
	p := NewAuthZoneCreateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAuthZoneCreateParamsWithContext(t *testing.T) {
	p := NewAuthZoneCreateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAuthZoneCreateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAuthZoneCreateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAuthZoneCreateParams_WithDefaults(t *testing.T) {
	p := NewAuthZoneCreateParams()
	p = p.WithDefaults()
}

func TestAuthZoneCreateParams_WithTimeout(t *testing.T) {
	p := NewAuthZoneCreateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAuthZoneCreateParams_WithContext(t *testing.T) {
	p := NewAuthZoneCreateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAuthZoneCreateParams_WithHTTPClient(t *testing.T) {
	p := NewAuthZoneCreateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAuthZoneCreateParams_WithBody(t *testing.T) {
	p := NewAuthZoneCreateParams()
	b := &models.ConfigAuthZone{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}
