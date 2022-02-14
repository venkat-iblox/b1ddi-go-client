package fixed_address

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

func TestNewFixedAddressCreateParams(t *testing.T) {
	p := NewFixedAddressCreateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewFixedAddressCreateParamsWithTimeout(t *testing.T) {
	p := NewFixedAddressCreateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewFixedFixedAddressCreateParamsWithContext(t *testing.T) {
	p := NewFixedAddressCreateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewFixedAddressCreateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewFixedAddressCreateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestFixedAddressCreateParams_WithDefaults(t *testing.T) {
	p := NewFixedAddressCreateParams()
	p = p.WithDefaults()
}

func TestFixedAddressCreateParams_WithTimeout(t *testing.T) {
	p := NewFixedAddressCreateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestFixedAddressCreateParams_WithContext(t *testing.T) {
	p := NewFixedAddressCreateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestFixedAddressCreateParams_WithHTTPClient(t *testing.T) {
	p := NewFixedAddressCreateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestFixedAddressCreateParams_WithBody(t *testing.T) {
	p := NewFixedAddressCreateParams()
	b := &models.IpamsvcFixedAddress{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}
