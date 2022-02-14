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

func TestNewSubnetListParams(t *testing.T) {
	p := NewSubnetListParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewSubnetListParamsWithTimeout(t *testing.T) {
	p := NewSubnetListParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewSubnetListParamsWithContext(t *testing.T) {
	p := NewSubnetListParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewSubnetListParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewSubnetListParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestSubnetListParams_WithDefaults(t *testing.T) {
	p := NewSubnetListParams()
	p = p.WithDefaults()
}

func TestSubnetListParams_WithTimeout(t *testing.T) {
	p := NewSubnetListParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestSubnetListParams_WithContext(t *testing.T) {
	p := NewSubnetListParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestSubnetListParams_WithHTTPClient(t *testing.T) {
	p := NewSubnetListParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestSubnetListParams_WithFields(t *testing.T) {
	p := NewSubnetListParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestSubnetListParams_WithFilter(t *testing.T) {
	p := NewSubnetListParams()
	p = p.WithFilter(swag.String("key==value"))
	require.NotNil(t, p.Filter)
	assert.Equal(t, "key==value", *p.Filter)
}

func TestSubnetListParams_WithLimit(t *testing.T) {
	p := NewSubnetListParams()
	p = p.WithLimit(swag.Int64(20))
	require.NotNil(t, p.Limit)
	assert.Equal(t, int64(20), *p.Limit)
}
