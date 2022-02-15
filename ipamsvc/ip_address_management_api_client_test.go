package ipamsvc

import (
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHTTPClient(t *testing.T) {
	cli := NewHTTPClient(strfmt.Default)

	assert.NotNil(t, cli.Transport)
	assert.NotNil(t, cli.Address)
	assert.NotNil(t, cli.AddressBlock)
	assert.NotNil(t, cli.Asm)
	assert.NotNil(t, cli.DhcpHost)
	assert.NotNil(t, cli.DNSUsage)
	assert.NotNil(t, cli.Filter)
	assert.NotNil(t, cli.FixedAddress)
	assert.NotNil(t, cli.Global)
	assert.NotNil(t, cli.HaGroup)
	assert.NotNil(t, cli.HardwareFilter)
	assert.NotNil(t, cli.IPSpace)
	assert.NotNil(t, cli.IpamHost)
	assert.NotNil(t, cli.OptionCode)
	assert.NotNil(t, cli.OptionFilter)
	assert.NotNil(t, cli.OptionGroup)
	assert.NotNil(t, cli.OptionSpace)
	assert.NotNil(t, cli.RangeOperations)
	assert.NotNil(t, cli.Server)
	assert.NotNil(t, cli.Subnet)
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
