package option_space

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

func TestNewOptionSpaceListParams(t *testing.T) {
	p := NewOptionSpaceListParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewOptionSpaceListParamsWithTimeout(t *testing.T) {
	p := NewOptionSpaceListParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewOptionSpaceListParamsWithContext(t *testing.T) {
	p := NewOptionSpaceListParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewOptionSpaceListParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewOptionSpaceListParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionSpaceListParams_WithDefaults(t *testing.T) {
	p := NewOptionSpaceListParams()
	p = p.WithDefaults()
}

func TestOptionSpaceListParams_WithTimeout(t *testing.T) {
	p := NewOptionSpaceListParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestOptionSpaceListParams_WithContext(t *testing.T) {
	p := NewOptionSpaceListParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestOptionSpaceListParams_WithHTTPClient(t *testing.T) {
	p := NewOptionSpaceListParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionSpaceListParams_WithFields(t *testing.T) {
	p := NewOptionSpaceListParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestOptionSpaceListParams_WithFilter(t *testing.T) {
	p := NewOptionSpaceListParams()
	p = p.WithFilter(swag.String("key==value"))
	require.NotNil(t, p.Filter)
	assert.Equal(t, "key==value", *p.Filter)
}

func TestOptionSpaceListParams_WithLimit(t *testing.T) {
	p := NewOptionSpaceListParams()
	p = p.WithLimit(swag.Int64(20))
	require.NotNil(t, p.Limit)
	assert.Equal(t, int64(20), *p.Limit)
}

func TestOptionSpaceListParams_WithOffset(t *testing.T) {
	p := NewOptionSpaceListParams()
	p = p.WithOffset(swag.Int64(20))
	require.NotNil(t, p.Offset)
	assert.Equal(t, int64(20), *p.Offset)
}

func TestOptionSpaceListParams_WithOrderBy(t *testing.T) {
	p := NewOptionSpaceListParams()
	p = p.WithOrderBy(swag.String("desc"))
	require.NotNil(t, p.OrderBy)
	assert.Equal(t, "desc", *p.OrderBy)
}

func TestOptionSpaceListParams_WithPageToken(t *testing.T) {
	p := NewOptionSpaceListParams()
	p = p.WithPageToken(swag.String("token"))
	require.NotNil(t, p.PageToken)
	assert.Equal(t, "token", *p.PageToken)
}

func TestOptionSpaceListParams_WithTfilter(t *testing.T) {
	p := NewOptionSpaceListParams()
	p = p.WithTfilter(swag.String("tfilter"))
	require.NotNil(t, p.Tfilter)
	assert.Equal(t, "tfilter", *p.Tfilter)
}

func TestOptionSpaceListParams_WithTorderBy(t *testing.T) {
	p := NewOptionSpaceListParams()
	p = p.WithTorderBy(swag.String("desc"))
	require.NotNil(t, p.TorderBy)
	assert.Equal(t, "desc", *p.TorderBy)
}
