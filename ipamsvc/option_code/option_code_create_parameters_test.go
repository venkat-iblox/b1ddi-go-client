package option_code

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

func TestNewOptionCodeCreateParams(t *testing.T) {
	p := NewOptionCodeCreateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewOptionCodeCreateParamsWithTimeout(t *testing.T) {
	p := NewOptionCodeCreateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewOptionCodeCreateParamsWithContext(t *testing.T) {
	p := NewOptionCodeCreateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewOptionCodeCreateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewOptionCodeCreateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionCodeCreateParams_WithDefaults(t *testing.T) {
	p := NewOptionCodeCreateParams()
	p = p.WithDefaults()
}

func TestOptionCodeCreateParams_WithTimeout(t *testing.T) {
	p := NewOptionCodeCreateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestOptionCodeCreateParams_WithContext(t *testing.T) {
	p := NewOptionCodeCreateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestOptionCodeCreateParams_WithHTTPClient(t *testing.T) {
	p := NewOptionCodeCreateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionCodeCreateParams_WithBody(t *testing.T) {
	p := NewOptionCodeCreateParams()
	b := &models.IpamsvcOptionCode{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}
