package ha_group

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewHaGroupDeleteParams(t *testing.T) {
	p := NewHaGroupDeleteParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewHaGroupDeleteParamsWithTimeout(t *testing.T) {
	p := NewHaGroupDeleteParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewHaGroupDeleteParamsWithContext(t *testing.T) {
	p := NewHaGroupDeleteParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewHaGroupDeleteParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewHaGroupDeleteParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestHaGroupDeleteParams_WithDefaults(t *testing.T) {
	p := NewHaGroupDeleteParams()
	p = p.WithDefaults()
}

func TestHaGroupDeleteParams_WithTimeout(t *testing.T) {
	p := NewHaGroupDeleteParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestHaGroupDeleteParams_WithContext(t *testing.T) {
	p := NewHaGroupDeleteParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestHaGroupDeleteParams_WithHTTPClient(t *testing.T) {
	p := NewHaGroupDeleteParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestHaGroupDeleteParams_WithID(t *testing.T) {
	p := NewHaGroupDeleteParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
