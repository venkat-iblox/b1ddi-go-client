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

func TestNewRangeCreateNextAvailableIPParams(t *testing.T) {
	p := NewRangeCreateNextAvailableIPParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewRangeCreateNextAvailableIPParamsWithTimeout(t *testing.T) {
	p := NewRangeCreateNextAvailableIPParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewRangeCreateNextAvailableIPParamsWithContext(t *testing.T) {
	p := NewRangeCreateNextAvailableIPParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewRangeCreateNextAvailableIPParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewRangeCreateNextAvailableIPParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestRangeCreateNextAvailableIPParams_WithDefaults(t *testing.T) {
	p := NewRangeCreateNextAvailableIPParams()
	p = p.WithDefaults()
}

func TestRangeCreateNextAvailableIPParams_WithTimeout(t *testing.T) {
	p := NewRangeCreateNextAvailableIPParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestRangeCreateNextAvailableIPParams_WithContext(t *testing.T) {
	p := NewRangeCreateNextAvailableIPParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestRangeCreateNextAvailableIPParams_WithHTTPClient(t *testing.T) {
	p := NewRangeCreateNextAvailableIPParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestRangeCreateNextAvailableIPParams_WithID(t *testing.T) {
	p := NewRangeCreateNextAvailableIPParams()
	p = p.WithID("test-id")
	assert.Equal(t, "test-id", p.ID)
}
