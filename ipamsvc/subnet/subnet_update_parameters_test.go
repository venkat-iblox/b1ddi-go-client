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

func TestNewSubnetUpdateParams(t *testing.T) {
	p := NewSubnetUpdateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewSubnetUpdateParamsWithTimeout(t *testing.T) {
	p := NewSubnetUpdateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewSubnetUpdateParamsWithContext(t *testing.T) {
	p := NewSubnetUpdateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewSubnetUpdateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewSubnetUpdateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestSubnetUpdateParams_WithDefaults(t *testing.T) {
	p := NewSubnetUpdateParams()
	p = p.WithDefaults()
}

func TestSubnetUpdateParams_WithTimeout(t *testing.T) {
	p := NewSubnetUpdateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestSubnetUpdateParams_WithContext(t *testing.T) {
	p := NewSubnetUpdateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestSubnetUpdateParams_WithHTTPClient(t *testing.T) {
	p := NewSubnetUpdateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestSubnetUpdateParams_WithBody(t *testing.T) {
	p := NewSubnetUpdateParams()
	b := &models.IpamsvcSubnet{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}

func TestSubnetUpdateParams_WithID(t *testing.T) {
	p := NewSubnetUpdateParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
