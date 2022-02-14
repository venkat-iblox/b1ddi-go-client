package auth_zone

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

func TestNewAuthZoneListParams(t *testing.T) {
	p := NewAuthZoneListParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAuthZoneListParamsWithTimeout(t *testing.T) {
	p := NewAuthZoneListParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAuthZoneListParamsWithContext(t *testing.T) {
	p := NewAuthZoneListParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAuthZoneListParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAuthZoneListParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAuthZoneListParams_WithDefaults(t *testing.T) {
	p := NewAuthZoneListParams()
	p = p.WithDefaults()
}

func TestAuthZoneListParams_WithTimeout(t *testing.T) {
	p := NewAuthZoneListParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAuthZoneListParams_WithContext(t *testing.T) {
	p := NewAuthZoneListParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAuthZoneListParams_WithHTTPClient(t *testing.T) {
	p := NewAuthZoneListParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAuthZoneListParams_WithFields(t *testing.T) {
	p := NewAuthZoneListParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestAuthZoneListParams_WithFilter(t *testing.T) {
	p := NewAuthZoneListParams()
	p = p.WithFilter(swag.String("key==value"))
	require.NotNil(t, p.Filter)
	assert.Equal(t, "key==value", *p.Filter)
}

func TestAuthZoneListParams_WithLimit(t *testing.T) {
	p := NewAuthZoneListParams()
	p = p.WithLimit(swag.Int64(20))
	require.NotNil(t, p.Limit)
	assert.Equal(t, int64(20), *p.Limit)
}

func TestAuthZoneListParams_WithOffset(t *testing.T) {
	p := NewAuthZoneListParams()
	p = p.WithOffset(swag.Int64(20))
	require.NotNil(t, p.Offset)
	assert.Equal(t, int64(20), *p.Offset)
}

func TestAuthZoneListParams_WithOrderBy(t *testing.T) {
	p := NewAuthZoneListParams()
	p = p.WithOrderBy(swag.String("desc"))
	require.NotNil(t, p.OrderBy)
	assert.Equal(t, "desc", *p.OrderBy)
}

func TestAuthZoneListParams_WithPageToken(t *testing.T) {
	p := NewAuthZoneListParams()
	p = p.WithPageToken(swag.String("token"))
	require.NotNil(t, p.PageToken)
	assert.Equal(t, "token", *p.PageToken)
}

func TestAuthZoneListParams_WithTfilter(t *testing.T) {
	p := NewAuthZoneListParams()
	p = p.WithTfilter(swag.String("tfilter"))
	require.NotNil(t, p.Tfilter)
	assert.Equal(t, "tfilter", *p.Tfilter)
}

func TestAuthZoneListParams_WithTorderBy(t *testing.T) {
	p := NewAuthZoneListParams()
	p = p.WithTorderBy(swag.String("desc"))
	require.NotNil(t, p.TorderBy)
	assert.Equal(t, "desc", *p.TorderBy)
}
