package address

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

func TestNewAddressCreateParams(t *testing.T) {
	p := NewAddressCreateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAddressCreateParamsWithTimeout(t *testing.T) {
	p := NewAddressCreateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAddressCreateParamsWithContext(t *testing.T) {
	p := NewAddressCreateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAddressCreateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAddressCreateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressCreateParams_WithDefaults(t *testing.T) {
	p := NewAddressCreateParams()
	p = p.WithDefaults()
}

func TestAddressCreateParams_WithTimeout(t *testing.T) {
	p := NewAddressCreateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAddressCreateParams_WithContext(t *testing.T) {
	p := NewAddressCreateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAddressCreateParams_WithHTTPClient(t *testing.T) {
	p := NewAddressCreateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressCreateParams_WithBody(t *testing.T) {
	p := NewAddressCreateParams()
	b := &models.IpamsvcAddress{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}
