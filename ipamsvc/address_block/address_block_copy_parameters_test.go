package address_block

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

func TestNewAddressBlockCopyParams(t *testing.T) {
	p := NewAddressBlockCopyParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAddressBlockCopyParamsWithTimeout(t *testing.T) {
	p := NewAddressBlockCopyParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAddressBlockCopyParamsWithContext(t *testing.T) {
	p := NewAddressBlockCopyParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAddressBlockCopyParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAddressBlockCopyParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressBlockCopyParams_WithDefaults(t *testing.T) {
	p := NewAddressBlockCopyParams()
	p = p.WithDefaults()
}

func TestAddressBlockCopyParams_WithTimeout(t *testing.T) {
	p := NewAddressBlockCopyParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAddressBlockCopyParams_WithContext(t *testing.T) {
	p := NewAddressBlockCopyParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAddressBlockCopyParams_WithHTTPClient(t *testing.T) {
	p := NewAddressBlockCopyParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressBlockCopyParams_WithBody(t *testing.T) {
	p := NewAddressBlockCopyParams()
	b := &models.IpamsvcCopyAddressBlock{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}

func TestAddressBlockCopyParams_WithID(t *testing.T) {
	p := NewAddressBlockCopyParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
