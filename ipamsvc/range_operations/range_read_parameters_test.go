package range_operations

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewRangeReadParams(t *testing.T) {
	p := NewRangeReadParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewRangeReadParamsWithTimeout(t *testing.T) {
	p := NewRangeReadParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewRangeReadParamsWithContext(t *testing.T) {
	p := NewRangeReadParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewRangeReadParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewRangeReadParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestRangeReadParams_WithDefaults(t *testing.T) {
	p := NewRangeReadParams()
	p = p.WithDefaults()
}

func TestRangeReadParams_WithTimeout(t *testing.T) {
	p := NewRangeReadParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestRangeReadParams_WithContext(t *testing.T) {
	p := NewRangeReadParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestRangeReadParams_WithHTTPClient(t *testing.T) {
	p := NewRangeReadParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestRangeReadParams_WithFields(t *testing.T) {
	p := NewRangeReadParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestRangeReadParams_WithID(t *testing.T) {
	p := NewRangeReadParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
