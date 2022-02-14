package option_space

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

func TestNewOptionSpaceCreateParams(t *testing.T) {
	p := NewOptionSpaceCreateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewOptionSpaceCreateParamsWithTimeout(t *testing.T) {
	p := NewOptionSpaceCreateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewOptionSpaceCreateParamsWithContext(t *testing.T) {
	p := NewOptionSpaceCreateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewOptionSpaceCreateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewOptionSpaceCreateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionSpaceCreateParams_WithDefaults(t *testing.T) {
	p := NewOptionSpaceCreateParams()
	p = p.WithDefaults()
}

func TestOptionSpaceCreateParams_WithTimeout(t *testing.T) {
	p := NewOptionSpaceCreateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestOptionSpaceCreateParams_WithContext(t *testing.T) {
	p := NewOptionSpaceCreateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestOptionSpaceCreateParams_WithHTTPClient(t *testing.T) {
	p := NewOptionSpaceCreateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionSpaceCreateParams_WithBody(t *testing.T) {
	p := NewOptionSpaceCreateParams()
	b := &models.IpamsvcOptionSpace{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}
