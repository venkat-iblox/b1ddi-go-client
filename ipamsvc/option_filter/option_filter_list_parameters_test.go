package option_filter

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

func TestNewOptionFilterListParams(t *testing.T) {
	p := NewOptionFilterListParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewOptionFilterListParamsWithTimeout(t *testing.T) {
	p := NewOptionFilterListParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewOptionFilterListParamsWithContext(t *testing.T) {
	p := NewOptionFilterListParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewOptionFilterListParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewOptionFilterListParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionFilterListParams_WithDefaults(t *testing.T) {
	p := NewOptionFilterListParams()
	p = p.WithDefaults()
}

func TestOptionFilterListParams_WithTimeout(t *testing.T) {
	p := NewOptionFilterListParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestOptionFilterListParams_WithContext(t *testing.T) {
	p := NewOptionFilterListParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestOptionFilterListParams_WithHTTPClient(t *testing.T) {
	p := NewOptionFilterListParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionFilterListParams_WithFields(t *testing.T) {
	p := NewOptionFilterListParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestOptionFilterListParams_WithFilter(t *testing.T) {
	p := NewOptionFilterListParams()
	p = p.WithFilter(swag.String("key==value"))
	require.NotNil(t, p.Filter)
	assert.Equal(t, "key==value", *p.Filter)
}

func TestOptionFilterListParams_WithLimit(t *testing.T) {
	p := NewOptionFilterListParams()
	p = p.WithLimit(swag.Int64(20))
	require.NotNil(t, p.Limit)
	assert.Equal(t, int64(20), *p.Limit)
}

func TestOptionFilterListParams_WithOffset(t *testing.T) {
	p := NewOptionFilterListParams()
	p = p.WithOffset(swag.Int64(20))
	require.NotNil(t, p.Offset)
	assert.Equal(t, int64(20), *p.Offset)
}

func TestOptionFilterListParams_WithOrderBy(t *testing.T) {
	p := NewOptionFilterListParams()
	p = p.WithOrderBy(swag.String("desc"))
	require.NotNil(t, p.OrderBy)
	assert.Equal(t, "desc", *p.OrderBy)
}

func TestOptionFilterListParams_WithPageToken(t *testing.T) {
	p := NewOptionFilterListParams()
	p = p.WithPageToken(swag.String("token"))
	require.NotNil(t, p.PageToken)
	assert.Equal(t, "token", *p.PageToken)
}

func TestOptionFilterListParams_WithTfilter(t *testing.T) {
	p := NewOptionFilterListParams()
	p = p.WithTfilter(swag.String("tfilter"))
	require.NotNil(t, p.Tfilter)
	assert.Equal(t, "tfilter", *p.Tfilter)
}

func TestOptionFilterListParams_WithTorderBy(t *testing.T) {
	p := NewOptionFilterListParams()
	p = p.WithTorderBy(swag.String("desc"))
	require.NotNil(t, p.TorderBy)
	assert.Equal(t, "desc", *p.TorderBy)
}
