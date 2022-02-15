package dns_config

import (
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHTTPClient(t *testing.T) {
	cli := NewHTTPClient(strfmt.Default)

	assert.NotNil(t, cli.Transport)
	assert.NotNil(t, cli.ACL)
	assert.NotNil(t, cli.AuthNsg)
	assert.NotNil(t, cli.AuthZone)
	assert.NotNil(t, cli.CacheFlush)
	assert.NotNil(t, cli.ConvertDomainName)
	assert.NotNil(t, cli.ConvertRname)
	assert.NotNil(t, cli.Delegation)
	assert.NotNil(t, cli.ForwardNsg)
	assert.NotNil(t, cli.ForwardZone)
	assert.NotNil(t, cli.Global)
	assert.NotNil(t, cli.Host)
	assert.NotNil(t, cli.Server)
	assert.NotNil(t, cli.View)
}

func TestTransportConfig_WithHost(t *testing.T) {
	cfg := DefaultTransportConfig()
	cfg = cfg.WithHost("host")
	assert.Equal(t, "host", cfg.Host)
}

func TestTransportConfig_WithBasePath(t *testing.T) {
	cfg := DefaultTransportConfig()
	cfg = cfg.WithBasePath("base_path")
	assert.Equal(t, "base_path", cfg.BasePath)
}

func TestTransportConfig_WithSchemes(t *testing.T) {
	cfg := DefaultTransportConfig()
	cfg = cfg.WithSchemes([]string{"schemes"})
	assert.Equal(t, []string{"schemes"}, cfg.Schemes)
}

func TestDNSConfigurationAPI_SetTransport(t *testing.T) {
	cli := NewHTTPClient(strfmt.Default)
	transport := httptransport.New("", "", nil)
	cli.SetTransport(transport)
	assert.Equal(t, transport, cli.Transport)
}
