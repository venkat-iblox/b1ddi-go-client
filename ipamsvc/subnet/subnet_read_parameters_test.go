package subnet

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

func TestNewSubnetReadParams(t *testing.T) {
	p := NewSubnetReadParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewSubnetReadParamsWithTimeout(t *testing.T) {
	p := NewSubnetReadParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewSubnetReadParamsWithContext(t *testing.T) {
	p := NewSubnetReadParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewSubnetReadParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewSubnetReadParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestSubnetReadParams_WithDefaults(t *testing.T) {
	p := NewSubnetReadParams()
	p = p.WithDefaults()
}

func TestSubnetReadParams_WithTimeout(t *testing.T) {
	p := NewSubnetReadParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestSubnetReadParams_WithContext(t *testing.T) {
	p := NewSubnetReadParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestSubnetReadParams_WithHTTPClient(t *testing.T) {
	p := NewSubnetReadParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestSubnetReadParams_WithFields(t *testing.T) {
	p := NewSubnetReadParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestSubnetReadParams_WithID(t *testing.T) {
	p := NewSubnetReadParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
