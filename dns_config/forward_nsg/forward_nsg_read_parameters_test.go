package forward_nsg

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

func TestNewForwardNsgReadParams(t *testing.T) {
	p := NewForwardNsgReadParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewForwardNsgReadParamsWithTimeout(t *testing.T) {
	p := NewForwardNsgReadParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewForwardNsgReadParamsWithContext(t *testing.T) {
	p := NewForwardNsgReadParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewForwardNsgReadParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewForwardNsgReadParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestForwardNsgReadParams_WithDefaults(t *testing.T) {
	p := NewForwardNsgReadParams()
	p = p.WithDefaults()
}

func TestForwardNsgReadParams_WithTimeout(t *testing.T) {
	p := NewForwardNsgReadParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestForwardNsgReadParams_WithContext(t *testing.T) {
	p := NewForwardNsgReadParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestForwardNsgReadParams_WithHTTPClient(t *testing.T) {
	p := NewForwardNsgReadParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestForwardNsgReadParams_WithFields(t *testing.T) {
	p := NewForwardNsgReadParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestForwardNsgReadParams_WithID(t *testing.T) {
	p := NewForwardNsgReadParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
