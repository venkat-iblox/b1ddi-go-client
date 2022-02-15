package ipam_host

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

func TestNewIpamHostListParams(t *testing.T) {
	p := NewIpamHostListParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewIpamHostListParamsWithTimeout(t *testing.T) {
	p := NewIpamHostListParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewIpamHostListParamsWithContext(t *testing.T) {
	p := NewIpamHostListParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewIpamHostListParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewIpamHostListParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestIpamHostListParams_WithDefaults(t *testing.T) {
	p := NewIpamHostListParams()
	p = p.WithDefaults()
}

func TestIpamHostListParams_WithTimeout(t *testing.T) {
	p := NewIpamHostListParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestIpamHostListParams_WithContext(t *testing.T) {
	p := NewIpamHostListParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestIpamHostListParams_WithHTTPClient(t *testing.T) {
	p := NewIpamHostListParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestIpamHostListParams_WithFields(t *testing.T) {
	p := NewIpamHostListParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestIpamHostListParams_WithFilter(t *testing.T) {
	p := NewIpamHostListParams()
	p = p.WithFilter(swag.String("key==value"))
	require.NotNil(t, p.Filter)
	assert.Equal(t, "key==value", *p.Filter)
}

func TestIpamHostListParams_WithLimit(t *testing.T) {
	p := NewIpamHostListParams()
	p = p.WithLimit(swag.Int64(20))
	require.NotNil(t, p.Limit)
	assert.Equal(t, int64(20), *p.Limit)
}

func TestIpamHostListParams_WithOffset(t *testing.T) {
	p := NewIpamHostListParams()
	p = p.WithOffset(swag.Int64(20))
	require.NotNil(t, p.Offset)
	assert.Equal(t, int64(20), *p.Offset)
}

func TestIpamHostListParams_WithOrderBy(t *testing.T) {
	p := NewIpamHostListParams()
	p = p.WithOrderBy(swag.String("desc"))
	require.NotNil(t, p.OrderBy)
	assert.Equal(t, "desc", *p.OrderBy)
}

func TestIpamHostListParams_WithPageToken(t *testing.T) {
	p := NewIpamHostListParams()
	p = p.WithPageToken(swag.String("token"))
	require.NotNil(t, p.PageToken)
	assert.Equal(t, "token", *p.PageToken)
}

func TestIpamHostListParams_WithTfilter(t *testing.T) {
	p := NewIpamHostListParams()
	p = p.WithTfilter(swag.String("tfilter"))
	require.NotNil(t, p.Tfilter)
	assert.Equal(t, "tfilter", *p.Tfilter)
}

func TestIpamHostListParams_WithTorderBy(t *testing.T) {
	p := NewIpamHostListParams()
	p = p.WithTorderBy(swag.String("desc"))
	require.NotNil(t, p.TorderBy)
	assert.Equal(t, "desc", *p.TorderBy)
}
