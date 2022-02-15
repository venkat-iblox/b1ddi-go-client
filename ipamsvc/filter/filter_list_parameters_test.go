package filter

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

func TestNewFilterListParams(t *testing.T) {
	p := NewFilterListParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewFilterListParamsWithTimeout(t *testing.T) {
	p := NewFilterListParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewFilterListParamsWithContext(t *testing.T) {
	p := NewFilterListParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewFilterListParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewFilterListParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestFilterListParams_WithDefaults(t *testing.T) {
	p := NewFilterListParams()
	p = p.WithDefaults()
}

func TestFilterListParams_WithTimeout(t *testing.T) {
	p := NewFilterListParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestFilterListParams_WithContext(t *testing.T) {
	p := NewFilterListParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestFilterListParams_WithHTTPClient(t *testing.T) {
	p := NewFilterListParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestFilterListParams_WithFields(t *testing.T) {
	p := NewFilterListParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestFilterListParams_WithFilter(t *testing.T) {
	p := NewFilterListParams()
	p = p.WithFilter(swag.String("key==value"))
	require.NotNil(t, p.Filter)
	assert.Equal(t, "key==value", *p.Filter)
}

func TestFilterListParams_WithLimit(t *testing.T) {
	p := NewFilterListParams()
	p = p.WithLimit(swag.Int64(20))
	require.NotNil(t, p.Limit)
	assert.Equal(t, int64(20), *p.Limit)
}

func TestFilterListParams_WithOffset(t *testing.T) {
	p := NewFilterListParams()
	p = p.WithOffset(swag.Int64(20))
	require.NotNil(t, p.Offset)
	assert.Equal(t, int64(20), *p.Offset)
}

func TestFilterListParams_WithOrderBy(t *testing.T) {
	p := NewFilterListParams()
	p = p.WithOrderBy(swag.String("desc"))
	require.NotNil(t, p.OrderBy)
	assert.Equal(t, "desc", *p.OrderBy)
}

func TestFilterListParams_WithPageToken(t *testing.T) {
	p := NewFilterListParams()
	p = p.WithPageToken(swag.String("token"))
	require.NotNil(t, p.PageToken)
	assert.Equal(t, "token", *p.PageToken)
}

func TestFilterListParams_WithTfilter(t *testing.T) {
	p := NewFilterListParams()
	p = p.WithTfilter(swag.String("tfilter"))
	require.NotNil(t, p.Tfilter)
	assert.Equal(t, "tfilter", *p.Tfilter)
}

func TestFilterListParams_WithTOrderBy(t *testing.T) {
	p := NewFilterListParams()
	p = p.WithTorderBy(swag.String("desc"))
	require.NotNil(t, p.TorderBy)
	assert.Equal(t, "desc", *p.TorderBy)
}
