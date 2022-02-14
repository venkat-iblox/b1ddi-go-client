package option_filter

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

func TestNewOptionFilterCreateParams(t *testing.T) {
	p := NewOptionFilterCreateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewOptionFilterCreateParamsWithTimeout(t *testing.T) {
	p := NewOptionFilterCreateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewOptionFilterCreateParamsWithContext(t *testing.T) {
	p := NewOptionFilterCreateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewOptionFilterCreateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewOptionFilterCreateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionFilterCreateParams_WithDefaults(t *testing.T) {
	p := NewOptionFilterCreateParams()
	p = p.WithDefaults()
}

func TestOptionFilterCreateParams_WithTimeout(t *testing.T) {
	p := NewOptionFilterCreateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestOptionFilterCreateParams_WithContext(t *testing.T) {
	p := NewOptionFilterCreateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestOptionFilterCreateParams_WithHTTPClient(t *testing.T) {
	p := NewOptionFilterCreateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionFilterCreateParams_WithBody(t *testing.T) {
	p := NewOptionFilterCreateParams()
	b := &models.IpamsvcOptionFilter{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}
