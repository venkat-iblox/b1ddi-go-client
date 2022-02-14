package subnet

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

func TestNewSubnetCreateParams(t *testing.T) {
	p := NewSubnetCreateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewSubnetCreateParamsWithTimeout(t *testing.T) {
	p := NewSubnetCreateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewSubnetCreateParamsWithContext(t *testing.T) {
	p := NewSubnetCreateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewSubnetCreateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewSubnetCreateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestSubnetCreateParams_WithDefaults(t *testing.T) {
	p := NewSubnetCreateParams()
	p = p.WithDefaults()
}

func TestSubnetCreateParams_WithTimeout(t *testing.T) {
	p := NewSubnetCreateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestSubnetCreateParams_WithContext(t *testing.T) {
	p := NewSubnetCreateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestSubnetCreateParams_WithHTTPClient(t *testing.T) {
	p := NewSubnetCreateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestSubnetCreateParams_WithBody(t *testing.T) {
	p := NewSubnetCreateParams()
	b := &models.IpamsvcSubnet{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}
