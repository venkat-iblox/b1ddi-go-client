package ip_space

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

func TestNewIPSpaceListParams(t *testing.T) {
	p := NewIPSpaceListParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewIPSpaceListParamsWithTimeout(t *testing.T) {
	p := NewIPSpaceListParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewIPSpaceListParamsWithContext(t *testing.T) {
	p := NewIPSpaceListParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewIPSpaceListParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewIPSpaceListParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestIPSpaceListParams_WithDefaults(t *testing.T) {
	p := NewIPSpaceListParams()
	p = p.WithDefaults()
}

func TestIPSpaceListParams_WithTimeout(t *testing.T) {
	p := NewIPSpaceListParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestIPSpaceListParams_WithContext(t *testing.T) {
	p := NewIPSpaceListParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestIPSpaceListParams_WithHTTPClient(t *testing.T) {
	p := NewIPSpaceListParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestIPSpaceListParams_WithFields(t *testing.T) {
	p := NewIPSpaceListParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestIPSpaceListParams_WithFilter(t *testing.T) {
	p := NewIPSpaceListParams()
	p = p.WithFilter(swag.String("key==value"))
	require.NotNil(t, p.Filter)
	assert.Equal(t, "key==value", *p.Filter)
}

func TestIPSpaceListParams_WithLimit(t *testing.T) {
	p := NewIPSpaceListParams()
	p = p.WithLimit(swag.Int64(20))
	require.NotNil(t, p.Limit)
	assert.Equal(t, int64(20), *p.Limit)
}
