package forward_zone

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

func TestNewForwardZoneReadParams(t *testing.T) {
	p := NewForwardZoneReadParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewForwardZoneReadParamsWithTimeout(t *testing.T) {
	p := NewForwardZoneReadParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewForwardZoneReadParamsWithContext(t *testing.T) {
	p := NewForwardZoneReadParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewForwardZoneReadParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewForwardZoneReadParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestForwardZoneReadParams_WithDefaults(t *testing.T) {
	p := NewForwardZoneReadParams()
	p = p.WithDefaults()
}

func TestForwardZoneReadParams_WithTimeout(t *testing.T) {
	p := NewForwardZoneReadParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestForwardZoneReadParams_WithContext(t *testing.T) {
	p := NewForwardZoneReadParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestForwardZoneReadParams_WithHTTPClient(t *testing.T) {
	p := NewForwardZoneReadParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestForwardZoneReadParams_WithFields(t *testing.T) {
	p := NewForwardZoneReadParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestForwardZoneReadParams_WithID(t *testing.T) {
	p := NewForwardZoneReadParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
