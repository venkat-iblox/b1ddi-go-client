package server

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

func TestNewServerUpdateParams(t *testing.T) {
	p := NewServerUpdateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewServerUpdateParamsWithTimeout(t *testing.T) {
	p := NewServerUpdateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewServerUpdateParamsWithContext(t *testing.T) {
	p := NewServerUpdateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewServerUpdateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewServerUpdateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestServerUpdateParams_WithDefaults(t *testing.T) {
	p := NewServerUpdateParams()
	p = p.WithDefaults()
}

func TestServerUpdateParams_WithTimeout(t *testing.T) {
	p := NewServerUpdateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestServerUpdateParams_WithContext(t *testing.T) {
	p := NewServerUpdateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestServerUpdateParams_WithHTTPClient(t *testing.T) {
	p := NewServerUpdateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestServerUpdateParams_WithBody(t *testing.T) {
	p := NewServerUpdateParams()
	b := &models.ConfigServer{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}

func TestServerUpdateParams_WithID(t *testing.T) {
	p := NewServerUpdateParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
