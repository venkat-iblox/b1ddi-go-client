package forward_zone

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

func TestNewForwardZoneListParams(t *testing.T) {
	p := NewForwardZoneListParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewForwardZoneListParamsWithTimeout(t *testing.T) {
	p := NewForwardZoneListParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewForwardZoneListParamsWithContext(t *testing.T) {
	p := NewForwardZoneListParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewForwardZoneListParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewForwardZoneListParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestForwardZoneListParams_WithDefaults(t *testing.T) {
	p := NewForwardZoneListParams()
	p = p.WithDefaults()
}

func TestForwardZoneListParams_WithTimeout(t *testing.T) {
	p := NewForwardZoneListParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestForwardZoneListParams_WithContext(t *testing.T) {
	p := NewForwardZoneListParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestForwardZoneListParams_WithHTTPClient(t *testing.T) {
	p := NewForwardZoneListParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestForwardZoneListParams_WithFields(t *testing.T) {
	p := NewForwardZoneListParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestForwardZoneListParams_WithFilter(t *testing.T) {
	p := NewForwardZoneListParams()
	p = p.WithFilter(swag.String("key==value"))
	require.NotNil(t, p.Filter)
	assert.Equal(t, "key==value", *p.Filter)
}

func TestForwardZoneListParams_WithLimit(t *testing.T) {
	p := NewForwardZoneListParams()
	p = p.WithLimit(swag.Int64(20))
	require.NotNil(t, p.Limit)
	assert.Equal(t, int64(20), *p.Limit)
}

func TestForwardZoneListParams_WithOffset(t *testing.T) {
	p := NewForwardZoneListParams()
	p = p.WithOffset(swag.Int64(20))
	require.NotNil(t, p.Offset)
	assert.Equal(t, int64(20), *p.Offset)
}

func TestForwardZoneListParams_WithOrderBy(t *testing.T) {
	p := NewForwardZoneListParams()
	p = p.WithOrderBy(swag.String("desc"))
	require.NotNil(t, p.OrderBy)
	assert.Equal(t, "desc", *p.OrderBy)
}

func TestForwardZoneListParams_WithPageToken(t *testing.T) {
	p := NewForwardZoneListParams()
	p = p.WithPageToken(swag.String("token"))
	require.NotNil(t, p.PageToken)
	assert.Equal(t, "token", *p.PageToken)
}

func TestForwardZoneListParams_WithTfilter(t *testing.T) {
	p := NewForwardZoneListParams()
	p = p.WithTfilter(swag.String("tfilter"))
	require.NotNil(t, p.Tfilter)
	assert.Equal(t, "tfilter", *p.Tfilter)
}

func TestForwardZoneListParams_WithTorderBy(t *testing.T) {
	p := NewForwardZoneListParams()
	p = p.WithTorderBy(swag.String("desc"))
	require.NotNil(t, p.TorderBy)
	assert.Equal(t, "desc", *p.TorderBy)
}
