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

func TestNewForwardZoneCopyParams(t *testing.T) {
	p := NewForwardZoneCopyParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewForwardZoneCopyParamsWithTimeout(t *testing.T) {
	p := NewForwardZoneCopyParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewForwardZoneCopyParamsWithContext(t *testing.T) {
	p := NewForwardZoneCopyParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewForwardZoneCopyParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewForwardZoneCopyParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestForwardZoneCopyParams_WithDefaults(t *testing.T) {
	p := NewForwardZoneCopyParams()
	p = p.WithDefaults()
}

func TestForwardZoneCopyParams_WithTimeout(t *testing.T) {
	p := NewForwardZoneCopyParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestForwardZoneCopyParams_WithContext(t *testing.T) {
	p := NewForwardZoneCopyParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestForwardZoneCopyParams_WithHTTPClient(t *testing.T) {
	p := NewForwardZoneCopyParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestForwardZoneCopyParams_WithBody(t *testing.T) {
	p := NewForwardZoneCopyParams()
	b := &models.ConfigCopyForwardZone{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}
