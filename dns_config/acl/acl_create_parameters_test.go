package acl

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

func TestNewACLCreateParams(t *testing.T) {
	p := NewACLCreateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewACLCreateParamsWithTimeout(t *testing.T) {
	p := NewACLCreateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewACLCreateParamsWithContext(t *testing.T) {
	p := NewACLCreateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewACLCreateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewACLCreateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestACLCreateParams_WithDefaults(t *testing.T) {
	p := NewACLCreateParams()
	p = p.WithDefaults()
}

func TestACLCreateParams_WithTimeout(t *testing.T) {
	p := NewACLCreateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestACLCreateParams_WithContext(t *testing.T) {
	p := NewACLCreateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestACLCreateParams_WithHTTPClient(t *testing.T) {
	p := NewACLCreateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestACLCreateParams_WithBody(t *testing.T) {
	p := NewACLCreateParams()
	b := &models.ConfigACL{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}
