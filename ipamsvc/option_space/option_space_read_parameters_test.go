package option_space

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

func TestNewOptionSpaceReadParams(t *testing.T) {
	p := NewOptionSpaceReadParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewOptionSpaceReadParamsWithTimeout(t *testing.T) {
	p := NewOptionSpaceReadParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewOptionSpaceReadParamsWithContext(t *testing.T) {
	p := NewOptionSpaceReadParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewOptionSpaceReadParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewOptionSpaceReadParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionSpaceReadParams_WithDefaults(t *testing.T) {
	p := NewOptionSpaceReadParams()
	p = p.WithDefaults()
}

func TestOptionSpaceReadParams_WithTimeout(t *testing.T) {
	p := NewOptionSpaceReadParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestOptionSpaceReadParams_WithContext(t *testing.T) {
	p := NewOptionSpaceReadParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestOptionSpaceReadParams_WithHTTPClient(t *testing.T) {
	p := NewOptionSpaceReadParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionSpaceReadParams_WithFields(t *testing.T) {
	p := NewOptionSpaceReadParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestOptionSpaceReadParams_WithID(t *testing.T) {
	p := NewOptionSpaceReadParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
