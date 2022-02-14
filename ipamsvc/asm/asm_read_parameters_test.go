package asm

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

func TestNewAsmReadParams(t *testing.T) {
	p := NewAsmReadParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAsmReadParamsWithTimeout(t *testing.T) {
	p := NewAsmReadParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAsmReadParamsWithContext(t *testing.T) {
	p := NewAsmReadParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAsmReadParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAsmReadParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAsmReadParams_WithDefaults(t *testing.T) {
	p := NewAsmReadParams()
	p = p.WithDefaults()
}

func TestAsmReadParams_WithTimeout(t *testing.T) {
	p := NewAsmReadParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAsmReadParams_WithContext(t *testing.T) {
	p := NewAsmReadParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAsmReadParams_WithHTTPClient(t *testing.T) {
	p := NewAsmReadParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAsmReadParams_WithFields(t *testing.T) {
	p := NewAsmReadParams()
	p = p.WithFields(swag.String("fields"))
	require.NotNil(t, p.Fields)
	assert.Equal(t, "fields", *p.Fields)
}

func TestAsmReadParams_WithID(t *testing.T) {
	p := NewAsmReadParams()
	p = p.WithID("test-id")
	require.NotNil(t, p.ID)
	assert.Equal(t, "test-id", p.ID)
}
