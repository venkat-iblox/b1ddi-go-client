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

func TestNewHardwareFilterUpdateParams(t *testing.T) {
	p := NewHardwareFilterUpdateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewHardwareFilterUpdateParamsWithTimeout(t *testing.T) {
	p := NewHardwareFilterUpdateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewHardwareFilterUpdateParamsWithContext(t *testing.T) {
	p := NewHardwareFilterUpdateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewHardwareFilterUpdateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewHardwareFilterUpdateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestHardwareFilterUpdateParams_WithDefaults(t *testing.T) {
	p := NewHardwareFilterUpdateParams()
	p = p.WithDefaults()
}

func TestHardwareFilterUpdateParams_WithTimeout(t *testing.T) {
	p := NewHardwareFilterUpdateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestHardwareFilterUpdateParams_WithContext(t *testing.T) {
	p := NewHardwareFilterUpdateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestHardwareFilterUpdateParams_WithHTTPClient(t *testing.T) {
	p := NewHardwareFilterUpdateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestHardwareFilterUpdateParams_WithBody(t *testing.T) {
	p := NewHardwareFilterUpdateParams()
	b := &models.IpamsvcHardwareFilter{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}

func TestHardwareFilterUpdateParams_WithID(t *testing.T) {
	p := NewHardwareFilterUpdateParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
