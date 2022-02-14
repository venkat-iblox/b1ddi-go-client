package hardware_filter

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

func TestNewHardwareFilterReadParams(t *testing.T) {
	p := NewHardwareFilterReadParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewHardwareFilterReadParamsWithTimeout(t *testing.T) {
	p := NewHardwareFilterReadParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewHardwareFilterReadParamsWithContext(t *testing.T) {
	p := NewHardwareFilterReadParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewHardwareFilterReadParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewHardwareFilterReadParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestHardwareFilterReadParams_WithDefaults(t *testing.T) {
	p := NewHardwareFilterReadParams()
	p = p.WithDefaults()
}

func TestHardwareFilterReadParams_WithTimeout(t *testing.T) {
	p := NewHardwareFilterReadParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestHardwareFilterReadParams_WithContext(t *testing.T) {
	p := NewHardwareFilterReadParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestHardwareFilterReadParams_WithHTTPClient(t *testing.T) {
	p := NewHardwareFilterReadParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestHardwareFilterReadParams_WithFields(t *testing.T) {
	p := NewHardwareFilterReadParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestHardwareFilterReadParams_WithID(t *testing.T) {
	p := NewHardwareFilterReadParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
