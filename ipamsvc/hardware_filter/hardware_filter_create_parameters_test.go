package hardware_filter

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

func TestNewHardwareFilterCreateParams(t *testing.T) {
	p := NewHardwareFilterCreateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewHardwareFilterCreateParamsWithTimeout(t *testing.T) {
	p := NewHardwareFilterCreateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewHardwareFilterCreateParamsWithContext(t *testing.T) {
	p := NewHardwareFilterCreateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewHardwareFilterCreateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewHardwareFilterCreateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestHardwareFilterCreateParams_WithDefaults(t *testing.T) {
	p := NewHardwareFilterCreateParams()
	p = p.WithDefaults()
}

func TestHardwareFilterCreateParams_WithTimeout(t *testing.T) {
	p := NewHardwareFilterCreateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestHardwareFilterCreateParams_WithContext(t *testing.T) {
	p := NewHardwareFilterCreateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestHardwareFilterCreateParams_WithHTTPClient(t *testing.T) {
	p := NewHardwareFilterCreateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestHardwareFilterCreateParams_WithBody(t *testing.T) {
	p := NewHardwareFilterCreateParams()
	b := &models.IpamsvcHardwareFilter{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}
