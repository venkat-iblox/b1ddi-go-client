package ipam_host

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

func TestNewIpamHostUpdateParams(t *testing.T) {
	p := NewIpamHostUpdateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewIpamHostUpdateParamsWithTimeout(t *testing.T) {
	p := NewIpamHostUpdateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewIpamHostUpdateParamsWithContext(t *testing.T) {
	p := NewIpamHostUpdateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewIpamHostUpdateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewIpamHostUpdateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestIpamHostUpdateParams_WithDefaults(t *testing.T) {
	p := NewIpamHostUpdateParams()
	p = p.WithDefaults()
}

func TestIpamHostUpdateParams_WithTimeout(t *testing.T) {
	p := NewIpamHostUpdateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestIpamHostUpdateParams_WithContext(t *testing.T) {
	p := NewIpamHostUpdateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestIpamHostUpdateParams_WithHTTPClient(t *testing.T) {
	p := NewIpamHostUpdateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestIpamHostUpdateParams_WithBody(t *testing.T) {
	p := NewIpamHostUpdateParams()
	b := &models.IpamsvcIpamHost{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}

func TestIpamHostUpdateParams_WithID(t *testing.T) {
	p := NewIpamHostUpdateParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
