package address_block

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

func TestNewAddressBlockListParams(t *testing.T) {
	p := NewAddressBlockListParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAddressBlockListParamsWithTimeout(t *testing.T) {
	p := NewAddressBlockListParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAddressBlockListParamsWithContext(t *testing.T) {
	p := NewAddressBlockListParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAddressBlockListParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAddressBlockListParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressBlockListParams_WithDefaults(t *testing.T) {
	p := NewAddressBlockListParams()
	p = p.WithDefaults()
}

func TestAddressBlockListParams_WithTimeout(t *testing.T) {
	p := NewAddressBlockListParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAddressBlockListParams_WithContext(t *testing.T) {
	p := NewAddressBlockListParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAddressBlockListParams_WithHTTPClient(t *testing.T) {
	p := NewAddressBlockListParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressBlockListParams_WithFields(t *testing.T) {
	p := NewAddressBlockListParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestAddressBlockListParams_WithFilter(t *testing.T) {
	p := NewAddressBlockListParams()
	p = p.WithFilter(swag.String("key==value"))
	require.NotNil(t, p.Filter)
	assert.Equal(t, "key==value", *p.Filter)
}

func TestAddressBlockListParams_WithLimit(t *testing.T) {
	p := NewAddressBlockListParams()
	p = p.WithLimit(swag.Int64(20))
	require.NotNil(t, p.Limit)
	assert.Equal(t, int64(20), *p.Limit)
}
