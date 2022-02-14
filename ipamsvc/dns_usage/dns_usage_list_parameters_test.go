package dns_usage

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

func TestNewDNSUsageListParams(t *testing.T) {
	p := NewDNSUsageListParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewDNSUsageListParamsWithTimeout(t *testing.T) {
	p := NewDNSUsageListParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewDNSUsageListParamsWithContext(t *testing.T) {
	p := NewDNSUsageListParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewDNSUsageListParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewDNSUsageListParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestDNSUsageListParams_WithDefaults(t *testing.T) {
	p := NewDNSUsageListParams()
	p = p.WithDefaults()
}

func TestDNSUsageListParams_WithTimeout(t *testing.T) {
	p := NewDNSUsageListParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestDNSUsageListParams_WithContext(t *testing.T) {
	p := NewDNSUsageListParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestDNSUsageListParams_WithHTTPClient(t *testing.T) {
	p := NewDNSUsageListParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestDNSUsageListParams_WithFields(t *testing.T) {
	p := NewDNSUsageListParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestDNSUsageListParams_WithFilter(t *testing.T) {
	p := NewDNSUsageListParams()
	p = p.WithFilter(swag.String("key==value"))
	require.NotNil(t, p.Filter)
	assert.Equal(t, "key==value", *p.Filter)
}

func TestDNSUsageListParams_WithLimit(t *testing.T) {
	p := NewDNSUsageListParams()
	p = p.WithLimit(swag.Int64(20))
	require.NotNil(t, p.Limit)
	assert.Equal(t, int64(20), *p.Limit)
}

func TestDNSUsageListParams_WithOffset(t *testing.T) {
	p := NewDNSUsageListParams()
	p = p.WithOffset(swag.Int64(20))
	require.NotNil(t, p.Offset)
	assert.Equal(t, int64(20), *p.Offset)
}

func TestDNSUsageListParams_WithOrderBy(t *testing.T) {
	p := NewDNSUsageListParams()
	p = p.WithOrderBy(swag.String("desc"))
	require.NotNil(t, p.OrderBy)
	assert.Equal(t, "desc", *p.OrderBy)
}

func TestDNSUsageListParams_WithPageToken(t *testing.T) {
	p := NewDNSUsageListParams()
	p = p.WithPageToken(swag.String("page-token"))
	require.NotNil(t, p.PageToken)
	assert.Equal(t, "page-token", *p.PageToken)
}
