package hardware_filter

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

func TestNewHardwareFilterListParams(t *testing.T) {
	p := NewHardwareFilterListParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewHardwareFilterListParamsWithTimeout(t *testing.T) {
	p := NewHardwareFilterListParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewHardwareFilterListParamsWithContext(t *testing.T) {
	p := NewHardwareFilterListParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewHardwareFilterListParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewHardwareFilterListParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestHardwareFilterListParams_WithDefaults(t *testing.T) {
	p := NewHardwareFilterListParams()
	p = p.WithDefaults()
}

func TestHardwareFilterListParams_WithTimeout(t *testing.T) {
	p := NewHardwareFilterListParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestHardwareFilterListParams_WithContext(t *testing.T) {
	p := NewHardwareFilterListParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestHardwareFilterListParams_WithHTTPClient(t *testing.T) {
	p := NewHardwareFilterListParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestHardwareFilterListParams_WithFields(t *testing.T) {
	p := NewHardwareFilterListParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestHardwareFilterListParams_WithFilter(t *testing.T) {
	p := NewHardwareFilterListParams()
	p = p.WithFilter(swag.String("key==value"))
	require.NotNil(t, p.Filter)
	assert.Equal(t, "key==value", *p.Filter)
}

func TestHardwareFilterListParams_WithLimit(t *testing.T) {
	p := NewHardwareFilterListParams()
	p = p.WithLimit(swag.Int64(20))
	require.NotNil(t, p.Limit)
	assert.Equal(t, int64(20), *p.Limit)
}

func TestHardwareFilterListParams_WithOffset(t *testing.T) {
	p := NewHardwareFilterListParams()
	p = p.WithOffset(swag.Int64(20))
	require.NotNil(t, p.Offset)
	assert.Equal(t, int64(20), *p.Offset)
}

func TestHardwareFilterListParams_WithOrderBy(t *testing.T) {
	p := NewHardwareFilterListParams()
	p = p.WithOrderBy(swag.String("desc"))
	require.NotNil(t, p.OrderBy)
	assert.Equal(t, "desc", *p.OrderBy)
}

func TestHardwareFilterListParams_WithPageToken(t *testing.T) {
	p := NewHardwareFilterListParams()
	p = p.WithPageToken(swag.String("page-token"))
	require.NotNil(t, p.PageToken)
	assert.Equal(t, "page-token", *p.PageToken)
}

func TestHardwareFilterListParams_WithTfilter(t *testing.T) {
	p := NewHardwareFilterListParams()
	p = p.WithTfilter(swag.String("tfilter"))
	require.NotNil(t, p.Tfilter)
	assert.Equal(t, "tfilter", *p.Tfilter)
}

func TestHardwareFilterListParams_WithTOrderBy(t *testing.T) {
	p := NewHardwareFilterListParams()
	p = p.WithTorderBy(swag.String("desc"))
	require.NotNil(t, p.TorderBy)
	assert.Equal(t, "desc", *p.TorderBy)
}
