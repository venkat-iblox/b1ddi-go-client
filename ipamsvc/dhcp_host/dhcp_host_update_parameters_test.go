package dhcp_host

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewDhcpHostUpdateParams(t *testing.T) {
	p := NewDhcpHostUpdateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewDhcpHostUpdateParamsWithTimeout(t *testing.T) {
	p := NewDhcpHostUpdateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewDhcpHostUpdateParamsWithContext(t *testing.T) {
	p := NewDhcpHostUpdateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewDhcpHostUpdateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewDhcpHostUpdateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestDhcpHostUpdateParams_WithDefaults(t *testing.T) {
	p := NewDhcpHostUpdateParams()
	p = p.WithDefaults()
}

func TestDhcpHostUpdateParams_WithTimeout(t *testing.T) {
	p := NewDhcpHostUpdateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestDhcpHostUpdateParams_WithContext(t *testing.T) {
	p := NewDhcpHostUpdateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestDhcpHostUpdateParams_WithHTTPClient(t *testing.T) {
	p := NewDhcpHostUpdateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestDhcpHostUpdateParams_WithBody(t *testing.T) {
	p := NewDhcpHostUpdateParams()
	b := &models.IpamsvcHost{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}

func TestDhcpHostUpdateParams_WithID(t *testing.T) {
	p := NewDhcpHostUpdateParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
