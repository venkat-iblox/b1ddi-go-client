package ipam_host

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

func TestNewIpamHostReadParams(t *testing.T) {
	p := NewIpamHostReadParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewIpamHostReadParamsWithTimeout(t *testing.T) {
	p := NewIpamHostReadParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewIpamHostReadParamsWithContext(t *testing.T) {
	p := NewIpamHostReadParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewIpamHostReadParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewIpamHostReadParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestIpamHostReadParams_WithDefaults(t *testing.T) {
	p := NewIpamHostReadParams()
	p = p.WithDefaults()
}

func TestIpamHostReadParams_WithTimeout(t *testing.T) {
	p := NewIpamHostReadParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestIpamHostReadParams_WithContext(t *testing.T) {
	p := NewIpamHostReadParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestIpamHostReadParams_WithHTTPClient(t *testing.T) {
	p := NewIpamHostReadParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestIpamHostReadParams_WithFields(t *testing.T) {
	p := NewIpamHostReadParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestIpamHostReadParams_WithID(t *testing.T) {
	p := NewIpamHostReadParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
