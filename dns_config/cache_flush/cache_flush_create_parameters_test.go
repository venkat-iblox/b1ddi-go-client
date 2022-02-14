package cache_flush

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

func TestNewCacheFlushCreateParams(t *testing.T) {
	p := NewCacheFlushCreateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewCacheFlushCreateParamsWithTimeout(t *testing.T) {
	p := NewCacheFlushCreateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewCacheFlushCreateParamsWithContext(t *testing.T) {
	p := NewCacheFlushCreateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewCacheFlushCreateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewCacheFlushCreateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestCacheFlushCreateParams_WithDefaults(t *testing.T) {
	p := NewCacheFlushCreateParams()
	p = p.WithDefaults()
}

func TestCacheFlushCreateParams_WithTimeout(t *testing.T) {
	p := NewCacheFlushCreateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestCacheFlushCreateParams_WithContext(t *testing.T) {
	p := NewCacheFlushCreateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestCacheFlushCreateParams_WithHTTPClient(t *testing.T) {
	p := NewCacheFlushCreateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestCacheFlushCreateParams_WithBody(t *testing.T) {
	p := NewCacheFlushCreateParams()
	b := &models.ConfigCacheFlush{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}
