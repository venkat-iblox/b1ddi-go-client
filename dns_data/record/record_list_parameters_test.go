package record

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

func TestNewRecordListParams(t *testing.T) {
	p := NewRecordListParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewRecordListParamsWithTimeout(t *testing.T) {
	p := NewRecordListParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewRecordListParamsWithContext(t *testing.T) {
	p := NewRecordListParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewRecordListParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewRecordListParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestRecordListParams_WithDefaults(t *testing.T) {
	p := NewRecordListParams()
	p = p.WithDefaults()
}

func TestRecordListParams_WithTimeout(t *testing.T) {
	p := NewRecordListParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestRecordListParams_WithContext(t *testing.T) {
	p := NewRecordListParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestRecordListParams_WithHTTPClient(t *testing.T) {
	p := NewRecordListParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestRecordListParams_WithFields(t *testing.T) {
	p := NewRecordListParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestRecordListParams_WithFilter(t *testing.T) {
	p := NewRecordListParams()
	p = p.WithFilter(swag.String("key==value"))
	require.NotNil(t, p.Filter)
	assert.Equal(t, "key==value", *p.Filter)
}

func TestRecordListParams_WithLimit(t *testing.T) {
	p := NewRecordListParams()
	p = p.WithLimit(swag.Int64(20))
	require.NotNil(t, p.Limit)
	assert.Equal(t, int64(20), *p.Limit)
}

func TestRecordListParams_WithOffset(t *testing.T) {
	p := NewRecordListParams()
	p = p.WithOffset(swag.Int64(20))
	require.NotNil(t, p.Offset)
	assert.Equal(t, int64(20), *p.Offset)
}

func TestRecordListParams_WithOrderBy(t *testing.T) {
	p := NewRecordListParams()
	p = p.WithOrderBy(swag.String("desc"))
	require.NotNil(t, p.OrderBy)
	assert.Equal(t, "desc", *p.OrderBy)
}

func TestRecordListParams_WithPageToken(t *testing.T) {
	p := NewRecordListParams()
	p = p.WithPageToken(swag.String("token"))
	require.NotNil(t, p.PageToken)
	assert.Equal(t, "token", *p.PageToken)
}

func TestRecordListParams_WithTfilter(t *testing.T) {
	p := NewRecordListParams()
	p = p.WithTfilter(swag.String("tfilter"))
	require.NotNil(t, p.Tfilter)
	assert.Equal(t, "tfilter", *p.Tfilter)
}

func TestRecordListParams_WithTorderBy(t *testing.T) {
	p := NewRecordListParams()
	p = p.WithTorderBy(swag.String("desc"))
	require.NotNil(t, p.TorderBy)
	assert.Equal(t, "desc", *p.TorderBy)
}
