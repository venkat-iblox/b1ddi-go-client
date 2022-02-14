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

func TestNewAddressBlockCreateParams(t *testing.T) {
	p := NewAddressBlockCreateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAddressBlockCreateParamsWithTimeout(t *testing.T) {
	p := NewAddressBlockCreateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAddressBlockCreateParamsWithContext(t *testing.T) {
	p := NewAddressBlockCreateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAddressBlockCreateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAddressBlockCreateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressBlockCreateParams_WithDefaults(t *testing.T) {
	p := NewAddressBlockCreateParams()
	p = p.WithDefaults()
}

func TestAddressBlockCreateParams_WithTimeout(t *testing.T) {
	p := NewAddressBlockCreateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAddressBlockCreateParams_WithContext(t *testing.T) {
	p := NewAddressBlockCreateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAddressBlockCreateParams_WithHTTPClient(t *testing.T) {
	p := NewAddressBlockCreateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressBlockCreateParams_WithBody(t *testing.T) {
	p := NewAddressBlockCreateParams()
	b := &models.IpamsvcAddressBlock{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}
