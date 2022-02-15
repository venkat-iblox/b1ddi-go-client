package address

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

func TestNewAddressListParams(t *testing.T) {
	p := NewAddressListParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAddressListParamsWithTimeout(t *testing.T) {
	p := NewAddressListParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAddressListParamsWithContext(t *testing.T) {
	p := NewAddressListParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAddressListParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAddressListParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressListParams_WithDefaults(t *testing.T) {
	p := NewAddressListParams()
	p = p.WithDefaults()
}

func TestAddressListParams_WithTimeout(t *testing.T) {
	p := NewAddressListParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAddressListParams_WithContext(t *testing.T) {
	p := NewAddressListParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAddressListParams_WithHTTPClient(t *testing.T) {
	p := NewAddressListParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressListParams_WithFields(t *testing.T) {
	p := NewAddressListParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestAddressListParams_WithFilter(t *testing.T) {
	p := NewAddressListParams()
	p = p.WithFilter(swag.String("key==value"))
	require.NotNil(t, p.Filter)
	assert.Equal(t, "key==value", *p.Filter)
}

func TestAddressListParams_WithLimit(t *testing.T) {
	p := NewAddressListParams()
	p = p.WithLimit(swag.Int64(20))
	require.NotNil(t, p.Limit)
	assert.Equal(t, int64(20), *p.Limit)
}

func TestAddressListParams_WithOffset(t *testing.T) {
	p := NewAddressListParams()
	p = p.WithOffset(swag.Int64(20))
	require.NotNil(t, p.Offset)
	assert.Equal(t, int64(20), *p.Offset)
}

func TestAddressListParams_WithOrderBy(t *testing.T) {
	p := NewAddressListParams()
	p = p.WithOrderBy(swag.String("desc"))
	require.NotNil(t, p.OrderBy)
	assert.Equal(t, "desc", *p.OrderBy)
}

func TestAddressListParams_WithPageToken(t *testing.T) {
	p := NewAddressListParams()
	p = p.WithPageToken(swag.String("token"))
	require.NotNil(t, p.PageToken)
	assert.Equal(t, "token", *p.PageToken)
}

func TestAddressListParams_WithTfilter(t *testing.T) {
	p := NewAddressListParams()
	p = p.WithTfilter(swag.String("tfilter"))
	require.NotNil(t, p.Tfilter)
	assert.Equal(t, "tfilter", *p.Tfilter)
}

func TestAddressListParams_WithTorderBy(t *testing.T) {
	p := NewAddressListParams()
	p = p.WithTorderBy(swag.String("desc"))
	require.NotNil(t, p.TorderBy)
	assert.Equal(t, "desc", *p.TorderBy)
}
