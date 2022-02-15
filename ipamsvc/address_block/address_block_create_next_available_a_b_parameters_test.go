package address_block

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewAddressBlockCreateNextAvailableABParams(t *testing.T) {
	p := NewAddressBlockCreateNextAvailableABParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAddressBlockCreateNextAvailableABParamsWithTimeout(t *testing.T) {
	p := NewAddressBlockCreateNextAvailableABParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAddressBlockCreateNextAvailableABParamsWithContext(t *testing.T) {
	p := NewAddressBlockCreateNextAvailableABParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAddressBlockCreateNextAvailableABParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAddressBlockCreateNextAvailableABParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressBlockCreateNextAvailableABParams_WithDefaults(t *testing.T) {
	p := NewAddressBlockCreateNextAvailableABParams()
	p = p.WithDefaults()
}

func TestAddressBlockCreateNextAvailableABParams_WithTimeout(t *testing.T) {
	p := NewAddressBlockCreateNextAvailableABParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAddressBlockCreateNextAvailableABParams_WithContext(t *testing.T) {
	p := NewAddressBlockCreateNextAvailableABParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAddressBlockCreateNextAvailableABParams_WithHTTPClient(t *testing.T) {
	p := NewAddressBlockCreateNextAvailableABParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressBlockCreateNextAvailableABParams_WithID(t *testing.T) {
	p := NewAddressBlockCreateNextAvailableABParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
