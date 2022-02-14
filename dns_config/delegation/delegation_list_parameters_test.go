package delegation

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

func TestNewDelegationListParams(t *testing.T) {
	p := NewDelegationListParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewDelegationListParamsWithTimeout(t *testing.T) {
	p := NewDelegationListParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewDelegationListParamsWithContext(t *testing.T) {
	p := NewDelegationListParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewDelegationListParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewDelegationListParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestDelegationListParams_WithDefaults(t *testing.T) {
	p := NewDelegationListParams()
	p = p.WithDefaults()
}

func TestDelegationListParams_WithTimeout(t *testing.T) {
	p := NewDelegationListParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestDelegationListParams_WithContext(t *testing.T) {
	p := NewDelegationListParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestDelegationListParams_WithHTTPClient(t *testing.T) {
	p := NewDelegationListParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestDelegationListParams_WithFields(t *testing.T) {
	p := NewDelegationListParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestDelegationListParams_WithFilter(t *testing.T) {
	p := NewDelegationListParams()
	p = p.WithFilter(swag.String("key==value"))
	require.NotNil(t, p.Filter)
	assert.Equal(t, "key==value", *p.Filter)
}

func TestDelegationListParams_WithLimit(t *testing.T) {
	p := NewDelegationListParams()
	p = p.WithLimit(swag.Int64(20))
	require.NotNil(t, p.Limit)
	assert.Equal(t, int64(20), *p.Limit)
}

func TestDelegationListParams_WithOffset(t *testing.T) {
	p := NewDelegationListParams()
	p = p.WithOffset(swag.Int64(20))
	require.NotNil(t, p.Offset)
	assert.Equal(t, int64(20), *p.Offset)
}

func TestDelegationListParams_WithOrderBy(t *testing.T) {
	p := NewDelegationListParams()
	p = p.WithOrderBy(swag.String("desc"))
	require.NotNil(t, p.OrderBy)
	assert.Equal(t, "desc", *p.OrderBy)
}

func TestDelegationListParams_WithPageToken(t *testing.T) {
	p := NewDelegationListParams()
	p = p.WithPageToken(swag.String("token"))
	require.NotNil(t, p.PageToken)
	assert.Equal(t, "token", *p.PageToken)
}

func TestDelegationListParams_WithTfilter(t *testing.T) {
	p := NewDelegationListParams()
	p = p.WithTfilter(swag.String("tfilter"))
	require.NotNil(t, p.Tfilter)
	assert.Equal(t, "tfilter", *p.Tfilter)
}

func TestDelegationListParams_WithTorderBy(t *testing.T) {
	p := NewDelegationListParams()
	p = p.WithTorderBy(swag.String("desc"))
	require.NotNil(t, p.TorderBy)
	assert.Equal(t, "desc", *p.TorderBy)
}
