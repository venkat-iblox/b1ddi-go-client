package forward_zone

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewForwardZoneDeleteParams(t *testing.T) {
	p := NewForwardZoneDeleteParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewForwardZoneDeleteParamsWithTimeout(t *testing.T) {
	p := NewForwardZoneDeleteParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewForwardZoneDeleteParamsWithContext(t *testing.T) {
	p := NewForwardZoneDeleteParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewForwardZoneDeleteParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewForwardZoneDeleteParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestForwardZoneDeleteParams_WithDefaults(t *testing.T) {
	p := NewForwardZoneDeleteParams()
	p = p.WithDefaults()
}

func TestForwardZoneDeleteParams_WithTimeout(t *testing.T) {
	p := NewForwardZoneDeleteParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestForwardZoneDeleteParams_WithContext(t *testing.T) {
	p := NewForwardZoneDeleteParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestForwardZoneDeleteParams_WithHTTPClient(t *testing.T) {
	p := NewForwardZoneDeleteParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestForwardZoneDeleteParams_WithID(t *testing.T) {
	p := NewForwardZoneDeleteParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
