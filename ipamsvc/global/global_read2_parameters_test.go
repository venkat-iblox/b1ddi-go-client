package global

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

func TestNewGlobalRead2Params(t *testing.T) {
	p := NewGlobalRead2Params()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewGlobalRead2ParamsWithTimeout(t *testing.T) {
	p := NewGlobalRead2ParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewGlobalRead2ParamsWithContext(t *testing.T) {
	p := NewGlobalRead2ParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewGlobalRead2ParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewGlobalRead2ParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestGlobalRead2Params_WithDefaults(t *testing.T) {
	p := NewGlobalRead2Params()
	p = p.WithDefaults()
}

func TestGlobalRead2Params_WithTimeout(t *testing.T) {
	p := NewGlobalRead2Params()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestGlobalRead2Params_WithContext(t *testing.T) {
	p := NewGlobalRead2Params()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestGlobalRead2Params_WithHTTPClient(t *testing.T) {
	p := NewGlobalRead2Params()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestGlobalRead2Params_WithFields(t *testing.T) {
	p := NewGlobalRead2Params()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestGlobalRead2Params_WithID(t *testing.T) {
	p := NewGlobalRead2Params()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
