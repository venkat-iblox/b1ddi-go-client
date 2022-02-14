package global

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

func TestNewGlobalUpdateParams(t *testing.T) {
	p := NewGlobalUpdateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewGlobalUpdateParamsWithTimeout(t *testing.T) {
	p := NewGlobalUpdateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewGlobalUpdateParamsWithContext(t *testing.T) {
	p := NewGlobalUpdateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewGlobalUpdateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewGlobalUpdateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestGlobalUpdateParams_WithDefaults(t *testing.T) {
	p := NewGlobalUpdateParams()
	p = p.WithDefaults()
}

func TestGlobalUpdateParams_WithTimeout(t *testing.T) {
	p := NewGlobalUpdateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestGlobalUpdateParams_WithContext(t *testing.T) {
	p := NewGlobalUpdateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestGlobalUpdateParams_WithHTTPClient(t *testing.T) {
	p := NewGlobalUpdateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestGlobalUpdateParams_WithBody(t *testing.T) {
	p := NewGlobalUpdateParams()
	b := &models.IpamsvcGlobal{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}
