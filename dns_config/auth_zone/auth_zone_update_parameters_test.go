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

func TestNewAuthZoneUpdateParams(t *testing.T) {
	p := NewAuthZoneUpdateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAuthZoneUpdateParamsWithTimeout(t *testing.T) {
	p := NewAuthZoneUpdateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAuthZoneUpdateParamsWithContext(t *testing.T) {
	p := NewAuthZoneUpdateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAuthZoneUpdateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAuthZoneUpdateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAuthZoneUpdateParams_WithDefaults(t *testing.T) {
	p := NewAuthZoneUpdateParams()
	p = p.WithDefaults()
}

func TestAuthZoneUpdateParams_WithTimeout(t *testing.T) {
	p := NewAuthZoneUpdateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAuthZoneUpdateParams_WithContext(t *testing.T) {
	p := NewAuthZoneUpdateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAuthZoneUpdateParams_WithHTTPClient(t *testing.T) {
	p := NewAuthZoneUpdateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAuthZoneUpdateParams_WithBody(t *testing.T) {
	p := NewAuthZoneUpdateParams()
	b := &models.ConfigAuthZone{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}

func TestAuthZoneUpdateParams_WithID(t *testing.T) {
	p := NewAuthZoneUpdateParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
