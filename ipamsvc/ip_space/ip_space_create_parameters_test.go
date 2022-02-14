package ip_space

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

func TestNewIPSpaceCreateParams(t *testing.T) {
	p := NewIPSpaceCreateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewIPSpaceCreateParamsWithTimeout(t *testing.T) {
	p := NewIPSpaceCreateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewIPSpaceCreateParamsWithContext(t *testing.T) {
	p := NewIPSpaceCreateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewIPSpaceCreateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewIPSpaceCreateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestIPSpaceCreateParams_WithDefaults(t *testing.T) {
	p := NewIPSpaceCreateParams()
	p = p.WithDefaults()
}

func TestIPSpaceCreateParams_WithTimeout(t *testing.T) {
	p := NewIPSpaceCreateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestIPSpaceCreateParams_WithContext(t *testing.T) {
	p := NewIPSpaceCreateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestIPSpaceCreateParams_WithHTTPClient(t *testing.T) {
	p := NewIPSpaceCreateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestIPSpaceCreateParams_WithBody(t *testing.T) {
	p := NewIPSpaceCreateParams()
	b := &models.IpamsvcIPSpace{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}
