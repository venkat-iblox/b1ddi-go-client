package option_filter

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

func TestNewOptionFilterReadParams(t *testing.T) {
	p := NewOptionFilterReadParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewOptionFilterReadParamsWithTimeout(t *testing.T) {
	p := NewOptionFilterReadParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewOptionFilterReadParamsWithContext(t *testing.T) {
	p := NewOptionFilterReadParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewOptionFilterReadParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewOptionFilterReadParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionFilterReadParams_WithDefaults(t *testing.T) {
	p := NewOptionFilterReadParams()
	p = p.WithDefaults()
}

func TestOptionFilterReadParams_WithTimeout(t *testing.T) {
	p := NewOptionFilterReadParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestOptionFilterReadParams_WithContext(t *testing.T) {
	p := NewOptionFilterReadParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestOptionFilterReadParams_WithHTTPClient(t *testing.T) {
	p := NewOptionFilterReadParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionFilterReadParams_WithFields(t *testing.T) {
	p := NewOptionFilterReadParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestOptionFilterReadParams_WithID(t *testing.T) {
	p := NewOptionFilterReadParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
