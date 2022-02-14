package option_group

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

func TestNewOptionGroupReadParams(t *testing.T) {
	p := NewOptionGroupReadParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewOptionGroupReadParamsWithTimeout(t *testing.T) {
	p := NewOptionGroupReadParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewOptionGroupReadParamsWithContext(t *testing.T) {
	p := NewOptionGroupReadParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewOptionGroupReadParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewOptionGroupReadParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionGroupReadParams_WithDefaults(t *testing.T) {
	p := NewOptionGroupReadParams()
	p = p.WithDefaults()
}

func TestOptionGroupReadParams_WithTimeout(t *testing.T) {
	p := NewOptionGroupReadParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestOptionGroupReadParams_WithContext(t *testing.T) {
	p := NewOptionGroupReadParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestOptionGroupReadParams_WithHTTPClient(t *testing.T) {
	p := NewOptionGroupReadParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionGroupReadParams_WithFields(t *testing.T) {
	p := NewOptionGroupReadParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestOptionGroupReadParams_WithID(t *testing.T) {
	p := NewOptionGroupReadParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
