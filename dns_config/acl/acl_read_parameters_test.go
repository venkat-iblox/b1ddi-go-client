package acl

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

func TestNewACLReadParams(t *testing.T) {
	p := NewACLReadParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewACLReadParamsWithTimeout(t *testing.T) {
	p := NewACLReadParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewACLReadParamsWithContext(t *testing.T) {
	p := NewACLReadParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewACLReadParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewACLReadParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestACLReadParams_WithDefaults(t *testing.T) {
	p := NewACLReadParams()
	p = p.WithDefaults()
}

func TestACLReadParams_WithTimeout(t *testing.T) {
	p := NewACLReadParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestACLReadParams_WithContext(t *testing.T) {
	p := NewACLReadParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestACLReadParams_WithHTTPClient(t *testing.T) {
	p := NewACLReadParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestACLReadParams_WithFields(t *testing.T) {
	p := NewACLReadParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestACLReadParams_WithID(t *testing.T) {
	p := NewACLReadParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
