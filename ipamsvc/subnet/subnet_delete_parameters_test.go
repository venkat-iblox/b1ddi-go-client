package subnet

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewSubnetDeleteParams(t *testing.T) {
	p := NewSubnetDeleteParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewSubnetDeleteParamsWithTimeout(t *testing.T) {
	p := NewSubnetDeleteParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewSubnetDeleteParamsWithContext(t *testing.T) {
	p := NewSubnetDeleteParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewSubnetDeleteParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewSubnetDeleteParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestSubnetDeleteParams_WithDefaults(t *testing.T) {
	p := NewSubnetDeleteParams()
	p = p.WithDefaults()
}

func TestSubnetDeleteParams_WithTimeout(t *testing.T) {
	p := NewSubnetDeleteParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestSubnetDeleteParams_WithContext(t *testing.T) {
	p := NewSubnetDeleteParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestSubnetDeleteParams_WithHTTPClient(t *testing.T) {
	p := NewSubnetDeleteParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestSubnetDeleteParams_WithID(t *testing.T) {
	p := NewSubnetDeleteParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
