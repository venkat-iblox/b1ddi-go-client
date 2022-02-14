package dhcp_host

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

func TestNewDhcpHostReadParams(t *testing.T) {
	p := NewDhcpHostReadParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewDhcpHostReadParamsWithTimeout(t *testing.T) {
	p := NewDhcpHostReadParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewDhcpHostReadParamsWithContext(t *testing.T) {
	p := NewDhcpHostReadParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewDhcpHostReadParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewDhcpHostReadParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestDhcpHostReadParams_WithDefaults(t *testing.T) {
	p := NewDhcpHostReadParams()
	p = p.WithDefaults()
}

func TestDhcpHostReadParams_WithTimeout(t *testing.T) {
	p := NewDhcpHostReadParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestDhcpHostReadParams_WithContext(t *testing.T) {
	p := NewDhcpHostReadParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestDhcpHostReadParams_WithHTTPClient(t *testing.T) {
	p := NewDhcpHostReadParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestDhcpHostReadParams_WithFields(t *testing.T) {
	p := NewDhcpHostReadParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestDhcpHostReadParams_WithID(t *testing.T) {
	p := NewDhcpHostReadParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
