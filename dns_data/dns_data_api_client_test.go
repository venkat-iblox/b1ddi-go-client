package dns_data

import (
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHTTPClient(t *testing.T) {
	cli := NewHTTPClient(strfmt.Default)

	assert.NotNil(t, cli.Transport)
	assert.NotNil(t, cli.Record)
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
