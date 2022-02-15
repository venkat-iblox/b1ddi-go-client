package subnet

import (
	"context"
	"github.com/go-openapi/runtime/client"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewSubnetCopyParams(t *testing.T) {
	p := NewSubnetCopyParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewSubnetCopyParamsWithTimeout(t *testing.T) {
	p := NewSubnetCopyParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewSubnetCopyParamsWithContext(t *testing.T) {
	p := NewSubnetCopyParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewSubnetCopyParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewSubnetCopyParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestSubnetCopyParams_WithDefaults(t *testing.T) {
	p := NewSubnetCopyParams()
	p = p.WithDefaults()
}

func TestSubnetCopyParams_WithTimeout(t *testing.T) {
	p := NewSubnetCopyParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestSubnetCopyParams_WithContext(t *testing.T) {
	p := NewSubnetCopyParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestSubnetCopyParams_WithHTTPClient(t *testing.T) {
	p := NewSubnetCopyParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestSubnetCopyParams_WithBody(t *testing.T) {
	p := NewSubnetCopyParams()
	b := &models.IpamsvcCopySubnet{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}

func TestSubnetCopyParams_WithID(t *testing.T) {
	p := NewSubnetCopyParams()
	p = p.WithID("test-id")
	assert.Equal(t, "test-id", p.ID)
}
