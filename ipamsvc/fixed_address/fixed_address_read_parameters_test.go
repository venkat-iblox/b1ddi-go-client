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

func TestNewFixedAddressReadParams(t *testing.T) {
	p := NewFixedAddressReadParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewFixedAddressReadParamsWithTimeout(t *testing.T) {
	p := NewFixedAddressReadParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewFixedAddressReadParamsWithContext(t *testing.T) {
	p := NewFixedAddressReadParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewFixedAddressReadParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewFixedAddressReadParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestFixedAddressReadParams_WithDefaults(t *testing.T) {
	p := NewFixedAddressReadParams()
	p = p.WithDefaults()
}

func TestFixedAddressReadParams_WithTimeout(t *testing.T) {
	p := NewFixedAddressReadParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestFixedAddressReadParams_WithContext(t *testing.T) {
	p := NewFixedAddressReadParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestFixedAddressReadParams_WithHTTPClient(t *testing.T) {
	p := NewFixedAddressReadParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestFixedAddressReadParams_WithFields(t *testing.T) {
	p := NewFixedAddressReadParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestFixedAddressReadParams_WithID(t *testing.T) {
	p := NewFixedAddressReadParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
