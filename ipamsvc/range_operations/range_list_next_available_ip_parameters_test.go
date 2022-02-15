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

func TestNewRangeListNextAvailableIPParams(t *testing.T) {
	p := NewRangeListNextAvailableIPParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewRangeListNextAvailableIPParamsWithTimeout(t *testing.T) {
	p := NewRangeListNextAvailableIPParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewRangeListNextAvailableIPParamsWithContext(t *testing.T) {
	p := NewRangeListNextAvailableIPParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewRangeListNextAvailableIPParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewRangeListNextAvailableIPParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestRangeListNextAvailableIPParams_WithDefaults(t *testing.T) {
	p := NewRangeListNextAvailableIPParams()
	p = p.WithDefaults()
}

func TestRangeListNextAvailableIPParams_WithTimeout(t *testing.T) {
	p := NewRangeListNextAvailableIPParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestRangeListNextAvailableIPParams_WithContext(t *testing.T) {
	p := NewRangeListNextAvailableIPParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestRangeListNextAvailableIPParams_WithHTTPClient(t *testing.T) {
	p := NewRangeListNextAvailableIPParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestRangeListNextAvailableIPParams_WithContiguous(t *testing.T) {
	p := NewRangeListNextAvailableIPParams()
	p = p.WithContiguous(swag.Bool(true))
	require.NotNil(t, p.Contiguous)
	assert.Equal(t, true, *p.Contiguous)
}

func TestRangeListNextAvailableIPParams_WithCount(t *testing.T) {
	p := NewRangeListNextAvailableIPParams()
	p = p.WithCount(swag.Int32(5))
	require.NotNil(t, p.Count)
	assert.Equal(t, int32(5), *p.Count)
}

func TestRangeListNextAvailableIPParams_WithID(t *testing.T) {
	p := NewRangeListNextAvailableIPParams()
	p = p.WithID("test-id")
	assert.Equal(t, "test-id", p.ID)
}
