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

func TestNewForwardZoneUpdateParams(t *testing.T) {
	p := NewForwardZoneUpdateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewForwardZoneUpdateParamsWithTimeout(t *testing.T) {
	p := NewForwardZoneUpdateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewForwardZoneUpdateParamsWithContext(t *testing.T) {
	p := NewForwardZoneUpdateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewForwardZoneUpdateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewForwardZoneUpdateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestForwardZoneUpdateParams_WithDefaults(t *testing.T) {
	p := NewForwardZoneUpdateParams()
	p = p.WithDefaults()
}

func TestForwardZoneUpdateParams_WithTimeout(t *testing.T) {
	p := NewForwardZoneUpdateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestForwardZoneUpdateParams_WithContext(t *testing.T) {
	p := NewForwardZoneUpdateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestForwardZoneUpdateParams_WithHTTPClient(t *testing.T) {
	p := NewForwardZoneUpdateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestForwardZoneUpdateParams_WithBody(t *testing.T) {
	p := NewForwardZoneUpdateParams()
	b := &models.ConfigForwardZone{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}

func TestForwardZoneUpdateParams_WithID(t *testing.T) {
	p := NewForwardZoneUpdateParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
