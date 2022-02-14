package option_group

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

func TestNewOptionGroupListParams(t *testing.T) {
	p := NewOptionGroupListParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewOptionGroupListParamsWithTimeout(t *testing.T) {
	p := NewOptionGroupListParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewOptionGroupListParamsWithContext(t *testing.T) {
	p := NewOptionGroupListParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewOptionGroupListParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewOptionGroupListParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionGroupListParams_WithDefaults(t *testing.T) {
	p := NewOptionGroupListParams()
	p = p.WithDefaults()
}

func TestOptionGroupListParams_WithTimeout(t *testing.T) {
	p := NewOptionGroupListParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestOptionGroupListParams_WithContext(t *testing.T) {
	p := NewOptionGroupListParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestOptionGroupListParams_WithHTTPClient(t *testing.T) {
	p := NewOptionGroupListParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionGroupListParams_WithFields(t *testing.T) {
	p := NewOptionGroupListParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestOptionGroupListParams_WithFilter(t *testing.T) {
	p := NewOptionGroupListParams()
	p = p.WithFilter(swag.String("key==value"))
	require.NotNil(t, p.Filter)
	assert.Equal(t, "key==value", *p.Filter)
}

func TestOptionGroupListParams_WithLimit(t *testing.T) {
	p := NewOptionGroupListParams()
	p = p.WithLimit(swag.Int64(20))
	require.NotNil(t, p.Limit)
	assert.Equal(t, int64(20), *p.Limit)
}

func TestOptionGroupListParams_WithOffset(t *testing.T) {
	p := NewOptionGroupListParams()
	p = p.WithOffset(swag.Int64(20))
	require.NotNil(t, p.Offset)
	assert.Equal(t, int64(20), *p.Offset)
}

func TestOptionGroupListParams_WithOrderBy(t *testing.T) {
	p := NewOptionGroupListParams()
	p = p.WithOrderBy(swag.String("desc"))
	require.NotNil(t, p.OrderBy)
	assert.Equal(t, "desc", *p.OrderBy)
}

func TestOptionGroupListParams_WithPageToken(t *testing.T) {
	p := NewOptionGroupListParams()
	p = p.WithPageToken(swag.String("token"))
	require.NotNil(t, p.PageToken)
	assert.Equal(t, "token", *p.PageToken)
}

func TestOptionGroupListParams_WithTfilter(t *testing.T) {
	p := NewOptionGroupListParams()
	p = p.WithTfilter(swag.String("tfilter"))
	require.NotNil(t, p.Tfilter)
	assert.Equal(t, "tfilter", *p.Tfilter)
}

func TestOptionGroupListParams_WithTorderBy(t *testing.T) {
	p := NewOptionGroupListParams()
	p = p.WithTorderBy(swag.String("desc"))
	require.NotNil(t, p.TorderBy)
	assert.Equal(t, "desc", *p.TorderBy)
}
