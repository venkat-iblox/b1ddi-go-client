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

func TestNewIpamHostCreateParams(t *testing.T) {
	p := NewIpamHostCreateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewIpamHostCreateParamsWithTimeout(t *testing.T) {
	p := NewIpamHostCreateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewIpamHostCreateParamsWithContext(t *testing.T) {
	p := NewIpamHostCreateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewIpamHostCreateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewIpamHostCreateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestIpamHostCreateParams_WithDefaults(t *testing.T) {
	p := NewIpamHostCreateParams()
	p = p.WithDefaults()
}

func TestIpamHostCreateParams_WithTimeout(t *testing.T) {
	p := NewIpamHostCreateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestIpamHostCreateParams_WithContext(t *testing.T) {
	p := NewIpamHostCreateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestIpamHostCreateParams_WithHTTPClient(t *testing.T) {
	p := NewIpamHostCreateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestIpamHostCreateParams_WithBody(t *testing.T) {
	p := NewIpamHostCreateParams()
	b := &models.IpamsvcIpamHost{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}
