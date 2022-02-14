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

func TestNewAddressReadParams(t *testing.T) {
	p := NewAddressReadParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAddressReadParamsWithTimeout(t *testing.T) {
	p := NewAddressReadParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAddressReadParamsWithContext(t *testing.T) {
	p := NewAddressReadParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAddressReadParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAddressReadParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressReadParams_WithDefaults(t *testing.T) {
	p := NewAddressReadParams()
	p = p.WithDefaults()
}

func TestAddressReadParams_WithTimeout(t *testing.T) {
	p := NewAddressReadParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAddressReadParams_WithContext(t *testing.T) {
	p := NewAddressReadParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAddressReadParams_WithHTTPClient(t *testing.T) {
	p := NewAddressReadParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAddressReadParams_WithFields(t *testing.T) {
	p := NewAddressReadParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestAddressReadParams_WithID(t *testing.T) {
	p := NewAddressReadParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
