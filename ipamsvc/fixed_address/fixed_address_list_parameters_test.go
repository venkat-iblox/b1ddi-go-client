package fixed_address

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

func TestNewFixedAddressListParams(t *testing.T) {
	p := NewFixedAddressListParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewFixedAddressListParamsWithTimeout(t *testing.T) {
	p := NewFixedAddressListParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewFixedAddressListParamsWithContext(t *testing.T) {
	p := NewFixedAddressListParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewFixedAddressListParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewFixedAddressListParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestFixedAddressListParams_WithDefaults(t *testing.T) {
	p := NewFixedAddressListParams()
	p = p.WithDefaults()
}

func TestFixedAddressListParams_WithTimeout(t *testing.T) {
	p := NewFixedAddressListParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestFixedAddressListParams_WithContext(t *testing.T) {
	p := NewFixedAddressListParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestFixedAddressListParams_WithHTTPClient(t *testing.T) {
	p := NewFixedAddressListParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestFixedAddressListParams_WithFields(t *testing.T) {
	p := NewFixedAddressListParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestFixedAddressListParams_WithFilter(t *testing.T) {
	p := NewFixedAddressListParams()
	p = p.WithFilter(swag.String("key==value"))
	require.NotNil(t, p.Filter)
	assert.Equal(t, "key==value", *p.Filter)
}

func TestFixedAddressListParams_WithLimit(t *testing.T) {
	p := NewFixedAddressListParams()
	p = p.WithLimit(swag.Int64(20))
	require.NotNil(t, p.Limit)
	assert.Equal(t, int64(20), *p.Limit)
}

func TestFixedAddressListParams_WithOffset(t *testing.T) {
	p := NewFixedAddressListParams()
	p = p.WithOffset(swag.Int64(20))
	require.NotNil(t, p.Offset)
	assert.Equal(t, int64(20), *p.Offset)
}

func TestFixedAddressListParams_WithOrderBy(t *testing.T) {
	p := NewFixedAddressListParams()
	p = p.WithOrderBy(swag.String("desc"))
	require.NotNil(t, p.OrderBy)
	assert.Equal(t, "desc", *p.OrderBy)
}

func TestFixedAddressListParams_WithPageToken(t *testing.T) {
	p := NewFixedAddressListParams()
	p = p.WithPageToken(swag.String("token"))
	require.NotNil(t, p.PageToken)
	assert.Equal(t, "token", *p.PageToken)
}

func TestFixedAddressListParams_WithTfilter(t *testing.T) {
	p := NewFixedAddressListParams()
	p = p.WithTfilter(swag.String("tfilter"))
	require.NotNil(t, p.Tfilter)
	assert.Equal(t, "tfilter", *p.Tfilter)
}

func TestFixedAddressListParams_WithTorderBy(t *testing.T) {
	p := NewFixedAddressListParams()
	p = p.WithTorderBy(swag.String("desc"))
	require.NotNil(t, p.TorderBy)
	assert.Equal(t, "desc", *p.TorderBy)
}
