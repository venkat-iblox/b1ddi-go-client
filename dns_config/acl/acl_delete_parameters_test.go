package acl

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewACLDeleteParams(t *testing.T) {
	p := NewACLDeleteParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewACLDeleteParamsWithTimeout(t *testing.T) {
	p := NewACLDeleteParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewACLDeleteParamsWithContext(t *testing.T) {
	p := NewACLDeleteParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewACLDeleteParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewACLDeleteParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestACLDeleteParams_WithDefaults(t *testing.T) {
	p := NewACLDeleteParams()
	p = p.WithDefaults()
}

func TestACLDeleteParams_WithTimeout(t *testing.T) {
	p := NewACLDeleteParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestACLDeleteParams_WithContext(t *testing.T) {
	p := NewACLDeleteParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestACLDeleteParams_WithHTTPClient(t *testing.T) {
	p := NewACLDeleteParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestACLDeleteParams_WithID(t *testing.T) {
	p := NewACLDeleteParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
