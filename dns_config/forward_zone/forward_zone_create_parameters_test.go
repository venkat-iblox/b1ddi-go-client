package forward_zone

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

func TestNewForwardZoneCreateParams(t *testing.T) {
	p := NewForwardZoneCreateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewForwardZoneCreateParamsWithTimeout(t *testing.T) {
	p := NewForwardZoneCreateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewForwardZoneCreateParamsWithContext(t *testing.T) {
	p := NewForwardZoneCreateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewForwardZoneCreateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewForwardZoneCreateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestForwardZoneCreateParams_WithDefaults(t *testing.T) {
	p := NewForwardZoneCreateParams()
	p = p.WithDefaults()
}

func TestForwardZoneCreateParams_WithTimeout(t *testing.T) {
	p := NewForwardZoneCreateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestForwardZoneCreateParams_WithContext(t *testing.T) {
	p := NewForwardZoneCreateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestForwardZoneCreateParams_WithHTTPClient(t *testing.T) {
	p := NewForwardZoneCreateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestForwardZoneCreateParams_WithBody(t *testing.T) {
	p := NewForwardZoneCreateParams()
	b := &models.ConfigForwardZone{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}
