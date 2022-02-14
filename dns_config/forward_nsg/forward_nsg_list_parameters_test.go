package forward_nsg

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

func TestNewForwardNsgListParams(t *testing.T) {
	p := NewForwardNsgListParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewForwardNsgListParamsWithTimeout(t *testing.T) {
	p := NewForwardNsgListParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewForwardNsgListParamsWithContext(t *testing.T) {
	p := NewForwardNsgListParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewForwardNsgListParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewForwardNsgListParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestForwardNsgListParams_WithDefaults(t *testing.T) {
	p := NewForwardNsgListParams()
	p = p.WithDefaults()
}

func TestForwardNsgListParams_WithTimeout(t *testing.T) {
	p := NewForwardNsgListParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestForwardNsgListParams_WithContext(t *testing.T) {
	p := NewForwardNsgListParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestForwardNsgListParams_WithHTTPClient(t *testing.T) {
	p := NewForwardNsgListParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestForwardNsgListParams_WithFields(t *testing.T) {
	p := NewForwardNsgListParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestForwardNsgListParams_WithFilter(t *testing.T) {
	p := NewForwardNsgListParams()
	p = p.WithFilter(swag.String("key==value"))
	require.NotNil(t, p.Filter)
	assert.Equal(t, "key==value", *p.Filter)
}

func TestForwardNsgListParams_WithLimit(t *testing.T) {
	p := NewForwardNsgListParams()
	p = p.WithLimit(swag.Int64(20))
	require.NotNil(t, p.Limit)
	assert.Equal(t, int64(20), *p.Limit)
}

func TestForwardNsgListParams_WithOffset(t *testing.T) {
	p := NewForwardNsgListParams()
	p = p.WithOffset(swag.Int64(20))
	require.NotNil(t, p.Offset)
	assert.Equal(t, int64(20), *p.Offset)
}

func TestForwardNsgListParams_WithOrderBy(t *testing.T) {
	p := NewForwardNsgListParams()
	p = p.WithOrderBy(swag.String("desc"))
	require.NotNil(t, p.OrderBy)
	assert.Equal(t, "desc", *p.OrderBy)
}

func TestForwardNsgListParams_WithPageToken(t *testing.T) {
	p := NewForwardNsgListParams()
	p = p.WithPageToken(swag.String("token"))
	require.NotNil(t, p.PageToken)
	assert.Equal(t, "token", *p.PageToken)
}

func TestForwardNsgListParams_WithTfilter(t *testing.T) {
	p := NewForwardNsgListParams()
	p = p.WithTfilter(swag.String("tfilter"))
	require.NotNil(t, p.Tfilter)
	assert.Equal(t, "tfilter", *p.Tfilter)
}

func TestForwardNsgListParams_WithTorderBy(t *testing.T) {
	p := NewForwardNsgListParams()
	p = p.WithTorderBy(swag.String("desc"))
	require.NotNil(t, p.TorderBy)
	assert.Equal(t, "desc", *p.TorderBy)
}
