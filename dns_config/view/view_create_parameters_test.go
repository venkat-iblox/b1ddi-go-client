package view

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

func TestNewViewCreateParams(t *testing.T) {
	p := NewViewCreateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewViewCreateParamsWithTimeout(t *testing.T) {
	p := NewViewCreateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewViewCreateParamsWithContext(t *testing.T) {
	p := NewViewCreateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewViewCreateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewViewCreateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestViewCreateParams_WithDefaults(t *testing.T) {
	p := NewViewCreateParams()
	p = p.WithDefaults()
}

func TestViewCreateParams_WithTimeout(t *testing.T) {
	p := NewViewCreateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestViewCreateParams_WithContext(t *testing.T) {
	p := NewViewCreateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestViewCreateParams_WithHTTPClient(t *testing.T) {
	p := NewViewCreateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestViewCreateParams_WithBody(t *testing.T) {
	p := NewViewCreateParams()
	b := &models.ConfigView{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}
