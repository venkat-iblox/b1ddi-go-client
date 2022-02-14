package convert_domain_name

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewConvertDomainNameConvertParams(t *testing.T) {
	p := NewConvertDomainNameConvertParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewConvertDomainNameConvertParamsWithTimeout(t *testing.T) {
	p := NewConvertDomainNameConvertParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewConvertDomainNameConvertParamsWithContext(t *testing.T) {
	p := NewConvertDomainNameConvertParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewConvertDomainNameConvertParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewConvertDomainNameConvertParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestConvertDomainNameConvertParams_WithDefaults(t *testing.T) {
	p := NewConvertDomainNameConvertParams()
	p = p.WithDefaults()
}

func TestConvertDomainNameConvertParams_WithTimeout(t *testing.T) {
	p := NewConvertDomainNameConvertParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestConvertDomainNameConvertParams_WithContext(t *testing.T) {
	p := NewConvertDomainNameConvertParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestConvertDomainNameConvertParams_WithHTTPClient(t *testing.T) {
	p := NewConvertDomainNameConvertParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestConvertDomainNameConvertParams_WithDomainName(t *testing.T) {
	p := NewConvertDomainNameConvertParams()
	p = p.WithDomainName("domain-name")
	require.NotNil(t, p.DomainName)
	assert.Equal(t, "domain-name", p.DomainName)
}
