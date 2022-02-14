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

func TestNewAddressBlockUpdateParams(t *testing.T) {
	p := NewAddressBlockUpdateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAddressBlockUpdateParamsWithTimeout(t *testing.T) {
	p := NewAddressBlockUpdateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAddressBlockUpdateParamsWithContext(t *testing.T) {
	p := NewAddressBlockUpdateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAddressBlockUpdateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAddressBlockUpdateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressBlockUpdateParams_WithDefaults(t *testing.T) {
	p := NewAddressBlockUpdateParams()
	p = p.WithDefaults()
}

func TestAddressBlockUpdateParams_WithTimeout(t *testing.T) {
	p := NewAddressBlockUpdateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAddressBlockUpdateParams_WithContext(t *testing.T) {
	p := NewAddressBlockUpdateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAddressBlockUpdateParams_WithHTTPClient(t *testing.T) {
	p := NewAddressBlockUpdateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressBlockUpdateParams_WithBody(t *testing.T) {
	p := NewAddressBlockUpdateParams()
	b := &models.IpamsvcAddressBlock{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}

func TestAddressBlockUpdateParams_WithID(t *testing.T) {
	p := NewAddressBlockUpdateParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
