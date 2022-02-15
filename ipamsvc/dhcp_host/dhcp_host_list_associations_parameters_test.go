package dhcp_host

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewDhcpHostListAssociationsParams(t *testing.T) {
	p := NewDhcpHostListAssociationsParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewDhcpHostListAssociationsParamsWithTimeout(t *testing.T) {
	p := NewDhcpHostListAssociationsParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewDhcpHostListAssociationsParamsWithContext(t *testing.T) {
	p := NewDhcpHostListAssociationsParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewDhcpHostListAssociationsParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewDhcpHostListAssociationsParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestDhcpHostListAssociationsParams_WithDefaults(t *testing.T) {
	p := NewDhcpHostListAssociationsParams()
	p = p.WithDefaults()
}

func TestDhcpHostListAssociationsParams_WithTimeout(t *testing.T) {
	p := NewDhcpHostListAssociationsParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestDhcpHostListAssociationsParams_WithContext(t *testing.T) {
	p := NewDhcpHostListAssociationsParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestDhcpHostListAssociationsParams_WithHTTPClient(t *testing.T) {
	p := NewDhcpHostListAssociationsParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestDhcpHostListAssociationsParams_WithID(t *testing.T) {
	p := NewDhcpHostListAssociationsParams()
	p = p.WithID("test-id")
	assert.Equal(t, "test-id", p.ID)
}
