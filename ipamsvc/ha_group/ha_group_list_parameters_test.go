package ha_group

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

func TestNewHaGroupListParams(t *testing.T) {
	p := NewHaGroupListParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewHaGroupListParamsWithTimeout(t *testing.T) {
	p := NewHaGroupListParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewHaGroupListParamsWithContext(t *testing.T) {
	p := NewHaGroupListParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewHaGroupListParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewHaGroupListParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestHaGroupListParams_WithDefaults(t *testing.T) {
	p := NewHaGroupListParams()
	p = p.WithDefaults()
}

func TestHaGroupListParams_WithTimeout(t *testing.T) {
	p := NewHaGroupListParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestHaGroupListParams_WithContext(t *testing.T) {
	p := NewHaGroupListParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestHaGroupListParams_WithHTTPClient(t *testing.T) {
	p := NewHaGroupListParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestHaGroupListParams_WithFields(t *testing.T) {
	p := NewHaGroupListParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestHaGroupListParams_WithFilter(t *testing.T) {
	p := NewHaGroupListParams()
	p = p.WithFilter(swag.String("key==value"))
	require.NotNil(t, p.Filter)
	assert.Equal(t, "key==value", *p.Filter)
}

func TestHaGroupListParams_WithLimit(t *testing.T) {
	p := NewHaGroupListParams()
	p = p.WithLimit(swag.Int64(20))
	require.NotNil(t, p.Limit)
	assert.Equal(t, int64(20), *p.Limit)
}

func TestHaGroupListParams_WithOffset(t *testing.T) {
	p := NewHaGroupListParams()
	p = p.WithOffset(swag.Int64(20))
	require.NotNil(t, p.Offset)
	assert.Equal(t, int64(20), *p.Offset)
}

func TestHaGroupListParams_WithOrderBy(t *testing.T) {
	p := NewHaGroupListParams()
	p = p.WithOrderBy(swag.String("desc"))
	require.NotNil(t, p.OrderBy)
	assert.Equal(t, "desc", *p.OrderBy)
}

func TestHaGroupListParams_WithPageToken(t *testing.T) {
	p := NewHaGroupListParams()
	p = p.WithPageToken(swag.String("page-token"))
	require.NotNil(t, p.PageToken)
	assert.Equal(t, "page-token", *p.PageToken)
}

func TestHaGroupListParams_WithTfilter(t *testing.T) {
	p := NewHaGroupListParams()
	p = p.WithTfilter(swag.String("tfilter"))
	require.NotNil(t, p.Tfilter)
	assert.Equal(t, "tfilter", *p.Tfilter)
}

func TestHaGroupListParams_WithTOrderBy(t *testing.T) {
	p := NewHaGroupListParams()
	p = p.WithTorderBy(swag.String("desc"))
	require.NotNil(t, p.TorderBy)
	assert.Equal(t, "desc", *p.TorderBy)
}
