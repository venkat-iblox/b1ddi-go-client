package forward_nsg

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

func TestNewForwardNsgCreateParams(t *testing.T) {
	p := NewForwardNsgCreateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewForwardNsgCreateParamsWithTimeout(t *testing.T) {
	p := NewForwardNsgCreateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewForwardNsgCreateParamsWithContext(t *testing.T) {
	p := NewForwardNsgCreateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewForwardNsgCreateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewForwardNsgCreateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestForwardNsgCreateParams_WithDefaults(t *testing.T) {
	p := NewForwardNsgCreateParams()
	p = p.WithDefaults()
}

func TestForwardNsgCreateParams_WithTimeout(t *testing.T) {
	p := NewForwardNsgCreateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestForwardNsgCreateParams_WithContext(t *testing.T) {
	p := NewForwardNsgCreateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestForwardNsgCreateParams_WithHTTPClient(t *testing.T) {
	p := NewForwardNsgCreateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestForwardNsgCreateParams_WithBody(t *testing.T) {
	p := NewForwardNsgCreateParams()
	b := &models.ConfigForwardNSG{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}
