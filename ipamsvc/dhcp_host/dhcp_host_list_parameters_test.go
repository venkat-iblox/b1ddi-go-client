package dhcp_host

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

func TestNewDhcpHostListParams(t *testing.T) {
	p := NewDhcpHostListParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewDhcpHostListParamsWithTimeout(t *testing.T) {
	p := NewDhcpHostListParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewDhcpHostListParamsWithContext(t *testing.T) {
	p := NewDhcpHostListParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewDhcpHostListParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewDhcpHostListParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestDhcpHostListParams_WithDefaults(t *testing.T) {
	p := NewDhcpHostListParams()
	p = p.WithDefaults()
}

func TestDhcpHostListParams_WithTimeout(t *testing.T) {
	p := NewDhcpHostListParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestDhcpHostListParams_WithContext(t *testing.T) {
	p := NewDhcpHostListParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestDhcpHostListParams_WithHTTPClient(t *testing.T) {
	p := NewDhcpHostListParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestDhcpHostListParams_WithFields(t *testing.T) {
	p := NewDhcpHostListParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestDhcpHostListParams_WithFilter(t *testing.T) {
	p := NewDhcpHostListParams()
	p = p.WithFilter(swag.String("key==value"))
	require.NotNil(t, p.Filter)
	assert.Equal(t, "key==value", *p.Filter)
}

func TestDhcpHostListParams_WithLimit(t *testing.T) {
	p := NewDhcpHostListParams()
	p = p.WithLimit(swag.Int64(20))
	require.NotNil(t, p.Limit)
	assert.Equal(t, int64(20), *p.Limit)
}

func TestDhcpHostListParams_WithOffset(t *testing.T) {
	p := NewDhcpHostListParams()
	p = p.WithOffset(swag.Int64(20))
	require.NotNil(t, p.Offset)
	assert.Equal(t, int64(20), *p.Offset)
}

func TestDhcpHostListParams_WithOrderBy(t *testing.T) {
	p := NewDhcpHostListParams()
	p = p.WithOrderBy(swag.String("desc"))
	require.NotNil(t, p.OrderBy)
	assert.Equal(t, "desc", *p.OrderBy)
}

func TestDhcpHostListParams_WithPageToken(t *testing.T) {
	p := NewDhcpHostListParams()
	p = p.WithPageToken(swag.String("page-token"))
	require.NotNil(t, p.PageToken)
	assert.Equal(t, "page-token", *p.PageToken)
}

func TestDhcpHostListParams_WithTfilter(t *testing.T) {
	p := NewDhcpHostListParams()
	p = p.WithTfilter(swag.String("tfilter"))
	require.NotNil(t, p.Tfilter)
	assert.Equal(t, "tfilter", *p.Tfilter)
}

func TestDhcpHostListParams_WithTorderBy(t *testing.T) {
	p := NewDhcpHostListParams()
	p = p.WithTorderBy(swag.String("desc"))
	require.NotNil(t, p.TorderBy)
	assert.Equal(t, "desc", *p.TorderBy)
}
