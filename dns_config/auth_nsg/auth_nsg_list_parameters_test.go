package auth_nsg

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

func TestNewAuthNsgListParams(t *testing.T) {
	p := NewAuthNsgListParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAuthNsgListParamsWithTimeout(t *testing.T) {
	p := NewAuthNsgListParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAuthNsgListParamsWithContext(t *testing.T) {
	p := NewAuthNsgListParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAuthNsgListParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAuthNsgListParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAuthNsgListParams_WithDefaults(t *testing.T) {
	p := NewAuthNsgListParams()
	p = p.WithDefaults()
}

func TestAuthNsgListParams_WithTimeout(t *testing.T) {
	p := NewAuthNsgListParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAuthNsgListParams_WithContext(t *testing.T) {
	p := NewAuthNsgListParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAuthNsgListParams_WithHTTPClient(t *testing.T) {
	p := NewAuthNsgListParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAuthNsgListParams_WithFields(t *testing.T) {
	p := NewAuthNsgListParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestAuthNsgListParams_WithFilter(t *testing.T) {
	p := NewAuthNsgListParams()
	p = p.WithFilter(swag.String("key==value"))
	require.NotNil(t, p.Filter)
	assert.Equal(t, "key==value", *p.Filter)
}

func TestAuthNsgListParams_WithLimit(t *testing.T) {
	p := NewAuthNsgListParams()
	p = p.WithLimit(swag.Int64(20))
	require.NotNil(t, p.Limit)
	assert.Equal(t, int64(20), *p.Limit)
}

func TestAuthNsgListParams_WithOffset(t *testing.T) {
	p := NewAuthNsgListParams()
	p = p.WithOffset(swag.Int64(20))
	require.NotNil(t, p.Offset)
	assert.Equal(t, int64(20), *p.Offset)
}

func TestAuthNsgListParams_WithOrderBy(t *testing.T) {
	p := NewAuthNsgListParams()
	p = p.WithOrderBy(swag.String("desc"))
	require.NotNil(t, p.OrderBy)
	assert.Equal(t, "desc", *p.OrderBy)
}

func TestAuthNsgListParams_WithPageToken(t *testing.T) {
	p := NewAuthNsgListParams()
	p = p.WithPageToken(swag.String("token"))
	require.NotNil(t, p.PageToken)
	assert.Equal(t, "token", *p.PageToken)
}

func TestAuthNsgListParams_WithTfilter(t *testing.T) {
	p := NewAuthNsgListParams()
	p = p.WithTfilter(swag.String("tfilter"))
	require.NotNil(t, p.Tfilter)
	assert.Equal(t, "tfilter", *p.Tfilter)
}

func TestAuthNsgListParams_WithTorderBy(t *testing.T) {
	p := NewAuthNsgListParams()
	p = p.WithTorderBy(swag.String("desc"))
	require.NotNil(t, p.TorderBy)
	assert.Equal(t, "desc", *p.TorderBy)
}
