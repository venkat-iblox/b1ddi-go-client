package ip_space

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

func TestNewIPSpaceReadParams(t *testing.T) {
	p := NewIPSpaceReadParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewIPSpaceReadParamsWithTimeout(t *testing.T) {
	p := NewIPSpaceReadParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewIPSpaceReadParamsWithContext(t *testing.T) {
	p := NewIPSpaceReadParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewIPSpaceReadParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewIPSpaceReadParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestIPSpaceReadParams_WithDefaults(t *testing.T) {
	p := NewIPSpaceReadParams()
	p = p.WithDefaults()
}

func TestIPSpaceReadParams_WithTimeout(t *testing.T) {
	p := NewIPSpaceReadParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestIPSpaceReadParams_WithContext(t *testing.T) {
	p := NewIPSpaceReadParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestIPSpaceReadParams_WithHTTPClient(t *testing.T) {
	p := NewIPSpaceReadParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestIPSpaceReadParams_WithFields(t *testing.T) {
	p := NewIPSpaceReadParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestIPSpaceReadParams_WithID(t *testing.T) {
	p := NewIPSpaceReadParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
