package dns_usage

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

func TestNewDNSUsageReadParams(t *testing.T) {
	p := NewDNSUsageReadParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewDNSUsageReadParamsWithTimeout(t *testing.T) {
	p := NewDNSUsageReadParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewDNSUsageReadParamsWithContext(t *testing.T) {
	p := NewDNSUsageReadParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewDNSUsageReadParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewDNSUsageReadParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestDNSUsageReadParams_WithDefaults(t *testing.T) {
	p := NewDNSUsageReadParams()
	p = p.WithDefaults()
}

func TestDNSUsageReadParams_WithTimeout(t *testing.T) {
	p := NewDNSUsageReadParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestDNSUsageReadParams_WithContext(t *testing.T) {
	p := NewDNSUsageReadParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestDNSUsageReadParams_WithHTTPClient(t *testing.T) {
	p := NewDNSUsageReadParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestDNSUsageReadParams_WithFields(t *testing.T) {
	p := NewDNSUsageReadParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestDNSUsageReadParams_WithID(t *testing.T) {
	p := NewDNSUsageReadParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
