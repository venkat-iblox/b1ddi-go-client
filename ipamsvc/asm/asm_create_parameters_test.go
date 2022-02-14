package asm

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

func TestNewAsmCreateParams(t *testing.T) {
	p := NewAsmCreateParams()
	require.NotNil(t, p.timeout)
	assert.Equal(t, client.DefaultTimeout, p.timeout)
}

func TestNewAsmCreateParamsWithTimeout(t *testing.T) {
	p := NewAsmCreateParamsWithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestNewAsmCreateParamsWithContext(t *testing.T) {
	p := NewAsmCreateParamsWithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestNewAsmCreateParamsWithHTTPClient(t *testing.T) {
	cli := &http.Client{}
	p := NewAsmCreateParamsWithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAsmCreateParams_WithDefaults(t *testing.T) {
	p := NewAsmCreateParams()
	p = p.WithDefaults()
}

func TestAsmCreateParams_WithTimeout(t *testing.T) {
	p := NewAsmCreateParams()
	p = p.WithTimeout(time.Minute * 5)
	require.NotNil(t, p.timeout)
	assert.Equal(t, time.Minute*5, p.timeout)
}

func TestAsmCreateParams_WithContext(t *testing.T) {
	p := NewAsmCreateParams()
	p = p.WithContext(context.TODO())
	require.NotNil(t, p.Context)
	assert.Equal(t, context.TODO(), p.Context)
}

func TestAsmCreateParams_WithHTTPClient(t *testing.T) {
	p := NewAsmCreateParams()
	cli := &http.Client{}
	p = p.WithHTTPClient(cli)
	require.NotNil(t, p.HTTPClient)
	assert.Equal(t, cli, p.HTTPClient)
}

func TestAsmCreateParams_WithBody(t *testing.T) {
	p := NewAsmCreateParams()
	b := &models.IpamsvcASM{}
	p = p.WithBody(b)
	require.NotNil(t, p.Body)
	assert.Equal(t, b, p.Body)
}
