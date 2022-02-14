package host

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

func TestNewHostReadParams(t *testing.T) {
	p := NewHostReadParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewHostReadParamsWithTimeout(t *testing.T) {
	p := NewHostReadParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewHostReadParamsWithContext(t *testing.T) {
	p := NewHostReadParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewHostReadParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewHostReadParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestHostReadParams_WithDefaults(t *testing.T) {
	p := NewHostReadParams()
	p = p.WithDefaults()
}

func TestHostReadParams_WithTimeout(t *testing.T) {
	p := NewHostReadParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestHostReadParams_WithContext(t *testing.T) {
	p := NewHostReadParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestHostReadParams_WithHTTPClient(t *testing.T) {
	p := NewHostReadParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestHostReadParams_WithFields(t *testing.T) {
	p := NewHostReadParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestHostReadParams_WithID(t *testing.T) {
	p := NewHostReadParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
