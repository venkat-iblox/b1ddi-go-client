package range_operations

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

func TestNewRangeListParams(t *testing.T) {
	p := NewRangeListParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewRangeListParamsWithTimeout(t *testing.T) {
	p := NewRangeListParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewRangeListParamsWithContext(t *testing.T) {
	p := NewRangeListParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewRangeListParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewRangeListParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestRangeListParams_WithDefaults(t *testing.T) {
	p := NewRangeListParams()
	p = p.WithDefaults()
}

func TestRangeListParams_WithTimeout(t *testing.T) {
	p := NewRangeListParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestRangeListParams_WithContext(t *testing.T) {
	p := NewRangeListParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestRangeListParams_WithHTTPClient(t *testing.T) {
	p := NewRangeListParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestRangeListParams_WithFields(t *testing.T) {
	p := NewRangeListParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestRangeListParams_WithFilter(t *testing.T) {
	p := NewRangeListParams()
	p = p.WithFilter(swag.String("key==value"))
	require.NotNil(t, p.Filter)
	assert.Equal(t, "key==value", *p.Filter)
}

func TestRangeListParams_WithLimit(t *testing.T) {
	p := NewRangeListParams()
	p = p.WithLimit(swag.Int64(20))
	require.NotNil(t, p.Limit)
	assert.Equal(t, int64(20), *p.Limit)
}

func TestRangeListParams_WithOffset(t *testing.T) {
	p := NewRangeListParams()
	p = p.WithOffset(swag.Int64(20))
	require.NotNil(t, p.Offset)
	assert.Equal(t, int64(20), *p.Offset)
}

func TestRangeListParams_WithOrderBy(t *testing.T) {
	p := NewRangeListParams()
	p = p.WithOrderBy(swag.String("desc"))
	require.NotNil(t, p.OrderBy)
	assert.Equal(t, "desc", *p.OrderBy)
}

func TestRangeListParams_WithPageToken(t *testing.T) {
	p := NewRangeListParams()
	p = p.WithPageToken(swag.String("token"))
	require.NotNil(t, p.PageToken)
	assert.Equal(t, "token", *p.PageToken)
}

func TestRangeListParams_WithTfilter(t *testing.T) {
	p := NewRangeListParams()
	p = p.WithTfilter(swag.String("tfilter"))
	require.NotNil(t, p.Tfilter)
	assert.Equal(t, "tfilter", *p.Tfilter)
}

func TestRangeListParams_WithTorderBy(t *testing.T) {
	p := NewRangeListParams()
	p = p.WithTorderBy(swag.String("desc"))
	require.NotNil(t, p.TorderBy)
	assert.Equal(t, "desc", *p.TorderBy)
}
