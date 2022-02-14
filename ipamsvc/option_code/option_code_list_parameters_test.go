package option_code

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

func TestNewOptionCodeListParams(t *testing.T) {
	p := NewOptionCodeListParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewOptionCodeListParamsWithTimeout(t *testing.T) {
	p := NewOptionCodeListParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewOptionCodeListParamsWithContext(t *testing.T) {
	p := NewOptionCodeListParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewOptionCodeListParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewOptionCodeListParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionCodeListParams_WithDefaults(t *testing.T) {
	p := NewOptionCodeListParams()
	p = p.WithDefaults()
}

func TestOptionCodeListParams_WithTimeout(t *testing.T) {
	p := NewOptionCodeListParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestOptionCodeListParams_WithContext(t *testing.T) {
	p := NewOptionCodeListParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestOptionCodeListParams_WithHTTPClient(t *testing.T) {
	p := NewOptionCodeListParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionCodeListParams_WithFields(t *testing.T) {
	p := NewOptionCodeListParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestOptionCodeListParams_WithFilter(t *testing.T) {
	p := NewOptionCodeListParams()
	p = p.WithFilter(swag.String("key==value"))
	require.NotNil(t, p.Filter)
	assert.Equal(t, "key==value", *p.Filter)
}

func TestOptionCodeListParams_WithLimit(t *testing.T) {
	p := NewOptionCodeListParams()
	p = p.WithLimit(swag.Int64(20))
	require.NotNil(t, p.Limit)
	assert.Equal(t, int64(20), *p.Limit)
}

func TestOptionCodeListParams_WithOffset(t *testing.T) {
	p := NewOptionCodeListParams()
	p = p.WithOffset(swag.Int64(20))
	require.NotNil(t, p.Offset)
	assert.Equal(t, int64(20), *p.Offset)
}

func TestOptionCodeListParams_WithOrderBy(t *testing.T) {
	p := NewOptionCodeListParams()
	p = p.WithOrderBy(swag.String("desc"))
	require.NotNil(t, p.OrderBy)
	assert.Equal(t, "desc", *p.OrderBy)
}

func TestOptionCodeListParams_WithPageToken(t *testing.T) {
	p := NewOptionCodeListParams()
	p = p.WithPageToken(swag.String("page-token"))
	require.NotNil(t, p.PageToken)
	assert.Equal(t, "page-token", *p.PageToken)
}
