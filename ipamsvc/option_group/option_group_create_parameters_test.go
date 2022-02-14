package option_group

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

func TestNewOptionGroupCreateParams(t *testing.T) {
	p := NewOptionGroupCreateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewOptionGroupCreateParamsWithTimeout(t *testing.T) {
	p := NewOptionGroupCreateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewOptionGroupCreateParamsWithContext(t *testing.T) {
	p := NewOptionGroupCreateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewOptionGroupCreateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewOptionGroupCreateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionGroupCreateParams_WithDefaults(t *testing.T) {
	p := NewOptionGroupCreateParams()
	p = p.WithDefaults()
}

func TestOptionGroupCreateParams_WithTimeout(t *testing.T) {
	p := NewOptionGroupCreateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestOptionGroupCreateParams_WithContext(t *testing.T) {
	p := NewOptionGroupCreateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestOptionGroupCreateParams_WithHTTPClient(t *testing.T) {
	p := NewOptionGroupCreateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionGroupCreateParams_WithBody(t *testing.T) {
	p := NewOptionGroupCreateParams()
	b := &models.IpamsvcOptionGroup{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}
