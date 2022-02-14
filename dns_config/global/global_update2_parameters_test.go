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

func TestNewGlobalUpdate2Params(t *testing.T) {
	p := NewGlobalUpdate2Params()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewGlobalUpdate2ParamsWithTimeout(t *testing.T) {
	p := NewGlobalUpdate2ParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewGlobalUpdate2ParamsWithContext(t *testing.T) {
	p := NewGlobalUpdate2ParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewGlobalUpdate2ParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewGlobalUpdate2ParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestGlobalUpdate2Params_WithDefaults(t *testing.T) {
	p := NewGlobalUpdate2Params()
	p = p.WithDefaults()
}

func TestGlobalUpdate2Params_WithTimeout(t *testing.T) {
	p := NewGlobalUpdate2Params()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestGlobalUpdate2Params_WithContext(t *testing.T) {
	p := NewGlobalUpdate2Params()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestGlobalUpdate2Params_WithHTTPClient(t *testing.T) {
	p := NewGlobalUpdate2Params()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestGlobalUpdate2Params_WithBody(t *testing.T) {
	p := NewGlobalUpdate2Params()
	b := &models.ConfigGlobal{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}

func TestGlobalUpdate2Params_WithID(t *testing.T) {
	p := NewGlobalUpdate2Params()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
