package option_code

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

func TestNewOptionCodeReadParams(t *testing.T) {
	p := NewOptionCodeReadParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewOptionCodeReadParamsWithTimeout(t *testing.T) {
	p := NewOptionCodeReadParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewOptionCodeReadParamsWithContext(t *testing.T) {
	p := NewOptionCodeReadParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewOptionCodeReadParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewOptionCodeReadParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionCodeReadParams_WithDefaults(t *testing.T) {
	p := NewOptionCodeReadParams()
	p = p.WithDefaults()
}

func TestOptionCodeReadParams_WithTimeout(t *testing.T) {
	p := NewOptionCodeReadParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestOptionCodeReadParams_WithContext(t *testing.T) {
	p := NewOptionCodeReadParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestOptionCodeReadParams_WithHTTPClient(t *testing.T) {
	p := NewOptionCodeReadParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestOptionCodeReadParams_WithFields(t *testing.T) {
	p := NewOptionCodeReadParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestOptionCodeReadParams_WithID(t *testing.T) {
	p := NewOptionCodeReadParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
