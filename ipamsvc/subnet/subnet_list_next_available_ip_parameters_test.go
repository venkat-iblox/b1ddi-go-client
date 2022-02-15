package subnet

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

func TestNewSubnetListNextAvailableIPParams(t *testing.T) {
	p := NewSubnetListNextAvailableIPParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewSubnetListNextAvailableIPParamsWithTimeout(t *testing.T) {
	p := NewSubnetListNextAvailableIPParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewSubnetListNextAvailableIPParamsWithContext(t *testing.T) {
	p := NewSubnetListNextAvailableIPParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewSubnetListNextAvailableIPParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewSubnetListNextAvailableIPParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestSubnetListNextAvailableIPParams_WithDefaults(t *testing.T) {
	p := NewSubnetListNextAvailableIPParams()
	p = p.WithDefaults()
}

func TestSubnetListNextAvailableIPParams_WithTimeout(t *testing.T) {
	p := NewSubnetListNextAvailableIPParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestSubnetListNextAvailableIPParams_WithContext(t *testing.T) {
	p := NewSubnetListNextAvailableIPParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestSubnetListNextAvailableIPParams_WithHTTPClient(t *testing.T) {
	p := NewSubnetListNextAvailableIPParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestSubnetListNextAvailableIPParams_WithContiguous(t *testing.T) {
	p := NewSubnetListNextAvailableIPParams()
	p = p.WithContiguous(swag.Bool(true))
	require.NotNil(t, p.Contiguous)
	assert.Equal(t, true, *p.Contiguous)
}

func TestSubnetListNextAvailableIPParams_WithCount(t *testing.T) {
	p := NewSubnetListNextAvailableIPParams()
	p = p.WithCount(swag.Int32(5))
	require.NotNil(t, p.Count)
	assert.Equal(t, int32(5), *p.Count)
}

func TestSubnetListNextAvailableIPParams_WithID(t *testing.T) {
	p := NewSubnetListNextAvailableIPParams()
	p = p.WithID("test-id")
	assert.Equal(t, "test-id", p.ID)
}
