package range_operations

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewRangeDeleteParams(t *testing.T) {
	p := NewRangeDeleteParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewRangeDeleteParamsWithTimeout(t *testing.T) {
	p := NewRangeDeleteParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewRangeDeleteParamsWithContext(t *testing.T) {
	p := NewRangeDeleteParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewRangeDeleteParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewRangeDeleteParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestRangeDeleteParams_WithDefaults(t *testing.T) {
	p := NewRangeDeleteParams()
	p = p.WithDefaults()
}

func TestRangeDeleteParams_WithTimeout(t *testing.T) {
	p := NewRangeDeleteParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestRangeDeleteParams_WithContext(t *testing.T) {
	p := NewRangeDeleteParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestRangeDeleteParams_WithHTTPClient(t *testing.T) {
	p := NewRangeDeleteParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestRangeDeleteParams_WithID(t *testing.T) {
	p := NewRangeDeleteParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
