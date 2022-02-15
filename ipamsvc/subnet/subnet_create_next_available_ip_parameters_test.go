package subnet

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewSubnetCreateNextAvailableIPParams(t *testing.T) {
	p := NewSubnetCreateNextAvailableIPParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewSubnetCreateNextAvailableIPParamsWithTimeout(t *testing.T) {
	p := NewSubnetCreateNextAvailableIPParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewSubnetCreateNextAvailableIPParamsWithContext(t *testing.T) {
	p := NewSubnetCreateNextAvailableIPParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewSubnetCreateNextAvailableIPParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewSubnetCreateNextAvailableIPParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestSubnetCreateNextAvailableIPParams_WithDefaults(t *testing.T) {
	p := NewSubnetCreateNextAvailableIPParams()
	p = p.WithDefaults()
}

func TestSubnetCreateNextAvailableIPParams_WithTimeout(t *testing.T) {
	p := NewSubnetCreateNextAvailableIPParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestSubnetCreateNextAvailableIPParams_WithContext(t *testing.T) {
	p := NewSubnetCreateNextAvailableIPParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestSubnetCreateNextAvailableIPParams_WithHTTPClient(t *testing.T) {
	p := NewSubnetCreateNextAvailableIPParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestSubnetCreateNextAvailableIPParams_WithID(t *testing.T) {
	p := NewSubnetCreateNextAvailableIPParams()
	p = p.WithID("test-id")
	assert.Equal(t, "test-id", p.ID)
}
