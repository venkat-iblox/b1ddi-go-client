package hardware_filter

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewHardwareFilterDeleteParams(t *testing.T) {
	p := NewHardwareFilterDeleteParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewHardwareFilterDeleteParamsWithTimeout(t *testing.T) {
	p := NewHardwareFilterDeleteParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewHardwareFilterDeleteParamsWithContext(t *testing.T) {
	p := NewHardwareFilterDeleteParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewHardwareFilterDeleteParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewHardwareFilterDeleteParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestHardwareFilterDeleteParams_WithDefaults(t *testing.T) {
	p := NewHardwareFilterDeleteParams()
	p = p.WithDefaults()
}

func TestHardwareFilterDeleteParams_WithTimeout(t *testing.T) {
	p := NewHardwareFilterDeleteParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestHardwareFilterDeleteParams_WithContext(t *testing.T) {
	p := NewHardwareFilterDeleteParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestHardwareFilterDeleteParams_WithHTTPClient(t *testing.T) {
	p := NewHardwareFilterDeleteParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestHardwareFilterDeleteParams_WithID(t *testing.T) {
	p := NewHardwareFilterDeleteParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
