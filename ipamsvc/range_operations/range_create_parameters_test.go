package range_operations

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

func TestNewRangeCreateParams(t *testing.T) {
	p := NewRangeCreateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewRangeCreateParamsWithTimeout(t *testing.T) {
	p := NewRangeCreateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewRangeCreateParamsWithContext(t *testing.T) {
	p := NewRangeCreateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewRangeCreateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewRangeCreateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestRangeCreateParams_WithDefaults(t *testing.T) {
	p := NewRangeCreateParams()
	p = p.WithDefaults()
}

func TestRangeCreateParams_WithTimeout(t *testing.T) {
	p := NewRangeCreateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestRangeCreateParams_WithContext(t *testing.T) {
	p := NewRangeCreateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestRangeCreateParams_WithHTTPClient(t *testing.T) {
	p := NewRangeCreateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestRangeCreateParams_WithBody(t *testing.T) {
	p := NewRangeCreateParams()
	b := &models.IpamsvcRange{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}
