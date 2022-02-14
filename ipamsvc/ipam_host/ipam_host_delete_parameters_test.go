package ipam_host

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewIpamHostDeleteParams(t *testing.T) {
	p := NewIpamHostDeleteParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewIpamHostDeleteParamsWithTimeout(t *testing.T) {
	p := NewIpamHostDeleteParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewIpamHostDeleteParamsWithContext(t *testing.T) {
	p := NewIpamHostDeleteParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewIpamHostDeleteParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewIpamHostDeleteParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestIpamHostDeleteParams_WithDefaults(t *testing.T) {
	p := NewIpamHostDeleteParams()
	p = p.WithDefaults()
}

func TestIpamHostDeleteParams_WithTimeout(t *testing.T) {
	p := NewIpamHostDeleteParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestIpamHostDeleteParams_WithContext(t *testing.T) {
	p := NewIpamHostDeleteParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestIpamHostDeleteParams_WithHTTPClient(t *testing.T) {
	p := NewIpamHostDeleteParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestIpamHostDeleteParams_WithID(t *testing.T) {
	p := NewIpamHostDeleteParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
