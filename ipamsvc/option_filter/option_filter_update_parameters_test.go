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

func TestNewOptionFilterUpdateParams(t *testing.T) {
	p := NewOptionFilterUpdateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewOptionFilterUpdateParamsWithTimeout(t *testing.T) {
	p := NewOptionFilterUpdateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewOptionFilterUpdateParamsWithContext(t *testing.T) {
	p := NewOptionFilterUpdateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewOptionFilterUpdateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewOptionFilterUpdateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionFilterUpdateParams_WithDefaults(t *testing.T) {
	p := NewOptionFilterUpdateParams()
	p = p.WithDefaults()
}

func TestOptionFilterUpdateParams_WithTimeout(t *testing.T) {
	p := NewOptionFilterUpdateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestOptionFilterUpdateParams_WithContext(t *testing.T) {
	p := NewOptionFilterUpdateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestOptionFilterUpdateParams_WithHTTPClient(t *testing.T) {
	p := NewOptionFilterUpdateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionFilterUpdateParams_WithBody(t *testing.T) {
	p := NewOptionFilterUpdateParams()
	b := &models.IpamsvcOptionFilter{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}

func TestOptionFilterUpdateParams_WithID(t *testing.T) {
	p := NewOptionFilterUpdateParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
