package address_block

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

func TestNewAddressBlockListNextAvailableABParams(t *testing.T) {
	p := NewAddressBlockListNextAvailableABParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAddressBlockListNextAvailableABParamsWithTimeout(t *testing.T) {
	p := NewAddressBlockListNextAvailableABParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAddressBlockListNextAvailableABParamsWithContext(t *testing.T) {
	p := NewAddressBlockListNextAvailableABParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAddressBlockListNextAvailableABParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAddressBlockListNextAvailableABParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressBlockListNextAvailableABParams_WithDefaults(t *testing.T) {
	p := NewAddressBlockListNextAvailableABParams()
	p = p.WithDefaults()
}

func TestAddressBlockListNextAvailableABParams_WithTimeout(t *testing.T) {
	p := NewAddressBlockListNextAvailableABParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAddressBlockListNextAvailableABParams_WithContext(t *testing.T) {
	p := NewAddressBlockListNextAvailableABParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAddressBlockListNextAvailableABParams_WithHTTPClient(t *testing.T) {
	p := NewAddressBlockListNextAvailableABParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressBlockListNextAvailableABParams_WithID(t *testing.T) {
	p := NewAddressBlockListNextAvailableABParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}

func TestAddressBlockListNextAvailableABParams_WithName(t *testing.T) {
	p := NewAddressBlockListNextAvailableABParams()
	p = p.WithName(swag.String("name"))
	require.NotNil(t, p.Name)
	assert.Equal(t, "name", *p.Name)
}
